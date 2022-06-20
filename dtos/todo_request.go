package dtos

import (
	"go-iris/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRequest struct {
	Title string `json:"title" validate:"required"`
}

func CreateTodoEntity(request TodoRequest) models.Todo {
	return models.Todo{
		ID:    primitive.NewObjectID(),
		Title: request.Title,
	}
}
