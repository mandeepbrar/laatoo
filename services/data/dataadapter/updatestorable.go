package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
)

type updateStorable struct {
	core.Service
	fac          *DataAdapterFactory
	DataStore    data.DataComponent
	updateFields []string
}

func (gi *updateStorable) Initialize(ctx core.ServerContext) error {
	gi.SetDescription("Update object using underlying data component. Expects an entity id. Value should be storable object")
	gi.SetRequestType(gi.DataStore.GetObject(), false, false)
	gi.AddConfigurations(map[string]string{CONF_SVC_UPDATE_FIELDS: config.CONF_OBJECT_STRINGARR})
	gi.AddStringParam(CONF_DATA_ID)

	return nil
}
func (es *updateStorable) Start(ctx core.ServerContext) error {
	v, _ := es.GetConfiguration(CONF_SVC_UPDATE_FIELDS)
	uf, ok := v.([]string)
	if !ok {
		return errors.BadConf(ctx, CONF_SVC_UPDATE_FIELDS)
	}
	es.updateFields = uf
	return nil
}

func (es *updateStorable) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("UPDATE_WITH_STORABLE")
	id, _ := req.GetStringParam(CONF_DATA_ID)
	stor := req.GetBody().(data.Storable)
	vals := utils.GetObjectFields(stor, es.updateFields)
	log.Debug(ctx, "Coverted storable to fields", "field map", vals, "fields", es.updateFields, "stor", stor)
	return updateVals(ctx, id, vals, es.DataStore)
}
