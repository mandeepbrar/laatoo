// +build appengine

package laatoodata

import (
	"google.golang.org/appengine/datastore"
	"laatoocore"
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

func DatastoreServiceFactory(conf map[string]interface{}) (interface{}, error) {
	/*connectionStringInt, ok := conf[CONF_MONGO_CONNECTIONSTRING]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_CONNECTION_STRING)
	}
	sess, err := mgo.Dial(connectionStringInt.(string))
	if err != nil {
		return nil, errors.RethrowError(DATA_ERROR_CONNECTION, err)
	}
	databaseInt, ok := conf[CONF_MONGO_DATABASE]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_DATABASE)
	}*/
	objectsInt, ok := conf[CONF_DATASTORE_OBJECTS]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_OBJECTS)
	}
	objs, ok := objectsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_OBJECTS)
	}

	datastoreSvc := &DatastoreDataService{name: "Datastore Data Service"}
	datastoreSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		datastoreSvc.objects[obj] = collection.(string)
	}
	log.Logger.Debugf("Mongo service configured for objects ", datastoreSvc.objects)
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
func (svc *DatastoreDataService) Serve() error {
	return nil
}

func (ms *DatastoreDataService) Save(objectType string, item interface{}) error {
	log.Logger.Debugf("Saving object", objectType, item)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave()
	key, err := datastore.Put(laatoocore.APPENGINE_CONTEXT, datastore.NewKey(laatoocore.APPENGINE_CONTEXT, collection, stor.GetId(), 0, nil), item)
	if err != nil {
		return err
	}
	log.Logger.Errorf("saved with key", objectType, key)
	return nil
}

func (ms *DatastoreDataService) Put(objectType string, id string, item interface{}) error {
	log.Logger.Debugf("Saving object", objectType, item)
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave()
	key, err := datastore.Put(laatoocore.APPENGINE_CONTEXT, datastore.NewKey(laatoocore.APPENGINE_CONTEXT, collection, id, 0, nil), item)
	if err != nil {
		return err
	}
	log.Logger.Debugf("saved with key", objectType, key)
	return nil
	/*connCopy := ms.connection.Copy()
	defer connCopy.Close()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	condition := bson.M{}
	stor := item.(data.Storable)
	idkey := stor.GetIdField()
	condition[idkey] = id
	stor.PreSave()
	err := connCopy.DB(ms.database).C(collection).Update(condition, item)
	if err != nil {
		return err
	}
	return nil*/
}

func (ms *DatastoreDataService) GetById(objectType string, id string) (interface{}, error) {
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	object, err := ms.context.CreateObject(objectType, nil)
	if err != nil {
		return nil, err
	}
	key := datastore.NewKey(laatoocore.APPENGINE_CONTEXT, collection, id, 0, nil)
	err = datastore.Get(laatoocore.APPENGINE_CONTEXT, key, object)
	stor := object.(data.Storable)
	/*object, err := ms.context.CreateObject(objectType, nil)
	if err != nil {
		return nil, err
	}
	collection, ok := ms.objects[objectType]

	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	stor := object.(data.Storable)
	idkey := stor.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	err = connCopy.DB(ms.database).C(collection).Find(condition).One(object)*/
	stor.PostLoad()
	if err != nil {
		log.Logger.Debugf("Error in getting object with id %s %s", id, err)
		return nil, err
	}
	log.Logger.Debugf("Got the object with id %s", id)
	return object, nil
}

func (ms *DatastoreDataService) Get(objectType string, queryCond interface{}, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	totalrecs = -1
	recsreturned = -1
	results, err := ms.context.CreateCollection(objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Debugf("Got the object ", results)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	query := datastore.NewQuery(collection)
	if pageSize > 0 {
		totalrecs, err = query.Count(laatoocore.APPENGINE_CONTEXT)
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	_, err = query.GetAll(laatoocore.APPENGINE_CONTEXT, results)
	arr := reflect.ValueOf(results).Elem()
	length := arr.Len()
	i := 0
	for i = 0; i < length; i++ {
		stor := arr.Index(i).Addr().Interface().(data.Storable)
		stor.PostLoad()
	}
	recsreturned = i
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	return results, totalrecs, recsreturned, nil
}

func (ms *DatastoreDataService) Delete(objectType string, id string) error {
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	key := datastore.NewKey(laatoocore.APPENGINE_CONTEXT, collection, id, 0, nil)
	return datastore.Delete(laatoocore.APPENGINE_CONTEXT, key)
}

func (ms *DatastoreDataService) GetList(objectType string, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	totalrecs = -1
	recsreturned = -1
	results, err := ms.context.CreateCollection(objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Debugf("Got the object ", results)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION, objectType)
	}
	query := datastore.NewQuery(collection)
	if pageSize > 0 {
		totalrecs, err = query.Count(laatoocore.APPENGINE_CONTEXT)
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	_, err = query.GetAll(laatoocore.APPENGINE_CONTEXT, results)
	arr := reflect.ValueOf(results).Elem()
	length := arr.Len()
	i := 0
	for i = 0; i < length; i++ {
		stor := arr.Index(i).Addr().Interface().(data.Storable)
		stor.PostLoad()
	}
	recsreturned = i
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	return results, totalrecs, recsreturned, nil
}

//Execute method
func (svc *DatastoreDataService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
