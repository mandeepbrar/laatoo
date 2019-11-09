package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"reflect"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type gaeDataService struct {
	*data.BaseComponent

	conf        config.Config
	name        string
	collection  string
	serviceType string
}

func newGaeDataService(ctx core.ServerContext, name string) (*gaeDataService, error) {
	gaeDataSvc := &gaeDataService{BaseComponent: &data.BaseComponent{}, name: name, serviceType: "Gae Datastore"}
	return gaeDataSvc, nil
}

func (svc *gaeDataService) Describe(ctx core.ServerContext) error {
	err := svc.BaseComponent.Describe(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.AddOptionalConfigurations(ctx, map[string]string{data.CONF_DATA_COLLECTION: config.OBJECTTYPE_STRING}, nil)

	svc.SetDescription(ctx, "GAE data component")
	return nil
}

func (svc *gaeDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initialize gae datastore service")

	err := svc.BaseComponent.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	collection, ok := svc.GetConfiguration(ctx, data.CONF_DATA_COLLECTION)
	if ok {
		svc.collection = collection.(string)
	} else {
		svc.collection = svc.ObjectConfig.Collection
	}

	if svc.collection == "" {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", data.CONF_DATA_COLLECTION)
	}
	return nil
}

func (svc *gaeDataService) Start(ctx core.ServerContext) error {

	err := svc.BaseComponent.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (svc *gaeDataService) GetName() string {
	return svc.name
}

func (svc *gaeDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (svc *gaeDataService) GetObject() string {
	return svc.Object
}

func (svc *gaeDataService) Supports(feature data.Feature) bool {
	switch feature {
	case data.InQueries:
		return false
	case data.Ancestors:
		return true
	}
	return false
}

func (svc *gaeDataService) Save(ctx core.RequestContext, item data.Storable) error {
	ctx = ctx.SubContext("Save")
	appEngineContext := ctx.GetAppengineContext()
	log.Trace(ctx, "Saving object", "Object", svc.Object)
	id := item.GetId()
	if id == "" {
		item.Initialize(ctx, nil)
	}
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if svc.Multitenant {
		item.SetTenant(ctx.GetUser().GetTenant())
	}
	if svc.Auditable {
		data.Audit(ctx, item)
	}
	key := datastore.NewKey(appEngineContext, svc.collection, item.GetId(), 0, nil)

	err := datastore.Get(appEngineContext, key, item)
	if err == nil {
		return errors.ThrowError(ctx, data.DATA_ERROR_OPERATION, "Entity exists ", svc.Object+id)
	}

	_, err = datastore.Put(appEngineContext, key, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.PostSave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
		err = ctx.SendSynchronousMessage(data.CONF_NEWOBJ_MSG, item)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *gaeDataService) CreateMulti(ctx core.RequestContext, items []data.Storable) error {
	return svc.PutMulti(ctx, items)
}

func (svc *gaeDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	ctx = ctx.SubContext("PutMulti")
	if svc.Multitenant {
		for _, item := range items {
			if item.GetTenant() != ctx.GetUser().GetTenant() {
				return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.GetTenant(), "Item", item.GetId())
			}
		}
	}
	appEngineContext := ctx.GetAppengineContext()
	log.Trace(ctx, "Saving multiple objects", "ObjectType", svc.Object)
	keys := make([]*datastore.Key, len(items))
	for ind, item := range items {
		if item == nil {
			return errors.BadRequest(ctx)
		}
		id := item.GetId()
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
		if svc.PreSave {
			err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}
		if svc.Auditable {
			data.Audit(ctx, item)
		}
	}
	_, err := datastore.PutMulti(appEngineContext, keys, items)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.PostSave {
		for _, item := range items {
			if svc.PostSave {
				err = item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
		}
	}
	log.Trace(ctx, "Saved multiple objects")
	return nil
}

func (svc *gaeDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	ctx = ctx.SubContext("Put")
	if svc.Multitenant && (item.GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.GetTenant(), "Item", item.GetId())
	}
	appEngineContext := ctx.GetAppengineContext()
	log.Trace(ctx, "Putting object", "ObjectType", svc.Object, "id", id)
	item.SetId(id)
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if svc.Auditable {
		data.Audit(ctx, item)
	}
	_, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, svc.collection, id, 0, nil), item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.PostSave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	return nil
}

//upsert an object ...insert if not there... update if there
func (svc *gaeDataService) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpsertId")
	return svc.update(ctx, id, newVals, true)
}

func (svc *gaeDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("Update")
	upsert, _ := ctx.GetBool("upsert")
	return svc.update(ctx, id, newVals, upsert)
}

func (svc *gaeDataService) update(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	appEngineContext := ctx.GetAppengineContext()
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
		if err != nil {
			return err
		}
	}
	if svc.Auditable {
		data.Audit(ctx, newVals)
	}

	object, _ := ctx.CreateObject(svc.Object)
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	err := datastore.Get(appEngineContext, key, object)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	stor := object.(data.Storable)
	if svc.Multitenant && (stor.GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", ctx.GetUser().GetTenant(), "Item", id)
	}
	log.Info(ctx, "Going to set values", "stor", stor, "newVals", newVals)
	stor.SetValues(object, newVals)
	if upsert {
		err = svc.Save(ctx, stor)
	} else {
		err = svc.Put(ctx, stor.GetId(), stor)
	}

	if err != nil {
		return err
	}
	if svc.PostUpdate {
		err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *gaeDataService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, getids bool) ([]string, error) {
	ctx = ctx.SubContext("Upsert")
	return svc.updateAll(ctx, queryCond, newVals, true, getids)
}

func (svc *gaeDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, getids bool) ([]string, error) {
	ctx = ctx.SubContext("UpdateAll")
	return svc.updateAll(ctx, queryCond, newVals, false, getids)
}

func (svc *gaeDataService) updateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, upsert bool, getids bool) ([]string, error) {
	appEngineContext := ctx.GetAppengineContext()
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": svc.Object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}
	if svc.Auditable {
		data.Audit(ctx, newVals)
	}
	query := datastore.NewQuery(svc.collection)
	query, err := svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, err
	}
	results, _ := ctx.CreateCollection(svc.Object, 0)

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	keys, err := query.GetAll(appEngineContext, results)
	if len(keys) == 0 {
		if upsert {
			object, _ := ctx.CreateObject(svc.Object)
			stor := object.(data.Storable)
			if svc.Multitenant && (stor.GetTenant() != ctx.GetUser().GetTenant()) {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", ctx.GetUser().GetTenant(), "Item", stor.GetId())
			}
			stor.SetValues(object, newVals)
			err := svc.Save(ctx, stor)
			if err != nil {
				return nil, err
			}
			return []string{stor.GetId()}, nil
		}
		return []string{}, nil
	}
	resultStor, ids, err := data.CastToStorableCollection(ctx, results)
	for ind, item := range resultStor {
		item.SetValues(reflect.ValueOf(results).Index(ind).Interface(), newVals)
		/*if svc.PreSave {
			err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
			if err != nil {
				return nil, err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return nil, err
			}
		}*/
	}
	_, err = datastore.PutMulti(appEngineContext, keys, results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	/*if svc.PostSave {
		for _, stor := range resultStor {
			err = stor.PostSave(ctx)
			if err != nil {
				errors.WrapError(ctx, err)
			}
		}
	}*/
	if svc.PostUpdate {
		for _, id := range ids {
			err := ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
			if err != nil {
				return nil, errors.WrapError(ctx, err)
			}
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *gaeDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpdateMulti")
	appEngineContext := ctx.GetAppengineContext()
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	if svc.Auditable {
		data.Audit(ctx, newVals)
	}
	lenids := len(ids)
	results, _ := ctx.CreateCollection(svc.Object, lenids)
	keys := make([]*datastore.Key, lenids)
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	err := datastore.GetMulti(appEngineContext, keys, utils.ElementPtr(results))
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Debug(ctx, "Geting object", "err", err)
			return err
		}
	}
	resultStor, ids, err := data.CastToStorableCollection(ctx, results)
	for ind, item := range resultStor {
		if svc.Multitenant && (item.GetTenant() != ctx.GetUser().GetTenant()) {
			return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", ctx.GetUser().GetTenant(), "Item", item.GetId())
		}
		item.SetValues(reflect.ValueOf(results).Index(ind).Interface(), newVals)
		/*if svc.PreSave {
			err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}*/
	}
	_, err = datastore.PutMulti(appEngineContext, keys, results)
	if err != nil {
		return err
	}
	/*if svc.PostSave {
		for _, stor := range resultStor {
			err = stor.PostSave(ctx)
			if err != nil {
				errors.WrapError(ctx, err)
			}
		}
	}*/
	if svc.PostUpdate {
		for _, id := range ids {
			err := ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (svc *gaeDataService) Delete(ctx core.RequestContext, id string) error {
	ctx = ctx.SubContext("Delete")
	appEngineContext := ctx.GetAppengineContext()

	if svc.SoftDelete {
		return svc.Update(ctx, id, map[string]interface{}{svc.SoftDeleteField: true})
	}
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	return datastore.Delete(appEngineContext, key)
}

//Delete object by ids
func (svc *gaeDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	ctx = ctx.SubContext("DeleteMulti")
	appEngineContext := ctx.GetAppengineContext()

	if svc.SoftDelete {
		return svc.UpdateMulti(ctx, ids, map[string]interface{}{svc.SoftDeleteField: true})
	}
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	return datastore.DeleteMulti(appEngineContext, keys)
}

//Delete object by condition
func (svc *gaeDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}, getids bool) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if svc.SoftDelete {
		return svc.UpdateAll(ctx, queryCond, map[string]interface{}{svc.SoftDeleteField: true})
	}

	appEngineContext := ctx.GetAppengineContext()
	results, _ := ctx.CreateCollection(svc.Object, 0)

	query := datastore.NewQuery(svc.collection)
	query, err := svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, err
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	keys, err := query.KeysOnly().GetAll(appEngineContext, results)
	ids := make([]string, len(keys))
	for i, val := range keys {
		ids[i] = val.StringID()
	}
	err = datastore.DeleteMulti(appEngineContext, keys)
	return ids, err
}
