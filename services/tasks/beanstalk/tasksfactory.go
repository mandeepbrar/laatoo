package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type TasksFactory struct {
	core.ServiceFactory
}

const (
	CONF_BEANSTALK_TASKS_FACTORY  = "beanstalkfactory"
	CONF_TASKS_PRODUCER           = "publisher"
	CONF_TASKS_CONSUMER           = "consumer"
	CONF_TASKS_BEANSTALK_PRODUCER = "beanstalktaskpublisher"
	CONF_TASKS_BEANSTALK_CONSUMER = "beanstalktaskconsumer"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_TASKS_BEANSTALK_PRODUCER, Object: BeanstalkProducer{}},
		core.PluginComponent{Name: CONF_BEANSTALK_TASKS_FACTORY, Object: TasksFactory{}},
		core.PluginComponent{Name: CONF_TASKS_BEANSTALK_CONSUMER, Object: BeanstalkConsumer{}}}
}

//Create the services configured for factory.
func (tf *TasksFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case CONF_TASKS_PRODUCER:
		return &BeanstalkProducer{}, nil
	case CONF_TASKS_CONSUMER:
		return &BeanstalkConsumer{}, nil
	}
	return nil, nil
}
