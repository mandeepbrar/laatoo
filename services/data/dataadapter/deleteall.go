package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type deleteAll struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *deleteAll) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Delete all objects specified by criteria. Criteria should be map containing field values")
	return gi.AddParamWithType(ctx, "argsMap", config.OBJECTTYPE_STRINGMAP)
	//	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}
func (svc *deleteAll) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}

func (es *deleteAll) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("DELETEALL")
	argsMap, _ := ctx.GetStringMapValue("argsMap")
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	retval, err := es.DataStore.DeleteAll(ctx, condition)
	if err == nil {
		ctx.SetResponse(core.SuccessResponse(retval))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
