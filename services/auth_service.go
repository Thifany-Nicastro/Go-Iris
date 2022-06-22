package services

import (
	"go-iris/dtos"
	"go-iris/repositories"
	"go-iris/utils"
	"time"

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

	tokenString := GenerateToken(user.ID.Hex())

	return tokenString, nil
}

func GenerateToken(userId string) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userId,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(utils.GetEnvVar("SECRET")))

	return tokenString
}
