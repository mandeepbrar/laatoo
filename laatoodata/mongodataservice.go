package laatoodata

import (
	"github.com/labstack/echo"
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

func MongoServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
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
func (svc *mongoDataService) Initialize(ctx *echo.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *mongoDataService) Serve(ctx *echo.Context) error {
	return nil
}

func (ms *mongoDataService) Save(ctx *echo.Context, objectType string, item interface{}) error {
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saving object", "Object", objectType)
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
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
	err := connCopy.DB(ms.database).C(collection).Insert(item)
	if err != nil {
		return err
	}
	stor.PostSave(ctx)
	return nil
}

func (ms *mongoDataService) PutMulti(ctx *echo.Context, objectType string, ids []string, items interface{}) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Saving multiple objects", "ObjectType", objectType)
	arr := reflect.ValueOf(items)
	length := arr.Len()
	bulk := connCopy.DB(ms.database).C(collection).Bulk()
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor := valPtr.(data.Storable)
		stor.PreSave(ctx)
		bulk.Upsert(bson.M{stor.GetIdField(): stor.GetId()}, stor)
	}

	r, err := bulk.Run()
	if err != nil {
		return err
	}
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor := valPtr.(data.Storable)
		stor.PostSave(ctx)
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved multiple objects", "r", r, "err", err)
	return nil
}

func (ms *mongoDataService) Put(ctx *echo.Context, objectType string, id string, item interface{}) error {
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
	stor.PreSave(ctx)
	err := connCopy.DB(ms.database).C(collection).Update(condition, item)
	if err != nil {
		return err
	}
	stor.PostSave(ctx)
	return nil
}

func (ms *mongoDataService) GetById(ctx *echo.Context, objectType string, id string) (interface{}, error) {
	object, err := laatoocore.CreateObject(ctx, objectType, nil)
	if err != nil {
		return nil, err
	}
	collection, ok := ms.objects[objectType]
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Got the object ", "objectType", objectType)

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
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Got the object ", "condition", condition)
	stor.PostLoad(ctx)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error in getting object with", "ID", id, "Error", err)
		return nil, err
	}
	return object, nil
}

//Get multiple objects by id
func (ms *mongoDataService) GetMulti(ctx *echo.Context, objectType string, ids []string) (map[string]interface{}, error) {
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(ctx, DATA_ERROR_MISSING_COLLECTION, "ObjectType", objectType)
	}
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, err
	}
	object, err := laatoocore.CreateObject(ctx, objectType, nil)
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	stor := object.(data.Storable)
	idkey := stor.GetIdField()
	condition := bson.M{}
	operatorCond := bson.M{}
	operatorCond["$in"] = ids
	condition[idkey] = operatorCond
	err = connCopy.DB(ms.database).C(collection).Find(condition).All(results)
	if err != nil {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error in getting multiple objects", "ids", ids, "Error", err)
		return nil, err
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	retVal := make(map[string]interface{}, len(ids))
	arr := reflect.ValueOf(results).Elem()
	length := arr.Len()
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor := valPtr.(data.Storable)
		retVal[stor.GetId()] = valPtr
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

func (ms *mongoDataService) Get(ctx *echo.Context, objectType string, queryCond interface{}, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
	totalrecs = -1
	recsreturned = -1
	results, err := laatoocore.CreateCollection(ctx, objectType)
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Got the object ", "Result", results, "objectType", objectType)
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
		stor.PostLoad(ctx)
	}
	recsreturned = i
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Returning multiple objects ", "conditions", conditions, "objectType", objectType, "recsreturned", recsreturned)
	return results, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) Delete(ctx *echo.Context, objectType string, id string) error {
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

func (ms *mongoDataService) GetList(ctx *echo.Context, objectType string, pageSize int, pageNum int, mode string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error) {
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
		stor.PostLoad(ctx)
	}
	recsreturned = i
	if err != nil {
		return nil, totalrecs, recsreturned, err
	}
	return results, totalrecs, recsreturned, nil
}

//Execute method
func (svc *mongoDataService) Execute(ctx *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
