package dataadapter

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type updatemultiple struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *updatemultiple) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Update multiple objects using data component. Input a string map containing 'ids' as well as 'data' containing string map of field value updates")
	return gi.AddParamWithType(ctx, "argsMap", config.OBJECTTYPE_STRINGMAP)
	//	gi.SetRequestType(ctx, config.OBJECTTYPE_STRINGMAP, false, false)
}
func (svc *updatemultiple) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}
func (es *updatemultiple) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATEMULTIPLE")
	vals, _ := ctx.GetStringMapParam("argsMap")
	ids, ok := vals["ids"]
	if !ok {
		log.Error(ctx, "Missing argument", "Name", "ids")
		ctx.SetResponse(core.BadRequestResponse("Missing ids in args map"))
		return nil
	}
	idsArr, ok := ids.([]interface{})
	if !ok {
		log.Error(ctx, "Bad argument", "Name", "ids")
		ctx.SetResponse(core.BadRequestResponse("Ids should be an array of strings"))
		return nil
	}
	stringIds := make([]string, len(idsArr))
	for i, val := range idsArr {
		stringIds[i], ok = val.(string)
		if !ok {
			log.Error(ctx, "Bad argument")
			ctx.SetResponse(core.BadRequestResponse("Ids should be an array of strings"))
			return nil
		}
	}
	data, ok := vals["data"]
	if !ok {
		log.Error(ctx, "Missing argument", "Name", "data")
		ctx.SetResponse(core.BadRequestResponse("Missing Argument data in args map"))
		return nil
	}
	updatesMap, ok := data.(map[string]interface{})
	if !ok {
		log.Error(ctx, "Bad argument", "Name", "data")
		ctx.SetResponse(core.BadRequestResponse("Argument data should be map"))
		return nil
	}
	err := es.DataStore.UpdateMulti(ctx, stringIds, updatesMap)
	if err != nil {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
	return nil
}
