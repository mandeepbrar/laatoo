package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONF_MONGO_CONNECTIONSTRING = "mongoconnectionstring"
)

type MongoDataServicesFactory struct {
	core.ServiceFactory
	connection *mongo.Client
	database   string
}

/*
func (ms *mongoDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	ms.AddStringConfiguration(ctx, CONF_MONGO_CONNECTIONSTRING)
	ms.AddStringConfiguration(ctx, CONF_MONGO_DATABASE)

	return nil
}*/

//Create the services configured for factory.
func (ms *MongoDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newMongoDataService(ctx, name, ms)
}

//"mongodb://localhost:27017"
//The services start serving when this method is called
func (ms *MongoDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	connectionString, _ := ms.GetStringConfiguration(ctx, CONF_MONGO_CONNECTIONSTRING)
	database, _ := ms.GetStringConfiguration(ctx, CONF_MONGO_DATABASE)

	regbuilder := bson.NewRegistryBuilder()
	timeenc, err := bson.DefaultRegistry.LookupEncoder(reflect.TypeOf(time.Now()))
	if err != nil {
		return err
	}
	timedec, err := bson.DefaultRegistry.LookupDecoder(reflect.TypeOf(time.Now()))
	if err != nil {
		return err
	}
	srl := &StorableSerializer{bson.DefaultRegistry, ctx, timeenc, timedec}
	regbuilder = regbuilder.RegisterDefaultEncoder(reflect.Struct, srl).RegisterDefaultDecoder(reflect.Struct, srl)
	regbuilder = regbuilder.RegisterDefaultEncoder(reflect.Ptr, srl).RegisterDefaultDecoder(reflect.Ptr, srl)
	//regbuilder.RegisterEncoder(reflect.TypeOf(val), srl)
	//regbuilder.RegisterDecoder(reflect.TypeOf(val), srl)
	reg := regbuilder.Build()
	//srl.dc =

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString).SetRegistry(reg))
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

func (ms *MongoDataServicesFactory) getConnection(ctx core.RequestContext) *mongo.Client {
	return ms.connection
}
