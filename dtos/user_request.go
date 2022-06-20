package dtos

import (
	"go-iris/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func CreateUserEntity(request UserRequest) models.User {
	return models.User{
		ID:        primitive.NewObjectID(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}
}
