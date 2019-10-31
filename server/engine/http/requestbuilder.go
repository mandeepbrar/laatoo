package http

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/engine/http/net"
)

type RequestBuilder func(engineContext net.WebContext) (core.RequestContext, map[string]interface{}, error)

func (channel *httpChannel) getRequestBuilder(ctx core.ServerContext, method string, routename string, svc elements.Service,
	routeParams map[string]string, staticValues map[string]interface{}, allowedQParams map[string]bool,
	bodyParamName, bodyParamType string) (RequestBuilder, error) {
	includeBody := bodyParamType != ""
	multipart := bodyParamType == config.OBJECTTYPE_FILES
	return func(engineContext net.WebContext) (core.RequestContext, map[string]interface{}, error) {

		reqCtx, err := ctx.CreateNewRequest(channel.svcName, channel.engine.proxy, engineContext, "")
		if err != nil {
			return nil, nil, errors.WrapError(ctx, err)
		}

		httpreq := engineContext.GetRequest()
		reqCtx.SetGaeReq(httpreq)
		log.Info(reqCtx, "Got request", "Path", httpreq.URL.RequestURI(), "channel", channel.name, "method", httpreq.Method)

		vals := make(map[string]interface{})

		log.Trace(reqCtx, "Invoking service ", "router", routename, "routeParams", routeParams, "staticValues", staticValues, "allowedQParams", allowedQParams)

		if includeBody {
			if multipart {
				files, err := engineContext.GetFiles()
				if err != nil {
					return reqCtx, vals, err
				}
				vals[bodyParamName] = files
			} else {
				bytes, err := engineContext.GetBody()
				if err != nil {
					return reqCtx, vals, err
				}
				vals[bodyParamName] = bytes
			}
		}

		if routeParams != nil {
			for param, routeParamName := range routeParams {
				paramVal := engineContext.GetRouteParam(routeParamName)
				vals[param] = channel.encodeString(ctx, paramVal)
			}
		}

		if channel.allowedHeaders != nil {
			for _, headerName := range channel.allowedHeaders {
				headerVal := engineContext.GetHeader(headerName)
				vals[headerName] = channel.encodeString(ctx, headerVal)
			}
		}

		if channel.allowedCookies != nil {
			for _, cookieName := range channel.allowedCookies {
				cookie, _ := engineContext.GetCookie(cookieName)
				if cookie != nil {
					log.Trace(reqCtx, "Found cookies", "cookieName", cookie)
					vals[cookieName] = channel.encodeString(ctx, cookie.Value)
				}
			}
		}

		queryParams := engineContext.GetQueryParams()
		for param, _ := range queryParams {
			_, found := allowedQParams[param]
			if found {
				vals[param] = channel.encodeString(ctx, engineContext.GetQueryParam(param))
			} else {
				log.Info(reqCtx, "Parameter not allowed in request", "parameter", param)
			}
		}

		if staticValues != nil {
			for name, val := range staticValues {
				//already encoded by codec
				vals[name] = val
			}
		}

		return reqCtx, vals, nil
	}, nil
}

func (channel *httpChannel) encodeString(ctx core.ServerContext, val string) []byte {
	encVal, err := channel.codec.Marshal(ctx, val)
	if err != nil {
		log.Error(ctx, "Codec could not encode string", "val", val, "err", err)
	}
	return encVal
}

/*
func (router *routerImpl) processStreamServiceRequest(ctx core.ServerContext, respHandler elements.ServiceResponseHandler, method string, routename string,
	svc core.Service, serverElement core.ServerElement, routeParams map[string]string, staticValues map[string]string, headers map[string]string) core.ServiceFunc {
	return func(webctx core.RequestContext) error {
		log.Trace(webctx, "Received request ", "route", routename, "service", serverElement.GetName())
		engineContext := webctx.EngineContext().(net.WebContext)
		reqData, err := engineContext.GetBody()
		if err != nil {
			log.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
			webctx.SetResponse(core.StatusBadRequestResponse)
			return respHandler.HandleResponse(webctx)
		}
		return router.processRequest(webctx, reqData, engineContext, respHandler, routename, svc, serverElement, routeParams, staticValues, headers)
	}
}*/
