package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

func GetAuthenticatedUser(ctx iris.Context) string {
	claims := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	return claims["user"].(string)
}
