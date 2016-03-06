package laatoocore

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/rs/cors"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
)

const (
	CONF_AUTHORIZATION = "authorization"
)

type Router struct {
	eRouter     *echo.Group
	environment *Environment
}

func (router *Router) Group(ctx core.Context, path string, conf map[string]interface{}) core.Router {
	retRouter := &Router{eRouter: router.eRouter.Group(path), environment: router.environment}

	env := router.environment

	_, ok := conf[CONF_SERVICE_USECORS]
	if ok {
		corsHostsInt, _ := conf[CONF_SERVICE_CORSHOSTS]
		if corsHostsInt != nil {
			allowedOrigins := corsHostsInt.([]string)
			/*allowedOrigins := make([]string, len(corsHostsInt))
			i := 0
			for _, k := range corsHostsInt {
				allowedOrigins[i] = k.(string)
				i++
			}*/
			corsMw := cors.New(cors.Options{
				AllowedOrigins:   allowedOrigins,
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{env.AuthHeader},
				AllowCredentials: true,
			}).Handler
			log.Logger.Info(ctx, "core.env", "CORS enabled for hosts ", "hosts", allowedOrigins)
			retRouter.Use(ctx, corsMw)
		}
	}

	bypassauth := false
	//authentication required by default unless explicitly turned off
	bypassauthInt, ok := conf[CONF_SERVICE_AUTHBYPASS]
	if ok {
		bypassauth = (bypassauthInt == "true")
	}

	//provide environment context to every request using middleware
	retRouter.Use(ctx, func(ctx core.Context) error {
		//ctx.Set(CONF_ENV_CONTEXT, env)
		if bypassauth {
			ctx.Set(CONF_SERVICE_AUTHBYPASS, true)
		}
		return nil
	})

	retRouter.setupAuthMiddleware(ctx, bypassauth)
	if !bypassauth {
		_, confok := conf[CONF_AUTHORIZATION]
		if confok {
			retRouter.Use(ctx, func(permCtx core.Context) error {
				authorized, err := retRouter.authorize(permCtx, conf)
				if !authorized {
					return errors.ThrowError(ctx, AUTH_ERROR_SECURITY)
				}
				return err
			})
		}
	}
	return retRouter
}

func (router *Router) Get(ctx core.Context, path string, conf map[string]interface{}, handler core.HandlerFunc) error {
	router.eRouter.Get(path, func(pathCtx *echo.Context) error {
		corectx := &Context{Context: pathCtx, Conf: conf, environment: router.environment}
		authorized, err := router.authorize(corectx, conf)
		if authorized {
			err = handler(corectx)
		}
		return err
	})
	return nil

}

func (router *Router) Put(ctx core.Context, path string, conf map[string]interface{}, handler core.HandlerFunc) error {
	router.eRouter.Put(path, func(pathCtx *echo.Context) error {
		corectx := &Context{Context: pathCtx, Conf: conf, environment: router.environment}
		authorized, err := router.authorize(corectx, conf)
		if authorized {
			err = handler(corectx)
		}
		return err
	})
	return nil

}

func (router *Router) Post(ctx core.Context, path string, conf map[string]interface{}, handler core.HandlerFunc) error {
	router.eRouter.Post(path, func(pathCtx *echo.Context) error {
		corectx := &Context{Context: pathCtx, Conf: conf, environment: router.environment}
		authorized, err := router.authorize(corectx, conf)
		if authorized {
			err = handler(corectx)
		}
		return err
	})
	return nil

}

func (router *Router) Delete(ctx core.Context, path string, conf map[string]interface{}, handler core.HandlerFunc) error {
	router.eRouter.Delete(path, func(pathCtx *echo.Context) error {
		corectx := &Context{Context: pathCtx, Conf: conf, environment: router.environment}
		authorized, err := router.authorize(corectx, conf)
		if authorized {
			err = handler(corectx)
		}
		return err
	})
	return nil

}

func (router *Router) Use(ctx core.Context, handler interface{}) {
	switch handler := handler.(type) {
	case func(core.Context) error:
		router.eRouter.Use(func(pathCtx *echo.Context) error {
			core_ctx := &Context{Context: pathCtx, environment: router.environment}
			return handler(core_ctx)
		})
		return
	case core.HandlerFunc:
		router.eRouter.Use(func(pathCtx *echo.Context) error {
			core_ctx := &Context{Context: pathCtx}
			return handler(core_ctx)
		})
		return
	default:
		router.eRouter.Use(handler)
		return
	}
}

func (router *Router) Static(ctx core.Context, path string, conf map[string]interface{}, dir string) error {
	router.eRouter.Static(path, dir)
	return nil
}

func (router *Router) ServeFile(ctx core.Context, pagePath string, conf map[string]interface{}, dest string) {
	authorized, _ := router.authorize(ctx, conf)
	if authorized {
		router.eRouter.ServeFile(pagePath, dest)
	}
}

//authentication middleware that assigns the roles permissions and users
// for all service requests to be authenticated
func (router *Router) setupAuthMiddleware(ctx core.Context, bypassauth bool) error {
	env := router.environment
	//create an anonymous user for unauthenticated requests
	auserInt, err := CreateEmptyObject(ctx, env.SystemUser)
	if err != nil {
		return err
	}
	anonymousUser := auserInt.(auth.RbacUser)
	anonymousUser.AddRole("Anonymous")

	router.Use(ctx, func(ctx core.Context) error {
		context := ctx.(*Context)
		//get the authentication header for request
		headerVal := ctx.Request().Header.Get(env.AuthHeader)
		log.Logger.Trace(ctx, "core.router.authmiddleware", "Testing header val", "auth header", headerVal)

		if headerVal != "" {
			//parse the header to check validity
			token, err := jwt.Parse(headerVal, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					if bypassauth {
						context.SetUser(anonymousUser)
						return nil, nil
					}
					return nil, errors.ThrowError(ctx, AUTH_ERROR_WRONG_SIGNING_METHOD)
				}
				return []byte(env.JWTSecret), nil
			})

			if err == nil && token.Valid {
				log.Logger.Trace(ctx, "core.router.authmiddleware", "valid token")

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
				isAdmin := token.Claims["Admin"]
				if isAdmin != nil {
					ctx.Set("Admin", isAdmin.(bool))
				}
				//set the user in the context object
				context.SetUser(user)
				//get the roles of the user
				roles, _ := user.GetRoles()
				log.Logger.Trace(ctx, "core.router", "Testing auth", "roles", roles, "user", user)
				//set the roles and token in the context
				ctx.Set("Roles", roles)
				ctx.Set("JWT_Token", token)
				log.Logger.Info(ctx, "core.router.authmiddleware", "Set roles", "roles", roles)
				//fire auth coplete event
				utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_AUTH_COMPLETE, ctx})
				return nil
			} else {
				if bypassauth {
					context.SetUser(anonymousUser)
					return nil
				}
				//if the token is invalid throw security error
				if token == nil || !token.Valid {
					return errors.RethrowError(ctx, AUTH_ERROR_INVALID_TOKEN, err)
				}
				return err
			}
		} else {
			//if there is no header, set the user as anonymous
			context.SetUser(anonymousUser)
		}
		return nil
	})
	return nil
}
