package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	jwt "github.com/dgrijalva/jwt-go"
)

type VerifyEmailService struct {
	core.Service
	key string
}

func (svc *VerifyEmailService) Initialize(ctx core.ServerContext, conf config.Config) error {
	key, _ := conf.GetString(ctx, "Key")
	svc.key = key
	return nil
}

func (svc *VerifyEmailService) Invoke(ctx core.RequestContext) error {
	token, ok := ctx.GetStringParam("token")
	if !ok {
		return errors.MissingArg(ctx, "token")
	}
	email, err := svc.verifyToken(ctx, token)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Debug(ctx, "email token verified", "email", email)
	return nil
}

func (svc *VerifyEmailService) verifyToken(ctx core.RequestContext, tokenVal string) (string, error) {

	token, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS512 {
			log.Error(ctx, "Invalid Token", "method", method)
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
		}
		return []byte(svc.key), nil
	})
	if err == nil && token.Valid {
		log.Error(ctx, "Token validated")
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"]
		return email.(string), nil
	}
	return "", err
}
