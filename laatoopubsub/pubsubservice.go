package laatoopubsub

import (
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

type RedisPubSubService struct {
	context service.ServiceContext
}

const (
	CONF_REDISPUBSUB_NAME = "redis_pubsub"
	CONF_MONGO_DATABASE   = "database"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_REDISPUBSUB_NAME, RedisPubSubServiceFactory)
}

func RedisPubSubServiceFactory(conf map[string]interface{}) (interface{}, error) {
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
	mongoSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		mongoSvc.objects[obj] = collection.(string)
	}
	log.Logger.Debugf("Mongo service configured for objects ", mongoSvc.objects)
	return mongoSvc, nil
}

func (ms *RedisPubSubService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ms *RedisPubSubService) GetServiceType() string {
	return service.SERVICE_TYPE_DATA
}

//name of the service
func (svc *RedisPubSubService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *RedisPubSubService) Initialize(ctx service.ServiceContext) error {
	svc.context = ctx
	return nil
}

//The service starts serving when this method is called
func (svc *RedisPubSubService) Serve() error {
	return nil
}

func (svc *RedisPubSubService) PublishMessage(topic string, message interface{}) error {
	return nil
}
func (svc *RedisPubSubService) Subscribe(topic string, lstnr TopicListener) {

}

//Execute method
func (svc *RedisPubSubService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
