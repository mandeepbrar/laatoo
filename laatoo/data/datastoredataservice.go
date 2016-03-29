// +build appengine

package data

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"reflect"
)

type DatastoreDataService struct {
	name           string
	objects        map[string]string
	deleteRefOpers map[string][]*refKeyOperation
	getRefOpers    map[string][]*refKeyOperation
	putRefOpers    map[string][]*refKeyOperation
	updateRefOpers map[string][]*refKeyOperation
}

type DatastoreCondition struct {
	operation data.ConditionType
	arg1      interface{}
	arg2      interface{}
	arg3      interface{}
}

const (
	CONF_DATASTORE_OBJECTS     = "objects"
	CONF_DATASTORE_SERVICENAME = "datastore_data_service"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_DATASTORE_SERVICENAME, DatastoreServiceFactory)
}

func DatastoreServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	objectsInt, ok := conf[CONF_DATASTORE_OBJECTS]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_OBJECTS)
	}
	objs, ok := objectsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_OBJECTS)
	}

	datastoreSvc := &DatastoreDataService{name: "Datastore Data Service"}
	datastoreSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		datastoreSvc.objects[obj] = collection.(string)
	}
	deleteOps, _, _, _, err := buildRefOps(ctx, conf)
	if err != nil {
		return nil, err
	}
	datastoreSvc.deleteRefOpers = deleteOps
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Datastore service configured for objects ", "Objects", datastoreSvc.objects)
	return datastoreSvc, nil
}

func (ds *DatastoreDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ds *DatastoreDataService) GetServiceType() string {
	return core.SERVICE_TYPE_DATA
}

//name of the service
func (svc *DatastoreDataService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *DatastoreDataService) Initialize(ctx core.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *DatastoreDataService) Serve(ctx core.Context) error {
	return nil
}

func (ms *DatastoreDataService) Supports(feature data.Feature) bool {
	switch feature {
	case data.InQueries:
		return false
	case data.Ancestors:
		return true
	}
	return false
}

func (ms *DatastoreDataService) Save(ctx core.Context, objectType string, item interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Saving object", "ObjectType", objectType, "Item", item)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave(ctx)
	id := stor.GetId()
	if id == "" {
		return errors.ThrowError(ctx, DATA_ERROR_ID_NOT_FOUND, "ObjectType", objectType)
	}
	key, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, collection, stor.GetId(), 0, nil), item)
	if err != nil {
		return err
	}
	stor.PostSave(ctx)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved with key", "ObjectType", objectType, "Key", key)
	return nil
}

func (ms *DatastoreDataService) Put(ctx core.Context, objectType string, id string, item interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Saving object", "ObjectType", objectType, "Item", item)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave(ctx)
	key, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, collection, id, 0, nil), item)
	if err != nil {
		return err
	}
	stor.PostSave(ctx)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved with key", "ObjectType", objectType, "Key", key)
	return nil
}

func (ms *DatastoreDataService) PutMulti(ctx core.Context, objectType string, ids []string, items interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Saving multiple objects", "ObjectType", objectType)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
		keys[ind] = key
	}
	arr := reflect.ValueOf(items)
	length := arr.Len()
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor := valPtr.(data.Storable)
		stor.PreSave(ctx)
	}
	_, err := datastore.PutMulti(appEngineContext, keys, items)
	if err != nil {
		return err
	}
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor := valPtr.(data.Storable)
		stor.PostSave(ctx)
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved multiple objects", "ObjectType", objectType)
	return nil
}

func (ms *DatastoreDataService) GetById(ctx core.Context, objectType string, id string) (interface{}, error) {
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	object, err := laatoocore.CreateObject(ctx, objectType, nil)
	if err != nil {
		return nil, err
	}
	key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
	err = datastore.Get(appEngineContext, key, object)
	stor := object.(data.Storable)
	stor.PostLoad(ctx)
	if err != nil {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error in getting object", "ID", id, "Error", err)
		return nil, err
	}
	return object, nil
}

//Get multiple objects by id
func (ms *DatastoreDataService) GetMulti(ctx core.Context, objectType string, ids []string, orderBy string) (map[string]interface{}, error) {
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	ctype, err := laatoocore.GetCollectionType(ctx, objectType)
	if err != nil {
		return nil, err
	}
	lenids := len(ids)
	arr := reflect.MakeSlice(ctype, lenids, lenids)
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
		keys[ind] = key
	}
	/*if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}*/

	err = datastore.GetMulti(appEngineContext, keys, arr.Interface())
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Logger.Debug(ctx, LOGGING_CONTEXT, "Geting object", "err", err)
			return nil, err
		}
	}
	retVal := make(map[string]interface{}, lenids)
	length := arr.Len()
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor := valPtr.(data.Storable)
		id := stor.GetId()
		if id != "" {
			retVal[ids[i]] = valPtr
			stor.PostLoad(ctx)
		} else {
			retVal[ids[i]] = nil
		}
	}
	return retVal, nil
}

func (ms *DatastoreDataService) Get(ctx core.Context, objectType string, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	appEngineContext := ctx.GetAppengineContext()
	totalrecs = -1
	recsreturned = -1
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Get Resuls with condition", "Object Type", objectType, "QueryCond", queryCond)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	query := datastore.NewQuery(collection)
	query, err = ms.processCondition(ctx, query, queryCond)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	if pageSize > 0 {
		totalrecs, err = query.Limit(500).Count(appEngineContext)
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}
	if len(orderBy) > 0 {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Order query by", "orderBy", orderBy)
		query = query.Order(orderBy)
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	_, err = query.GetAll(appEngineContext, results)
	arr := reflect.ValueOf(results).Elem()
	length := arr.Len()
	i := 0
	for i = 0; i < length; i++ {
		stor := arr.Index(i).Addr().Interface().(data.Storable)
		stor.PostLoad(ctx)
	}
	recsreturned = i
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	return results, totalrecs, recsreturned, nil
}

//create condition for passing to data service
func (ms *DatastoreDataService) CreateCondition(ctx core.Context, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	switch operation {
	case data.MATCHANCESTOR:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			collection, ok := ms.objects[args[0].(string)]
			if !ok {
				return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
			}
			return &DatastoreCondition{operation: operation, arg1: collection, arg2: args[1]}, nil
		}
	case data.MATCHMULTIPLEVALUES:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return &DatastoreCondition{operation: operation, arg1: args[0], arg2: args[1]}, nil
		}
	case data.FIELDVALUE:
		{
			if len(args) < 1 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return &DatastoreCondition{operation: operation, arg1: args[0]}, nil
		}
	}
	return nil, nil
}

func (ms *DatastoreDataService) processCondition(ctx core.Context, query *datastore.Query, condition interface{}) (*datastore.Query, error) {
	if condition == nil {
		return query, nil
	}
	dqCondition := condition.(*DatastoreCondition)
	switch dqCondition.operation {
	case data.MATCHANCESTOR:
		collection, ok := dqCondition.arg1.(string)
		id, ok := dqCondition.arg2.(string)
		if id == "" {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
		}
		key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
		query = query.Ancestor(key), nil
	case data.MATCHMULTIPLEVALUES:
		return nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_NOT_IMPLEMENTED)
	case data.FIELDVALUE:
		queryCondMap, ok := dqCondition.arg1.(map[string]interface{})
		if ok {
			for k, v := range queryCondMap {
				query = query.Filter(fmt.Sprintf("%s =", k), v)
			}
			return query, nil
		} else {
			return nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_TYPE_MISMATCH)
		}

	}
	return query, nil
}

func (ms *DatastoreDataService) Update(ctx core.Context, objectType string, id string, newVals map[string]interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	object, err := laatoocore.CreateObject(ctx, objectType, nil)
	if err != nil {
		return err
	}
	key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
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
	return nil
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *DatastoreDataService) UpdateMulti(ctx core.Context, objectType string, ids []string, newVals map[string]interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	ctype, err := laatoocore.GetCollectionType(ctx, objectType)
	if err != nil {
		return err
	}
	lenids := len(ids)
	arr := reflect.MakeSlice(ctype, lenids, lenids)
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
		keys[ind] = key
	}
	err = datastore.GetMulti(appEngineContext, keys, arr.Interface())
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Logger.Debug(ctx, LOGGING_CONTEXT, "Geting object", "err", err)
			return err
		}
	}
	length := arr.Len()
	for i := 0; i < length; i++ {
		val := arr.Index(i)
		for k, v := range newVals {
			f := val.FieldByName(k)
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
	_, err = datastore.PutMulti(appEngineContext, keys, arr.Interface())
	if err != nil {
		return err
	}
	return nil
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *DatastoreDataService) UpdateAll(ctx core.Context, objectType string, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, err
	}
	query := datastore.NewQuery(collection)
	query, err = ms.processCondition(ctx, query, queryCond)
	if err != nil {
		return nil, err
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	keys, err := query.GetAll(appEngineContext, results)
	arr := reflect.ValueOf(results).Elem()
	length := arr.Len()
	ids := make([]string, length)
	for i := 0; i < length; i++ {
		val := arr.Index(i)
		stor := val.Addr().Interface().(data.Storable)
		ids[i] = stor.GetId()
		for k, v := range newVals {
			f := val.FieldByName(k)
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
	_, err = datastore.PutMulti(appEngineContext, keys, arr.Interface())
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (ms *DatastoreDataService) Delete(ctx core.Context, objectType string, id string, softdelete bool) error {
	if softdelete {
		err := ms.Update(ctx, objectType, id, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, []string{id})
		}
		return err
	}
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
	err := datastore.Delete(appEngineContext, key)
	if err == nil {
		err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, []string{id})
	}
	return err
}

//Delete object by ids
func (ms *DatastoreDataService) DeleteMulti(ctx core.Context, objectType string, ids []string, softdelete bool) error {
	if softdelete {
		err := ms.UpdateMulti(ctx, objectType, ids, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		}
		return err
	}
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	keys := make([]*datastore.Key, len(ids))
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
		keys[ind] = key
	}
	err := datastore.DeleteMulti(appEngineContext, keys)
	if err == nil {
		err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
	}
	return err
}

//Delete object by condition
func (ms *DatastoreDataService) DeleteAll(ctx core.Context, objectType string, queryCond interface{}, softdelete bool) ([]string, error) {
	if softdelete {
		ids, err := ms.UpdateAll(ctx, objectType, queryCond, map[string]interface{}{"Deleted": true})
		if err == nil {
			err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
		}
		return ids, err
	}
	appEngineContext := ctx.GetAppengineContext()
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, err
	}
	query := datastore.NewQuery(collection)
	query, err = ms.processCondition(ctx, query, queryCond)
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
	if err == nil {
		err = deleteRefOps(ctx, ms, ms.deleteRefOpers, objectType, ids)
	}
	return ids, err
}

func (ms *DatastoreDataService) GetList(ctx core.Context, objectType string, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	appEngineContext := ctx.GetAppengineContext()
	totalrecs = -1
	recsreturned = -1
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	query := datastore.NewQuery(collection)
	if pageSize > 0 {
		totalrecs, err = query.Count(appEngineContext)
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}
	if len(orderBy) > 0 {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Order query by", "orderBy", orderBy)
		query = query.Order(orderBy)
	}

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	_, err = query.GetAll(appEngineContext, results)
	arr := reflect.ValueOf(results).Elem()
	length := arr.Len()
	i := 0
	for i = 0; i < length; i++ {
		stor := arr.Index(i).Addr().Interface().(data.Storable)
		stor.PostLoad(ctx)
	}
	recsreturned = i
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	return results, totalrecs, recsreturned, nil
}

//Execute method
func (svc *DatastoreDataService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
