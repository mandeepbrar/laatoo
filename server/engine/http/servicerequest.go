package http

import (
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/engine/http/net"
)

type objectType int

const (
	stringmap objectType = iota
	bytes
	files
	stringtype
	custom
)

func (channel *httpChannel) processServiceRequest(ctx core.ServerContext, respHandler server.ServiceResponseHandler, method string, routename string,
	svc server.Service, otype objectType, dataObjectName string, isdataObject bool, isdataCollection bool, dataObjectCreator core.ObjectCreator,
	dataObjectCollectionCreator core.ObjectCollectionCreator, routeParams map[string]string, staticValues map[string]interface{},
	headers map[string]string, allowedQParams map[string]bool) (core.ServiceFunc, error) {
	return func(webctx core.RequestContext) error {
		var reqData interface{}
		var err error
		engineContext := webctx.EngineRequestContext().(net.WebContext)
		/*if method == CONF_ROUTE_METHOD_INVOKE {
			service := engineContext.GetQueryParam(CONF_ROUTE_SERVICE)
			if len(service) == 0 {
				return errors.ThrowError(webctx, errors.CORE_ERROR_MISSING_ARG, "Missing argument", CONF_ROUTE_SERVICE)
			}
			serverCtx := webctx.ParentContext().(core.ServerContext)
			svc, serverElement, err = serverCtx.GetService(service)
			if err != nil || svc == nil {
				return errors.RethrowError(webctx, errors.CORE_ERROR_BAD_ARG, err, "No such service has been created", service)
			}
			respHandler = svc.GetResponseHandler()
			if respHandler == nil {
				respHandler = router
			}
		}*/
		log.Trace(webctx, "Received request ", "route", routename, "dataObjectName", dataObjectName)
		if isdataObject {
			switch otype {
			case stringmap:
				mapobj := make(map[string]interface{}, 10)
				reqData = &mapobj
				if method != "GET" {
					err = engineContext.Bind(reqData)
					if err != nil {
						log.Trace(webctx, "Could not unmarshal Json ", "data", reqData, "err", err)
						webctx.SetResponse(core.StatusBadRequestResponse)
						return respHandler.HandleResponse(webctx)
					}
					log.Trace(webctx, "Request data bound ", "data", reqData)
				}
			case bytes:
				reqData, err = engineContext.GetBody()
				if err != nil {
					log.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
					webctx.SetResponse(core.StatusBadRequestResponse)
					return respHandler.HandleResponse(webctx)
				}
			case stringtype:
				reqDataBytes, err := engineContext.GetBody()
				if err != nil {
					log.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
					webctx.SetResponse(core.StatusBadRequestResponse)
					return respHandler.HandleResponse(webctx)
				}
				reqData = string(reqDataBytes)
			case files:
				fileObjs, err := engineContext.GetFiles()
				reqData = &fileObjs
				if err != nil {
					log.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
					webctx.SetResponse(core.StatusBadRequestResponse)
					return respHandler.HandleResponse(webctx)
				}
			default:
				if isdataCollection {
					reqData = dataObjectCollectionCreator(5)
				} else {
					reqData = dataObjectCreator()
				}
				if method != "GET" {
					err = engineContext.Bind(reqData)
					if err != nil {
						log.Trace(webctx, "Could not unmarshal Json ", "data", reqData, "err", err)
						webctx.SetResponse(core.StatusBadRequestResponse)
						return respHandler.HandleResponse(webctx)
					}
					log.Trace(webctx, "Request data bound ", "data", reqData)
				}
			}
		} else {
			reqData, err = engineContext.GetRequestStream()
			if err != nil {
				log.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
				webctx.SetResponse(core.StatusBadRequestResponse)
				return respHandler.HandleResponse(webctx)
			}
		}
		return channel.processRequest(webctx, reqData, engineContext, respHandler, routename, svc, routeParams, staticValues, headers, allowedQParams)
	}, nil
}

/*
func (router *routerImpl) processStreamServiceRequest(ctx core.ServerContext, respHandler server.ServiceResponseHandler, method string, routename string,
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
func (channel *httpChannel) processRequest(webctx core.RequestContext, reqData interface{}, engineContext net.WebContext, respHandler server.ServiceResponseHandler, routename string,
	svc server.Service, routeParams map[string]string, staticValues map[string]interface{}, headers map[string]string, allowedQParams map[string]bool) error {
	var err error
	log.Trace(webctx, "Invoking service ", "router", routename, "routeParams", routeParams, "staticValues", staticValues, "headers", headers, "allowedQParams", allowedQParams)
	reqctx := webctx.SubContext(svc.GetName())
	defer reqctx.CompleteRequest()
	reqctx.SetRequest(reqData)
	if routeParams != nil {
		for param, routeParamName := range routeParams {
			paramVal := engineContext.GetRouteParam(routeParamName)
			reqctx.Set(param, paramVal)
		}
	}
	log.Trace(webctx, "Headers provided", "headers", headers)
	if headers != nil {
		for param, header := range headers {
			headerVal := engineContext.GetHeader(header)
			reqctx.Set(param, headerVal)
			log.Trace(webctx, "Setting header param ", "headers", param, "value ", headerVal)
		}
	}

	queryParams := engineContext.GetQueryParams()
	for param, _ := range queryParams {
		_, found := allowedQParams[param]
		if found {
			reqctx.Set(param, engineContext.GetQueryParam(param))
		} else {
			log.Info(webctx, "Parameter not allowed in request", "parameter", param)
		}
	}

	if staticValues != nil {
		for name, val := range staticValues {
			reqctx.Set(name, val)
		}
	}
	err = svc.Invoke(reqctx)
	if err != nil {
		return errors.WrapError(reqctx, err)
	}
	log.Trace(webctx, "Completed request for service. Handling Response")
	return respHandler.HandleResponse(reqctx)

}
