package http

import (
	//	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
	//"strconv"
)

func (channel *httpChannel) serve(ctx core.ServerContext, svc server.Service, routeConf config.Config) error {
	ctx = ctx.SubContext("Serve")
	disabled, _ := routeConf.GetBool(constants.CONF_HTTPENGINE_DISABLEROUTE)
	if disabled {
		return nil
	}
	path, ok := routeConf.GetString(constants.CONF_HTTPENGINE_PATH)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_HTTPENGINE_PATH)
	}
	method, ok := routeConf.GetString(constants.CONF_HTTPENGINE_METHOD)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_HTTPENGINE_METHOD)
	}
	var respHandler server.ServiceResponseHandler
	handler := ctx.GetServerElement(core.ServerElementServiceResponseHandler)
	if handler != nil {
		respHandler = handler.(server.ServiceResponseHandler)
	} else {
		respHandler = DefaultResponseHandler(ctx)
	}

	svcParams := svc.ParamsConfig()

	var err error

	////build value parameters
	var routeParams map[string]string
	routeParamValuesConf, ok := routeConf.GetSubConfig(constants.CONF_HTTPENGINE_ROUTEPARAMVALUES)
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
	allowedQParamsFunc(routeConf)
	allowedQParamsFunc(svcParams)

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
	staticValuesFunc(routeConf)
	staticValuesFunc(svcParams)

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
	headersFunc(routeConf)
	headersFunc(svcParams)

	//get any data creators for body objects that need to be bound
	var dataObjectCreator core.ObjectCreator
	var dataObjectCollectionCreator core.ObjectCollectionCreator
	dataObjectName, isdataObject := routeConf.GetString(constants.CONF_ENGINE_DATA_OBJECT)
	_, isdataCollection := routeConf.GetString(constants.CONF_ENGINE_DATA_COLLECTION)

	var otype objectType
	switch dataObjectName {
	case constants.CONF_ENGINE_STRINGMAP_DATA_OBJECT:
		otype = stringmap
	case constants.CONF_ENGINE_BYTES_DATA_OBJECT:
		otype = bytes
	case constants.CONF_ENGINE_STRING_DATA_OBJECT:
		otype = stringtype
	case constants.CONF_ENGINE_FILES_DATA_OBJECT:
		otype = files
	default:
		otype = custom
	}

	if isdataObject && (otype == custom) {
		if isdataCollection {
			dataObjectCollectionCreator, err = ctx.GetObjectCollectionCreator(dataObjectName)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", dataObjectName)
			}
		} else {
			dataObjectCreator, err = ctx.GetObjectCreator(dataObjectName)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", dataObjectName)
			}
		}
	}
	log.Trace(ctx, "Service mapping for route", "name", channel.name, "method", method, "dataObjectName", dataObjectName, "isdataObject", isdataObject, "isdataCollection", isdataCollection)

	webReqHandler, err := channel.processServiceRequest(ctx, respHandler, method, channel.name, svc, otype, dataObjectName, isdataObject,
		isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParams, staticValues, headers, allowedQueryParams)
	if err != nil {
		return err
	}
	switch method {
	case "GET":
		channel.get(ctx, path, svc.GetName(), webReqHandler, svc)
	case "POST":
		channel.post(ctx, path, svc.GetName(), webReqHandler, svc)
	case "PUT":
		channel.put(ctx, path, svc.GetName(), webReqHandler, svc)
	case "DELETE":
		channel.delete(ctx, path, svc.GetName(), webReqHandler, svc)
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
