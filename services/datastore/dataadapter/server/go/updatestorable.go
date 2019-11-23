package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
)

type updateStorable struct {
	core.Service
	fac          *DataAdapterFactory
	DataStore    data.DataComponent
	updateFields []string
}

func (gi *updateStorable) Describe(ctx core.ServerContext) error {
	gi.SetDescription(ctx, "Update object using underlying data component. Expects an entity id. Value should be storable object")
	gi.AddConfigurations(ctx, map[string]string{CONF_SVC_UPDATE_FIELDS: config.OBJECTTYPE_STRINGARR})
	gi.AddStringParam(ctx, CONF_DATA_ID)
	return nil
}

func (svc *updateStorable) Start(ctx core.ServerContext) error {
	svc.DataStore = svc.fac.DataStore
	//svc.SetRequestType(ctx, svc.DataStore.GetObject(), false, false)
	v, _ := svc.GetConfiguration(ctx, CONF_SVC_UPDATE_FIELDS)
	uf, ok := v.([]string)
	if !ok {
		return errors.BadConf(ctx, CONF_SVC_UPDATE_FIELDS)
	}
	svc.updateFields = uf
	/**** TODO***/
	return svc.AddParamWithType(ctx, "object", svc.DataStore.GetObject())
}

func (es *updateStorable) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATE_WITH_STORABLE")
	id, _ := ctx.GetStringParam(CONF_DATA_ID)
	obj, _ := ctx.GetParamValue("object")
	stor := obj.(data.Storable)
	vals := utils.GetObjectFields(stor, es.updateFields)
	log.Debug(ctx, "Coverted storable to fields", "field map", vals, "fields", es.updateFields, "stor", stor)
	res, err := updateVals(ctx, id, vals, es.DataStore)
	ctx.SetResponse(res)
	return err
}