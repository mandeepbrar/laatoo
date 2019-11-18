package dataadapter

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type count struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *count) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Count objects meeting selection criteria")
	return gi.AddParamWithType(ctx, "argsMap", config.OBJECTTYPE_STRINGMAP)
	//	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}
func (svc *count) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}
func (es *count) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("COUNT")
	argsMap, _ := ctx.GetStringMapParam("argsMap")
	//	body := ctx.GetBody().(*map[string]interface{})
	//	argsMap := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	count, err := es.DataStore.Count(ctx, condition)
	if err == nil {
		ctx.SetResponse(core.SuccessResponse(count))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
