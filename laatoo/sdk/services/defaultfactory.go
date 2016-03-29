package services

import (
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type DefaultFactory struct {
	Conf config.Config
}

//Create the services configured for factory.
func (df *DefaultFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	df.Conf = conf
	method, err := df.GetMethod(ctx, name, conf)
	if err != nil {
		return nil, err
	}
	return services.NewService(ctx, method, conf), nil
}

func (df *DefaultFactory) GetMethod(ctx core.ServerContext, name string, conf config.Config) (core.ServiceFunc, error) {
	return nil, nil
}

//The services start serving when this method is called
func (ds *DefaultFactory) StartServices(ctx core.ServerContext) error {
	return nil
}
