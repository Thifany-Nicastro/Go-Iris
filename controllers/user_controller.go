package controllers

import (
	"go-iris/dtos"
	"go-iris/services"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func List(ctx iris.Context) {
	//
}

func Show(service services.UserService, ctx iris.Context) {
	id := ctx.Params().Get("id")

	user, _ := service.FindUserById(id)

	ctx.JSON(user)
}

func Create(service services.UserService, ctx iris.Context) {
	var request dtos.UserRequest

	ctx.ReadJSON(&request)

	id, err := service.CreateUser(request)

	if err != nil {
		ctx.JSON(iris.Map{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(iris.Map{
		"message": "User #" + id + " created",
	})
}

func Delete(ctx iris.Context) {
	//
}

func Login(service services.UserService, ctx iris.Context) {
	var request dtos.UserRequest

	ctx.ReadJSON(&request)

	user, _ := service.FindUserByEmail(request.Email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"message": "Incorrect credentials",
		})

		return
	}

	tokenString := GenerateToken()

	ctx.JSON(iris.Map{
		"token": tokenString,
	})
}

func GenerateToken() string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	tokenString, _ := token.SignedString([]byte("secret_key"))

	return tokenString
}
