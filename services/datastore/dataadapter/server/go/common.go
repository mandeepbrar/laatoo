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

	var orderByCond interface{}
	orderBy, ok := ctx.GetStringParam(CONF_FIELD_ORDERBY)
	if ok {
		if orderBy != "" {
			orderByCond, err = datastore.CreateCondition(ctx, data.SORTASC, orderBy)
			if err != nil {
				return nil, nil, -1, -1, errors.WrapError(ctx, err)
			}
		}
	}

	argsMap, _ := ctx.GetStringMapParam("argsMap")
	log.Trace(ctx, "select", "argsMap", argsMap, "pagesize", pagesize, "pagenum", pagenum)
	condition, err := datastore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return nil, nil, -1, -1, errors.WrapError(ctx, err)
	}
	return datastore.Get(ctx, condition, pagesize, pagenum, "", orderByCond)
}

func updateVals(ctx core.RequestContext, id string, vals map[string]interface{}, datastore data.DataComponent) (*core.Response, error) {
	err := datastore.Update(ctx, id, vals)
	if err != nil {
		return core.InternalErrorResponse("Could not update values"), errors.WrapError(ctx, err)
	}
	return nil, nil
}