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

func (gi *upsert) Describe(ctx core.ServerContext) {
	gi.SetDescription(ctx, "Upsert object using underlying data component. Expects a map containing condition and value. Value should be map containing field values")
	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}

func (es *upsert) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPSERT")
	body := ctx.GetBody().(*map[string]interface{})
	vals := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, vals["condition"].(map[string]interface{}))
	if err != nil {
		ctx.SetResponse(core.StatusBadRequestResponse)
		return errors.WrapError(ctx, err)
	}
	_, err = es.DataStore.Upsert(ctx, condition, vals["value"].(map[string]interface{}))
	if err != nil {
		ctx.SetResponse(core.StatusInternalErrorResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
