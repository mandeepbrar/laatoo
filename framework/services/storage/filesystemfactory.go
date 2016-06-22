package storage

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

const (
	CONF_FILES_SERVICEFACTORY = "filesystem"
)

type FileServiceFactory struct {
}

func init() {
	objects.RegisterObject(CONF_FILES_SERVICEFACTORY, createFileServiceFactory, nil)
}

func createFileServiceFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &FileServiceFactory{}, nil
}

//The services start serving when this method is called
func (fs *FileServiceFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (fs *FileServiceFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	log.Logger.Trace(ctx, "Creating service for file system factory", "name", name, "method", method)
	return &FileSystemSvc{}, nil
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
func (fs *FileServiceFactory) Start(ctx core.ServerContext) error {
	return nil
}
