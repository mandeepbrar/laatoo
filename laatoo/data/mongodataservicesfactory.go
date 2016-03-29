package data

import (
	"gopkg.in/mgo.v2"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//"laatoosdk/log"
)

type mongoDataServicesFactory struct {
	connection *mgo.Session
}

const (
	CONF_MONGO_CONNECTIONSTRING = "connectionstring"
	CONF_MONGO_SERVICES         = "mongo_services"
	CONF_MONGO_DATA_SVCS        = "dataservices"
)

func init() {
	registry.RegisterServiceFactoryProvider(CONF_MONGO_SERVICES, MongoServicesFactory)
}

func MongoServicesFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	connectionString, ok := conf.GetString(CONF_MONGO_CONNECTIONSTRING)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_CONNECTIONSTRING)
	}
	sess, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, errors.RethrowError(ctx, DATA_ERROR_CONNECTION, err, "Connection String", connectionString)
	}

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
	mongoSvcFactory := &mongoDataServicesFactory{connection: sess}
	return mongoSvcFactory, nil
}

//Create the services configured for factory.
func (ms *mongoDataServicesFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	switch name {
	case CONF_MONGO_DATA_SVCS:
		{

			return newMongoDataService(ctx, ms, conf)

		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (ms *mongoDataServicesFactory) StartServices(ctx core.ServerContext) error {
	return nil
}
