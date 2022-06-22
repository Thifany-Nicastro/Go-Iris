package services

import (
	"go-iris/dtos"
	"go-iris/repositories"

	"github.com/iris-contrib/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	Repository repositories.UserRepository
}

type AuthService interface {
	Login(request dtos.UserRequest) (string, error)
}

func NewAuthService(Repository repositories.UserRepository) AuthService {
	return &authService{
		Repository: Repository,
	}
}

func (s *authService) Login(request dtos.UserRequest) (string, error) {
	user, _ := s.Repository.FindByEmail(request.Email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", err
	}

	tokenString := GenerateToken()

	return tokenString, nil
}

func GenerateToken() string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	tokenString, _ := token.SignedString([]byte("secret_key"))

	return tokenString
}
