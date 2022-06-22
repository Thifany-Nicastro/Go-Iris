package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func JWTMiddleware() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func GenerateToken(ctx iris.Context) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	tokenString, _ := token.SignedString([]byte("secret_key"))

	ctx.JSON(tokenString)
}
