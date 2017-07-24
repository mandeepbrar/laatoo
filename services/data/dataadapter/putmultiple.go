package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type putmultiple struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *putmultiple) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Put multiple objects using data component. Input an array of objects")
	gi.SetRequestType(gi.DataStore.GetObject(), true, false)
	return nil
}

func (es *putmultiple) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("PUTMULTIPLE")
	arr := req.GetBody()
	log.Trace(ctx, "Collection ", "arr", arr)
	storables, _, err := data.CastToStorableCollection(arr)
	if err != nil {
		return core.StatusInternalErrorResponse, err
	}
	err = es.DataStore.PutMulti(ctx, storables)
	if err != nil {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
	return nil, nil
}
