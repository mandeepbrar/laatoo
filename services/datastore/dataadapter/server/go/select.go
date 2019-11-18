package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type selectSvc struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *selectSvc) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Get multiple element by criteria. Criteria is specified in stringmap.")
	gi.AddStringParams(ctx, []string{CONF_FIELD_ORDERBY}, []string{""})
	return gi.AddParams(ctx, map[string]string{"argsMap": config.OBJECTTYPE_STRINGMAP, data.DATA_PAGESIZE: config.OBJECTTYPE_INT, data.DATA_PAGENUM: config.OBJECTTYPE_INT}, false)
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
		ctx.SetResponse(core.SuccessResponseWithInfo(retdata, requestinfo))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
