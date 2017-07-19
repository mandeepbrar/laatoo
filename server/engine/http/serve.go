package http

import (
	//	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
	"laatoo/server/constants"
	//"strconv"
)

func (channel *httpChannel) serve(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Serve")

	disabled, _ := channel.config.GetBool(constants.CONF_HTTPENGINE_DISABLEROUTE)
	if disabled {
		return nil
	}

	svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(server.ServiceManager)
	svc, err := svcManager.GetService(ctx, channel.svcName)
	if err != nil {
		return err
	}

	method, ok := channel.config.GetString(constants.CONF_HTTPENGINE_METHOD)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_HTTPENGINE_METHOD)
	}

	var respHandler server.ServiceResponseHandler
	handler := ctx.GetServerElement(core.ServerElementServiceResponseHandler)
	if handler != nil {
		respHandler = handler.(server.ServiceResponseHandler)
	} else {
		respHandler = DefaultResponseHandler(ctx)
	}

	//svcParams := svc.ParamsConfig()

	////build value parameters
	var routeParams map[string]string
	routeParamValuesConf, ok := channel.config.GetSubConfig(constants.CONF_HTTPENGINE_ROUTEPARAMVALUES)
	if ok {
		values := routeParamValuesConf.AllConfigurations()
		routeParams = make(map[string]string, len(values))
		for _, paramname := range values {
			routeParams[paramname], _ = routeParamValuesConf.GetString(paramname)
		}
	}

	allowedQueryParams := make(map[string]bool)

	if channel.allowedQParams != nil {
		for _, p := range channel.allowedQParams {
			allowedQueryParams[p] = true
		}
	}

	allowedQParamsFunc := func(confElem config.Config) {
		if confElem == nil {
			return
		}
		allowedParams, ok := confElem.GetStringArray(constants.CONF_HTTPENGINE_ALLOWEDQUERYPARAMS)
		if ok {
			for _, p := range allowedParams {
				allowedQueryParams[p] = true
			}
		}
	}
	allowedQParamsFunc(channel.config)
	//allowedQParamsFunc(svcParams)

	////build value parameters
	staticValues := make(map[string]interface{})
	staticValuesFunc := func(confElem config.Config) {
		if confElem == nil {
			return
		}
		staticValuesConf, ok := confElem.GetSubConfig(constants.CONF_HTTPENGINE_STATICVALUES)
		if ok {
			values := staticValuesConf.AllConfigurations()
			for _, paramname := range values {
				staticValues[paramname], _ = staticValuesConf.Get(paramname)
			}
		}
	}
	staticValuesFunc(channel.config)
	//staticValuesFunc(svcParams)

	//build header param mappings
	headers := make(map[string]string, 0)
	headersFunc := func(confElem config.Config) {
		if confElem == nil {
			return
		}
		headersConf, ok := confElem.GetSubConfig(constants.CONF_HTTPENGINE_HEADERSTOINCLUDE)
		if ok {
			headersToInclude := headersConf.AllConfigurations()
			for _, paramName := range headersToInclude {
				header, _ := headersConf.GetString(paramName)
				headers[paramName] = header
			}
		}
	}
	headersFunc(channel.config)
	//headersFunc(svcParams)

	webReqHandler, err := channel.processServiceRequest(ctx, method, channel.name, svc, routeParams, staticValues, headers, allowedQueryParams)
	if err != nil {
		return err
	}
	switch method {
	case "GET":
		channel.get(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
	case "POST":
		channel.post(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
	case "PUT":
		channel.put(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
	case "DELETE":
		channel.delete(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
		/*	case CONF_ROUTE_METHOD_INVOKE:
					router.Post(ctx, path, router.processServiceRequest(ctx, respHandler, method, router.name, svc, serverElement, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParams, staticValues, headers))
			case CONF_ROUTE_METHOD_GETSTREAM:
				{
					router.Post(ctx, path, router.processStreamServiceRequest(ctx, respHandler, method, router.name, svc, routeParams, staticValues, headers))
				}
			case CONF_ROUTE_METHOD_POSTSTREAM:
				{
					router.Post(ctx, path, router.processStreamServiceRequest(ctx, respHandler, method, router.name, svc, routeParams, staticValues, headers))
				}*/
	}
	return nil
}
