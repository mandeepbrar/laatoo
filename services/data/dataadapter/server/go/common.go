package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

func selectMethod(ctx core.RequestContext, datastore data.DataComponent) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	pagesize, _ := ctx.GetIntParam(data.DATA_PAGESIZE)
	pagenum, _ := ctx.GetIntParam(data.DATA_PAGENUM)
	orderBy, _ := ctx.GetStringParam(CONF_FIELD_ORDERBY)
	argsMap, _ := ctx.GetStringMapValue("argsMap")
	log.Trace(ctx, "select", "argsMap", argsMap, "pagesize", pagesize, "pagenum", pagenum)
	condition, err := datastore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return nil, nil, -1, -1, errors.WrapError(ctx, err)
	}
	return datastore.Get(ctx, condition, pagesize, pagenum, "", orderBy)
}

func updateVals(ctx core.RequestContext, id string, vals map[string]interface{}, datastore data.DataComponent) (*core.Response, error) {
	err := datastore.Update(ctx, id, vals)
	if err != nil {
		return core.StatusInternalErrorResponse, errors.WrapError(ctx, err)
	}
	return nil, nil
}
