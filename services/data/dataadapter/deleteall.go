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

func (gi *deleteAll) Initialize(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Delete all objects specified by criteria. Criteria should be map containing field values")
	gi.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *deleteAll) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("DELETEALL")
	body := ctx.GetBody().(*map[string]interface{})
	argsMap := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	retval, err := es.DataStore.DeleteAll(ctx, condition)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retval, nil))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
