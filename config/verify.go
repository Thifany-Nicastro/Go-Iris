package config

import (
	"time"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type FooClaims struct {
	Foo string `json:"foo"`
}

func Verify() (context.Handler, *jwt.Signer) {
	secret := []byte("signature_hmac_secret_shared_key")
	signer := jwt.NewSigner(
		jwt.HS256,
		secret,
		10*time.Minute,
	)
	verifier := jwt.NewVerifier(jwt.HS256, secret)
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(FooClaims)
	})

	return verifyMiddleware, signer
}
