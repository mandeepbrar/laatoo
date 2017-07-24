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

func (gi *put) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Put a storable using data component. Takes object id as the parameter")
	gi.SetRequestType(gi.DataStore.GetObject(), false, false)
	gi.AddStringParam(CONF_DATA_ID)
	return nil
}

func (es *put) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("PUT")
	id, _ := req.GetStringParam(CONF_DATA_ID)
	ent := req.GetBody()
	stor := ent.(data.Storable)
	err := es.DataStore.Put(ctx, id, stor)
	if err != nil {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
	return nil, nil
}
