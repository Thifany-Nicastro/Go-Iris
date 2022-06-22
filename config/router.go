package config

import (
	"go-iris/controllers"
	"go-iris/middlewares"
	"go-iris/repositories"
	"go-iris/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Router(app *iris.Application) {
	db := Connect()
	jwt := middlewares.JWTMiddleware()

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

	mvc.Configure(app.Party("/todos"), func(m *mvc.Application) {
		m.Register(
			db,
			repositories.NewTodoRepository,
			services.NewTodoService,
		)

		m.Handle(new(controllers.TodoController))
	})

	app.Get("/protected", jwt.Serve, func(ctx iris.Context) {
		ctx.JSON("ok")
	})

	app.Get("/token", middlewares.GenerateToken)
}
