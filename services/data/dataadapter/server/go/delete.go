package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type deleteSvc struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *deleteSvc) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Delete an entity represented by id")
	gi.AddStringParam(ctx, CONF_DATA_ID)
	return nil
}

func (svc *deleteSvc) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}

func (es *deleteSvc) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("DELETE")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	err := es.DataStore.Delete(ctx, id)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
