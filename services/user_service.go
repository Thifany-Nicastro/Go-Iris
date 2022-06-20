package services

import (
	"go-iris/dtos"
	"go-iris/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	Repository repositories.UserRepository
}

type UserService interface {
	GetUsers() []dtos.UserResponse
	FindUser(id string) (dtos.UserResponse, error)
	CreateUser(request dtos.UserRequest) (interface{}, error)
	// UpdateUser(id string, request dtos.UserRequest) int64
	DeleteUser(id string) int64
}

func NewUserService(Repository repositories.UserRepository) UserService {
	return &userService{
		Repository: Repository,
	}
}

func (s *userService) GetUsers() []dtos.UserResponse {
	users := s.Repository.All()

	return dtos.CreateUserListResponse(users)
}

func (s *userService) FindUser(id string) (dtos.UserResponse, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	user, err := s.Repository.FindById(objId)

	return dtos.CreateUserResponse(user), err
}

func (s *userService) CreateUser(request dtos.UserRequest) (interface{}, error) {
	user := dtos.CreateUserEntity(request)

	return s.Repository.Create(user)
}

// func (s *userService) UpdateUser(id string, request dtos.UserRequest) int64 {
// 	fields := bson.M{
// 		"title": request.Title,
// 	}

// 	objId, _ := primitive.ObjectIDFromHex(id)

// 	return s.Repository.Update(objId, fields)
// }

func (s *userService) DeleteUser(id string) int64 {
	objId, _ := primitive.ObjectIDFromHex(id)

	return s.Repository.Delete(objId)
}
