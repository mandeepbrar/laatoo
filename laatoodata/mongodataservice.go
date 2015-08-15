package datastores

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
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

func MongoServiceFactory(conf map[string]interface{}) (interface{}, error) {
	connectionStringInt, ok := conf[CONF_MONGO_CONNECTIONSTRING]
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
	}
	objectsInt, ok := conf[CONF_MONGO_OBJECTS]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_OBJECTS)
	}
	objs, ok := objectsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_OBJECTS)
	}

	mongoSvc := &mongoDataService{name: "Mongo Data Service", connection: sess, database: databaseInt.(string)}
	mongoSvc.objects = make(map[string]string, 50)
	for obj, collection := range objs {
		mongoSvc.objects[obj] = collection.(string)
	}

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
func (svc *mongoDataService) Serve() error {
	return nil
}

func (ms *mongoDataService) Save(objectType string, item interface{}) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(DATA_ERROR_MISSING_COLLECTION)
	}
	err := connCopy.DB(ms.database).C(collection).Insert(item)
	if err != nil {
		return err
	}
	return nil
}

func (ms *mongoDataService) Put(objectType string, id string, item interface{}) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	collection, ok := ms.objects[objectType]
	if !ok {
		return errors.ThrowError(DATA_ERROR_MISSING_COLLECTION)
	}
	condition := bson.M{}
	stor := item.(data.Storable)
	idkey := stor.GetIdField()
	condition[idkey] = id
	err := connCopy.DB(ms.database).C(collection).Update(condition, item)
	if err != nil {
		return err
	}
	return nil
}

func (ms *mongoDataService) GetById(objectType string, id string) (interface{}, error) {
	object, err := ms.context.CreateObject(objectType, nil)
	if err != nil {
		return nil, err
	}
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION)
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	stor := object.(data.Storable)
	idkey := stor.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	err = connCopy.DB(ms.database).C(collection).Find(condition).One(object)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		log.Logger.Debugf("Error in getting object with id %s %s", id, err)
		return nil, err
	}
	log.Logger.Debugf("Got the object with id %s", id)
	return object, nil
}

func (ms *mongoDataService) Get(objectType string, conditions interface{}) (interface{}, error) {
	results, err := ms.context.CreateCollection(objectType)
	if err != nil {
		return nil, err
	}
	log.Logger.Debugf("Got the object ", results)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION)
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	err = connCopy.DB(ms.database).C(collection).Find(condition).All(results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *mongoDataService) Delete(objectType string, id string) error {
	return nil
}

func (ms *mongoDataService) GetList(objectType string) (interface{}, error) {
	results, err := ms.context.CreateCollection(objectType)
	if err != nil {
		return nil, err
	}
	log.Logger.Debugf("Got the object ", results)
	collection, ok := ms.objects[objectType]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION)
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	err = connCopy.DB(ms.database).C(collection).Find(condition).All(results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

//Execute method
func (svc *mongoDataService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
