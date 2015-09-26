package laatoodata

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"reflect"
)

type mongoDataService struct {
	name       string
	connection *mgo.Session
	database   string
	objects    map[string]string
	context    service.ServiceContext
}

const (
	CONF_MONGO_CONNECTIONSTRING = "connectionstring"
	CONF_MONGO_DATABASE         = "database"
	CONF_MONGO_OBJECTS          = "objects"
	CONF_MONGO_SERVICENAME      = "mongo_data_service"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_MONGO_SERVICENAME, MongoServiceFactory)
}

func MongoServiceFactory(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	connectionStringInt, ok := conf[CONF_MONGO_CONNECTIONSTRING]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_CONNECTION_STRING)
	}
	sess, err := mgo.Dial(connectionStringInt.(string))
	if err != nil {
		return nil, errors.RethrowError(ctx, DATA_ERROR_CONNECTION, err, "Connection String", connectionStringInt)
	}
	databaseInt, ok := conf[CONF_MONGO_DATABASE]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_DATABASE)
	}
	objectsInt, ok := conf[CONF_MONGO_OBJECTS]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_OBJECTS)
	}
	objs, ok := objectsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_OBJECTS)
	}

	mongoSvc := &mongoDataService{name: "Mongo Data Service", connection: sess, database: databaseInt.(string)}
	mongoSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		mongoSvc.objects[obj] = collection.(string)
	}
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Mongo service configured for objects ", "Objects", mongoSvc.objects)
	return mongoSvc, nil
}

func (ms *mongoDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ms *mongoDataService) GetServiceType() string {
	return service.SERVICE_TYPE_DATA
}

//name of the service
func (svc *mongoDataService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *mongoDataService) Initialize(ctx service.ServiceContext) error {
	svc.context = ctx
	return nil
}

//The service starts serving when this method is called
func (svc *mongoDataService) Serve(ctx interface{}) error {
	return nil
}

func (ms *mongoDataService) Save(ctx interface{}, objectType string, item interface{}) error {
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saving object", "Object", objectType, "Item", item)
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	stor := item.(data.Storable)
	stor.PreSave()
	id := stor.GetId()
	if id == "" {
		return errors.ThrowError(ctx, DATA_ERROR_ID_NOT_FOUND, "ObjectType", objectType)
	}
	err := connCopy.DB(ms.database).C(collection).Insert(item)
	if err != nil {
		return err
	}
	return nil
}

func (ms *mongoDataService) Put(ctx interface{}, objectType string, id string, item interface{}) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
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
	return nil
}

func (ms *mongoDataService) GetById(ctx interface{}, objectType string, id string) (interface{}, error) {
	object, err := laatoocore.CreateObject(ctx, objectType, nil)
	if err != nil {
		return nil, err
	}
	collection, ok := ms.objects[objectType]

	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	stor := object.(data.Storable)
	idkey := stor.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	err = connCopy.DB(ms.database).C(collection).Find(condition).One(object)
	stor.PostLoad()
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error in getting object with", "ID", id, "Error", err)
		return nil, err
	}
	return object, nil
}

func (ms *mongoDataService) Get(ctx interface{}, objectType string, queryCond interface{}, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	totalrecs = -1
	recsreturned = -1
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Got the object ", "Result", results)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, totalrecs, recsreturned, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	var conditions map[string]interface{}
	if queryCond != nil {
		conditions = queryCond.(map[string]interface{})
	} else {
		conditions = bson.M{}
	}
	query := connCopy.DB(ms.database).C(collection).Find(conditions)
	if pageSize > 0 {
		totalrecs, err = query.Count()
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Skip(recsToSkip)
	}
	err = query.All(results)
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

func (ms *mongoDataService) Delete(ctx interface{}, objectType string, id string) error {
	object, err := laatoocore.CreateObject(ctx, objectType, nil)
	if err != nil {
		return nil
	}
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	stor := object.(data.Storable)
	idkey := stor.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	return connCopy.DB(ms.database).C(collection).Remove(condition)
}

func (ms *mongoDataService) GetList(ctx interface{}, objectType string, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
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
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	query := connCopy.DB(ms.database).C(collection).Find(condition)
	if pageSize > 0 {
		totalrecs, err = query.Count()
		if err != nil {
			return nil, totalrecs, recsreturned, err
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Skip(recsToSkip)
	}
	err = query.All(results)
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
func (svc *mongoDataService) Execute(ctx interface{}, name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
