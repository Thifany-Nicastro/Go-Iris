package services

import (
	"go-iris/dtos"
	"go-iris/models"
	"go-iris/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todoService struct {
	Repository repositories.TodoRepository
}

type TodoService interface {
	GetTodos() []dtos.TodoResponse
	FindTodo(id string) (dtos.TodoResponse, error)
	CreateTodo(request dtos.TodoRequest) (interface{}, error)
	UpdateTodo(id string, request dtos.TodoRequest) int64
	DeleteTodo(id string) int64
	CompleteTodo(id string) int64
}

func NewTodoService(Repository repositories.TodoRepository) TodoService {
	return &todoService{
		Repository: Repository,
	}
}

func (s *todoService) GetTodos() []dtos.TodoResponse {
	return s.Repository.All()
}

func (s *todoService) FindTodo(id string) (dtos.TodoResponse, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	todo, err := s.Repository.FindById(objId)

	return dtos.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		IsCompleted: todo.IsCompleted,
	}, err
}

func (s *todoService) CreateTodo(request dtos.TodoRequest) (interface{}, error) {
	todo := models.Todo{
		ID:    primitive.NewObjectID(),
		Title: request.Title,
	}

	return s.Repository.Create(todo)
}

func (s *todoService) UpdateTodo(id string, request dtos.TodoRequest) int64 {
	fields := bson.M{
		"title": request.Title,
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.Update(objId, fields)
}

func (s *todoService) DeleteTodo(id string) int64 {
	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.Delete(objId)
}

func (s *todoService) CompleteTodo(id string) int64 {
	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.Update(objId, bson.M{"is_completed": true})
}
