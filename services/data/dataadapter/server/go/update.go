package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
)

type update struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *update) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Update object using underlying data component. Expects an entity id. Value should be map containing field values")
	gi.AddStringParam(ctx, CONF_DATA_ID)
	return gi.AddParamWithType(ctx, "argsMap", config.OBJECTTYPE_STRINGMAP)
}
func (svc *update) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}
func (es *update) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATE")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	vals, _ := ctx.GetStringMapParam("argsMap")
	//vals := *body
	res, err := updateVals(ctx, id, vals, es.DataStore)
	ctx.SetResponse(res)
	return err
}
