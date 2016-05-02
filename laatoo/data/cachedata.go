package data

/*
import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/services"
	"reflect"
)

const (
	CONF_CACHED_SETS                = "cachedsets"
	CONF_CACHED_SET                 = "cachedset"
	CONF_CACHED_SET_PARAMS          = "params"
	CONF_CACHED_SET_LOOKUP          = "lookup"
	CONF_SVC_LOOKUPSVC              = "lookupsvc"
	CONF_SVC_REMOTE_LOOKUP_FIELD    = "remote"
	CONF_SVC_LOCAL_LOOKUP_FIELD     = "local"
	CONF_LOOKUP_RETURN              = "returndata"
	CONF_LOOKUP_RETURN_ONLYLOOKEDUP = "lookedup"
	CONF_LOOKUP_RETURN_ALL          = "all"
)

type cachedSet struct {
	paramMap     map[string]interface{}
	local_lookup string
	//remote_lookup string
	remoteSvcName    string
	remoteSvc        data.DataService
	onlylookedupdata bool
	lookup           bool
}

type cachedDataService struct {
	dataSvc    data.DataService
	cachedSets map[string]*cachedSet
	fac        *DataAdapterFactory
}

func (cs *cachedDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	cs.cachedSets = make(map[string]*cachedSet, 5)
	cachedsetsConf, ok := conf.GetSubConfig(CONF_CACHED_SETS)
	if ok {
		cachedsetNames := cachedsetsConf.AllConfigurations()
		for _, name := range cachedsetNames {
			cSet := &cachedSet{paramMap: make(map[string]interface{}, 5)}
			setConf, _ := cachedsetsConf.GetSubConfig(name)
			paramsConf, ok := setConf.GetSubConfig(CONF_CACHED_SET_PARAMS)
			paramNames := paramsConf.AllConfigurations()
			for _, paramname := range paramNames {
				cSet.paramMap[paramname], _ = paramsConf.Get(paramname)
			}
			lookupConf, ok := setConf.GetSubConfig(CONF_CACHED_SET_LOOKUP)
			if ok {
				remoteSvcName, _ := lookupConf.GetString(CONF_SVC_LOOKUPSVC)
				//				remote_lookup, _ := lookupConf.GetString(CONF_SVC_REMOTE_LOOKUP_FIELD)
				local_lookup, _ := lookupConf.GetString(CONF_SVC_LOCAL_LOOKUP_FIELD)
				datareturn, ok := lookupConf.GetString(CONF_LOOKUP_RETURN)
				if !ok {
					if datareturn == CONF_LOOKUP_RETURN_ALL {
						cset.onlylookedupdata = false
					} else {
						cset.onlylookedupdata = true
					}
				} else {
					cset.onlylookedupdata = true
				}
				cSet.remoteSvcName = remoteSvcName
				//cSet.remote_lookup = remote_lookup
				cSet.local_lookup = local_lookup
				cSet.lookup = true
			}
			cs.cachedSets[name] = cSet
		}
	}
	return nil
}

func (cs *cachedDataService) Start(ctx core.ServerContext) error {
	cs.dataSvc = cs.fac.DataStore
	for _, cset := range cs.cachedSets {
		if cset.lookup {
			remoteSvcInt, err := ctx.GetService(cset.remoteSvcName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			remoteSvc, ok := remoteSvcInt.(data.DataService)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
			}
			cset.remoteSvc = remoteSvc
		}
	}
	return nil
}

func (cs *cachedDataService) Invoke(ctx core.RequestContext) error {
	cachedSet, ok := ctx.GetString(CONF_CACHED_SET)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "Arg", CONF_CACHED_SET)
	}
	var err error
	var retdata interface{}
	var argsMap map[string]interface{}
	body := ctx.GetRequest().(*map[string]interface{})
	argsMap = *body
	cacheKey := services.GetCacheKey(cachedSet, argsMap)
	log.Logger.Trace(ctx, "looking up key", "key", cacheKey)
	prs := ctx.GetFromCache(cacheKey, &retdata)
	if prs {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retdata, make(map[string]interface{}, 0)))
		return nil
	}
	cset, ok := cs.cachedSets[cachedSet]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Arg", CONF_CACHED_SET)
	}
	for k, v := range cset.paramMap {
		argsMap[k] = v
	}
	condition, err := cs.dataSvc.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	data, _, _, err := cs.dataSvc.Get(ctx, condition, -1, -1, "", "")
	if err == nil {
		var retdata interface{}
		if cset.lookup {
			remoteIds := make([]string, len(data))
			for ind, item := range data {
				entVal := reflect.ValueOf(item).Elem()
				f := entVal.FieldByName(cset.local_lookup)
				if !f.IsValid() {
					return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
				}
				remoteIds[ind] = f.String()
			}
			lookedupData, err := cset.remoteSvc.GetMulti(ctx, remoteIds, "")
			if err != nil {
				return err
			}
			retdata = lookedupData
		} else {
			retdata = data
		}
		err := ctx.PutInCache(cacheKey, retdata)
		if err != nil {
			log.Logger.Error(ctx, err.Error())
		}
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retdata, make(map[string]interface{}, 0)))
	}
	return err
}
*/
