package http

import (
	//	"laatoo/core/common"

	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
	//"strconv"
)

func (channel *httpChannel) serve(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Serve")

	log.Trace(ctx, "Channel config", "name", channel.name, "config", channel.config)

	disabled, _ := channel.config.GetBool(ctx, constants.CONF_HTTPENGINE_DISABLEROUTE)
	if disabled {
		return nil
	}

	svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
	svc, err := svcManager.GetService(ctx, channel.svcName)
	if err != nil {
		return err
	}

	bodyParam := "Data"
	body, ok := channel.config.GetString(ctx, constants.CONF_HTTPENGINE_BODY)
	if ok {
		bodyParam = body
	}

	var respHandler elements.ServiceResponseHandler
	handler := ctx.GetServerElement(core.ServerElementServiceResponseHandler)
	if handler != nil {
		respHandler = handler.(elements.ServiceResponseHandler)
	} else {
		respHandler = DefaultResponseHandler(ctx)
	}

	//svcParams := svc.ParamsConfig()

	////build value parameters
	var routeParams map[string]string
	routeParamValuesConf, ok := channel.config.GetSubConfig(ctx, constants.CONF_HTTPENGINE_ROUTEPARAMVALUES)
	if ok {
		values := routeParamValuesConf.AllConfigurations(ctx)
		routeParams = make(map[string]string, len(values))
		for _, paramname := range values {
			routeParams[paramname], _ = routeParamValuesConf.GetString(ctx, paramname)
		}
	}

	allowedQueryParams := make(map[string]bool)

	if channel.allowedQParams != nil {
		for _, p := range channel.allowedQParams {
			allowedQueryParams[p] = true
		}
	}

	////build value parameters
	staticValues := make(map[string]interface{})
	staticValuesConf, ok := channel.config.GetSubConfig(ctx, constants.CONF_HTTPENGINE_STATICVALUES)
	if ok {
		values := staticValuesConf.AllConfigurations(ctx)
		for _, paramname := range values {
			staticValues[paramname], _ = staticValuesConf.Get(ctx, paramname)
		}
	}

	//build header param mappings
	headers := make(map[string]string, 0)
	headersConf, ok := channel.config.GetSubConfig(ctx, constants.CONF_HTTPENGINE_HEADERSTOINCLUDE)
	if ok {
		headersToInclude := headersConf.AllConfigurations(ctx)
		for _, paramName := range headersToInclude {
			header, _ := headersConf.GetString(ctx, paramName)
			headers[paramName] = header
		}
	}

	webReqHandler, err := channel.processServiceRequest(ctx, channel.method, channel.name, svc, routeParams, staticValues, headers, allowedQueryParams, bodyParam)
	if err != nil {
		return err
	}

	switch channel.method {
	case "GET":
		err = channel.get(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
	case "POST":
		err = channel.post(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
	case "PUT":
		err = channel.put(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
	case "DELETE":
		err = channel.delete(ctx, channel.path, svc.GetName(), webReqHandler, respHandler, svc)
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
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
