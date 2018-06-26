package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type putmultiple struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *putmultiple) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Put multiple objects using data component. Input an array of objects")
	//	gi.SetRequestType(ctx, gi.DataStore.GetObject(), true, false)
	return nil
}

func (svc *putmultiple) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	/****************TODO******/
	return svc.AddParamWithType(ctx, "object", svc.DataStore.GetObject())
}
func (es *putmultiple) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("PUTMULTIPLE")
	arr, _ := ctx.GetParamValue("object")
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
