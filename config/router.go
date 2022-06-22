package config

import (
	"go-iris/controllers"
	"go-iris/repositories"
	"go-iris/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRoutes(app *iris.Application, db *mongo.Client) {
	app.Party("/users").ConfigureContainer(func(r *iris.APIContainer) {
		r.Use(iris.Compression)

		userRepository := repositories.NewUserRepository(db)
		userService := services.NewUserService(userRepository)
		r.RegisterDependency(userService)

		r.Get("/", controllers.List)
		r.Get("/{id:string}", controllers.Show)
		r.Post("/", controllers.Create)
		r.Delete("/", controllers.Delete)
	})
}

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
