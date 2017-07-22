package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"laatoosdk/auth"
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
	"laatoo/server/engine/http/net"
	"net/http"

	"github.com/rs/cors"
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
	name           string
	Router         net.Router
	config         config.Config
	svcName        string
	engine         *httpEngine
	allowedQParams []string
	skipAuth       bool
	path           string
}

func newHttpChannel(ctx core.ServerContext, name string, conf config.Config, engine *httpEngine, parentChannel *httpChannel) *httpChannel {
	var routername string
	var router net.Router
	path, ok := conf.GetString(constants.CONF_HTTPENGINE_PATH)
	if !ok {
		path = engine.path
	}

	if parentChannel == nil {
		routername = name
		router = engine.framework.GetParentRouter(path)
	} else {
		routername = fmt.Sprintf("%s > %s", parentChannel.name, name)
		router = parentChannel.Router.Group(path)
	}

	svc, _ := conf.GetString(constants.CONF_CHANNEL_SERVICE)

	skipAuth, ok := conf.GetBool(constants.CONF_HTTPENGINE_SKIPAUTH)
	if !ok && parentChannel != nil {
		skipAuth = parentChannel.skipAuth
	}

	channel := &httpChannel{name: routername, Router: router, config: conf, engine: engine, skipAuth: skipAuth, svcName: svc, path: path}

	allowedQParams, ok := conf.GetStringArray(constants.CONF_HTTPENGINE_ALLOWEDQUERYPARAMS)
	if ok {
		if parentChannel != nil && parentChannel.allowedQParams != nil {
			channel.allowedQParams = append(parentChannel.allowedQParams, allowedQParams...)
		} else {
			channel.allowedQParams = allowedQParams
		}
	} else {
		if parentChannel != nil {
			channel.allowedQParams = parentChannel.allowedQParams
		}
	}

	usecors, _ := conf.GetBool(constants.CONF_HTTPENGINE_USECORS)

	log.Debug(ctx, "Created group router", "name", channel.name, "path", path, "using cors", usecors, " skipauth", skipAuth)
	if usecors {
		allowedOrigins, ok := conf.GetStringArray(constants.CONF_HTTPENGINE_CORSHOSTS)
		if ok {
			corsOptionsPath, _ := conf.GetString(constants.CONF_HTTPENGINE_CORSOPTIONSPATH)
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
			log.Info(ctx, "CORS enabled for hosts ", "hosts", allowedOrigins)
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
	log.Trace(ctx, "Creating child channel", "Parent", channel.name, "New channel", name)
	return newHttpChannel(ctx, name, conf, channel.engine, channel)
}

func (channel *httpChannel) httpAdapter(ctx core.ServerContext, serviceName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) net.HandlerFunc {
	var shandler server.SecurityHandler
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	authtoken := ""
	if sh != nil {
		shandler = sh.(server.SecurityHandler)
		val := sh.GetProperty(config.AUTHHEADER)
		authtoken = val.(string)
	}

	processRequest := func(reqctx core.RequestContext, vals map[string]interface{}, body interface{}) (*core.Response, error) {
		if (!channel.skipAuth) && (shandler != nil) {
			_, err := shandler.AuthenticateRequest(reqctx, false)
			if err != nil {
				return core.StatusUnauthorizedResponse, nil
			}
		} else {
			log.Trace(reqctx, "Auth skipped")
		}
		return handler(reqctx, vals, body)
	}

	return func(pathCtx net.WebContext) error {
		errChannel := make(chan error)
		corectx := ctx.CreateNewRequest(serviceName, pathCtx)
		webreq := pathCtx.GetRequest()
		corectx.SetGaeReq(webreq)
		log.Info(corectx, "Got request", "Path", webreq.URL.RequestURI(), "Method", webreq.Method)
		defer corectx.CompleteRequest()
		go func(reqctx core.RequestContext, webctx net.WebContext) {
			bytes, err := webreq.GetBody()
			if err != nil {
				errChannel <- err
				return
			}
			vals := make(map[string]interface{})
			vals[authtoken] = pathCtx.GetHeader(authtoken)
			resp, err := processRequest(reqctx, vals, bytes)
			log.Trace(reqctx, "Completed request for service. Handling Response")
			if err == nil {
				err = respHandler.HandleResponse(reqctx, resp)
			}
			errChannel <- err
		}(corectx, pathCtx)
		err := <-errChannel
		if err != nil {
			log.Info(corectx, "Got error in the request", "error", err)
		}
		return err
	}
}

func (channel *httpChannel) get(ctx core.ServerContext, path string, serviceName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Get(path, channel.httpAdapter(ctx, serviceName, handler, respHandler, svc))
}

func (channel *httpChannel) options(ctx core.ServerContext, path string, serviceName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Options")
	channel.Router.Options(path, channel.httpAdapter(ctx, serviceName, handler, respHandler, svc))
}

func (channel *httpChannel) put(ctx core.ServerContext, path string, serviceName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Put(path, channel.httpAdapter(ctx, serviceName, handler, respHandler, svc))
}

func (channel *httpChannel) post(ctx core.ServerContext, path string, serviceName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Post(path, channel.httpAdapter(ctx, serviceName, handler, respHandler, svc))
}

func (channel *httpChannel) delete(ctx core.ServerContext, path string, serviceName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	channel.Router.Delete(path, channel.httpAdapter(ctx, serviceName, handler, respHandler, svc))
}

func (channel *httpChannel) use(ctx core.ServerContext, middlewareName string, handler ServiceInvoker, respHandler server.ServiceResponseHandler, svc server.Service) {
	channel.Router.Use(channel.httpAdapter(ctx, middlewareName, handler, respHandler, svc))
}

func (channel *httpChannel) useMW(ctx core.ServerContext, handler func(http.Handler) http.Handler) {
	channel.Router.UseMW(handler)
}
func (channel *httpChannel) useMiddleware(ctx core.ServerContext, handler http.HandlerFunc) {
	channel.Router.UseMiddleware(handler)
}
