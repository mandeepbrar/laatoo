package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type save struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (svc *save) Describe(ctx core.ServerContext) error {
	svc.SetDescription(ctx, "Saves a storable using data component.")
	return nil
}
func (svc *save) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	//	svc.SetRequestType(ctx, svc.DataStore.GetObject(), false, false)
	/****TODO test*****/
	return svc.AddParamWithType(ctx, "object", svc.DataStore.GetObject())
}
func (es *save) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("SAVE")
	ent, _ := ctx.GetParamValue("object")
	stor := ent.(data.Storable)
	err := es.DataStore.Save(ctx, stor)
	if err == nil {
		ctx.Set("Id", stor.GetId())
		ctx.SetResponse(core.SuccessResponse(stor.GetId()))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
