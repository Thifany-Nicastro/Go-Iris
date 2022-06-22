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

	id, err := service.CreateUser(request)

	if err != nil {
		ctx.JSON(iris.Map{
			"message": err.Error(),
		})
	}

	ctx.JSON(iris.Map{
		"message": "User #" + id + " created",
	})
}

func Delete(ctx iris.Context) {
	//
}

func Auth(service services.UserService, ctx iris.Context) {
	//
}
