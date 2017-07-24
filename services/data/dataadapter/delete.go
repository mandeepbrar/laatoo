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

func (gi *deleteSvc) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Delete an entity represented by id")
	gi.AddStringParam(CONF_DATA_ID)
	return nil
}

func (es *deleteSvc) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("DELETE")
	id, _ := req.GetStringParam(CONF_DATA_ID)
	err := es.DataStore.Delete(ctx, id)
	if err != nil {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
	return nil, nil
}
