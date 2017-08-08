package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type save struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *save) Initialize(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Saves a storable using data component.")
	gi.SetRequestType(ctx, gi.DataStore.GetObject(), false, false)
	return nil
}

func (es *save) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("SAVE")
	ent := ctx.GetBody()
	stor := ent.(data.Storable)
	err := es.DataStore.Save(ctx, stor)
	if err == nil {
		ctx.Set("Id", stor.GetId())
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, stor.GetId(), nil))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
