package data

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"reflect"
)

type gaeDataService struct {
	conf                    config.Config
	name                    string
	auditable               bool
	softdelete              bool
	cacheable               bool
	presave                 bool
	postsave                bool
	postload                bool
	notifynew               bool
	notifyupdates           bool
	collection              string
	object                  string
	objectid                string
	objectCollectionCreator core.ObjectCollectionCreator
	objectCreator           core.ObjectCreator
	deleteRefOpers          []*refOperation
	/*getRefOpers    map[string][]*refKeyOperation
	putRefOpers    map[string][]*refKeyOperation
	updateRefOpers map[string][]*refKeyOperation*/
}

func newGaeDataService(ctx core.ServerContext, name string) (*gaeDataService, error) {
	gaeDataSvc := &gaeDataService{}
	return gaeDataSvc, nil
}

func (svc *gaeDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	collection, ok := conf.GetString(CONF_DATA_COLLECTION)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_DATA_COLLECTION)
	}
	object, ok := conf.GetString(CONF_DATA_OBJECT)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_DATA_OBJECT)
	}
	objectid, ok := conf.GetString(CONF_DATA_OBJECT_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_DATA_OBJECT_ID)
	}
	objectCreator, err := ctx.GetObjectCreator(object)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object creator for", object)
	}
	objectCollectionCreator, err := ctx.GetObjectCollectionCreator(object)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object Collection creator for", object)
	}
	svc.conf = conf
	svc.object = object
	svc.objectCreator = objectCreator
	svc.collection = collection
	svc.objectid = objectid
	svc.objectCollectionCreator = objectCollectionCreator

	cacheable, ok := conf.GetBool(CONF_DATA_CACHEABLE)
	if ok {
		svc.cacheable = cacheable
	}
	softdelete, ok := conf.GetBool(CONF_DATA_SOFTDELETE)
	if ok {
		svc.softdelete = softdelete
	} else {
		svc.softdelete = true
	}

	auditable, ok := conf.GetBool(CONF_DATA_AUDITABLE)
	if ok {
		svc.auditable = auditable
	}
	postsave, ok := conf.GetBool(CONF_DATA_POSTSAVE)
	if ok {
		svc.postsave = postsave
	}
	presave, ok := conf.GetBool(CONF_DATA_PRESAVE)
	if ok {
		svc.presave = presave
	}
	postload, ok := conf.GetBool(CONF_DATA_POSTLOAD)
	if ok {
		svc.postload = postload
	}
	notifyupdates, ok := conf.GetBool(CONF_DATA_NOTIFYUPDATES)
	if ok {
		svc.notifyupdates = notifyupdates
	}
	notifynew, ok := conf.GetBool(CONF_DATA_NOTIFYNEW)
	if ok {
		svc.notifynew = notifynew
	}
	deleteOps, _, _, _, err := buildRefOps(ctx, conf)
	if err != nil {
		return err
	}
	svc.deleteRefOpers = deleteOps
	if svc.deleteRefOpers != nil {
		for _, refop := range svc.deleteRefOpers {
			refop.Initialize(ctx)
		}
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
	return svc.object
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
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Trace(ctx, "Saving object", "Object", svc.object)
	if svc.presave {
		err := ctx.SendSynchronousMessage(CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		item.PreSave(ctx)
	}
	if svc.auditable {
		data.Audit(ctx, item)
	}
	id := item.GetId()
	if id == "" {
		return errors.ThrowError(ctx, DATA_ERROR_ID_NOT_FOUND, "ObjectType", svc.object)
	}
	_, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, svc.collection, id, 0, nil), item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
		item.PostSave(ctx)
	}
	return nil
}

func (svc *gaeDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Trace(ctx, "Saving multiple objects", "ObjectType", svc.object)
	keys := make([]*datastore.Key, len(items))
	for ind, item := range items {
		id := item.GetId()
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
		invalidateCache(ctx, svc.object, id)
		if svc.presave {
			err := ctx.SendSynchronousMessage(CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			item.PreSave(ctx)
		}
		if svc.auditable {
			data.Audit(ctx, item)
		}
	}
	_, err := datastore.PutMulti(appEngineContext, keys, items)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave || svc.notifyupdates {
		for _, item := range items {
			if svc.postsave {
				item.PostSave(ctx)
			}
			if svc.notifyupdates {
				notifyUpdate(ctx, svc.object, item.GetId())
			}
		}
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (svc *gaeDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	appEngineContext := ctx.GetAppengineContext()
	invalidateCache(ctx, svc.object, id)
	log.Logger.Trace(ctx, "Putting object", "ObjectType", svc.object, "id", id)
	item.SetId(id)
	if svc.presave {
		err := ctx.SendSynchronousMessage(CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		item.PreSave(ctx)
	}
	if svc.auditable {
		data.Audit(ctx, item)
	}
	_, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, svc.collection, id, 0, nil), item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
		item.PostSave(ctx)
	}
	if svc.notifyupdates {
		notifyUpdate(ctx, svc.object, id)
	}
	return nil
}

func (svc *gaeDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	if svc.auditable {
		data.Audit(ctx, newVals)
	}

	object, err := svc.objectCreator(ctx, nil)
	if err != nil {
		return err
	}
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	err = datastore.Get(appEngineContext, key, object)
	if err != nil {
		return err
	}
	entVal := reflect.ValueOf(object).Elem()
	for k, v := range newVals {
		f := entVal.FieldByName(k)
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				f.Set(reflect.ValueOf(v))
			}
		}
	}
	key, err = datastore.Put(appEngineContext, key, object)
	if err != nil {
		return err
	}
	invalidateCache(ctx, svc.object, id)
	if svc.notifyupdates {
		notifyUpdate(ctx, svc.object, id)
	}
	return nil
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *gaeDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	appEngineContext := ctx.GetAppengineContext()
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	query := datastore.NewQuery(svc.collection)
	query, err := svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, err
	}
	results, err := svc.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	keys, err := query.GetAll(appEngineContext, results)
	resultStor, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	ids := make([]string, length)
	for i, item := range resultStor {
		entVal := reflect.ValueOf(item)
		ids[i] = item.GetId()
		invalidateCache(ctx, svc.object, ids[i])
		for k, v := range newVals {
			f := entVal.FieldByName(k)
			if f.IsValid() {
				// A Value can be changed only if it is
				// addressable and was not obtained by
				// the use of unexported struct fields.
				if f.CanSet() {
					f.Set(reflect.ValueOf(v))
				}
			}
		}
	}
	_, err = datastore.PutMulti(appEngineContext, keys, results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	if svc.notifyupdates {
		for _, id := range ids {
			notifyUpdate(ctx, svc.object, id)
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *gaeDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	for _, id := range ids {
		invalidateCache(ctx, svc.object, id)
	}
	lenids := len(ids)
	results, _ := svc.objectCollectionCreator(ctx, lenids, nil)
	keys := make([]*datastore.Key, lenids)
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	err := datastore.GetMulti(appEngineContext, keys, reflect.ValueOf(results).Elem().Interface())
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Logger.Debug(ctx, "Geting object", "err", err)
			return err
		}
	}
	resultStor, err := data.CastToStorableCollection(results)
	for _, item := range resultStor {
		entVal := reflect.ValueOf(item)
		for k, v := range newVals {
			f := entVal.FieldByName(k)
			if f.IsValid() {
				// A Value can be changed only if it is
				// addressable and was not obtained by
				// the use of unexported struct fields.
				if f.CanSet() {
					f.Set(reflect.ValueOf(v))
				}
			}
		}
	}
	_, err = datastore.PutMulti(appEngineContext, keys, results)
	if err != nil {
		return err
	}
	if svc.notifyupdates {
		for _, id := range ids {
			notifyUpdate(ctx, svc.object, id)
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (svc *gaeDataService) Delete(ctx core.RequestContext, id string) error {
	appEngineContext := ctx.GetAppengineContext()

	if svc.softdelete {
		err := svc.Update(ctx, id, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, svc.deleteRefOpers, []string{id})
		}
		return err
	}
	invalidateCache(ctx, svc.object, id)
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	err := datastore.Delete(appEngineContext, key)
	if err == nil {
		err = deleteRefOps(ctx, svc.deleteRefOpers, []string{id})
	}
	return err
}

//Delete object by ids
func (svc *gaeDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	appEngineContext := ctx.GetAppengineContext()

	if svc.softdelete {
		err := svc.UpdateMulti(ctx, ids, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, svc.deleteRefOpers, ids)
		}
		return err
	}
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		invalidateCache(ctx, svc.object, id)
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	err := datastore.DeleteMulti(appEngineContext, keys)
	if err == nil {
		err = deleteRefOps(ctx, svc.deleteRefOpers, ids)
	}
	return err
}

//Delete object by condition
func (svc *gaeDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	if svc.softdelete {
		ids, err := svc.UpdateAll(ctx, queryCond, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, svc.deleteRefOpers, ids)
		}
		return ids, err
	}

	appEngineContext := ctx.GetAppengineContext()
	results, err := svc.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	query := datastore.NewQuery(svc.collection)
	query, err = svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, err
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	keys, err := query.KeysOnly().GetAll(appEngineContext, results)
	ids := make([]string, len(keys))
	for i, val := range keys {
		ids[i] = val.StringID()
		invalidateCache(ctx, svc.object, ids[i])
	}
	err = datastore.DeleteMulti(appEngineContext, keys)

	if err == nil {
		err = deleteRefOps(ctx, svc.deleteRefOpers, ids)
		if err != nil {
			return ids, err
		}
	} else {
		return nil, err
	}
	return ids, err
}
