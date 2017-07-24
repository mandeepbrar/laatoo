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
	gi.SetDescription("Update multiple objects using data component. Input a string map containing 'ids' as well as 'data' containing string map of field value updates")
	gi.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	return nil
}

func (es *updatemultiple) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("UPDATEMULTIPLE")
	body := req.GetBody().(*map[string]interface{})
	vals := *body
	ids, ok := vals["ids"]
	if !ok {
		log.Error(ctx, "Missing argument", "Name", "ids")
		return core.StatusBadRequestResponse, nil
	}
	idsArr, ok := ids.([]interface{})
	if !ok {
		log.Error(ctx, "Bad argument", "Name", "ids")
		return core.StatusBadRequestResponse, nil
	}
	stringIds := make([]string, len(idsArr))
	for i, val := range idsArr {
		stringIds[i], ok = val.(string)
		if !ok {
			log.Error(ctx, "Bad argument")
			return core.StatusBadRequestResponse, nil
		}
	}
	data, ok := vals["data"]
	if !ok {
		log.Error(ctx, "Missing argument", "Name", "data")
		return core.StatusBadRequestResponse, nil
	}
	updatesMap, ok := data.(map[string]interface{})
	if !ok {
		log.Error(ctx, "Bad argument", "Name", "data")
		return core.StatusBadRequestResponse, nil
	}
	err := es.DataStore.UpdateMulti(ctx, stringIds, updatesMap)
	if err != nil {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
	return nil, nil
}
