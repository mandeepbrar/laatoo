package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type getById struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *getById) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Get element by Id from the underlying data component")
	gi.AddStringParam(CONF_DATA_ID)
	return nil
}

func (es *getById) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("GETBYID")
	id, ok := req.GetStringParam(CONF_DATA_ID)
	if !ok {
		return core.StatusNotFoundResponse, errors.BadArg(ctx, CONF_DATA_ID)
	}
	result, err := es.DataStore.GetById(ctx, id)
	if err == nil {
		if result == nil {
			return core.StatusNotFoundResponse, nil
		} else {
			return core.NewServiceResponse(core.StatusSuccess, result, nil), nil
		}
	} else {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
}
