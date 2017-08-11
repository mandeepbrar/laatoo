package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	//"log"

	"gopkg.in/mgo.v2"
)

type mongoDataServicesFactory struct {
	core.ServiceFactory
	connection *mgo.Session
	database   string
}

const (
	CONF_MONGO_CONNECTIONSTRING = "mongoconnectionstring"
	CONF_MONGO_SERVICES         = "mongo_services"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_MONGO_SERVICES, Object: mongoDataServicesFactory{}}}
}

func (ms *mongoDataServicesFactory) Initialize(ctx core.ServerContext) error {
	ms.AddStringConfiguration(ctx, CONF_MONGO_CONNECTIONSTRING)
	ms.AddStringConfiguration(ctx, CONF_MONGO_DATABASE)
	/*mongoSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		mongoSvc.objects[obj] = collection.(string)
	}
	deleteOps, _, _, _, err := buildRefOps(ctx, conf)
	if err != nil {
		return nil, err
	}
	mongoSvc.deleteRefOpers = deleteOps
	log.Debug(ctx, LOGGING_CONTEXT, "Mongo service configured for objects ", "Objects", mongoSvc.objects)*/
	return nil
}

//Create the services configured for factory.
func (ms *mongoDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newMongoDataService(ctx, name, ms)
}

//The services start serving when this method is called
func (ms *mongoDataServicesFactory) Start(ctx core.ServerContext) error {
	connectionString, _ := ms.GetStringConfiguration(ctx, CONF_MONGO_CONNECTIONSTRING)
	database, _ := ms.GetStringConfiguration(ctx, CONF_MONGO_DATABASE)

	sess, err := mgo.Dial(connectionString)
	if err != nil {
		return errors.RethrowError(ctx, data.DATA_ERROR_CONNECTION, err, "Connection String", connectionString)
	}
	log.Info(ctx, "Connection established to mongo database", "connectionString", connectionString)
	ms.connection = sess
	ms.database = database
	return nil
}
