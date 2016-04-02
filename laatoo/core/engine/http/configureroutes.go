package http

import (
	"laatoo/core/common"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"strconv"
)

func (router *Router) ConfigureRoutes(ctx core.ServerContext) error {
	router.processRoutesGrp(ctx, router.config)
	/*	routepairs := router.config.AllConfigurations()
		for _, routename := range routepairs {
			routesconf, err := common.ConfigFileAdapter(router.config, routename)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "Router Name", routename)
			}
			routeCtx := ctx.SubContext(routename, routesconf)
			err = router.configureRoutesConf(routeCtx, routesconf)
			if err != nil {
				return err
			}
		}*/
	return nil
}

func (router *Router) processRoutesGrp(ctx core.ServerContext, conf config.Config) error {
	allroutegroups, ok := conf.GetSubConfig(CONF_GROUPS)
	if ok {
		routegroups := allroutegroups.AllConfigurations()
		for _, routegroupname := range routegroups {
			log.Logger.Trace(ctx, "Process Route group", "Route group", routegroupname)
			routegrpConfig, err := common.ConfigFileAdapter(allroutegroups, routegroupname)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for Route group", routegroupname)
			}
			rtgrpCtx := ctx.SubContext("Route Group:"+routegroupname, routegrpConfig)
			path, ok := routegrpConfig.GetString(CONF_ROUTE_PATH)
			if !ok {
				path = ""
			}
			grpRouter := router.group(rtgrpCtx, path, routegroupname, routegrpConfig)

			err = grpRouter.processRoutesGrp(rtgrpCtx, routegrpConfig)
			if err != nil {
				return err
			}
		}
	}

	//get a map of all the services
	routes, ok := conf.GetSubConfig(CONF_ROUTES)
	if !ok {
		return nil
	}
	routeCfgs := routes.AllConfigurations()
	for _, routeName := range routeCfgs {
		log.Logger.Trace(ctx, "Process Route ", "Route name", routeName)
		routeConfig, err := common.ConfigFileAdapter(routes, routeName)
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for route", routeName)
		}
		routeCtx := ctx.SubContext("Route:"+routeName, routeConfig)
		router.createRoute(routeCtx, routeConfig)
	}
	return nil
}

/*
func (router *Router) configureRoutesConf(ctx core.ServerContext, conf config.Config) error {
	groups, ok := conf.GetConfigArray(CONF_GROUPS)
	if !ok {
		return nil
	}
	for _, groupConf := range groups {
		if err := router.createGroup(ctx, groupConf); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil

}

func (router *Router) createGroup(ctx core.ServerContext, groupConf config.Config) error {
	path, ok := groupConf.GetString(CONF_ROUTE_PATH)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing path variable in group", groupConf)
	}
	grpRouter := router.group(ctx, path, groupConf)
	routes, ok := groupConf.GetConfigArray(CONF_ROUTES)
	if !ok {
		return nil
	}
	for _, routeConf := range routes {
		if err := grpRouter.createRoute(ctx, routeConf); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}*/

func (router *Router) createRoute(ctx core.ServerContext, routeConf config.Config) error {
	path, ok := routeConf.GetString(CONF_ROUTE_PATH)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing path variable in route", routeConf)
	}
	method, ok := routeConf.GetString(CONF_ROUTE_METHOD)
	if !ok {
		return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing method variable in route", path)
	}
	var service string
	var svc core.Service
	var respHandler core.ServiceResponseHandler
	var err error
	if method != CONF_ROUTE_METHOD_INVOKE {
		service, ok = routeConf.GetString(CONF_ROUTE_SERVICE)
		if !ok {
			return errors.ThrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, "Missing service variable in route", path)
		}
		svc, err = ctx.GetService(service)
		if err != nil || svc == nil {
			return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err, "No such service has been created", path, "service", service)
		}
		respHandler = svc.GetResponseHandler()
		if respHandler == nil {
			respHandler = router
		}
	}

	////build value parameters
	var routeParamValues map[string]string
	routeParamValuesConf, ok := routeConf.GetSubConfig(CONF_ROUTE_ROUTEPARAMVALUES)
	if ok {
		values := routeParamValuesConf.AllConfigurations()
		routeParamValues = make(map[string]string, len(values))
		for _, paramname := range values {
			routeParamValues[paramname], _ = routeParamValuesConf.GetString(paramname)
		}
	}

	//////build index parameters
	///index parameter overrides all route parameters
	routeParamIndicesConf, _ := routeConf.GetSubConfig(CONF_ROUTE_ROUTEPARAMINDICES)
	routeParamIndices, err := createRouteParamIndices(ctx, routeParamIndicesConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	//build header param mappings
	headerParams := make(map[string]string, 0)
	headersConf, ok := routeConf.GetSubConfig(CONF_ROUTE_HEADERSTOINCLUDE)
	if ok {
		headersToInclude := headersConf.AllConfigurations()
		for _, paramName := range headersToInclude {
			header, _ := headersConf.GetString(paramName)
			headerParams[paramName] = header
		}
	}

	//get any data creators for body objects that need to be bound
	var dataObjectCreator core.ObjectCreator
	var dataObjectCollectionCreator core.ObjectCollectionCreator
	dataObjectName, isdataObject := routeConf.GetString(CONF_ROUTE_DATA_OBJECT)
	_, isdataCollection := routeConf.GetString(CONF_ROUTE_DATA_COLLECTION)
	if isdataObject && (dataObjectName != CONF_STRINGMAP_DATA_OBJECT) {
		if isdataCollection {
			dataObjectCollectionCreator, err = registry.GetObjectCollectionCreator(ctx, dataObjectName)
			if err != nil {
				return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err, "No such object", dataObjectName)
			}
		} else {
			dataObjectCreator, err = registry.GetObjectCreator(ctx, dataObjectName)
			if err != nil {
				return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err, "No such object", dataObjectName)
			}
		}
	}
	log.Logger.Trace(ctx, "Service got data object ", "isdataObject", isdataObject, "isdataCollection", isdataCollection, "service", service)

	switch method {
	case "GET":
		router.Get(ctx, path, routeConf, router.processServiceRequest(ctx, respHandler, method, router.name, svc, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParamIndices, routeParamValues, headerParams))
	case "POST":
		router.Post(ctx, path, routeConf, router.processServiceRequest(ctx, respHandler, method, router.name, svc, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParamIndices, routeParamValues, headerParams))
	case "PUT":
		router.Put(ctx, path, routeConf, router.processServiceRequest(ctx, respHandler, method, router.name, svc, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParamIndices, routeParamValues, headerParams))
	case "DELETE":
		router.Delete(ctx, path, routeConf, router.processServiceRequest(ctx, respHandler, method, router.name, svc, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParamIndices, routeParamValues, headerParams))
	case CONF_ROUTE_METHOD_INVOKE:
		router.Post(ctx, path, routeConf, router.processServiceRequest(ctx, respHandler, method, router.name, svc, service, dataObjectName, isdataObject, isdataCollection, dataObjectCreator, dataObjectCollectionCreator, routeParamIndices, routeParamValues, headerParams))
	case CONF_ROUTE_METHOD_GETSTREAM:
		{
			router.Post(ctx, path, routeConf, router.processStreamServiceRequest(ctx, respHandler, method, router.name, svc, service, routeParamIndices, routeParamValues, headerParams))
		}
	case CONF_ROUTE_METHOD_POSTSTREAM:
		{
			router.Post(ctx, path, routeConf, router.processStreamServiceRequest(ctx, respHandler, method, router.name, svc, service, routeParamIndices, routeParamValues, headerParams))
		}
	}
	return nil
}

func createRouteParamIndices(ctx core.ServerContext, paramsConf config.Config) (map[string]int, error) {
	if paramsConf != nil {
		retval := make(map[string]int, 5)
		paramNames := paramsConf.AllConfigurations()
		for _, paramName := range paramNames {
			indexStr, _ := paramsConf.GetString(paramName)
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return nil, err
			} else {
				retval[paramName] = index
			}
		}
		return retval, nil
	}
	return nil, nil
}
