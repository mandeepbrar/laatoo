package dataadapter

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type getById struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *getById) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Get element by Id from the underlying data component")
	gi.AddStringParam(ctx, CONF_DATA_ID)
	return nil
}
func (svc *getById) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}

func (es *getById) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("GETBYID")
	id, ok := ctx.GetStringParam(CONF_DATA_ID)
	if !ok {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.BadArg(ctx, CONF_DATA_ID)
	}
	result, err := es.DataStore.GetById(ctx, id)
	if err == nil {
		if result == nil {
			ctx.SetResponse(core.StatusNotFoundResponse)
			return nil
		} else {
			ctx.SetResponse(core.SuccessResponse(result))
			return nil
		}
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
