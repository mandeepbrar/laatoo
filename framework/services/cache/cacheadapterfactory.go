package cache

import (
	"laatoo/framework/core/objects"
	//"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	//"laatoo/sdk/log"
)

const (
	CONF_CACHEADAPTER_SERVICES = "cacheadapter"
	CONF_SVC_CACHINGSERVICE    = "cache"
	CONF_SVC_SERVICETOCACHE    = "service"
	CONF_SVC_CACHEBUCKET       = "bucket"
	/*CONF_SVC_CACHEADAPTER               = "cacheadapter"
	CONF_CACHED_VAL_PARAMS              = "params"
	CONF_CACHED_VAL_PARAMSMODE          = "mode"
	CONF_CACHED_VAL_PARAMSADDTOREQBODY  = "addtobody"
	CONF_CACHED_VAL_PARAMSPOSTTOREQBODY = "setbody"
	CONF_CACHED_VAL_PARAMSADDTOQUERY    = "addtoquery"
	CONF_CACHED_VALS                    = "cacheunits"
	CONF_CACHED_VAL                     = "cacheunit"*/
)

/*
type ParamsMode int

const (
	ADDTOBODY ParamsMode = iota
	SETBODY
	ADDTOQUERY
)*/

func init() {
	objects.RegisterObject(CONF_CACHEADAPTER_SERVICES, createCacheAdapterFactory, nil)
}

type cacheAdapterFactory struct {
}

func createCacheAdapterFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &cacheAdapterFactory{}, nil
}

func (es *cacheAdapterFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (es *cacheAdapterFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	switch method {
	case CONF_SVC_RESULTSCACHEMETHOD:
		{
			return &resultsCacheService{name: name}, nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (es *cacheAdapterFactory) Start(ctx core.ServerContext) error {
	return nil
}

/*
type cachedVal struct {
	configParams map[string]interface{}
	paramsMode   ParamsMode
	serviceName  string
	service      core.Service
}*/

/*
type cacheAdapterService struct {
	name         string
	cacheSvcName string
	cacheSvc     components.CacheComponent
	dataSvcName  string
	dataSvc      components.CacheComponent
	conf         config.Config
}

func (cs *cacheAdapterService) Initialize(ctx core.ServerContext, conf config.Config) error {
	svcName, ok := conf.GetString(CONF_SVC_CACHINGSERVICE)
	if !ok {
		return errors.MissingConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	cs.cacheSvcName = svcName
	svcName, ok = conf.GetString(CONF_SVC_SERVICETOCACHE)
	if !ok {
		return errors.MissingConf(ctx, CONF_SVC_SERVICETOCACHE)
	}
	cs.dataSvcName = svcName
	/*	cachedvalsConf, ok := conf.GetSubConfig(CONF_CACHED_VALS)
		if ok {
			cachedvalNames := cachedvalsConf.AllConfigurations()
			for _, name := range cachedvalNames {
				cVal := &cachedVal{configParams: make(map[string]interface{}, 5)}
				cvalConf, _ := cachedvalsConf.GetSubConfig(name)
				paramsConf, ok := cvalConf.GetSubConfig(CONF_CACHED_VAL_PARAMS)
				if ok {
					paramNames := paramsConf.AllConfigurations()
					for _, paramname := range paramNames {
						cVal.configParams[paramname], _ = paramsConf.Get(paramname)
					}
					cVal.paramsMode = SETBODY
					paramsMode, ok := cvalConf.GetString(CONF_CACHED_VAL_PARAMSMODE)
					if ok {
						switch paramsMode {
						case "addtobody":
							cVal.paramsMode = ADDTOBODY
						case "addtoquery":
							cVal.paramsMode = ADDTOQUERY
						}
					}
				}
				svcName, ok := cvalConf.GetString(CONF_SVC_SERVICETOCACHE)
				if !ok {
					return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "arg", CONF_SVC_SERVICETOCACHE)
				}
				cVal.serviceName = svcName
				cs.cachedVals[name] = cVal
			}
		}
	return nil
}

func (cs *cacheAdapterService) Start(ctx core.ServerContext) error {
	svc, err := ctx.GetService(cs.cacheSvcName)
	if err != nil {
		return errors.BadConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	cs.cacheSvc = svc
	svc, err := ctx.GetService(cs.dataSvcName)
	if err != nil {
		return errors.BadConf(ctx, CONF_SVC_SERVICETOCACHE)
	}
	cs.dataSvc = svc
	/*	for _, cVal := range cs.cachedVals {
		svc, err := ctx.GetService(cVal.serviceName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		cVal.service = svc
	}
	return nil
}
func (cs *cacheAdapterService) Invoke(ctx core.RequestContext) error {
	var err error
	var retResponse core.ServiceResponse
	var argsMap map[string]interface{}
	body := ctx.GetRequest()
	cacheKey := components.GetCacheKey(cachedVal, body)
	log.Logger.Trace(ctx, "Looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retResponse)
	if prs {
		log.Logger.Trace(ctx, "Cache Hit ")
		ctx.SetResponse(&retResponse)
		return nil
	}
}


func (cs *cacheAdapterService) Invoke(ctx core.RequestContext) error {
	cachedVal, ok := ctx.GetString(CONF_CACHED_VAL)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "Arg", CONF_CACHED_VAL)
	}
	var err error
	var retResponse core.ServiceResponse
	var argsMap map[string]interface{}
	body := ctx.GetRequest()
	cacheKey := components.GetCacheKey(cachedVal, body)
	log.Logger.Trace(ctx, "Looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retResponse)
	if prs {
		log.Logger.Trace(ctx, "Cache Hit ")
		ctx.SetResponse(&retResponse)
		return nil
	}
	cVal, ok := cs.cachedVals[cachedVal]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Arg", CONF_CACHED_VAL)
	}
	switch cVal.paramsMode {
	case SETBODY:
		ctx.SetRequest(&cVal.configParams)
	case ADDTOQUERY:
		for k, v := range cVal.configParams {
			ctx.Set(k, v)
		}
	case ADDTOBODY:
		argsMap = *body.(*map[string]interface{})
		for k, v := range cVal.configParams {
			argsMap[k] = v
		}
		ctx.SetRequest(&argsMap)
	}
	err = cVal.service.Invoke(ctx)
	if err != nil {
		return err
	}
	resp := ctx.GetResponse()
	err = ctx.PutInCache(cacheKey, resp)
	if err != nil {
		log.Logger.Error(ctx, err.Error())
	}
	return nil
}
*/
