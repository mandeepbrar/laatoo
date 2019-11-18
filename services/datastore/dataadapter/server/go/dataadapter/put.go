package dataadapter

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type put struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *put) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Put a storable using data component. Takes object id as the parameter")
	//gi.SetRequestType(ctx, gi.DataStore.GetObject(), false, false)
	gi.AddStringParam(ctx, CONF_DATA_ID)
	return nil
}

func (svc *put) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	/*********TODO test*****************/
	return svc.AddParamWithType(ctx, "object", svc.DataStore.GetObject())
}

func (es *put) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("PUT")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	ent, _ := ctx.GetParamValue("object")
	stor := ent.(data.Storable)
	err := es.DataStore.Put(ctx, id, stor)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
