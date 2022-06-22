package services

import (
	"go-iris/dtos"
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
	CreateTodo(request dtos.TodoRequest, userId string) (string, error)
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
	todos := s.Repository.All()

	return dtos.CreateTodoListResponse(todos)
}

func (s *todoService) FindTodo(id string) (dtos.TodoResponse, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	todo, err := s.Repository.FindById(objId)

	return dtos.CreateTodoResponse(todo), err
}

func (s *todoService) CreateTodo(request dtos.TodoRequest, userId string) (string, error) {
	objId, _ := primitive.ObjectIDFromHex(userId)

	todo := dtos.CreateTodoEntity(request, objId)

	id, err := s.Repository.Create(todo)

	return id.Hex(), err
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
