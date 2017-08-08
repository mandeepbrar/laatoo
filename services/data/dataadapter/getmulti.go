package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"strings"
)

type getMulti struct {
	core.Service
	fac       *DataAdapterFactory
	DataStore data.DataComponent
}

func (gi *getMulti) Initialize(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Get multiple element by Ids from the underlying data component. Ids are separated by comma")
	gi.AddStringParams(ctx, []string{CONF_DATA_IDS, CONF_FIELD_ORDERBY}, nil)
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
	orderBy, _ := ctx.GetStringParam(CONF_FIELD_ORDERBY)
	result, err := es.DataStore.GetMulti(ctx, ids, orderBy)
	if err == nil {
		log.Trace(ctx, "Returning results ", "result", result)
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, nil))
		return nil
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return errors.WrapError(ctx, err)
	}
}
