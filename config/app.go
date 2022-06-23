package config

import (
	"go-iris/utils"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func NewApp() *iris.Application {
	utils.LoadEnvironment()

	app := iris.New()
	v := validator.New()

	// app.Logger().SetLevel("debug")

	app.Validator = v

	Router(app)

	return app
}
