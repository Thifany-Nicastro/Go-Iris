package main

import (
	"go-iris/config"
	"go-iris/controllers"
	"go-iris/routes"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	config.Connect()

	mvc.Configure(app.Party("/todos"), configureMVC)

	routes.RegisterRoutes(app)

	app.Listen("localhost:8080")
}

func configureMVC(app *mvc.Application) {
	app.Handle(new(controllers.TodoController))
}
