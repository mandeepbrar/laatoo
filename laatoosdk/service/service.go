package service

import (
	"github.com/labstack/echo"
	"laatoosdk/config"
	"laatoosdk/data"
)

//Service interface that needs to be implemented by any service of a system
type Service interface {
	//Provides the name of the service
	GetName() string
	//Initialize the service. Consumer of a service passes the data
	Initialize(ctx *echo.Context) error
	//The service starts serving when this method is called
	//called on first request
	Serve(ctx *echo.Context) error
	//Type of service
	GetServiceType() string
	//Execute method
	Execute(*echo.Context, string, map[string]interface{}) (interface{}, error)
}

//service context object for initializing services
type Environment interface {
	GetVariable(variable string) interface{}
	GetService(ctx *echo.Context, alias string) (Service, error)
	GetCache() data.Cache
	//RegisterPermissions(ctx *echo.Context, perm []string)
	/*	ListAllPermissions() []string
		RegisterRoles(ctx *echo.Context, rolesInt interface{})
		RegisterRolePermissions(ctx *echo.Context, role auth.Role)*/
	IsAllowed(ctx *echo.Context, perm string) bool
	SubscribeTopic(ctx *echo.Context, topic string, handler TopicListener)
	PublishMessage(ctx *echo.Context, topic string, message interface{}) error
	GetConfig() config.Config
}

const (
	SERVICE_TYPE_APP  = "SERVICE_TYPE_APP"
	SERVICE_TYPE_DATA = "SERVICE_TYPE_DATA"
	SERVICE_TYPE_WEB  = "SERVICE_TYPE_WEB"
)
