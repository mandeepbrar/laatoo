package http

import (
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"laatoosdk/auth"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
	"laatoo/server/engine/http/common"
	"laatoo/server/engine/http/net"
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
	group          bool
	disabled       bool
	adapter        *common.WebFWAdapter
	Router         net.Router
	config         config.Config
	svcName        string
	engine         *httpEngine
	allowedQParams []string
	allowedCookies []string
	allowedHeaders []string
	skipAuth       bool
	path           string
	codec          core.Codec
	parentChannel  *httpChannel
	svrContext     core.ServerContext
}

func (channel *httpChannel) initialize(ctx core.ServerContext) error {
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

	allowedCookies, ok := channel.config.GetStringArray(ctx, constants.CONF_HTTPENGINE_ALLOWEDCOOKIES)
	if ok {
		if channel.parentChannel != nil && channel.parentChannel.allowedCookies != nil {
			channel.allowedCookies = append(channel.parentChannel.allowedCookies, allowedCookies...)
		} else {
			channel.allowedCookies = allowedCookies
		}
	} else {
		if channel.parentChannel != nil {
			channel.allowedCookies = channel.parentChannel.allowedCookies
		}
	}

	allowedHeaders, ok := channel.config.GetStringArray(ctx, constants.CONF_HTTPENGINE_HEADERSTOINCLUDE)
	if ok {
		if channel.parentChannel != nil && channel.parentChannel.allowedHeaders != nil {
			channel.allowedHeaders = append(channel.parentChannel.allowedHeaders, allowedHeaders...)
		} else {
			channel.allowedHeaders = allowedHeaders
		}
	} else {
		if channel.parentChannel != nil {
			channel.allowedHeaders = channel.parentChannel.allowedHeaders
		}
	}

	if !channel.group {
		method, ok := channel.config.GetString(ctx, constants.CONF_HTTPENGINE_METHOD)
		if !ok {
			return errors.MissingConf(ctx, constants.CONF_HTTPENGINE_METHOD)
		}
		channel.method = method
	}

	usecors, _ := channel.config.GetBool(ctx, constants.CONF_HTTPENGINE_USECORS)

	if usecors {
		allowedOrigins, ok := channel.config.GetStringArray(ctx, constants.CONF_HTTPENGINE_CORSHOSTS)
		if ok {
			corsOptionsPath, _ := channel.config.GetString(ctx, constants.CONF_HTTPENGINE_CORSOPTIONSPATH)
			if err := channel.adapter.SetupCors(channel.Router, allowedOrigins, corsOptionsPath); err != nil {
				return errors.WrapError(ctx, err)
			}
			log.Info(ctx, "CORS enabled for hosts ", "hosts", allowedOrigins)
		}
	}

	codecname := "json"
	/*co, ok := vals["encoding"]
	if ok {
		codecname = co.(string)
	}*/
	channel.codec, ok = ctx.GetCodec(codecname)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_CODEC_NOT_FOUND)
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
	group := false
	if found {
		router = channel.Router
		routername = fmt.Sprintf("%s > %s", channel.name, name)
	} else {
		routername = fmt.Sprintf("%s > %s", channel.name, name)
		router = channel.adapter.Group(channel.Router, path, channel.name)
		group = true
	}

	log.Trace(ctx, "Creating child channel ", "Parent", channel.name, "Name", name, "Service", svc, "Path", path)
	childChannel := &httpChannel{name: routername, parentChannel: channel, Router: router, adapter: channel.adapter,
		config: channelConfig, group: group, engine: channel.engine, svcName: svc, path: path, disabled: false, svrContext: ctx}
	err := childChannel.initialize(ctx)
	if err != nil {
		return nil, err
	}
	return childChannel, nil
}

func (channel *httpChannel) httpAdapter(ctx core.ServerContext, serviceName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) net.HandlerFunc {
	var shandler elements.SecurityHandler
	var authtoken string
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler = sh.(elements.SecurityHandler)
		val := sh.GetProperty(config.AUTHHEADER)
		authtoken = val.(string)
	}

	authRequest := func(reqctx core.RequestContext, vals map[string]interface{}) error {
		if (!channel.skipAuth) && (shandler != nil) {
			token, found := vals[authtoken]
			log.Error(reqctx, "Testing authentication", "vals", vals, "token", authtoken)
			if !found {
				return errors.Unauthorized(reqctx)
			} else {
				reqctx.Set(authtoken, token)
			}
			_, err := shandler.AuthenticateRequest(reqctx, false)
			if err != nil {
				return err
			}
			return nil
		} else {
			log.Info(reqctx, "Auth skipped")
			return nil
		}
	}

	return func(pathCtx net.WebContext) error {
		if channel.disabled {
			pathCtx.NoContent(400)
			return nil
		}
		errChannel := make(chan error)
		go func(pathCtx net.WebContext) {
			reqCtx, vals, err := reqBuilder(pathCtx)
			if err != nil {
				err = respHandler.HandleResponse(reqCtx, core.BadRequestResponse(err.Error()), err)
				errChannel <- err
				return
			}
			defer reqCtx.CompleteRequest()
			err = authRequest(reqCtx, vals)
			if err != nil {
				err = respHandler.HandleResponse(reqCtx, core.StatusUnauthorizedResponse, err)
				errChannel <- err
				return
			}
			log.Trace(reqCtx, "Invoking service ", "vals", vals)
			requestCtx := reqCtx.SubContext("Http request handler")
			resp, err := svc.HandleRequest(requestCtx, vals)

			log.Trace(reqCtx, "Completed request for service. Handling Response")
			resCtx := reqCtx.SubContext("Response Handler")
			err = respHandler.HandleResponse(resCtx, resp, err)

			errChannel <- err
		}(pathCtx)
		err := <-errChannel
		if err != nil {
			log.Info(ctx, "Got error in the request", "error", err)
		}
		return err
	}
}

func (channel *httpChannel) get(ctx core.ServerContext, path string, serviceName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) error {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	return channel.adapter.Get(channel.Router, path, channel.name, channel.httpAdapter(ctx, serviceName, reqBuilder, respHandler, svc))
}

func (channel *httpChannel) options(ctx core.ServerContext, path string, serviceName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) error {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Options")
	return channel.adapter.Options(channel.Router, path, channel.name, channel.httpAdapter(ctx, serviceName, reqBuilder, respHandler, svc))
}

func (channel *httpChannel) put(ctx core.ServerContext, path string, serviceName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) error {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	return channel.adapter.Put(channel.Router, path, channel.name, channel.httpAdapter(ctx, serviceName, reqBuilder, respHandler, svc))
}

func (channel *httpChannel) post(ctx core.ServerContext, path string, serviceName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) error {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	return channel.adapter.Post(channel.Router, path, channel.name, channel.httpAdapter(ctx, serviceName, reqBuilder, respHandler, svc))
}

func (channel *httpChannel) delete(ctx core.ServerContext, path string, serviceName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) error {
	log.Info(ctx, "Registering route", "channel", channel.name, "path", path, "method", "Get")
	return channel.adapter.Delete(channel.Router, path, channel.name, channel.httpAdapter(ctx, serviceName, reqBuilder, respHandler, svc))
}

func (channel *httpChannel) use(ctx core.ServerContext, middlewareName string, reqBuilder RequestBuilder, respHandler elements.ServiceResponseHandler, svc elements.Service) {
	channel.adapter.Use(channel.Router, channel.httpAdapter(ctx, middlewareName, reqBuilder, respHandler, svc))
}

func (channel *httpChannel) destruct(ctx core.ServerContext, parentChannel *httpChannel) error {
	return channel.adapter.RemovePath(channel.Router, channel.path, channel.name, channel.method)
}
