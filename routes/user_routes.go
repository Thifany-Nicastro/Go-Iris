package routes

import (
	"go-iris/controllers"
	"go-iris/repositories"
	"go-iris/services"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(app *iris.Application, db *mongo.Client) {
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
