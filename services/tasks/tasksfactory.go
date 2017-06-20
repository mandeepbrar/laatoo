package tasks

/*
import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	//	"laatoo/core/common"
	//	"laatoo/sdk/errors"
)


type tasksFactory struct {
}

const (
	CONF_TASKS_FACTORY            = "tasksfactory"
	CONF_TASKS_BEANSTALK_PRODUCER = "beanstalktaskpublisher"
	CONF_TASKS_BEANSTALK_CONSUMER = "beanstalktaskprocessor"
	CONF_TASKS_GAE_PRODUCER       = "gaetaskpublisher"
	CONF_TASKS_GAE_CONSUMER       = "gaetaskprocessor"
)

func init() {
	objects.Register(CONF_TASKS_FACTORY, tasksFactory{})
}

func createTaskServiceFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &tasksFactory{}, nil
}

//Create the services configured for factory.
func (tf *tasksFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case CONF_TASKS_BEANSTALK_PRODUCER:
		return &beanstalkProducer{}, nil
	case CONF_TASKS_BEANSTALK_CONSUMER:
		return &beanstalkConsumer{}, nil
	case CONF_TASKS_GAE_PRODUCER:
		return &gaeProducer{}, nil
	case CONF_TASKS_GAE_CONSUMER:
		return &gaeConsumer{queues: make(map[string]*taskQueue, 10)}, nil
	}
	return nil, nil
}

func (tf *tasksFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (tf *tasksFactory) Start(ctx core.ServerContext) error {
	return nil
}
*/