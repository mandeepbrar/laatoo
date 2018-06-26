package cacheadapter

import (
	"laatoo/sdk/server/log"
	//"laatoo/sdk/server/components"
	"laatoo/sdk/server/components"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	//"laatoo/sdk/server/errors"
	//"laatoo/sdk/server/log"
)

const (
	CONF_CACHEADAPTER_SERVICES  = "cacheadapter"
	CONF_SVC_CACHINGSERVICE     = "cache"
	CONF_SVC_SERVICETOCACHE     = "service"
	CONF_SVC_CACHEBUCKET        = "bucket"
	CONF_SVC_ENTITYCACHEMETHOD  = "entitycache"
	CONF_SVC_RESULTSCACHEMETHOD = "resultscache"
	/*CONF_SVC_CACHEADAPTER               = "cacheadapter"
	CONF_CACHED_VAL_PARAMS              = "params"
	CONF_CACHED_VAL_PARAMSMODE          = "mode"
	CONF_CACHED_VAL_PARAMSADDTOREQBODY  = "addtobody"
	CONF_CACHED_VAL_PARAMSPOSTTOREQBODY = "setbody"
	CONF_CACHED_VAL_PARAMSADDTOQUERY    = "addtoquery"
	CONF_CACHED_VALS                    = "cacheunits"
	CONF_CACHED_VAL                     = "cacheunit"*/
)

type ParamsMode int

const (
	ADDTOBODY ParamsMode = iota
	SETBODY
	ADDTOQUERY
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_CACHEADAPTER_SERVICES, Object: CacheAdapterFactory{}}}
}

type CacheAdapterFactory struct {
}

func (es *CacheAdapterFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (es *CacheAdapterFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	cs := &cacheService{name: name}
	switch method {
	case CONF_SVC_RESULTSCACHEMETHOD:
		{
			cs.svcFunc = cs.RESULTS
			break
		}
	case CONF_SVC_ENTITYCACHEMETHOD:
		{
			break
		}
	}
	return cs, nil
}

//The services start serving when this method is called
func (es *CacheAdapterFactory) Start(ctx core.ServerContext) error {
	return nil
}

type cacheService struct {
	name         string
	bucket       string
	cacheSvcName string
	cacheSvc     components.CacheComponent
	dataSvcName  string
	dataSvc      core.Service
	conf         config.Config
	svcFunc      core.ServiceFunc
}

func (cs *cacheService) Initialize(ctx core.ServerContext, conf config.Config) error {
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
	bucket, ok := conf.GetString(CONF_SVC_CACHEBUCKET)
	if !ok {
		return errors.MissingConf(ctx, CONF_SVC_CACHEBUCKET)
	}
	cs.bucket = bucket
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
		}*/
	return nil
}

func (cs *cacheService) Start(ctx core.ServerContext) error {
	var ok bool
	svc, err := ctx.GetService(cs.cacheSvcName)
	if err != nil {
		return errors.BadConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	cs.cacheSvc, ok = svc.(components.CacheComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	svc, err = ctx.GetService(cs.dataSvcName)
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
	}*/
	return nil
}

func (cs *cacheService) Invoke(ctx core.RequestContext) error {
	return cs.svcFunc(ctx)
}

func (cs *cacheService) SELECT(ctx core.RequestContext) error {
	//	var err error
	var retResponse core.Response
	//	var argsMap map[string]interface{}
	body := ctx.GetRequest()
	cacheKey := components.GetCacheKey(cs.bucket, body)
	log.Trace(ctx, "Looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retResponse)
	if prs {
		log.Trace(ctx, "Cache Hit ")
		ctx.SetResponse(&retResponse)
		return nil
	}
	log.Trace(ctx, "Cache Miss ")
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
	var retResponse core.Response
	var argsMap map[string]interface{}
	body := ctx.GetRequest()
	cacheKey := components.GetCacheKey(cachedVal, body)
	log.Trace(ctx, "Looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retResponse)
	if prs {
		log.Trace(ctx, "Cache Hit ")
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
	var retResponse core.Response
	var argsMap map[string]interface{}
	body := ctx.GetRequest()
	cacheKey := components.GetCacheKey(cachedVal, body)
	log.Trace(ctx, "Looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retResponse)
	if prs {
		log.Trace(ctx, "Cache Hit ")
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
		log.Error(ctx, err.Error())
	}
	return nil
}
*/
