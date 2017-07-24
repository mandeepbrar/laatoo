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
	gi.SetDescription("Get multiple element by Ids from the underlying data component. Ids are separated by comma")
	gi.AddStringParams([]string{CONF_DATA_IDS, CONF_FIELD_ORDERBY}, nil)
	return nil
}

func (es *getMulti) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("GETMULTI")
	idsstr, ok := req.GetStringParam(CONF_DATA_IDS)
	if !ok {
		return core.StatusNotFoundResponse, errors.BadArg(ctx, CONF_DATA_IDS)
	}
	ids := strings.Split(idsstr, ",")
	orderBy, _ := req.GetStringParam(CONF_FIELD_ORDERBY)
	result, err := es.DataStore.GetMulti(ctx, ids, orderBy)
	if err == nil {
		log.Trace(ctx, "Returning results ", "result", result)
		return core.NewServiceResponse(core.StatusSuccess, result, nil), nil
	} else {
		return core.StatusNotFoundResponse, errors.WrapError(ctx, err)
	}
}
