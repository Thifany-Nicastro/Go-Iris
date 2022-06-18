package services

import (
	"context"
	"fmt"
	"go-iris/dtos"
	"go-iris/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoService struct {
	db *mongo.Client
}

type TodoService interface {
	CreateTodo(todo models.Todo) (interface{}, error)
	DeleteTodo(id string) int64
	GetTodos() []dtos.TodoResponse
	FindTodo(id string) (models.Todo, error)
	UpdateTodo(id string, request models.Todo) int64
}

func NewTodoService(db *mongo.Client) TodoService {
	fmt.Println("hmm" + db.Database("go-iris").Name())
	return &todoService{
		db: db,
	}
}

func (s *todoService) GetTodos() []dtos.TodoResponse {
	// todos := []dtos.TodoResponse{
	// 	{Id: "1", Title: "Buy Milk"},
	// 	{Id: "2", Title: "Buy Milk"},
	// }

	fmt.Println(s.db.Database("go-iris").Name())

	var todos []dtos.TodoResponse

	todosCollection := s.db.Database("go-iris").Collection("todos")

	cursor, _ := todosCollection.Find(context.TODO(), bson.D{{}})

	cursor.All(context.TODO(), &todos)

	return todos
}

func (s *todoService) CreateTodo(todo models.Todo) (interface{}, error) {
	todosCollection := s.db.Database("go-iris").Collection("todos")

	result, err := todosCollection.InsertOne(context.TODO(), todo)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}

func (s *todoService) DeleteTodo(id string) int64 {
	todosCollection := s.db.Database("go-iris").Collection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, _ := todosCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	return result.DeletedCount
}

func (s *todoService) FindTodo(id string) (models.Todo, error) {
	var todo models.Todo

	todosCollection := s.db.Database("go-iris").Collection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	err := todosCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *todoService) UpdateTodo(id string, request models.Todo) int64 {
	todo := bson.M{
		"title": request.Title,
	}

	todosCollection := s.db.Database("go-iris").Collection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, _ := todosCollection.UpdateOne(context.TODO(),
		bson.M{"_id": objId},
		bson.M{"$set": todo},
	)

	return result.ModifiedCount
}
