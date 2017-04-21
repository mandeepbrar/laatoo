package storage

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

const (
	CONF_GOOGLESTORAGE_FACTORY = "googlestorage"
)

type GoogleStorageServiceFactory struct {
}

func init() {
	objects.Register(CONF_GOOGLESTORAGE_FACTORY, GoogleStorageServiceFactory{})
}

//The services start serving when this method is called
func (gs *GoogleStorageServiceFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (gs *GoogleStorageServiceFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	log.Logger.Trace(ctx, "Creating service for google storage factory", "name", name, "method", method)
	return &GoogleStorageSvc{}, nil
	/*
		switch method {
		case CONF_UPLOADFILE_SERVICENAME:
			{
				return &FileSystemSvc{fs}, nil

			}
		}
		return nil, nil*/
}

//The services start serving when this method is called
func (fs *GoogleStorageServiceFactory) Start(ctx core.ServerContext) error {
	return nil
}
