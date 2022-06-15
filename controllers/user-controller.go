package controllers

import (
	"go-iris/models"

	"github.com/kataras/iris/v12"
)

func List(ctx iris.Context) {
	users := []models.User{
		{Name: "Foo"},
		{Name: "Bar"},
	}

	ctx.JSON(users)
}

func Create(ctx iris.Context) {
	var u models.User
	err := ctx.ReadJSON(&u)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("User creation failure").DetailErr(err))

		return
	}

	println("Received Book: " + u.Name)

	ctx.StatusCode(iris.StatusCreated)
}
