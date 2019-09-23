package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type upsert struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (svc *upsert) Describe(ctx core.ServerContext) error {
	svc.SetDescription(ctx, "Upsert object using underlying data component. Expects a map containing condition and value. Value should be map containing field values")
	return svc.AddParamWithType(ctx, "argsMap", config.OBJECTTYPE_STRINGMAP)
	//	svc.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}
func (svc *upsert) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}
func (es *upsert) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPSERT")
	//	body := ctx.GetBody().(*map[string]interface{})
	//	vals := *body
	vals, _ := ctx.GetStringMapParam("argsMap")
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, vals["condition"].(map[string]interface{}))
	if err != nil {
		ctx.SetResponse(core.BadRequestResponse("Could not create condition" + err.Error()))
		return errors.WrapError(ctx, err)
	}
	_, err = es.DataStore.Upsert(ctx, condition, vals["value"].(map[string]interface{}))
	if err != nil {
		ctx.SetResponse(core.InternalErrorResponse("Could not upsert values with condition to datastore"))
		return errors.WrapError(ctx, err)
	}
	return nil
}
