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

func (gi *updateStorable) Describe(ctx core.ServerContext) {
	gi.SetDescription(ctx, "Update object using underlying data component. Expects an entity id. Value should be storable object")
	gi.SetRequestType(ctx, gi.DataStore.GetObject(), false, false)
	gi.AddConfigurations(ctx, map[string]string{CONF_SVC_UPDATE_FIELDS: config.OBJECTTYPE_STRINGARR})
	gi.AddStringParam(ctx, CONF_DATA_ID)

}

func (es *updateStorable) Start(ctx core.ServerContext) error {
	es.DataStore = es.fac.DataStore
	v, _ := es.GetConfiguration(ctx, CONF_SVC_UPDATE_FIELDS)
	uf, ok := v.([]string)
	if !ok {
		return errors.BadConf(ctx, CONF_SVC_UPDATE_FIELDS)
	}
	es.updateFields = uf
	return nil
}

func (es *updateStorable) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATE_WITH_STORABLE")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	stor := ctx.GetBody().(data.Storable)
	vals := utils.GetObjectFields(stor, es.updateFields)
	log.Debug(ctx, "Coverted storable to fields", "field map", vals, "fields", es.updateFields, "stor", stor)
	res, err := updateVals(ctx, id, vals, es.DataStore)
	ctx.SetResponse(res)
	return err
}
