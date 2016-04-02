package data

import (
	"gopkg.in/mgo.v2/bson"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type mongoDataService struct {
	conf                    config.Config
	database                string
	auditable               bool
	softdelete              bool
	cacheable               bool
	factory                 *mongoDataServicesFactory
	collection              string
	object                  string
	objectid                string
	objectCollectionCreator core.ObjectCollectionCreator
	objectCreator           core.ObjectCreator
	/*deleteRefOpers map[string][]*refKeyOperation
	getRefOpers    map[string][]*refKeyOperation
	putRefOpers    map[string][]*refKeyOperation
	updateRefOpers map[string][]*refKeyOperation*/
}

const (
	CONF_MONGO_DATABASE   = "database"
	CONF_MONGO_OBJECT     = "object"
	CONF_MONGO_OBJECT_ID  = "id"
	CONF_MONGO_CACHEABLE  = "cacheable"
	CONF_MONGO_AUDITABLE  = "auditable"
	CONF_MONGO_SOFTDELETE = "softdelete"
	CONF_MONGO_COLLECTION = "collection"
)

func newMongoDataService(ctx core.ServerContext, ms *mongoDataServicesFactory, conf config.Config) (data.DataService, error) {
	database, ok := conf.GetString(CONF_MONGO_DATABASE)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_DATABASE)
	}
	collection, ok := conf.GetString(CONF_MONGO_COLLECTION)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_COLLECTION)
	}
	object, ok := conf.GetString(CONF_MONGO_OBJECT)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_OBJECT)
	}
	objectid, ok := conf.GetString(CONF_MONGO_OBJECT_ID)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_OBJECT_ID)
	}
	objectCreator, err := registry.GetObjectCreator(ctx, object)
	if err != nil {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object creator for", object)
	}
	objectCollectionCreator, err := registry.GetObjectCollectionCreator(ctx, object)
	if err != nil {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object Collection creator for", object)
	}

	mongoSvc := &mongoDataService{factory: ms, conf: conf, object: object, objectCreator: objectCreator,
		collection: collection, objectid: objectid, objectCollectionCreator: objectCollectionCreator, database: database}
	cacheable, ok := conf.GetString(CONF_MONGO_CACHEABLE)
	if ok {
		mongoSvc.cacheable = (cacheable == "true")
	}
	softdelete, ok := conf.GetString(CONF_MONGO_SOFTDELETE)
	if ok {
		mongoSvc.softdelete = (softdelete == "true")
	}
	auditable, ok := conf.GetString(CONF_MONGO_AUDITABLE)
	if ok {
		mongoSvc.auditable = (auditable == "true")
	}
	/*mongoSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		mongoSvc.objects[obj] = collection.(string)
	}
	deleteOps, _, _, _, err := buildRefOps(ctx, conf)
	if err != nil {
		return nil, err
	}
	mongoSvc.deleteRefOpers = deleteOps
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Mongo service configured for objects ", "Objects", mongoSvc.objects)*/
	return mongoSvc, nil
}

func (ms *mongoDataService) Initialize(ctx core.ServerContext) error {
	return nil
}

func (ms *mongoDataService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (ms *mongoDataService) GetConf() config.Config {
	return ms.conf
}

func (ms *mongoDataService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}

func (ms *mongoDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
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
	item.PreSave(ctx)
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
	item.PostSave(ctx)
	if ms.cacheable {
		ctx.PutInCache(getCacheKey(ms.object, id), item)
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
		if ms.cacheable {
			ctx.DeleteFromCache(getCacheKey(ms.object, id))
		}
		item.PreSave(ctx)
		if ms.auditable {
			data.Audit(ctx, item)
		}
		bulk.Upsert(bson.M{ms.objectid: id}, item)
	}
	_, err := bulk.Run()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, item := range items {
		item.PostSave(ctx)
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (ms *mongoDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	if ms.cacheable {
		ctx.DeleteFromCache(getCacheKey(ms.object, id))
	}
	log.Logger.Trace(ctx, "Putting object", "ObjectType", ms.object, "id", id)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.objectid] = id
	item.PreSave(ctx)
	if ms.auditable {
		data.Audit(ctx, item)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Update(condition, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	item.PostSave(ctx)
	return nil
}

func (ms *mongoDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	if ms.cacheable {
		ctx.DeleteFromCache(getCacheKey(ms.object, id))
	}
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	return connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	results, err := ms.objectCollectionCreator(ctx, nil)
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
		if ms.cacheable {
			ctx.DeleteFromCache(getCacheKey(ms.object, ids[i]))
		}
	}
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).UpdateAll(queryCond, map[string]interface{}{"$set": newVals})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	if ms.cacheable {
		for _, id := range ids {
			ctx.DeleteFromCache(getCacheKey(ms.object, id))
		}
	}
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := ms.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, ms.objectid, ids)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	return connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
}

//item must support Deleted field for soft deletes
func (ms *mongoDataService) Delete(ctx core.RequestContext, id string) error {
	if ms.softdelete {
		err := ms.Update(ctx, id, map[string]interface{}{"Deleted": true})
		/*if err == nil
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, []string{id})
		}*/
		return err
	}
	if ms.cacheable {
		ctx.DeleteFromCache(getCacheKey(ms.object, id))
	}
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Remove(condition)
	/*if err == nil {
		err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, []string{id})
	}*/
	return err
}

//Delete object by ids
func (ms *mongoDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	if ms.softdelete {
		err := ms.UpdateMulti(ctx, ids, map[string]interface{}{"Deleted": true})
		/*if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		}*/
		return err
	}
	if ms.cacheable {
		for _, id := range ids {
			ctx.DeleteFromCache(getCacheKey(ms.object, id))
		}
	}
	conditionVal := bson.M{}
	conditionVal["$in"] = ids
	condition := bson.M{}
	condition[ms.objectid] = conditionVal
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).RemoveAll(condition)
	/*if err == nil {
		err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
	}*/
	return err
}

//Delete object by condition
func (ms *mongoDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	if ms.softdelete {
		ids, err := ms.UpdateAll(ctx, queryCond, map[string]interface{}{"Deleted": true})
		/*if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		}*/
		return ids, err
	}
	results, err := ms.objectCollectionCreator(ctx, nil)
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
		if ms.cacheable {
			ctx.DeleteFromCache(getCacheKey(ms.object, id))
		}
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).RemoveAll(queryCond)
	if err == nil {
		/*err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		if err != nil {
			return ids, err
		}*/
	} else {
		return nil, err
	}
	return ids, err
}
