package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"laatoosdk/auth"
	"fmt"
	"github.com/rs/cors"
	"laatoo/core/engine/http/net"
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"net/http"
)

const (
	CONF_AUTHORIZATION           = "authorization"
	CONF_GROUPS                  = "groups"
	CONF_ROUTEGROUPS             = "routegroups"
	CONF_ROUTES                  = "routes"
	CONF_ROUTE_USECORS           = "usecors"
	CONF_ROUTE_CORSOPTIONSPATH   = "corsoptionspath"
	CONF_ROUTE_CORSHOSTS         = "corshosts"
	CONF_ROUTE_PATH              = "path"
	CONF_ROUTE_STATICVALUES      = "staticvalues"
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
	name   string
	Router net.Router
	config config.Config
	engine *HttpEngine
}

func newRouter(ctx core.ServerContext, name string, conf config.Config, engine *HttpEngine, parentRouter *Router) *Router {
	var routername string
	var router net.Router
	path, ok := conf.GetString(CONF_ROUTE_PATH)
	if !ok {
		path = ""
	}
	if parentRouter == nil {
		routername = name
		router = engine.framework.GetParentRouter(path)
	} else {
		routername = fmt.Sprintf("%s > %s", parentRouter.name, name)
		router = parentRouter.Router.Group(path)
	}

	retRouter := &Router{name: routername, Router: router, config: conf, engine: engine}

	usecors, _ := conf.GetBool(CONF_ROUTE_USECORS)

	log.Logger.Info(ctx, "Created group router", "name", retRouter.name, "using cors", usecors)
	if usecors {
		allowedOrigins, ok := conf.GetStringArray(CONF_ROUTE_CORSHOSTS)
		if ok {
			corsOptionsPath, _ := conf.GetString(CONF_ROUTE_CORSOPTIONSPATH)
			corsMw := cors.New(cors.Options{
				AllowedOrigins:   allowedOrigins,
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{engine.authHeader},
				AllowCredentials: true,
			})
			switch engine.fwname {
			case "Echo":
				if corsOptionsPath == "" {
					corsOptionsPath = "/*"
				}
				retRouter.UseMW(ctx, corsMw.Handler)
			case "Gin":
				if corsOptionsPath == "" {
					corsOptionsPath = "/*f"
				}
				retRouter.UseMiddleware(ctx, corsMw.HandlerFunc)
			case "Goji":
				if corsOptionsPath == "" {
					corsOptionsPath = "/*"
				}
				retRouter.UseMW(ctx, corsMw.Handler)
			}
			retRouter.Router.Options(corsOptionsPath, func(webctx net.WebContext) error {
				webctx.NoContent(200)
				return nil
			})
			//retRouter.UseMiddleware(ctx, corsMw.HandlerFunc)
			log.Logger.Info(ctx, "CORS enabled for hosts ", "hosts", allowedOrigins)
		}
	}

	/*
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

func (router *Router) Group(ctx core.ServerContext, name string, conf config.Config) *Router {
	return newRouter(ctx, name, conf, router.engine, router)
}

func (router *Router) httpAdapater(ctx core.ServerContext, conf config.Config, handler core.ServiceFunc) net.HandlerFunc {
	return func(pathCtx net.WebContext) error {
		corectx := services.NewRequestContext(ctx.GetName(), conf, ctx, pathCtx)
		defer corectx.CompleteRequest()
		return handler(corectx)
	}
}

func (router *Router) Get(ctx core.ServerContext, path string, conf config.Config, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Get")
	router.Router.Get(path, router.httpAdapater(ctx, conf, handler))
}

func (router *Router) Options(ctx core.ServerContext, path string, conf config.Config, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Options")
	router.Router.Options(path, router.httpAdapater(ctx, conf, handler))
}

func (router *Router) Put(ctx core.ServerContext, path string, conf config.Config, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Get")
	router.Router.Put(path, router.httpAdapater(ctx, conf, handler))
}

func (router *Router) Post(ctx core.ServerContext, path string, conf config.Config, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Get")
	router.Router.Post(path, router.httpAdapater(ctx, conf, handler))
}

func (router *Router) Delete(ctx core.ServerContext, path string, conf config.Config, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "router", router.name, "path", path, "method", "Get")
	router.Router.Delete(path, router.httpAdapater(ctx, conf, handler))
}

func (router *Router) Use(ctx core.ServerContext, handler core.ServiceFunc) {
	router.Router.Use(router.httpAdapater(ctx, nil, handler))
}

func (router *Router) UseMW(ctx core.ServerContext, handler func(http.Handler) http.Handler) {
	router.Router.UseMW(handler)
}
func (router *Router) UseMiddleware(ctx core.ServerContext, handler http.HandlerFunc) {
	router.Router.UseMiddleware(handler)
}
