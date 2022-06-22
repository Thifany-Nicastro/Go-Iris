package controllers

import (
	"go-iris/dtos"
	"go-iris/services"

	"github.com/kataras/iris/v12"
)

func Login(service services.AuthService, ctx iris.Context) {
	var request dtos.UserRequest

	ctx.ReadJSON(&request)

	token, err := service.Login(request)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"message": "Incorrect credentials",
		})

		return
	}

	ctx.JSON(iris.Map{
		"token": token,
	})
}
