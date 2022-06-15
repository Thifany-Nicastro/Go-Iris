package routes

import (
	"go-iris/controllers"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application) {
	usersAPI := app.Party("/users")
	{
		usersAPI.Use(iris.Compression)

		usersAPI.Get("/", controllers.List)
		usersAPI.Post("/", controllers.Create)
	}
}
