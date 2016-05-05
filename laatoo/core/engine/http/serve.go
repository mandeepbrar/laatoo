package http

import (
	//	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	//"strconv"
)

func (channel *httpChannel) serve(ctx core.ServerContext, svc server.Service, routeConf config.Config) error {
	path, ok := routeConf.GetString(config.CONF_HTTPENGINE_PATH)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_HTTPENGINE_PATH)
	}
	method, ok := routeConf.GetString(config.CONF_HTTPENGINE_METHOD)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_HTTPENGINE_METHOD)
	}
	var respHandler server.ServiceResponseHandler
	handler := ctx.GetServerElement(core.ServerElementServiceResponseHandler)
	if handler != nil {
		respHandler = handler.(server.ServiceResponseHandler)
	} else {
		respHandler = DefaultResponseHandler(ctx)
	}

	var err error

	////build value parameters
	var routeParams map[string]string
	routeParamValuesConf, ok := routeConf.GetSubConfig(config.CONF_HTTPENGINE_ROUTEPARAMVALUES)
	if ok {
		values := routeParamValuesConf.AllConfigurations()
		routeParams = make(map[string]string, len(values))
		for _, paramname := range values {
			routeParams[paramname], _ = routeParamValuesConf.GetString(paramname)
		}
	}

	////build value parameters
	var staticValues map[string]string
	staticValuesConf, ok := routeConf.GetSubConfig(config.CONF_HTTPENGINE_STATICVALUES)
	if ok {
		values := staticValuesConf.AllConfigurations()
		staticValues = make(map[string]string, len(values))
		for _, paramname := range values {
			staticValues[paramname], _ = staticValuesConf.GetString(paramname)
		}
	}

	//build header param mappings
	headers := make(map[string]string, 0)
	headersConf, ok := routeConf.GetSubConfig(config.CONF_HTTPENGINE_HEADERSTOINCLUDE)
	if ok {
		headersToInclude := headersConf.AllConfigurations()
		for _, paramName := range headersToInclude {
			header, _ := headersConf.GetString(paramName)
			headers[paramName] = header
		}
	}

	//get any data creators for body objects that need to be bound
	var dataObjectCreator core.ObjectCreator
	var dataObjectCollectionCreator core.ObjectCollectionCreator
	dataObjectName, isdataObject := routeConf.GetString(config.CONF_ENGINE_DATA_OBJECT)
	_, isdataCollection := routeConf.GetString(config.CONF_ENGINE_DATA_COLLECTION)

	var otype objectType
	switch dataObjectName {
	case config.CONF_ENGINE_STRINGMAP_DATA_OBJECT:
		otype = stringmap
	case config.CONF_ENGINE_BYTES_DATA_OBJECT:
		otype = bytes
	case config.CONF_ENGINE_FILES_DATA_OBJECT:
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
	log.Logger.Trace(ctx, "Service mapping for route", "name", channel.name, "method", method, "dataObjectName", dataObjectName, "isdataObject", isdataObject, "isdataCollection", isdataCollection)

	webReqHandler, err := channel.processServiceRequest(ctx, respHandler, method, channel.name, svc, otype, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParams, staticValues, headers)
	if err != nil {
		return err
	}
	switch method {
	case "GET":
		channel.get(ctx, path, webReqHandler)
	case "POST":
		channel.post(ctx, path, webReqHandler)
	case "PUT":
		channel.put(ctx, path, webReqHandler)
	case "DELETE":
		channel.delete(ctx, path, webReqHandler)
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
