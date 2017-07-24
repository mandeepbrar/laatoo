package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

func selectMethod(ctx core.RequestContext, req core.Request, datastore data.DataComponent) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	pagesize, _ := req.GetIntParam(data.DATA_PAGESIZE)
	pagenum, _ := req.GetIntParam(data.DATA_PAGENUM)
	orderBy, _ := req.GetStringParam(CONF_FIELD_ORDERBY)
	var argsMap map[string]interface{}
	body := req.GetBody().(*map[string]interface{})
	argsMap = *body
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
