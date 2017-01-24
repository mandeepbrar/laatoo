package data

import (
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"reflect"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type gaeDataService struct {
	*data.BaseComponent

	conf                    config.Config
	objectConfig            *data.StorableConfig
	name                    string
	auditable               bool
	softdelete              bool
	presave                 bool
	postsave                bool
	postupdate              bool
	postload                bool
	collection              string
	object                  string
	objectid                string
	softDeleteField         string
	objectCollectionCreator core.ObjectCollectionCreator
	objectCreator           core.ObjectCreator
	serviceType             string
}

func newGaeDataService(ctx core.ServerContext, name string) (*gaeDataService, error) {
	gaeDataSvc := &gaeDataService{BaseComponent: &data.BaseComponent{}, name: name, serviceType: "Gae Datastore"}
	return gaeDataSvc, nil
}

func (svc *gaeDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initialize gae datastore service")

	err := svc.BaseComponent.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	svc.objectid = svc.ObjectConfig.IdField
	svc.softDeleteField = svc.ObjectConfig.SoftDeleteField

	if svc.softDeleteField == "" {
		svc.softdelete = false
	} else {
		svc.softdelete = true
	}

	collection, ok := conf.GetString(common.CONF_DATA_COLLECTION)
	if ok {
		svc.collection = collection
	} else {
		svc.collection = svc.ObjectConfig.Collection
	}

	if svc.collection == "" {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", common.CONF_DATA_COLLECTION)
	}

	auditable, ok := conf.GetBool(common.CONF_DATA_AUDITABLE)
	if ok {
		svc.auditable = auditable
	} else {
		svc.auditable = svc.ObjectConfig.Auditable
	}
	postsave, ok := conf.GetBool(common.CONF_DATA_POSTSAVE)
	if ok {
		svc.postsave = postsave
	} else {
		svc.postsave = svc.ObjectConfig.PostSave
	}
	presave, ok := conf.GetBool(common.CONF_DATA_PRESAVE)
	if ok {
		svc.presave = presave
	} else {
		svc.presave = svc.ObjectConfig.PreSave
	}
	postupdate, ok := conf.GetBool(common.CONF_DATA_POSTUPDATE)
	if ok {
		svc.postupdate = postupdate
	} else {
		svc.postupdate = svc.ObjectConfig.PostUpdate
	}
	postload, ok := conf.GetBool(common.CONF_DATA_POSTLOAD)
	if ok {
		svc.postload = postload
	} else {
		svc.postload = svc.ObjectConfig.PostLoad
	}

	return nil
}

func (svc *gaeDataService) Start(ctx core.ServerContext) error {
	return nil
}

func (svc *gaeDataService) Invoke(ctx core.RequestContext) error {
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
	log.Logger.Trace(ctx, "Saving object", "Object", svc.Object)
	id := item.GetId()
	if id == "" {
		item.Init(ctx, nil)
	}
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if svc.auditable {
		data.Audit(ctx, item)
	}
	key := datastore.NewKey(appEngineContext, svc.collection, item.GetId(), 0, nil)

	err := datastore.Get(appEngineContext, key, item)
	if err == nil {
		return errors.ThrowError(ctx, common.DATA_ERROR_OPERATION, "Entity exists ", svc.Object+id)
	}

	_, err = datastore.Put(appEngineContext, key, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
		err = ctx.SendSynchronousMessage(common.CONF_NEWOBJ_MSG, item)
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
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Trace(ctx, "Saving multiple objects", "ObjectType", svc.Object)
	keys := make([]*datastore.Key, len(items))
	for ind, item := range items {
		if item == nil {
			return errors.BadRequest(ctx)
		}
		id := item.GetId()
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
		if svc.presave {
			err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}
		if svc.auditable {
			data.Audit(ctx, item)
		}
	}
	_, err := datastore.PutMulti(appEngineContext, keys, items)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
		for _, item := range items {
			if svc.postsave {
				err = item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
		}
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (svc *gaeDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	ctx = ctx.SubContext("Put")
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Trace(ctx, "Putting object", "ObjectType", svc.Object, "id", id)
	item.SetId(id)
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if svc.auditable {
		data.Audit(ctx, item)
	}
	_, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, svc.collection, id, 0, nil), item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
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
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
		if err != nil {
			return err
		}
	}
	if svc.auditable {
		data.Audit(ctx, newVals)
	}

	object := svc.ObjectCreator()
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	err := datastore.Get(appEngineContext, key, object)
	stor := object.(data.Storable)
	log.Logger.Info(ctx, "Going to set values", "stor", stor, "newVals", newVals)
	stor.SetValues(object, newVals)
	if err != nil {
		if upsert {
			return svc.Save(ctx, stor)
		}
		return err
	}
	log.Logger.Info(ctx, "Set values", "object", object)
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, stor)
		if err != nil {
			return err
		}
		err = stor.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	log.Logger.Info(ctx, "Going to put object", "id", id, "object", object)
	_, err = datastore.Put(appEngineContext, key, object)
	if err != nil {
		return err
	}
	if svc.postupdate {
		err = ctx.SendSynchronousMessage(common.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *gaeDataService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ctx = ctx.SubContext("Upsert")
	return svc.updateAll(ctx, queryCond, newVals, true)
}

func (svc *gaeDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ctx = ctx.SubContext("UpdateAll")
	return svc.updateAll(ctx, queryCond, newVals, false)
}

func (svc *gaeDataService) updateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, upsert bool) ([]string, error) {
	appEngineContext := ctx.GetAppengineContext()
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": svc.Object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	query := datastore.NewQuery(svc.collection)
	query, err := svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, err
	}
	results := svc.ObjectCollectionCreator(0)

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	keys, err := query.GetAll(appEngineContext, results)
	if len(keys) == 0 {
		if upsert {
			object := svc.ObjectCreator()
			stor := object.(data.Storable)
			stor.SetValues(object, newVals)
			err := svc.Save(ctx, stor)
			if err != nil {
				return nil, err
			}
			return []string{stor.GetId()}, nil
		}
		return []string{}, nil
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	for ind, item := range resultStor {
		item.SetValues(reflect.ValueOf(results).Index(ind).Interface(), newVals)
		/*if svc.presave {
			err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
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

	/*if svc.postsave {
		for _, stor := range resultStor {
			err = stor.PostSave(ctx)
			if err != nil {
				errors.WrapError(ctx, err)
			}
		}
	}*/
	if svc.postupdate {
		for _, id := range ids {
			err := ctx.SendSynchronousMessage(common.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
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
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	lenids := len(ids)
	results := svc.ObjectCollectionCreator(lenids)
	keys := make([]*datastore.Key, lenids)
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	err := datastore.GetMulti(appEngineContext, keys, utils.ElementPtr(results))
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Logger.Debug(ctx, "Geting object", "err", err)
			return err
		}
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	for ind, item := range resultStor {
		item.SetValues(reflect.ValueOf(results).Index(ind).Interface(), newVals)
		/*if svc.presave {
			err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
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
	/*if svc.postsave {
		for _, stor := range resultStor {
			err = stor.PostSave(ctx)
			if err != nil {
				errors.WrapError(ctx, err)
			}
		}
	}*/
	if svc.postupdate {
		for _, id := range ids {
			err := ctx.SendSynchronousMessage(common.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
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

	if svc.softdelete {
		return svc.Update(ctx, id, map[string]interface{}{svc.softDeleteField: true})
	}
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	return datastore.Delete(appEngineContext, key)
}

//Delete object by ids
func (svc *gaeDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	ctx = ctx.SubContext("DeleteMulti")
	appEngineContext := ctx.GetAppengineContext()

	if svc.softdelete {
		return svc.UpdateMulti(ctx, ids, map[string]interface{}{svc.softDeleteField: true})
	}
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	return datastore.DeleteMulti(appEngineContext, keys)
}

//Delete object by condition
func (svc *gaeDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if svc.softdelete {
		return svc.UpdateAll(ctx, queryCond, map[string]interface{}{svc.softDeleteField: true})
	}

	appEngineContext := ctx.GetAppengineContext()
	results := svc.ObjectCollectionCreator(0)

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
