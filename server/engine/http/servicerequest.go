package http

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/engine/http/net"
)

type ServiceInvoker func(webctx core.RequestContext, vals map[string]interface{}, bytes interface{}) (*core.ServiceResponse, error)

func (channel *httpChannel) processServiceRequest(ctx core.ServerContext, method string, routename string,
	svc server.Service, routeParams map[string]string, staticValues map[string]interface{}, headers map[string]string, allowedQParams map[string]bool) (ServiceInvoker, error) {
	return func(webctx core.RequestContext, vals map[string]interface{}, body interface{}) (*core.ServiceResponse, error) {
		engineContext := webctx.EngineRequestContext().(net.WebContext)
		log.Trace(webctx, "Invoking service ", "router", routename, "routeParams", routeParams, "staticValues", staticValues, "headers", headers, "allowedQParams", allowedQParams)
		reqctx := webctx.SubContext(svc.GetName())
		defer reqctx.CompleteRequest()

		if routeParams != nil {
			for param, routeParamName := range routeParams {
				paramVal := engineContext.GetRouteParam(routeParamName)
				vals[param] = paramVal
			}
		}
		log.Trace(webctx, "Headers provided", "headers", headers)
		if headers != nil {
			for param, header := range headers {
				headerVal := engineContext.GetHeader(header)
				vals[param] = headerVal
				log.Trace(webctx, "Setting header param ", "headers", param, "value ", headerVal)
			}
		}

		queryParams := engineContext.GetQueryParams()
		for param, _ := range queryParams {
			_, found := allowedQParams[param]
			if found {
				vals[param] = engineContext.GetQueryParam(param)
			} else {
				log.Info(webctx, "Parameter not allowed in request", "parameter", param)
			}
		}

		if staticValues != nil {
			for name, val := range staticValues {
				vals[name] = val
			}
		}
		return svc.HandleRequest(reqctx, vals, body.([]byte))
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
