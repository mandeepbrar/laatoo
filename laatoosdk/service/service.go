package service

import (
	"github.com/labstack/echo"
	"laatoosdk/auth"
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
	Serve(ctx interface{}) error
	//Type of service
	GetServiceType() string
	//Execute method
	Execute(interface{}, string, map[string]interface{}) (map[string]interface{}, error)
}

//service context object for initializing services
type ServiceContext interface {
	GetVariable(variable string) interface{}
	GetService(ctx interface{}, alias string) (Service, error)
	RegisterPermissions(ctx interface{}, perm []string)
	ListAllPermissions() []string
	RegisterRoles(ctx interface{}, rolesInt interface{})
	RegisterRolePermissions(ctx interface{}, role auth.Role)
	IsAllowed(ctx *echo.Context, perm string) bool
	SubscribeTopic(ctx interface{}, topic string, handler TopicListener)
	PublishMessage(ctx interface{}, topic string, message interface{}) error
	GetConfig() config.Config
}

const (
	SERVICE_TYPE_APP  = "SERVICE_TYPE_APP"
	SERVICE_TYPE_DATA = "SERVICE_TYPE_DATA"
	SERVICE_TYPE_WEB  = "SERVICE_TYPE_WEB"
)
