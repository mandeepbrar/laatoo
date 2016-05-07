package laatoopubsub

import (
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type PubSubFactory struct {
}

const (
	CONF_PUBSUB_NAME     = "pubsub"
	CONF_REDISPUBSUB_SVC = "redis"
)

func init() {
	objects.RegisterObject(CONF_PUBSUB_NAME, createPubSubFactory, nil)
}

func createPubSubFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &PubSubFactory{}, nil
}

//Create the services configured for factory.
func (mf *PubSubFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	if method == CONF_REDISPUBSUB_SVC {
		return &RedisPubSubService{name: name}, nil
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *PubSubFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *PubSubFactory) Start(ctx core.ServerContext) error {
	return nil
}
