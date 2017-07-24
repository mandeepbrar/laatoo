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
	gi.SetDescription("Saves a storable using data component.")
	gi.SetRequestType(gi.DataStore.GetObject(), false, false)
	return nil
}

func (es *save) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("SAVE")
	ent := req.GetBody()
	stor := ent.(data.Storable)
	err := es.DataStore.Save(ctx, stor)
	if err == nil {
		ctx.Set("Id", stor.GetId())
		return core.NewServiceResponse(core.StatusSuccess, stor.GetId(), nil), nil
	} else {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
}
