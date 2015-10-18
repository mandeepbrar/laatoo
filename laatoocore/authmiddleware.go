package laatoocore

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"laatoosdk/auth"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
)

//authentication middleware that assigns the roles permissions and users
// for all service requests to be authenticated
func (env *Environment) setupAuthMiddleware(ctx *echo.Context, router *echo.Group) error {
	//create an anonymous user for unauthenticated requests
	auserInt, err := CreateEmptyObject(ctx, env.SystemUser)
	if err != nil {
		return err
	}
	anonymousUser := auserInt.(auth.RbacUser)
	anonymousUser.AddRole("Anonymous")

	router.Use(func(ctx *echo.Context) error {
		//get the authentication header for request
		headerVal := ctx.Request().Header.Get(env.AuthHeader)
		log.Logger.Trace(ctx, "core.env.authmiddleware", "Testing header val", "auth header", headerVal)

		if headerVal != "" {
			//parse the header to check validity
			token, err := jwt.Parse(headerVal, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.ThrowError(ctx, AUTH_ERROR_WRONG_SIGNING_METHOD)
				}
				return []byte(env.JWTSecret), nil
			})

			if err == nil && token.Valid {
				log.Logger.Trace(ctx, "core.env.authmiddleware", "valid token")

				//create empty user object
				userInt, err := CreateEmptyObject(ctx, env.SystemUser)
				if err != nil {
					return errors.RethrowError(ctx, AUTH_ERROR_WRONG_SIGNING_METHOD, err)
				}
				//check if the user is rbac user
				user, ok := userInt.(auth.RbacUser)
				if !ok {
					return errors.ThrowError(ctx, AUTH_ERROR_USEROBJECT_NOT_CREATED)
				}
				//load the jwt claims that were set int the token by the user object
				user.LoadJWTClaims(token)
				//set the id of the user object loaded from the header
				user.SetId(token.Claims["UserId"].(string))
				//set the user in the context object
				ctx.Set("User", userInt)
				//get the roles of the user
				roles, _ := user.GetRoles()
				//set the roles and token in the context
				ctx.Set("Roles", roles)
				ctx.Set("JWT_Token", token)
				log.Logger.Info(ctx, "core.env.authmiddleware", "Set roles", "roles", roles)
				//fire auth coplete event
				utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_AUTH_COMPLETE, ctx})
				return nil
			} else {
				//if the token is invalid throw security error
				if token == nil || !token.Valid {
					return errors.RethrowError(ctx, AUTH_ERROR_INVALID_TOKEN, err)
				}
				return err
			}
		} else {
			//if there is no header, set the user as anonymous
			ctx.Set("User", anonymousUser)
		}
		return nil
	})
	return nil
}
