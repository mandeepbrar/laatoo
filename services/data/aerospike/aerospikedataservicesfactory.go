package aerospike

import (
	"laatoo/framework/core/objects"
	"laatoo/framework/services/data/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"

	"gopkg.in/mgo.v2"
	//"laatoosdk/log"
)

type aerospikeDataServicesFactory struct {
	connection *mgo.Session
	database   string
}

const (
	CONF_AEROSPIKE_CONNECTIONSTRING = "connectionstring"
	CONF_AEROSPIKE_SERVICES         = "AEROSPIKE_services"
)

func init() {
	objects.RegisterObject(CONF_AEROSPIKE_SERVICES, createAerospikeDataServicesFactory, nil)
}

func createAerospikeDataServicesFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &aerospikeDataServicesFactory{}, nil
}

func (mf *aerospikeDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	connectionString, ok := conf.GetString(CONF_AEROSPIKE_CONNECTIONSTRING)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_AEROSPIKE_CONNECTIONSTRING)
	}
	database, ok := conf.GetString(CONF_AEROSPIKE_DATABASE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", CONF_AEROSPIKE_DATABASE)
	}
	sess, err := mgo.Dial(connectionString)
	if err != nil {
		return errors.RethrowError(ctx, common.DATA_ERROR_CONNECTION, err, "Connection String", connectionString)
	}

	/*aerospikeSvc.objects = make(map[string]string, len(objs))
	for obj, collection := range objs {

		aerospikeSvc.objects[obj] = collection.(string)
	}
	deleteOps, _, _, _, err := buildRefOps(ctx, conf)
	if err != nil {
		return nil, err
	}
	aerospikeSvc.deleteRefOpers = deleteOps
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Aerospike service configured for objects ", "Objects", aerospikeSvc.objects)*/
	mf.connection = sess
	mf.database = database
	return nil
}

//Create the services configured for factory.
func (ms *aerospikeDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case common.CONF_DATA_SVCS:
		{

			return newAerospikeDataService(ctx, name, ms)

		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (ms *aerospikeDataServicesFactory) Start(ctx core.ServerContext) error {
	return nil
}
