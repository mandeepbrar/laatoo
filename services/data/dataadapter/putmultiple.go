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

func (gi *putmultiple) Describe(ctx core.ServerContext) {
	gi.SetDescription(ctx, "Put multiple objects using data component. Input an array of objects")
	gi.SetRequestType(ctx, gi.DataStore.GetObject(), true, false)
}

func (svc *putmultiple) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}
func (es *putmultiple) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("PUTMULTIPLE")
	arr := ctx.GetBody()
	log.Trace(ctx, "Collection ", "arr", arr)
	storables, _, err := data.CastToStorableCollection(arr)
	if err != nil {
		ctx.SetResponse(core.StatusInternalErrorResponse)
		return err
	}
	err = es.DataStore.PutMulti(ctx, storables)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
