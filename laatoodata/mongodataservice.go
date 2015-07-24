package datastores

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/service"
)

type mongoDataService struct {
	name       string
	connection *mgo.Session
	collection string
	database   string
	object     string
	context    service.ServiceContext
}

const (
	CONF_MONGO_CONNECTIONSTRING = "connectionstring"
	CONF_MONGO_DATABASE         = "database"
	CONF_MONGO_COLLECTION       = "collection"
	CONF_MONGO_OBJECT           = "object"
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
	collectionInt, ok := conf[CONF_MONGO_COLLECTION]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_COLLECTION)
	}
	objectInt, ok := conf[CONF_MONGO_OBJECT]
	if !ok {
		return nil, errors.ThrowError(DATA_ERROR_MISSING_OBJECT)
	}
	mongoSvc := &mongoDataService{name: "Mongo Data Service", collection: collectionInt.(string), object: objectInt.(string), connection: sess, database: databaseInt.(string)}
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

func (ms *mongoDataService) Put(id string, item interface{}) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Insert(item)
	if err != nil {
		return err
	}
	return nil
}

func (ms *mongoDataService) GetById(id string) (interface{}, error) {
	object, err := ms.context.CreateObject(ms.object, nil)
	if err != nil {
		return nil, err
	}
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	stor := object.(data.Storable)
	idkey := stor.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	err = connCopy.DB(ms.database).C(ms.collection).Find(condition).One(object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (ms *mongoDataService) Get(conditions interface{}) (interface{}, error) {
	return nil, errors.ThrowError(DATA_ERROR_NOT_IMPLEMENTED)
}

func (ms *mongoDataService) Delete(id string) error {
	return nil
}

func (ms *mongoDataService) GetList() (interface{}, error) {
	return nil, nil
}
