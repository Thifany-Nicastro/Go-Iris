package controllers

import (
	"go-iris/dtos"
	"go-iris/services"

	"github.com/kataras/iris/v12"
)

func List(ctx iris.Context) {
	//
}

func Show(service services.UserService, ctx iris.Context) {
	id := ctx.Params().Get("id")

	user, _ := service.FindUser(id)

	ctx.JSON(user)
}

func Create(service services.UserService, ctx iris.Context) {
	var request dtos.UserRequest

	ctx.ReadJSON(&request)

	service.CreateUser(request)

	ctx.JSON(iris.Map{"message": "OK"})
}

func Delete(ctx iris.Context) {
	//
}
