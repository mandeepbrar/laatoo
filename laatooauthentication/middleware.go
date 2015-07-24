package laatooauthentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"laatoosdk/errors"
	"laatoosdk/user"
	"laatoosdk/utils"
)

func (svc *AuthService) Authenticate(ctx *echo.Context) error {
	headerVal := ctx.Request().Header.Get(svc.AuthHeader)
	if headerVal != "" {
		token, err := jwt.Parse(headerVal, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.ThrowError(AUTH_ERROR_WRONG_SIGNING_METHOD)
			}
			return []byte(svc.JWTSecret), nil
		})
		if err == nil && token.Valid {
			userInt, err := svc.CreateUser()
			if err != nil {
				return errors.RethrowHttpError(AUTH_ERROR_WRONG_SIGNING_METHOD, ctx, err)
			}
			user, ok := userInt.(user.User)
			if !ok {
				return errors.ThrowHttpError(AUTH_ERROR_USEROBJECT_NOT_CREATED, ctx)
			}
			user.LoadJWTClaims(token)
			user.SetId(token.Claims["UserId"].(string))
			ctx.Set("User", userInt)
			ctx.Set("JWT_Token", token)
			utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_AUTH_COMPLETE, ctx})
			return nil
		} else {
			return err
		}
	}
	return errors.ThrowHttpError(AUTH_ERROR_HEADER_NOT_FOUND, ctx)
}
