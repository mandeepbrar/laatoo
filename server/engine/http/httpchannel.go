package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"laatoosdk/auth"
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
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
	method         string
	Router         net.Router
	config         config.Config
	svcName        string
	engine         *httpEngine
	allowedQParams []string
	skipAuth       bool
	path           string
	parentChannel  *httpChannel
}

func (channel *httpChannel) configure(ctx core.ServerContext) error {
	skipAuth, ok := channel.config.GetBool(ctx, constants.CONF_HTTPENGINE_SKIPAUTH)
	if !ok && channel.parentChannel != nil {
		skipAuth = channel.parentChannel.skipAuth
	}
	channel.skipAuth = skipAuth

	allowedQParams, ok := channel.config.GetStringArray(ctx, constants.CONF_HTTPENGINE_ALLOWEDQUERYPARAMS)
	if ok {
		if channel.parentChannel != nil && channel.parentChannel.allowedQParams != nil {
			channel.allowedQParams = append(channel.parentChannel.allowedQParams, allowedQParams...)
		} else {
			channel.allowedQParams = allowedQParams
		}
	} else {
		if channel.parentChannel != nil {
			channel.allowedQParams = channel.parentChannel.allowedQParams
		}
	}

	usecors, _ := channel.config.GetBool(ctx, constants.CONF_HTTPENGINE_USECORS)

	if usecors {
		allowedOrigins, ok := channel.config.GetStringArray(ctx, constants.CONF_HTTPENGINE_CORSHOSTS)
		if ok {
			corsOptionsPath, _ := channel.config.GetString(ctx, constants.CONF_HTTPENGINE_CORSOPTIONSPATH)
			corsMw := cors.New(cors.Options{
				AllowedOrigins:     allowedOrigins,
				AllowedHeaders:     []string{"*"},
				AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"},
				ExposedHeaders:     []string{"*"},
				OptionsPassthrough: true,
				AllowCredentials:   true,
			})
			switch channel.engine.fwname {
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
	return nil
}

func (channel *httpChannel) child(ctx core.ServerContext, name string, channelConfig config.Config) (*httpChannel, error) {
	ctx = ctx.SubContext("Channel " + name)
	path, ok := channelConfig.GetString(ctx, constants.CONF_HTTPENGINE_PATH)
	if !ok {
		errors.BadConf(ctx, constants.CONF_HTTPENGINE_PATH)
	}
	svc, found := channelConfig.GetString(ctx, constants.CONF_CHANNEL_SERVICE)

	var routername string
	var router net.Router
	if found {
		router = channel.Router
		routername = fmt.Sprintf("%s > %s", channel.name, name)
	} else {
		routername = fmt.Sprintf("%s > %s", channel.name, name)
		router = channel.Router.Group(path)
	}

	log.Trace(ctx, "Creating child channel ", "Parent", channel.name, "Name", name, "Service", svc, "Path", path)
	childChannel := &httpChannel{name: routername, Router: router, config: channelConfig, engine: channel.engine, svcName: svc, path: path}
	err := childChannel.configure(ctx)
	if err != nil {
		return nil, err
	}
	return childChannel, nil
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
		corectx, err := ctx.CreateNewRequest(serviceName, pathCtx, "")
		if(err!=nil) {
			return errors.WrapError(ctx, err)
		}
	
		httpreq := pathCtx.GetRequest()
		corectx.SetGaeReq(httpreq)
		log.Info(corectx, "Got request", "Path", httpreq.URL.RequestURI(), "channel", channel.name, "method", httpreq.Method)
		defer corectx.CompleteRequest()
		go func(reqctx core.RequestContext, webctx net.WebContext) {
			bytes, err := webctx.GetBody()
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
		err = <-errChannel
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
