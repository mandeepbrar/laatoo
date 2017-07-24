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
	gi.SetDescription("Update object using underlying data component. Expects an entity id. Value should be map containing field values")
	gi.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	gi.AddStringParam(CONF_DATA_ID)
	return nil
}

func (es *update) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("UPDATE")
	id, _ := req.GetStringParam(CONF_DATA_ID)
	body := req.GetBody().(*map[string]interface{})
	vals := *body
	return updateVals(ctx, id, vals, es.DataStore)
}
