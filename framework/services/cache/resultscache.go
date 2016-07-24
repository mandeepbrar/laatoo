package cache

import (
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_SVC_RESULTSCACHEMETHOD = "resultscache"
)

type resultsCacheService struct {
	name         string
	bucket       string
	cacheSvcName string
	cacheSvc     components.CacheComponent
	dataSvcName  string
	dataSvc      core.Service
	conf         config.Config
}

func (rs *resultsCacheService) Initialize(ctx core.ServerContext, conf config.Config) error {
	svcName, ok := conf.GetString(CONF_SVC_CACHINGSERVICE)
	if !ok {
		return errors.MissingConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	rs.cacheSvcName = svcName
	svcName, ok = conf.GetString(CONF_SVC_SERVICETOCACHE)
	if !ok {
		return errors.MissingConf(ctx, CONF_SVC_SERVICETOCACHE)
	}
	rs.dataSvcName = svcName
	bucket, ok := conf.GetString(CONF_SVC_CACHEBUCKET)
	if !ok {
		return errors.MissingConf(ctx, CONF_SVC_CACHEBUCKET)
	}
	rs.bucket = bucket
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

func (rs *resultsCacheService) Start(ctx core.ServerContext) error {
	var ok bool
	svc, err := ctx.GetService(rs.cacheSvcName)
	if err != nil {
		return errors.BadConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	rs.cacheSvc, ok = svc.(components.CacheComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_SVC_CACHINGSERVICE)
	}
	svc, err = ctx.GetService(rs.dataSvcName)
	if err != nil {
		return errors.BadConf(ctx, CONF_SVC_SERVICETOCACHE)
	}
	rs.dataSvc = svc
	/*	for _, cVal := range cs.cachedVals {
		svc, err := ctx.GetService(cVal.serviceName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		cVal.service = svc
	}*/
	return nil
}
func (rs *resultsCacheService) Invoke(ctx core.RequestContext) error {
	//	var err error
	var retResponse core.ServiceResponse
	//	var argsMap map[string]interface{}
	body := ctx.GetRequest()
	cacheKey := components.GetCacheKey(rs.bucket, body)
	log.Logger.Trace(ctx, "Looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retResponse)
	if prs {
		log.Logger.Trace(ctx, "Cache Hit ")
		ctx.SetResponse(&retResponse)
		return nil
	}
	log.Logger.Trace(ctx, "Cache Miss ")
	return nil
}

/*
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
