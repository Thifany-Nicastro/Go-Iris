package config

import (
	"go-iris/controllers"
	"go-iris/repositories"
	"go-iris/routes"
	"go-iris/services"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func NewApp() *iris.Application {
	app := iris.New()
	v := validator.New()
	db := Connect()

	app.Logger().SetLevel("debug")

	app.Validator = v

	mvc.Configure(app.Party("/todos"), configureMVC)

	routes.RegisterRoutes(app, db)

	return app
}

func configureMVC(mvc *mvc.Application) {
	mvc.Register(
		Connect,
		repositories.NewTodoRepository,
		services.NewTodoService,
	)

	mvc.Handle(new(controllers.TodoController))
}
