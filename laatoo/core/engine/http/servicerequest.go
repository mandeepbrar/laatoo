package http

import (
	"laatoo/core/engine/http/net"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

func (router *Router) processServiceRequest(ctx core.ServerContext, respHandler core.ServiceResponseHandler, method string, routename string,
	svc core.Service, service string, dataObjectName string, isdataObject bool, isdataCollection bool, dataObjectCreator core.ObjectCreator,
	dataObjectCollectionCreator core.ObjectCollectionCreator, routeParams map[string]string, staticValues map[string]string, headers map[string]string) core.ServiceFunc {
	return func(webctx core.RequestContext) error {
		var reqData interface{}
		var err error
		engineContext := webctx.EngineContext().(net.WebContext)
		if method == CONF_ROUTE_METHOD_INVOKE {
			service = engineContext.GetQueryParam(CONF_ROUTE_SERVICE)
			if len(service) == 0 {
				return errors.ThrowError(webctx, errors.CORE_ERROR_MISSING_ARG, "Missing argument", CONF_ROUTE_SERVICE)
			}
			svc, err = webctx.GetService(service)
			if err != nil || svc == nil {
				return errors.RethrowError(webctx, errors.CORE_ERROR_BAD_ARG, err, "No such service has been created", service)
			}
			respHandler = svc.GetResponseHandler()
			if respHandler == nil {
				respHandler = router
			}
		}
		log.Logger.Trace(webctx, "Received request ", "route", routename, "service", service, "dataObjectName", dataObjectName)
		if isdataObject {
			if dataObjectName == CONF_STRINGMAP_DATA_OBJECT {
				mapobj := make(map[string]interface{}, 10)
				reqData = &mapobj
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
				log.Logger.Trace(webctx, "Could not unmarshal Json ", "data", reqData, "err", err)
				webctx.SetResponse(core.StatusBadRequestResponse)
				return respHandler.HandleResponse(webctx)
			}
			log.Logger.Trace(webctx, "Request data bound ", "data", reqData)
		} else {
			reqData, err = engineContext.GetBody()
			if err != nil {
				log.Logger.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
				webctx.SetResponse(core.StatusBadRequestResponse)
				return respHandler.HandleResponse(webctx)
			}
		}
		return router.processRequest(webctx, reqData, engineContext, respHandler, routename, svc, service, routeParams, staticValues, headers)
	}
}

func (router *Router) processStreamServiceRequest(ctx core.ServerContext, respHandler core.ServiceResponseHandler, method string, routename string,
	svc core.Service, service string, routeParams map[string]string, staticValues map[string]string, headers map[string]string) core.ServiceFunc {
	return func(webctx core.RequestContext) error {
		log.Logger.Trace(webctx, "Received request ", "route", routename, "service", service)
		engineContext := webctx.EngineContext().(net.WebContext)
		reqData, err := engineContext.GetBody()
		if err != nil {
			log.Logger.Trace(webctx, "Could not read stream", "data", reqData, "err", err)
			webctx.SetResponse(core.StatusBadRequestResponse)
			return respHandler.HandleResponse(webctx)
		}
		return router.processRequest(webctx, reqData, engineContext, respHandler, routename, svc, service, routeParams, staticValues, headers)
	}
}
func (router *Router) processRequest(webctx core.RequestContext, reqData interface{}, engineContext net.WebContext, respHandler core.ServiceResponseHandler, routename string,
	svc core.Service, service string, routeParams map[string]string, staticValues map[string]string, headers map[string]string) error {
	var err error
	log.Logger.Trace(webctx, "Invoking service ", "router", routename, "service", service, "routeParams", routeParams, "staticValues", staticValues, "headers", headers)
	reqctx := webctx.SubRequest(service, svc.GetConf())
	defer reqctx.CompleteRequest()
	reqctx.SetRequestBody(reqData)
	if routeParams != nil {
		for param, routeParamName := range routeParams {
			paramVal := engineContext.GetRouteParam(routeParamName)
			reqctx.Set(param, paramVal)
		}
	}
	if headers != nil {
		for param, header := range headers {
			headerVal := engineContext.GetHeader(header)
			reqctx.Set(param, headerVal)
		}
	}

	queryParams := engineContext.GetQueryParams()
	for param, _ := range queryParams {
		reqctx.Set(param, engineContext.GetQueryParam(param))
	}

	if staticValues != nil {
		for name, val := range staticValues {
			reqctx.Set(name, val)
		}
	}
	err = svc.Invoke(reqctx)
	if err != nil {
		return errors.WrapError(webctx, err)
	}
	log.Logger.Trace(webctx, "Completed request for service. Handling Response", "service", service)
	return respHandler.HandleResponse(reqctx)

}
