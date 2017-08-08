package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"reflect"
)

type getmulti_select struct {
	core.Service
	fac           *DataAdapterFactory
	DataStore     data.DataComponent
	hashmap       bool
	lookupSvcName string
	lookupSvc     data.DataComponent
	lookupField   string
}

const (
	HASHMAP_PARAM = "hashmap"
)

func (gi *getmulti_select) Initialize(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Get multiple element by criteria from the underlying data component. Criteria passed in the body")
	gi.AddStringConfigurations(ctx, []string{CONF_SVC_LOOKUP_FIELD, CONF_SVC_LOOKUPSVC}, nil)
	gi.AddOptionalConfigurations(ctx, map[string]string{HASHMAP_PARAM: config.CONF_OBJECT_BOOL}, map[string]interface{}{HASHMAP_PARAM: false})
	gi.AddStringParams(ctx, []string{CONF_FIELD_ORDERBY}, nil)
	gi.AddParams(ctx, map[string]string{data.DATA_PAGESIZE: config.CONF_OBJECT_INT, data.DATA_PAGENUM: config.CONF_OBJECT_INT})
	gi.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *getmulti_select) Start(ctx core.ServerContext) error {
	es.lookupField, _ = es.GetStringConfiguration(ctx, CONF_SVC_LOOKUP_FIELD)
	es.lookupSvcName, _ = es.GetStringConfiguration(ctx, CONF_SVC_LOOKUPSVC)
	es.hashmap, _ = es.GetBoolConfiguration(ctx, HASHMAP_PARAM)
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

func (es *getmulti_select) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("GETMULTI_SELECTIDS")
	retdata, _, totalrecs, recsreturned, err := selectMethod(ctx, es.lookupSvc)
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
	log.Trace(ctx, "GETMULTI_SELECTIDS: Looking up ids", "ids", lookupids)
	var result interface{}
	if es.hashmap {
		result, err = es.DataStore.GetMultiHash(ctx, lookupids)
	} else {
		result, err = es.DataStore.GetMulti(ctx, lookupids, "")
	}
	if err == nil {
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, requestinfo))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
