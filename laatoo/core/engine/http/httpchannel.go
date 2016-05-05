package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"laatoosdk/auth"
	"fmt"
	"github.com/rs/cors"
	"laatoo/core/engine/http/net"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"net/http"
)

const (

/*	CONF_AUTHORIZATION           = "authorization"
	CONF_GROUPS                  = "groups"
	CONF_ROUTEGROUPS             = "routegroups"
	CONF_ROUTES                  = "routes"
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
	CONF_STRINGMAP_DATA_OBJECT   = "__stringmap__"*/
)

type httpChannel struct {
	name   string
	Router net.Router
	config config.Config
	engine *httpEngine
}

func newHttpChannel(ctx core.ServerContext, name string, conf config.Config, engine *httpEngine, parentChannel *httpChannel) *httpChannel {
	var routername string
	var router net.Router
	path, ok := conf.GetString(config.CONF_HTTPENGINE_PATH)
	if !ok {
		path = ""
	}
	if parentChannel == nil {
		routername = name
		router = engine.framework.GetParentRouter(path)
	} else {
		routername = fmt.Sprintf("%s > %s", parentChannel.name, name)
		router = parentChannel.Router.Group(path)
	}

	channel := &httpChannel{name: routername, Router: router, config: conf, engine: engine}

	usecors, _ := conf.GetBool(config.CONF_HTTPENGINE_USECORS)

	log.Logger.Debug(ctx, "Created group router", "name", channel.name, "using cors", usecors)
	if usecors {
		allowedOrigins, ok := conf.GetStringArray(config.CONF_HTTPENGINE_CORSHOSTS)
		if ok {
			corsOptionsPath, _ := conf.GetString(config.CONF_HTTPENGINE_CORSOPTIONSPATH)
			corsMw := cors.New(cors.Options{
				AllowedOrigins:     allowedOrigins,
				AllowedHeaders:     []string{"*"},
				AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"},
				ExposedHeaders:     []string{"*"},
				OptionsPassthrough: true,
				AllowCredentials:   true,
			})
			switch engine.fwname {
			case "Echo":
				if corsOptionsPath == "" {
					corsOptionsPath = "/*"
				}
				channel.useMW(ctx, corsMw.Handler)
			case "Gin":
				if corsOptionsPath == "" {
					corsOptionsPath = "/*f"
				}
				channel.useMiddleware(ctx, corsMw.HandlerFunc)
			case "Goji":
				if corsOptionsPath == "" {
					corsOptionsPath = "/*"
				}
				channel.useMW(ctx, corsMw.Handler)
			}
			channel.Router.Options(corsOptionsPath, func(webctx net.WebContext) error {
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

	return channel
}

func (channel *httpChannel) group(ctx core.ServerContext, name string, conf config.Config) *httpChannel {
	return newHttpChannel(ctx, name, conf, channel.engine, channel)
}

func (channel *httpChannel) httpAdapter(ctx core.ServerContext, handler core.ServiceFunc) net.HandlerFunc {
	var shandler server.SecurityHandler
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler = sh.(server.SecurityHandler)
	}
	return func(pathCtx net.WebContext) error {
		corectx := ctx.CreateNewRequest(ctx.GetName(), pathCtx)
		defer corectx.CompleteRequest()
		if shandler != nil {
			err := shandler.AuthenticateRequest(corectx)
			if err != nil {
				pathCtx.NoContent(http.StatusUnauthorized)
				return err
			}
		}
		return handler(corectx)
	}
}

func (channel *httpChannel) get(ctx core.ServerContext, path string, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Get(path, channel.httpAdapter(ctx, handler))
}

func (channel *httpChannel) options(ctx core.ServerContext, path string, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Options")
	channel.Router.Options(path, channel.httpAdapter(ctx, handler))
}

func (channel *httpChannel) put(ctx core.ServerContext, path string, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Put(path, channel.httpAdapter(ctx, handler))
}

func (channel *httpChannel) post(ctx core.ServerContext, path string, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Post(path, channel.httpAdapter(ctx, handler))
}

func (channel *httpChannel) delete(ctx core.ServerContext, path string, handler core.ServiceFunc) {
	log.Logger.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Delete(path, channel.httpAdapter(ctx, handler))
}

func (channel *httpChannel) use(ctx core.ServerContext, handler core.ServiceFunc) {
	channel.Router.Use(channel.httpAdapter(ctx, handler))
}

func (channel *httpChannel) useMW(ctx core.ServerContext, handler func(http.Handler) http.Handler) {
	channel.Router.UseMW(handler)
}
func (channel *httpChannel) useMiddleware(ctx core.ServerContext, handler http.HandlerFunc) {
	channel.Router.UseMiddleware(handler)
}
