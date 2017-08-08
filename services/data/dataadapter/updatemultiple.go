package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type updatemultiple struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *updatemultiple) Initialize(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Update multiple objects using data component. Input a string map containing 'ids' as well as 'data' containing string map of field value updates")
	gi.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *updatemultiple) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATEMULTIPLE")
	body := ctx.GetBody().(*map[string]interface{})
	vals := *body
	ids, ok := vals["ids"]
	if !ok {
		log.Error(ctx, "Missing argument", "Name", "ids")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	idsArr, ok := ids.([]interface{})
	if !ok {
		log.Error(ctx, "Bad argument", "Name", "ids")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	stringIds := make([]string, len(idsArr))
	for i, val := range idsArr {
		stringIds[i], ok = val.(string)
		if !ok {
			log.Error(ctx, "Bad argument")
			ctx.SetResponse(core.StatusBadRequestResponse)
			return nil
		}
	}
	data, ok := vals["data"]
	if !ok {
		log.Error(ctx, "Missing argument", "Name", "data")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	updatesMap, ok := data.(map[string]interface{})
	if !ok {
		log.Error(ctx, "Bad argument", "Name", "data")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	err := es.DataStore.UpdateMulti(ctx, stringIds, updatesMap)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
