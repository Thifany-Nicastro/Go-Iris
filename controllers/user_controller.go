package controllers

import (
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

func Create(ctx iris.Context) {
	//
}

func Delete(ctx iris.Context) {
	//
}
