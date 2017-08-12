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
	MONGO_FACTORY               = "mongo_services"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: MONGO_FACTORY, Object: mongoDataServicesFactory{}}}
}

/*
func (ms *mongoDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	ms.AddStringConfiguration(ctx, CONF_MONGO_CONNECTIONSTRING)
	ms.AddStringConfiguration(ctx, CONF_MONGO_DATABASE)

	return nil
}*/

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
