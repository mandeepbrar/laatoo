package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	//"github.com/rs/cors"
	//	"laatoosdk/auth"
	"fmt"
	"io/ioutil"
	"laatoo/core/registry"
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net/http"
)

const (
	CONF_AUTHORIZATION         = "authorization"
	CONF_GROUPS                = "groups"
	CONF_ROUTES                = "routes"
	CONF_ROUTE_PATH            = "path"
	CONF_ROUTE_METHOD          = "method"
	CONF_ROUTE_SERVICE         = "service"
	CONF_ROUTE_DATA_OBJECT     = "dataobject"
	CONF_ROUTE_DATA_COLLECTION = "datacollection"
	CONF_STRINGMAP_DATA_OBJECT = "__stringmap__"
)

type Router struct {
	name    string
	eRouter *echo.Group
	config  config.Config
}

func (router *Router) ConfigureRoutes(ctx core.ServerContext) error {
	routepairs := router.config.AllConfigurations()
	for _, routename := range routepairs {
		routefile, _ := router.config.GetString(routename)
		routerFileName := fmt.Sprintf("%s/%s", ctx.GetServerName(), routefile)
		routesconf, err := config.NewConfigFromFile(routerFileName)
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Name", routename, "Router file ", routerFileName)
		}
		err = router.configureRoutesConf(ctx, routesconf)
		if err != nil {
			return err
		}
	}
	return nil
}
func (router *Router) configureRoutesConf(ctx core.ServerContext, conf config.Config) error {
	groups, ok := conf.GetConfigArray(CONF_GROUPS)
	if !ok {
		return nil
	}
	for _, groupConf := range groups {
		if err := router.createGroup(ctx, groupConf); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil

}

func (router *Router) createGroup(ctx core.ServerContext, groupConf config.Config) error {
	path, ok := groupConf.GetString(CONF_ROUTE_PATH)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing path variable in group", groupConf)
	}
	grpRouter := router.group(ctx, path, groupConf)
	routes, ok := groupConf.GetConfigArray(CONF_ROUTES)
	if !ok {
		return nil
	}
	for _, routeConf := range routes {
		if err := grpRouter.createRoute(ctx, routeConf); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (router *Router) createRoute(ctx core.ServerContext, routeConf config.Config) error {
	path, ok := routeConf.GetString(CONF_ROUTE_PATH)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing path variable in route", routeConf)
	}
	method, ok := routeConf.GetString(CONF_ROUTE_METHOD)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing method variable in route", path)
	}
	service, ok := routeConf.GetString(CONF_ROUTE_SERVICE)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing service variable in route", path)
	}
	var dataObjectCreator core.ObjectCreator
	var dataObjectCollectionCreator core.ObjectCollectionCreator
	var err error
	dataObjectName, isdataObject := routeConf.GetString(CONF_ROUTE_DATA_OBJECT)
	_, isdataCollection := routeConf.GetString(CONF_ROUTE_DATA_COLLECTION)
	if isdataObject && (dataObjectName != CONF_STRINGMAP_DATA_OBJECT) {
		if isdataCollection {
			dataObjectCollectionCreator, err = registry.GetObjectCollectionCreator(ctx, dataObjectName)
			if err != nil {
				return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err, "No such object", dataObjectName)
			}
		} else {
			dataObjectCreator, err = registry.GetObjectCreator(ctx, dataObjectName)
			if err != nil {
				return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err, "No such object", dataObjectName)
			}
		}
	}
	svc, err := ctx.GetService(service)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err, "No such service has been created", path, "service", service)
	}

	processServiceRequest := func(ctx core.ServerContext, routename string, service string, dataObjectName string, isdataObject bool, isdataCollection bool, dataObjectCreator core.ObjectCreator, dataObjectCollectionCreator core.ObjectCollectionCreator) core.HandlerFunc {
		return func(webctx core.RequestContext) error {
			log.Logger.Trace(webctx, "Received request ", "route", routename)
			engineContext := webctx.EngineContext().(echo.Context)
			var reqData interface{}
			var err error
			if isdataObject {
				if dataObjectName == CONF_STRINGMAP_DATA_OBJECT {
					reqData = make(map[string]interface{}, 10)
				} else {
					if isdataCollection {
						reqData, err = dataObjectCollectionCreator(webctx, nil)
						if err != nil {
							return errors.WrapError(webctx, err)
						}
					} else {
						reqData, err = dataObjectCreator(webctx, nil)
						if err != nil {
							return errors.WrapError(webctx, err)
						}
					}
				}
				err = engineContext.Bind(reqData)
				if err != nil {
					return errors.WrapError(webctx, err)
				}
			} else {
				reqData, err = ioutil.ReadAll(engineContext.Request().(engine.Request).Body())
				if err != nil {
					return errors.WrapError(webctx, err)
				}
			}
			log.Logger.Trace(webctx, "Invoking service ", "router", routename, "service", service)
			reqctx := webctx.SubContext(service, svc.GetConf())
			reqctx.SetRequestBody(reqData)
			paramNames := engineContext.ParamNames()
			for _, param := range paramNames {
				paramVal := engineContext.Param(param)
				reqctx.Set(param, paramVal)
			}
			queryParams := engineContext.QueryParams()
			for param, val := range queryParams {
				reqctx.Set(param, val)
			}
			err = svc.Invoke(reqctx)
			if err != nil {
				return errors.WrapError(webctx, err)
			}
			resp := reqctx.GetResponse()
			if resp != nil {
				return engineContext.JSON(http.StatusOK, resp)
			}
			return engineContext.NoContent(http.StatusOK)
		}
	}

	switch method {
	case "GET":
		router.Get(ctx, path, routeConf, processServiceRequest(ctx, router.name, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator))
	case "POST":
		router.Post(ctx, path, routeConf, processServiceRequest(ctx, router.name, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator))
	case "PUT":
		router.Put(ctx, path, routeConf, processServiceRequest(ctx, router.name, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator))
	case "DELETE":
		router.Delete(ctx, path, routeConf, processServiceRequest(ctx, router.name, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator))
	}
	return nil
}

func (router *Router) Group(ctx core.ServerContext, path string, conf config.Config) core.Router {
	return router.group(ctx, path, conf)
}

func (router *Router) group(ctx core.ServerContext, path string, conf config.Config) *Router {
	retRouter := &Router{name: fmt.Sprintf("%s  %s", router.name, path), eRouter: router.eRouter.Group(path), config: conf}
	log.Logger.Debug(ctx, "Created group router", "name", retRouter.name)

	/*env := router.environment

	_, ok := conf[CONF_SERVICE_USECORS]
	if ok {
		corsHostsInt, _ := conf[CONF_SERVICE_CORSHOSTS]
		if corsHostsInt != nil {
			allowedOrigins := corsHostsInt.([]string)
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
				return errors.WrapError(err)
			})
		}
	}*/
	return retRouter
}

func (router *Router) httpAdapater(ctx core.ServerContext, conf config.Config, handler core.HandlerFunc) echo.HandlerFunc {
	return func(pathCtx echo.Context) error {
		corectx := services.NewRequestContext(ctx.GetName(), conf, ctx, pathCtx)
		return handler(corectx)
	}
}

func (router *Router) Get(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Debug(ctx, "Registering route", "router", router.name, "path", path, "method", "Get")
	router.eRouter.Get(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

func (router *Router) Put(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Debug(ctx, "Registering route", "router", router.name, "path", path, "method", "Put")
	router.eRouter.Put(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

func (router *Router) Post(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Debug(ctx, "Registering route", "router", router.name, "path", path, "method", "Post")
	router.eRouter.Post(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

func (router *Router) Delete(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Debug(ctx, "Registering route", "router", router.name, "path", path, "method", "Delete")
	router.eRouter.Delete(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

type MiddlewareHandler struct {
	mwfunc        core.HandlerFunc
	serverContext core.ServerContext
}

func (mw *MiddlewareHandler) Handle(pathCtx echo.Context) error {
	corectx := services.NewRequestContext("middleware", nil, mw.serverContext, pathCtx)
	return mw.mwfunc(corectx)
}

func (router *Router) Use(ctx core.ServerContext, handler interface{}) {
	switch handler := handler.(type) {
	case func(core.RequestContext) error:
		var hf core.HandlerFunc
		hf = handler
		mware := echo.WrapMiddleware(&MiddlewareHandler{mwfunc: hf, serverContext: ctx})
		router.eRouter.Use(mware)
		return
	case core.HandlerFunc:
		mware := echo.WrapMiddleware(&MiddlewareHandler{mwfunc: handler, serverContext: ctx})
		router.eRouter.Use(mware)
		return
	}
}

func (router *Router) Static(ctx core.ServerContext, path string, conf config.Config, dir string) error {
	log.Logger.Debug(ctx, "Registering route", "router", router.name, "path", path, "method", "Static")
	//	router.eRouter.Static(path, dir)
	return nil
}

func (router *Router) ServeFile(ctx core.ServerContext, pagePath string, conf config.Config, dest string) {
	//router.eRouter.ServeFile(pagePath, dest)
}

/*
//authentication middleware that assigns the roles permissions and users
// for all service requests to be authenticated
func (router *Router) setupAuthMiddleware(ctx core.Context, bypassauth bool) error {
	env := router.environment
	//create an anonymous user for unauthenticated requests
	auserInt, err := CreateEmptyObject(ctx, env.SystemUser)
	if err != nil {
		return errors.WrapError(err)
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
				log.Logger.Trace(ctx, "core.router", "Admin logged in", "isAdmin", isAdmin)
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
				return errors.WrapError(err)
			}
		} else {
			//if there is no header, set the user as anonymous
			context.SetUser(anonymousUser)
		}
		return nil
	})
	return nil
}
*/
