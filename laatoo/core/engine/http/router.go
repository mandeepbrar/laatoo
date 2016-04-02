package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	//"github.com/rs/cors"
	//	"laatoosdk/auth"
	"fmt"
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

const (
	CONF_AUTHORIZATION           = "authorization"
	CONF_GROUPS                  = "groups"
	CONF_ROUTES                  = "routes"
	CONF_ROUTE_PATH              = "path"
	CONF_ROUTE_ROUTEPARAMINDICES = "paramindices"
	CONF_ROUTE_ROUTEPARAMVALUES  = "paramvalues"
	CONF_ROUTE_HEADERSTOINCLUDE  = "headers"
	CONF_ROUTE_METHOD            = "method"
	CONF_ROUTE_METHOD_INVOKE     = "INVOKE"
	CONF_ROUTE_METHOD_GETSTREAM  = "GETSTREAM"
	CONF_ROUTE_METHOD_POSTSTREAM = "POSTSTREAM"
	CONF_ROUTE_SERVICE           = "service"
	CONF_ROUTE_DATA_OBJECT       = "dataobject"
	CONF_ROUTE_DATA_COLLECTION   = "datacollection"
	CONF_STRINGMAP_DATA_OBJECT   = "__stringmap__"
)

type Router struct {
	name    string
	eRouter *echo.Group
	config  config.Config
	engine  *HttpEngine
}

func (router *Router) Group(ctx core.ServerContext, path string, name string, conf config.Config) core.Router {
	return router.group(ctx, path, name, conf)
}

func (router *Router) group(ctx core.ServerContext, path string, name string, conf config.Config) *Router {
	retRouter := &Router{name: fmt.Sprintf("%s > %s", router.name, name), eRouter: router.eRouter.Group(path), config: conf, engine: router.engine}
	log.Logger.Info(ctx, "Created group router", "name", retRouter.name)

	/*app := router.application

	_, ok := conf[CONF_SERVICE_USECORS]
	if ok {
		corsHostsInt, _ := conf[CONF_SERVICE_CORSHOSTS]
		if corsHostsInt != nil {
			allowedOrigins := corsHostsInt.([]string)
			corsMw := cors.New(cors.Options{
				AllowedOrigins:   allowedOrigins,
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{app.AuthHeader},
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

	//provide application context to every request using middleware
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
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Get")
	router.eRouter.Get(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

func (router *Router) Put(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Put")
	router.eRouter.Put(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

func (router *Router) Post(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Post")
	router.eRouter.Post(path, router.httpAdapater(ctx, conf, handler))
	return nil

}

func (router *Router) Delete(ctx core.ServerContext, path string, conf config.Config, handler core.HandlerFunc) error {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Delete")
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
