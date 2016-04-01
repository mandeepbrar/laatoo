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

func (ms *mongoDataService) GetById(ctx core.Context, id string) (data.Storable, error) {
	log.Logger.Trace(ctx, "Getting object by id ", "id", id, "object", ms.object)
	object, err := ms.objectCreator(ctx, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.objectid] = id
	err = connCopy.DB(ms.database).C(ms.collection).Find(condition).One(object)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		return nil, errors.RethrowError(ctx, DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	stor.PostLoad(ctx)
	return stor, nil
}

func (ms *mongoDataService) Save(ctx core.Context, item data.Storable) error {
	log.Logger.Trace(ctx, "Saving object", "Object", ms.object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	item.PreSave(ctx)
	id := item.GetId()
	if id == "" {
		return errors.ThrowError(ctx, DATA_ERROR_ID_NOT_FOUND, "ObjectType", ms.object)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Insert(item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	item.PostSave(ctx)
	return nil
}

func (ms *mongoDataService) PutMulti(ctx core.Context, ids []string, items []data.Storable) error {
	log.Logger.Trace(ctx, "Saving multiple objects", "ObjectType", ms.object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	bulk := connCopy.DB(ms.database).C(ms.collection).Bulk()
	for _, item := range items {
		item.PreSave(ctx)
		bulk.Upsert(bson.M{ms.objectid: item.GetId()}, item)
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

func (ms *mongoDataService) Put(ctx core.Context, id string, item data.Storable) error {
	log.Logger.Trace(ctx, "Putting object", "ObjectType", ms.object, "id", id)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.objectid] = id
	item.PreSave(ctx)
	err := connCopy.DB(ms.database).C(ms.collection).Update(condition, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	item.PostSave(ctx)
	return nil
}

//Get multiple objects by id
func (ms *mongoDataService) GetMulti(ctx core.Context, ids []string, orderBy string) (map[string]data.Storable, error) {
	results, err := ms.objectCollectionCreator(ctx, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	operatorCond := bson.M{}
	operatorCond["$in"] = ids
	condition[ms.objectid] = operatorCond
	query := connCopy.DB(ms.database).C(ms.collection).Find(condition)
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	retVal := make(map[string]data.Storable, len(ids))
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	resultStor, err := data.CastToStorableCollection(results)
	for _, stor := range resultStor {
		retVal[stor.GetId()] = stor
		stor.PostLoad(ctx)
	}
	for _, id := range ids {
		_, ok := retVal[id]
		if !ok {
			retVal[id] = nil
		}
	}
	return retVal, nil
}

func (ms *mongoDataService) GetList(ctx core.Context, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	/*totalrecs = -1
	recsreturned = -1
	results, err := ms.objectCollectionCreator(ctx, nil)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	query := connCopy.DB(ms.database).C(ms.collection).Find(condition)
	if pageSize > 0 {
		totalrecs, err = query.Count()
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Skip(recsToSkip)
	}
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	resultStor, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	recsreturned = len(resultStor)
	for _, stor := range resultStor {
		stor.PostLoad(ctx)
	}*/
	return ms.Get(ctx, bson.M{}, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) Get(ctx core.Context, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	totalrecs = -1
	recsreturned = -1
	results, err := ms.objectCollectionCreator(ctx, nil)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	if pageSize > 0 {
		totalrecs, err = query.Count()
		if err != nil {
			return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Skip(recsToSkip)
	}
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	resultStor, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	recsreturned = len(resultStor)
	for _, stor := range resultStor {
		stor.PostLoad(ctx)
	}
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", ms.object, "recsreturned", recsreturned)
	return resultStor, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) Update(ctx core.Context, id string, newVals map[string]interface{}) error {
	updateInterface := map[string]interface{}{"$set": newVals}
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	return connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateAll(ctx core.Context, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
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
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).UpdateAll(queryCond, map[string]interface{}{"$set": newVals})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateMulti(ctx core.Context, ids []string, newVals map[string]interface{}) error {
	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := ms.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, ms.objectid, ids)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	return connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
}

//create condition for passing to data service
func (ms *mongoDataService) CreateCondition(ctx core.Context, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	switch operation {
	case data.MATCHANCESTOR:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
	case data.MATCHMULTIPLEVALUES:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return bson.M{args[0].(string): bson.M{"$in": args[1]}}, nil
		}
	case data.FIELDVALUE:
		{
			if len(args) < 1 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return args[0], nil
		}
	}
	return nil, nil
}

func (ms *mongoDataService) Delete(ctx core.Context, id string, softdelete bool) error {
	/*if softdelete {
		err := ms.Update(ctx, objectType, id, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, []string{id})
		}
		return err
	}*/
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
func (ms *mongoDataService) DeleteMulti(ctx core.Context, ids []string, softdelete bool) error {
	/*if softdelete {
		err := ms.UpdateMulti(ctx, objectType, ids, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		}
		return err
	}*/
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
func (ms *mongoDataService) DeleteAll(ctx core.Context, queryCond interface{}, softdelete bool) ([]string, error) {
	/*if softdelete {
		ids, err := ms.UpdateAll(ctx, objectType, queryCond, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		}
		return ids, err
	}*/
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
		ids[i] = resultStor[i].GetId()
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
