package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type update struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *update) Initialize(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Update object using underlying data component. Expects an entity id. Value should be map containing field values")
	gi.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	gi.AddStringParam(ctx, CONF_DATA_ID)
	return nil
}

func (es *update) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATE")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	body := ctx.GetBody().(*map[string]interface{})
	vals := *body
	res, err := updateVals(ctx, id, vals, es.DataStore)
	ctx.SetResponse(res)
	return err
}
