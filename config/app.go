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
	v := validator.New()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Validator = v

	mvc.Configure(app.Party("/todos"), configureMVC)

	routes.RegisterRoutes(app)

	return app
}

func configureMVC(app *mvc.Application) {
	app.Register(
		Connect,
		repositories.NewTodoRepository,
		services.NewTodoService,
	)

	app.Handle(new(controllers.TodoController))
}
