package services

import (
	"context"
	"go-iris/config"
	"go-iris/dtos"
	"go-iris/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(todo models.Todo) (interface{}, error) {
	todosCollection := config.GetCollection("todos")

	result, err := todosCollection.InsertOne(context.TODO(), todo)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}

func DeleteTodo(id string) int64 {
	todosCollection := config.GetCollection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, _ := todosCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	return result.DeletedCount
}

func GetTodos() []dtos.TodoResponse {
	var todos []dtos.TodoResponse

	todosCollection := config.GetCollection("todos")

	cursor, _ := todosCollection.Find(context.TODO(), bson.D{{}})

	cursor.All(context.TODO(), &todos)

	return todos
}

func FindTodo(id string) (models.Todo, error) {
	var todo models.Todo

	todosCollection := config.GetCollection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	err := todosCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func UpdateTodo(id string, request models.Todo) int64 {
	todo := bson.M{
		"title": request.Title,
	}

	todosCollection := config.GetCollection("todos")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, _ := todosCollection.UpdateOne(context.TODO(),
		bson.M{"_id": objId},
		bson.M{"$set": todo},
	)

	return result.ModifiedCount
}
