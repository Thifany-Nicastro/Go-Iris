package services

import (
	"context"
	"go-iris/config"
	"go-iris/models"
)

func CreateTodo(todo models.Todo) (interface{}, error) {
	todosCollection := config.GetCollection("todos")

	result, err := todosCollection.InsertOne(context.TODO(), todo)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}
