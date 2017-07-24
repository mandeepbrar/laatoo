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

func (gi *selectSvc) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Get multiple element by criteria. Criteria is specified in stringmap.")
	gi.AddStringParams([]string{CONF_FIELD_ORDERBY}, nil)
	gi.AddParams(map[string]string{data.DATA_PAGESIZE: config.CONF_OBJECT_INT, data.DATA_PAGENUM: config.CONF_OBJECT_INT})
	gi.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *selectSvc) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("SELECT")
	retdata, _, totalrecs, recsreturned, err := selectMethod(ctx, req, es.DataStore)
	if err == nil {
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		return core.NewServiceResponse(core.StatusSuccess, retdata, requestinfo), nil
	} else {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
}
