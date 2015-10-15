// +build appengine

package laatoodata

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"laatoocore"
	"laatoosdk/context"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"reflect"
)

type DatastoreDataService struct {
	name    string
	objects map[string]string
	context service.ServiceContext
}

const (
	CONF_DATASTORE_OBJECTS     = "objects"
	CONF_DATASTORE_SERVICENAME = "datastore_data_service"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_DATASTORE_SERVICENAME, DatastoreServiceFactory)
}

func DatastoreServiceFactory(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
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
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Datastore service configured for objects ", "Objects", datastoreSvc.objects)
	return datastoreSvc, nil
}

func (ds *DatastoreDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ds *DatastoreDataService) GetServiceType() string {
	return service.SERVICE_TYPE_DATA
}

//name of the service
func (svc *DatastoreDataService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *DatastoreDataService) Initialize(ctx service.ServiceContext) error {
	svc.context = ctx
	return nil
}

//The service starts serving when this method is called
func (svc *DatastoreDataService) Serve(ctx interface{}) error {
	return nil
}

func (ms *DatastoreDataService) Save(ctx interface{}, objectType string, item interface{}) error {
	appEngineContext := context.GetAppengineContext(ctx)
	log.Logger.Debug(appEngineContext, LOGGING_CONTEXT, "Saving object", "ObjectType", objectType, "Item", item)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave(ctx)
	id := stor.GetId()
	if id == "" {
		return errors.ThrowError(appEngineContext, DATA_ERROR_ID_NOT_FOUND, "ObjectType", objectType)
	}
	key, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, collection, stor.GetId(), 0, nil), item)
	if err != nil {
		return err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved with key", "ObjectType", objectType, "Key", key)
	return nil
}

func (ms *DatastoreDataService) Put(ctx interface{}, objectType string, id string, item interface{}) error {
	appEngineContext := context.GetAppengineContext(ctx)
	log.Logger.Debug(appEngineContext, LOGGING_CONTEXT, "Saving object", "ObjectType", objectType, "Item", item)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave(ctx)
	key, err := datastore.Put(appEngineContext, datastore.NewKey(appEngineContext, collection, id, 0, nil), item)
	if err != nil {
		return err
	}
	log.Logger.Trace(appEngineContext, LOGGING_CONTEXT, "Saved with key", "ObjectType", objectType, "Key", key)
	return nil
}

func (ms *DatastoreDataService) GetById(ctx interface{}, objectType string, id string) (interface{}, error) {
	appEngineContext := context.GetAppengineContext(ctx)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	object, err := laatoocore.CreateObject(appEngineContext, objectType, nil)
	if err != nil {
		return nil, err
	}
	key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
	err = datastore.Get(appEngineContext, key, object)
	stor := object.(data.Storable)
	stor.PostLoad(ctx)
	if err != nil {
		log.Logger.Debug(appEngineContext, LOGGING_CONTEXT, "Error in getting object", "ID", id, "Error", err)
		return nil, err
	}
	log.Logger.Debug(appEngineContext, LOGGING_CONTEXT, "Got object", "Key", datastore.NewKey(appEngineContext, collection, "some id", 0, nil))
	return object, nil
}

//Get multiple objects by id
func (ms *DatastoreDataService) GetMulti(ctx interface{}, objectType string, ids []string) (map[string]interface{}, error) {
	appEngineContext := context.GetAppengineContext(ctx)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	ctype, err := laatoocore.GetCollectionType(appEngineContext, objectType)
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

	err = datastore.GetMulti(appEngineContext, keys, arr.Interface())
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Logger.Debug(appEngineContext, LOGGING_CONTEXT, "Geting object", "err", err)
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

func (ms *DatastoreDataService) Get(ctx interface{}, objectType string, queryCond interface{}, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	appEngineContext := context.GetAppengineContext(ctx)
	totalrecs = -1
	recsreturned = -1
	results, err := laatoocore.CreateCollection(appEngineContext, objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Trace(appEngineContext, LOGGING_CONTEXT, "Get Resuls with condition", "Object Type", objectType, "QueryCond", queryCond)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	query := datastore.NewQuery(collection)
	if queryCond != nil {
		queryCondMap, ok := queryCond.(map[string]interface{})
		if ok {
			for k, v := range queryCondMap {
				query = query.Filter(fmt.Sprintf("%s =", k), v)
				log.Logger.Trace(appEngineContext, LOGGING_CONTEXT, "Get Resuls with condition", "k", k, "v", v)

			}
		}
	}
	if pageSize > 0 {
		totalrecs, err = query.Limit(500).Count(appEngineContext)
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
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

func (ms *DatastoreDataService) Delete(ctx interface{}, objectType string, id string) error {
	appEngineContext := context.GetAppengineContext(ctx)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	key := datastore.NewKey(appEngineContext, collection, id, 0, nil)
	return datastore.Delete(appEngineContext, key)
}

func (ms *DatastoreDataService) GetList(ctx interface{}, objectType string, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	appEngineContext := context.GetAppengineContext(ctx)
	totalrecs = -1
	recsreturned = -1
	results, err := laatoocore.CreateCollection(appEngineContext, objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(appEngineContext, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
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
func (svc *DatastoreDataService) Execute(ctx interface{}, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
