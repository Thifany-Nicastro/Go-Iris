package repositories

import (
	"context"
	"go-iris/dtos"
	"go-iris/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

type TodoRepository interface {
	All() []dtos.TodoResponse
	FindById(id primitive.ObjectID) (models.Todo, error)
	Create(todo models.Todo) (interface{}, error)
	Update(id primitive.ObjectID, todo primitive.M) int64
	Delete(id primitive.ObjectID) int64
}

func NewTodoRepository(db *mongo.Client) TodoRepository {
	return &todoRepository{
		db:         db,
		collection: db.Database("go-iris").Collection("todos"),
	}
}

func (s *todoRepository) All() []dtos.TodoResponse {
	var todos []dtos.TodoResponse

	cursor, _ := s.collection.Find(context.TODO(), bson.D{{}})

	cursor.All(context.TODO(), &todos)

	return todos
}

func (s *todoRepository) FindById(id primitive.ObjectID) (models.Todo, error) {
	var todo models.Todo

	err := s.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)

	return todo, err
}

func (s *todoRepository) Create(todo models.Todo) (interface{}, error) {
	result, err := s.collection.InsertOne(context.TODO(), todo)

	return result.InsertedID, err
}

func (s *todoRepository) Update(id primitive.ObjectID, todo primitive.M) int64 {
	result, _ := s.collection.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": todo},
	)

	return result.ModifiedCount
}

func (s *todoRepository) Delete(id primitive.ObjectID) int64 {
	result, _ := s.collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	return result.DeletedCount
}
