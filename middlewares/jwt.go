package middlewares

import (
	"go-iris/utils"

	"github.com/iris-contrib/middleware/jwt"
)

func JWTMiddleware() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetEnvVar("Secret")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
