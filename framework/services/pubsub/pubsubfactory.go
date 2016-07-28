package pubsub

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type PubSubFactory struct {
}

const (
	CONF_PUBSUB_NAME     = "pubsub"
	CONF_REDISPUBSUB_SVC = "redis"
	CONF_APPPUBSUB_SVC   = "apppubsub"
)

func init() {
	objects.Register(CONF_PUBSUB_NAME, PubSubFactory{})
}

//Create the services configured for factory.
func (mf *PubSubFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case CONF_REDISPUBSUB_SVC:
		{
			return &RedisPubSubService{name: name}, nil
		}
	case CONF_APPPUBSUB_SVC:
		{
			return &RedisPubSubService{name: name}, nil
		}
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
