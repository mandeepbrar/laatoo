package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type put struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *put) Describe(ctx core.ServerContext) {
	gi.SetDescription(ctx, "Put a storable using data component. Takes object id as the parameter")
	gi.SetRequestType(ctx, gi.DataStore.GetObject(), false, false)
	gi.AddStringParam(ctx, CONF_DATA_ID)
}
func (svc *put) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}

func (es *put) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("PUT")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	ent := ctx.GetBody()
	stor := ent.(data.Storable)
	err := es.DataStore.Put(ctx, id, stor)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
