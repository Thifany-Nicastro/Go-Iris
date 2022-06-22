package config

import (
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

	RegisterTodoRoutes(app, db)
	RegisterUserRoutes(app, db)

	verifyMiddleware, signer := Verify()

	app.Get("/protected", protected).Use(verifyMiddleware)
	app.Get("/token", generateToken(signer))

	return app
}

func generateToken(signer *jwt.Signer) iris.Handler {
	return func(ctx iris.Context) {
		claims := FooClaims{Foo: "bar"}

		token, err := signer.Sign(claims)
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		ctx.Write(token)
	}
}

func protected(ctx iris.Context) {
	claims := jwt.Get(ctx).(*FooClaims)

	standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
	expiresAtString := standardClaims.ExpiresAt().
		Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
	timeLeft := standardClaims.Timeleft()

	ctx.Writef("foo=%s\nexpires at: %s\ntime left: %s\n", claims.Foo, expiresAtString, timeLeft)
}
