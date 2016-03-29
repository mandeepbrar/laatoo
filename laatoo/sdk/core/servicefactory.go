package core

import (
	"laatoo/sdk/config"
)

type ServiceFactoryProvider func(ctx ServerContext, config config.Config) (ServiceFactory, error)

//Service interface that needs to be implemented by any service of a system
type ServiceFactory interface {
	//Create the services configured for factory.
	CreateService(ctx ServerContext, name string, config config.Config) (Service, error)
	//The services start serving when this method is called
	StartServices(ctx ServerContext) error
}
