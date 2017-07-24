package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type upsert struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *upsert) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Upsert object using underlying data component. Expects a map containing condition and value. Value should be map containing field values")
	gi.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *upsert) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("UPSERT")
	body := req.GetBody().(*map[string]interface{})
	vals := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, vals["condition"].(map[string]interface{}))
	if err != nil {
		return core.StatusBadRequestResponse, errors.WrapError(ctx, err)
	}
	_, err = es.DataStore.Upsert(ctx, condition, vals["value"].(map[string]interface{}))
	if err != nil {
		return core.StatusInternalErrorResponse, errors.WrapError(ctx, err)
	}
	return nil, nil
}
