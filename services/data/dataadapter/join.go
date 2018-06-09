package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"reflect"
)

type join struct {
	core.Service
	fac           *DataAdapterFactory
	DataStore     data.DataComponent
	lookupSvcName string
	lookupSvc     data.DataComponent
	lookupField   string
}

func (gi *join) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Get multiple element by Ids from the underlying data component. Ids are separated by comma")
	gi.AddStringConfigurations(ctx, []string{CONF_SVC_LOOKUP_FIELD, CONF_SVC_LOOKUPSVC}, nil)
	gi.AddStringParams(ctx, []string{CONF_DATA_IDS, CONF_FIELD_ORDERBY}, nil)
	return gi.AddParams(ctx, map[string]string{"argsMap": config.OBJECTTYPE_STRINGMAP, data.DATA_PAGESIZE: config.OBJECTTYPE_INT, data.DATA_PAGENUM: config.OBJECTTYPE_INT}, false)
	//	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
	//return nil
}

func (es *join) Start(ctx core.ServerContext) error {
	es.DataStore = es.fac.DataStore
	es.lookupField, _ = es.GetStringConfiguration(ctx, CONF_SVC_LOOKUP_FIELD)
	es.lookupSvcName, _ = es.GetStringConfiguration(ctx, CONF_SVC_LOOKUPSVC)

	lookupSvcInt, err := ctx.GetService(es.lookupSvcName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	lookupSvc, ok := lookupSvcInt.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
	}
	es.lookupSvc = lookupSvc
	return nil
}

func (es *join) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("JOIN")
	retdata, _, totalrecs, recsreturned, err := selectMethod(ctx, es.DataStore)
	lookupids := make([]string, len(retdata))
	for ind, item := range retdata {
		entVal := reflect.ValueOf(item).Elem()
		f := entVal.FieldByName(es.lookupField)
		if !f.IsValid() {
			ctx.SetResponse(core.StatusNotFoundResponse)
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
		}
		lookupids[ind] = f.String()
	}
	log.Trace(ctx, "JOIN: Looking up ids", "ids", lookupids)
	result, err := es.lookupSvc.GetMultiHash(ctx, lookupids)
	if err == nil {
		for i, id := range lookupids {
			stor := retdata[i]
			item, ok := result[id]
			if ok {
				stor.Join(item)
			}
		}
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.SuccessResponseWithInfo(result, requestinfo))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
