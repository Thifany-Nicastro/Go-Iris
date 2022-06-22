package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
)

func JWTMiddleware() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
