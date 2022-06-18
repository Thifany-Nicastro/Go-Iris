package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoResponse struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	IsCompleted bool               `json:"is_completed" bson:"is_completed"`
}

type TodoRequest struct {
	Title string `json:"title" bson:"title" validate:"required"`
}
