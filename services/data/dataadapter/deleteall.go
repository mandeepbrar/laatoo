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
	gi.SetDescription("Delete all objects specified by criteria. Criteria should be map containing field values")
	gi.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *deleteAll) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("DELETEALL")
	body := req.GetBody().(*map[string]interface{})
	argsMap := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
	retval, err := es.DataStore.DeleteAll(ctx, condition)
	if err == nil {
		return core.NewServiceResponse(core.StatusSuccess, retval, nil), nil
	} else {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
}
