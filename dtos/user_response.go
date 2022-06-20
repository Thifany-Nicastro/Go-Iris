package dtos

import "go-iris/models"

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func CreateUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:   user.ID.Hex(),
		Name: user.FirstName + " " + user.LastName,
	}
}

func CreateUserListResponse(users []models.User) []UserResponse {
	var response []UserResponse

	for _, t := range users {
		user := CreateUserResponse(t)
		response = append(response, user)
	}

	return response
}
