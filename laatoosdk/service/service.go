package service

import (
	"laatoosdk/config"
)

//Service interface that needs to be implemented by any service of a system
type Service interface {
	//Provides the name of the service
	GetName() string
	//Initialize the service. Consumer of a service passes the data
	Initialize(ctx ServiceContext) error
	//The service starts serving when this method is called
	//called on first request
	Serve() error
	//Type of service
	GetServiceType() string
	//Execute method
	Execute(string, map[string]interface{}) (map[string]interface{}, error)
}

//service context object for initializing services
type ServiceContext interface {
	GetService(alias string) (Service, error)
	GetAllServices() []interface{}
	CreateObject(objName string, confData map[string]interface{}) (interface{}, error)
	CreateEmptyObject(objName string) (interface{}, error)
	CreateCollection(objName string) (interface{}, error)
	SubscribeTopic(topic string, handler TopicListener)
	PublishMessage(topic string, message interface{}) error
	GetConfig() config.Config
}

const (
	SERVICE_TYPE_APP  = "SERVICE_TYPE_APP"
	SERVICE_TYPE_DATA = "SERVICE_TYPE_DATA"
	SERVICE_TYPE_WEB  = "SERVICE_TYPE_WEB"
)
