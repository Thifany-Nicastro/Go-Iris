package routes

import (
	"go-iris/controllers"
	"go-iris/repositories"
	"go-iris/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterTodoRoutes(app *iris.Application, db *mongo.Client) {

	todoRouter := app.Party("/todos")
	{
		todoApp := mvc.New(todoRouter)

		todoApp.Register(
			db,
			repositories.NewTodoRepository,
			services.NewTodoService,
		)

		todoApp.Handle(new(controllers.TodoController))
	}
}
