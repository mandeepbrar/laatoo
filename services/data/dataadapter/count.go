package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type count struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *count) Describe(ctx core.ServerContext) {
	gi.SetDescription(ctx, "Count objects meeting selection criteria")
	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}

func (es *count) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("COUNT")
	body := ctx.GetBody().(*map[string]interface{})
	argsMap := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	count, err := es.DataStore.Count(ctx, condition)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, count, nil))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
