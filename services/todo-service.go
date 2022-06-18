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
	FindTodo(id string) (models.Todo, error)
	CreateTodo(todo models.Todo) (interface{}, error)
	UpdateTodo(id string, request models.Todo) int64
	DeleteTodo(id string) int64
}

func NewTodoService(Repository repositories.TodoRepository) TodoService {
	return &todoService{
		Repository: Repository,
	}
}

func (s *todoService) GetTodos() []dtos.TodoResponse {
	return s.Repository.All()
}

func (s *todoService) FindTodo(id string) (models.Todo, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.FindById(objId)
}

func (s *todoService) CreateTodo(request models.Todo) (interface{}, error) {
	return s.Repository.Create(request)
}

func (s *todoService) UpdateTodo(id string, request models.Todo) int64 {
	todo := bson.M{
		"title": request.Title,
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.Update(objId, todo)
}

func (s *todoService) DeleteTodo(id string) int64 {
	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.Delete(objId)
}
