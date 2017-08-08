package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type TasksFactory struct {
	core.ServiceFactory
}

const (
	CONF_GAE_TASKS_FACTORY  = "gaetasksfactory"
	CONF_TASKS_PRODUCER     = "publisher"
	CONF_TASKS_CONSUMER     = "consumer"
	CONF_TASKS_GAE_PRODUCER = "gaetaskpublisher"
	CONF_TASKS_GAE_CONSUMER = "gaetaskconsumer"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_TASKS_GAE_PRODUCER, Object: GaeProducer{}},
		core.PluginComponent{Name: CONF_GAE_TASKS_FACTORY, Object: TasksFactory{}},
		core.PluginComponent{Name: CONF_TASKS_GAE_CONSUMER, Object: GaeConsumer{}}}
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
