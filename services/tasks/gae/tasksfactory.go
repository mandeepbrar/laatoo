package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type TasksFactory struct {
}

const (
	CONF_GAE_TASKS_FACTORY  = "gaetasksfactory"
	CONF_TASKS_PRODUCER     = "publisher"
	CONF_TASKS_CONSUMER     = "consumer"
	CONF_TASKS_GAE_PRODUCER = "gaetaskpublisher"
	CONF_TASKS_GAE_CONSUMER = "gaetaskprocessor"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_TASKS_GAE_PRODUCER, Object: GaeProducer{}},
		core.PluginComponent{Name: CONF_GAE_TASKS_FACTORY, Object: core.NewFactory(func() interface{} { return &TasksFactory{} })},
		core.PluginComponent{Name: CONF_TASKS_GAE_CONSUMER, Object: GaeConsumer{}}}
}

func createTaskServiceFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &TasksFactory{}, nil
}

//Create the services configured for factory.
func (tf *TasksFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case CONF_TASKS_GAE_PRODUCER:
		return &GaeProducer{}, nil
	case CONF_TASKS_GAE_CONSUMER:
		return &GaeConsumer{queues: make(map[string]*taskQueue, 10)}, nil
	}
	return nil, nil
}

func (tf *TasksFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (tf *TasksFactory) Start(ctx core.ServerContext) error {
	return nil
}
