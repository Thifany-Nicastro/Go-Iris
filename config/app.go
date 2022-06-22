package config

import (
	"go-iris/middlewares"
	"go-iris/routes"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func NewApp() *iris.Application {
	app := iris.New()
	v := validator.New()
	db := Connect()

	app.Logger().SetLevel("debug")

	app.Validator = v

	routes.RegisterTodoRoutes(app, db)
	routes.RegisterUserRoutes(app, db)

	verifyMiddleware, signer := middlewares.Verify()

	app.Get("/protected", protected).Use(verifyMiddleware)
	app.Get("/token", generateToken(signer))

	return app
}

func generateToken(signer *jwt.Signer) iris.Handler {
	return func(ctx iris.Context) {
		claims := middlewares.FooClaims{Foo: "bar"}

		token, err := signer.Sign(claims)
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		ctx.Write(token)
	}
}

func protected(ctx iris.Context) {
	// Get the verified and decoded claims.
	claims := jwt.Get(ctx).(*middlewares.FooClaims)

	// Optionally, get token information if you want to work with them.
	// Just an example on how you can retrieve
	// all the standard claims (set by signer's max age, "exp").
	standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
	expiresAtString := standardClaims.ExpiresAt().
		Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
	timeLeft := standardClaims.Timeleft()

	ctx.Writef("foo=%s\nexpires at: %s\ntime left: %s\n", claims.Foo, expiresAtString, timeLeft)
}
