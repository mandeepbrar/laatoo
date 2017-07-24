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

func (gi *join) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Get multiple element by Ids from the underlying data component. Ids are separated by comma")
	gi.AddStringConfigurations([]string{CONF_SVC_LOOKUP_FIELD, CONF_SVC_LOOKUPSVC}, nil)
	gi.AddStringParams([]string{CONF_DATA_IDS, CONF_FIELD_ORDERBY}, nil)
	gi.AddParams(map[string]string{data.DATA_PAGESIZE: config.CONF_OBJECT_INT, data.DATA_PAGENUM: config.CONF_OBJECT_INT})
	gi.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *join) Start(ctx core.ServerContext) error {
	es.lookupField, _ = es.GetStringConfiguration(CONF_SVC_LOOKUP_FIELD)
	es.lookupSvcName, _ = es.GetStringConfiguration(CONF_SVC_LOOKUPSVC)

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

func (es *join) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("JOIN")
	retdata, _, totalrecs, recsreturned, err := selectMethod(ctx, req, es.DataStore)
	lookupids := make([]string, len(retdata))
	for ind, item := range retdata {
		entVal := reflect.ValueOf(item).Elem()
		f := entVal.FieldByName(es.lookupField)
		if !f.IsValid() {
			return core.StatusNotFoundResponse, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
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
		return core.NewServiceResponse(core.StatusSuccess, result, requestinfo), nil
	} else {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
}
