package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"strings"
)

type getMulti struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *getMulti) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Get multiple element by Ids from the underlying data component. Ids are separated by comma")
	gi.AddStringParams(ctx, []string{CONF_DATA_IDS, CONF_FIELD_ORDERBY}, nil)
	return nil
}

func (svc *getMulti) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	return nil
}

func (es *getMulti) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("GETMULTI")
	idsstr, ok := ctx.GetStringParam(CONF_DATA_IDS)
	if !ok {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.BadArg(ctx, CONF_DATA_IDS)
	}
	ids := strings.Split(idsstr, ",")
	orderBy, ok := ctx.GetStringParam(CONF_FIELD_ORDERBY)
	var orderByCond interface{}
	var err error
	if ok {
		orderByCond, err = es.DataStore.CreateCondition(ctx, data.SORTASC, orderBy)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	result, err := es.DataStore.GetMulti(ctx, ids, orderByCond)
	if err == nil {
		log.Trace(ctx, "Returning results ", "result", result)
		ctx.SetResponse(core.SuccessResponse(result))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
