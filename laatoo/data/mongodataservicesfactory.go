package data

import (
	"gopkg.in/mgo.v2"
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//"laatoosdk/log"
)

type mongoDataServicesFactory struct {
	connection *mgo.Session
	database   string
}

const (
	CONF_MONGO_CONNECTIONSTRING = "connectionstring"
	CONF_MONGO_SERVICES         = "mongo_services"
	CONF_MONGO_DATA_SVCS        = "dataservices"
)

func init() {
	objects.RegisterObject(CONF_MONGO_SERVICES, createMongoDataServicesFactory, nil)
}

func createMongoDataServicesFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &mongoDataServicesFactory{}, nil
}

func (mf *mongoDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	connectionString, ok := conf.GetString(CONF_MONGO_CONNECTIONSTRING)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_CONNECTIONSTRING)
	}
	database, ok := conf.GetString(CONF_MONGO_DATABASE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_MONGO_DATABASE)
	}
	sess, err := mgo.Dial(connectionString)
	if err != nil {
		return errors.RethrowError(ctx, DATA_ERROR_CONNECTION, err, "Connection String", connectionString)
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
	mf.connection = sess
	mf.database = database
	return nil
}

//Create the services configured for factory.
func (ms *mongoDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	switch method {
	case CONF_MONGO_DATA_SVCS:
		{

			return newMongoDataService(ctx, name, ms)

		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (ms *mongoDataServicesFactory) Start(ctx core.ServerContext) error {
	return nil
}
