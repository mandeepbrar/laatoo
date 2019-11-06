package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDataServicesFactory struct {
	core.ServiceFactory
	connection *mongo.Client
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

//"mongodb://localhost:27017"
//The services start serving when this method is called
func (ms *mongoDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	connectionString, _ := ms.GetStringConfiguration(ctx, CONF_MONGO_CONNECTIONSTRING)
	database, _ := ms.GetStringConfiguration(ctx, CONF_MONGO_DATABASE)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return errors.RethrowError(ctx, data.DATA_ERROR_CONNECTION, err, "Connection String", connectionString)
	}

	timeoutctx, _ := ctx.WithTimeout(5 * time.Second)
	/*	err = client.Ping(timeoutctx, readpref.Primary())
		if err != nil {
			return errors.WrapError(ctx, err)
		}*/
	err = client.Connect(timeoutctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Info(timeoutctx, "Connection established to mongo database", "connectionString", connectionString)
	ms.connection = client
	ms.database = database
	return nil
}

func (ms *mongoDataServicesFactory) getConnection(ctx core.RequestContext) *mongo.Client {
	return ms.connection
}
