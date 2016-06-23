package data

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"

	"gopkg.in/mgo.v2/bson"
)

type mongoDataService struct {
	conf                    config.Config
	name                    string
	database                string
	auditable               bool
	softdelete              bool
	cacheable               bool
	presave                 bool
	postsave                bool
	postload                bool
	notifynew               bool
	notifyupdates           bool
	factory                 *mongoDataServicesFactory
	collection              string
	object                  string
	objectid                string
	objectCollectionCreator core.ObjectCollectionCreator
	objectCreator           core.ObjectCreator
	deleteRefOpers          []*refOperation
	serviceType             string
	/*getRefOpers    map[string][]*refKeyOperation
	putRefOpers    map[string][]*refKeyOperation
	updateRefOpers map[string][]*refKeyOperation*/
}

const (
	CONF_MONGO_DATABASE = "database"
)

func newMongoDataService(ctx core.ServerContext, name string, ms *mongoDataServicesFactory) (*mongoDataService, error) {
	mongoSvc := &mongoDataService{name: name, factory: ms, serviceType: "Mongo"}
	return mongoSvc, nil
}

func (ms *mongoDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
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
	ms.database = ms.factory.database
	ms.conf = conf
	ms.object = object
	ms.objectCreator = objectCreator
	ms.collection = collection
	ms.objectid = objectid
	ms.objectCollectionCreator = objectCollectionCreator

	cacheable, ok := conf.GetBool(CONF_DATA_CACHEABLE)
	if ok {
		ms.cacheable = cacheable
	}
	softdelete, ok := conf.GetBool(CONF_DATA_SOFTDELETE)
	if ok {
		ms.softdelete = softdelete
	} else {
		ms.softdelete = true
	}

	auditable, ok := conf.GetBool(CONF_DATA_AUDITABLE)
	if ok {
		ms.auditable = auditable
	}
	postsave, ok := conf.GetBool(CONF_DATA_POSTSAVE)
	if ok {
		ms.postsave = postsave
	}
	presave, ok := conf.GetBool(CONF_DATA_PRESAVE)
	if ok {
		ms.presave = presave
	}
	postload, ok := conf.GetBool(CONF_DATA_POSTLOAD)
	if ok {
		ms.postload = postload
	}
	notifyupdates, ok := conf.GetBool(CONF_DATA_NOTIFYUPDATES)
	if ok {
		ms.notifyupdates = notifyupdates
	}
	notifynew, ok := conf.GetBool(CONF_DATA_NOTIFYNEW)
	if ok {
		ms.notifynew = notifynew
	}
	deleteOps, _, _, _, err := buildRefOps(ctx, conf)
	if err != nil {
		return err
	}
	ms.deleteRefOpers = deleteOps
	if ms.deleteRefOpers != nil {
		for _, refop := range ms.deleteRefOpers {
			refop.Initialize(ctx)
		}
	}
	return nil
}

func (ms *mongoDataService) Start(ctx core.ServerContext) error {
	return nil
}

func (ms *mongoDataService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (fs *mongoDataService) GetName() string {
	return fs.name
}

func (ms *mongoDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ms *mongoDataService) GetObject() string {
	return ms.object
}

func (ms *mongoDataService) Supports(feature data.Feature) bool {
	switch feature {
	case data.InQueries:
		return true
	case data.Ancestors:
		return false
	}
	return false
}

func (ms *mongoDataService) Save(ctx core.RequestContext, item data.Storable) error {
	log.Logger.Trace(ctx, "Saving object", "Object", ms.object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	if ms.presave {
		err := ctx.SendSynchronousMessage(CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if ms.auditable {
		data.Audit(ctx, item)
	}
	id := item.GetId()
	if id == "" {
		return errors.ThrowError(ctx, DATA_ERROR_ID_NOT_FOUND, "ObjectType", ms.object)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Insert(item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.postsave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (ms *mongoDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	log.Logger.Trace(ctx, "Saving multiple objects", "ObjectType", ms.object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	bulk := connCopy.DB(ms.database).C(ms.collection).Bulk()
	for _, item := range items {
		id := item.GetId()
		invalidateCache(ctx, ms.object, id)
		if ms.presave {
			err := ctx.SendSynchronousMessage(CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}
		if ms.auditable {
			data.Audit(ctx, item)
		}
		bulk.Upsert(bson.M{ms.objectid: id}, item)
	}
	_, err := bulk.Run()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.postsave || ms.notifyupdates {
		for _, item := range items {
			if ms.postsave {
				err = item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
			if ms.notifyupdates {
				notifyUpdate(ctx, ms.object, item.GetId())
			}
		}
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (ms *mongoDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	invalidateCache(ctx, ms.object, id)
	log.Logger.Trace(ctx, "Putting object", "ObjectType", ms.object, "id", id)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.objectid] = id
	item.SetId(id)
	if ms.presave {
		err := ctx.SendSynchronousMessage(CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		log.Logger.Trace(ctx, "Putting object", "err", err)
		if err != nil {
			return err
		}
	}
	if ms.auditable {
		data.Audit(ctx, item)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Update(condition, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.postsave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	if ms.notifyupdates {
		notifyUpdate(ctx, ms.object, id)
	}
	return nil
}

func (ms *mongoDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
	if err != nil {
		return err
	}
	invalidateCache(ctx, ms.object, id)
	if ms.notifyupdates {
		notifyUpdate(ctx, ms.object, id)
	}
	return nil
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	results, err := ms.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	err = query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resultStor, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	ids := make([]string, length)
	if length == 0 {
		return nil, nil
	}
	for i := 0; i < length; i++ {
		ids[i] = resultStor[i].GetId()
		invalidateCache(ctx, ms.object, ids[i])
	}
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).UpdateAll(queryCond, map[string]interface{}{"$set": newVals})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if ms.notifyupdates {
		for _, id := range ids {
			notifyUpdate(ctx, ms.object, id)
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := ms.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, ms.objectid, ids)
	for _, id := range ids {
		invalidateCache(ctx, ms.object, id)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).UpdateAll(condition, updateInterface)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.notifyupdates {
		for _, id := range ids {
			notifyUpdate(ctx, ms.object, id)
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (ms *mongoDataService) Delete(ctx core.RequestContext, id string) error {
	if ms.softdelete {
		err := ms.Update(ctx, id, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms.deleteRefOpers, []string{id})
		}
		return err
	}
	invalidateCache(ctx, ms.object, id)
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Remove(condition)
	if err == nil {
		err = deleteRefOps(ctx, ms.deleteRefOpers, []string{id})
	}
	return err
}

//Delete object by ids
func (ms *mongoDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	if ms.softdelete {
		err := ms.UpdateMulti(ctx, ids, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms.deleteRefOpers, ids)
		}
		return err
	}
	conditionVal := bson.M{}
	conditionVal["$in"] = ids
	condition := bson.M{}
	condition[ms.objectid] = conditionVal
	for _, id := range ids {
		invalidateCache(ctx, ms.object, id)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).RemoveAll(condition)
	if err == nil {
		err = deleteRefOps(ctx, ms.deleteRefOpers, ids)
	}
	return err
}

//Delete object by condition
func (ms *mongoDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	if ms.softdelete {
		ids, err := ms.UpdateAll(ctx, queryCond, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms.deleteRefOpers, ids)
		}
		return ids, err
	}
	results, err := ms.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	err = query.All(results)
	resultStor, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	ids := make([]string, length)
	if length == 0 {
		return nil, nil
	}
	for i := 0; i < length; i++ {
		id := resultStor[i].GetId()
		ids[i] = id
		invalidateCache(ctx, ms.object, id)
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).RemoveAll(queryCond)
	if err == nil {
		err = deleteRefOps(ctx, ms.deleteRefOpers, ids)
		if err != nil {
			return ids, err
		}
	} else {
		return nil, err
	}
	return ids, err
}
