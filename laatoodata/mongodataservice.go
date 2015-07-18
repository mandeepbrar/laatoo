package datastores

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"laatoocore"
	"laatoosdk/data"
)

type mongoDataService struct {
	name       string
	alias      string
	connection *mgo.Session
	collection string
	database   string
}

const (
	CONF_MONGO_CONNECTIONSTRING = "connectionstring"
	CONF_MONGO_DATABASE         = "database"
	CONF_MONGO_COLLECTION       = "collection"
	CONF_MONGO_SERVICENAME      = "mongo_data_service"
)

func init() {
	laatoocore.RegisterServiceProvider(CONF_MONGO_SERVICENAME, MongoServiceFactory)
}

func MongoServiceFactory(alias string, conf map[string]interface{}) (interface{}, error) {
	connectionStringInt, ok := conf[CONF_MONGO_CONNECTIONSTRING]
	if !ok {
		return nil, fmt.Errorf("Connection string not found for mongo service %s", alias)
	}
	sess, err := mgo.Dial(connectionStringInt.(string))
	if err != nil {
		return nil, fmt.Errorf("Alias %s: Connection not connect to the mongo server %s", alias, err)
	}
	databaseInt, ok := conf[CONF_MONGO_DATABASE]
	if !ok {
		return nil, fmt.Errorf("Database could not be found for mongo service %s", alias)
	}
	collectionInt, ok := conf[CONF_MONGO_COLLECTION]
	if !ok {
		return nil, fmt.Errorf("Collection could not be found for mongo service %s", alias)
	}
	mongoSvc := &mongoDataService{name: "Mongo Data Service", alias: alias, collection: collectionInt.(string), connection: sess, database: databaseInt.(string)}
	return mongoSvc, nil
}

func (ms *mongoDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

//name of the service
func (svc *mongoDataService) GetName() string {
	return svc.name
}

//alias by which it is used
func (svc *mongoDataService) GetAlias() string {
	return svc.alias
}

//Initialize the service. Consumer of a service passes the data
func (svc *mongoDataService) Initialize(ctx interface{}) error {
	/*env := ctx.(*laatoo.Environment)
	svc.Environment = env*/
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

func (ms *mongoDataService) GetById(id string, object data.Storable) error {
	connCopy := ms.connection.Copy()
	defer connCopy.Close()
	idkey := object.GetIdField()
	condition := bson.M{}
	condition[idkey] = id
	err := connCopy.DB(ms.database).C(ms.collection).Find(condition).One(object)
	if err != nil {
		return err
	}
	return nil
}

func (ms *mongoDataService) Get(conditions interface{}) (interface{}, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (ms *mongoDataService) Delete(id string) error {
	return nil
}

func (ms *mongoDataService) GetList() (interface{}, error) {
	return nil, nil
}
