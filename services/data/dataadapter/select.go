package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type selectSvc struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *selectSvc) Describe(ctx core.ServerContext) {
	gi.SetDescription(ctx, "Get multiple element by criteria. Criteria is specified in stringmap.")
	gi.AddStringParams(ctx, []string{CONF_FIELD_ORDERBY}, []string{""})
	gi.AddParams(ctx, map[string]string{data.DATA_PAGESIZE: config.OBJECTTYPE_INT, data.DATA_PAGENUM: config.OBJECTTYPE_INT}, false)
	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}
func (svc *selectSvc) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}
func (es *selectSvc) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("SELECT")
	retdata, _, totalrecs, recsreturned, err := selectMethod(ctx, es.DataStore)
	if err == nil {
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retdata, requestinfo))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
