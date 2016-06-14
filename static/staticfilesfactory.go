package static

import (
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

const (
	CONF_STATIC_SERVICEFACTORY = "static_files"
	CONF_STATICSVC_DIRECTORY   = "directory"
	CONF_STATIC_CACHE          = "cache"
	CONF_STATIC_DIR            = "directory"
)

type StaticServiceFactory struct {
	conf config.Config
}

func init() {
	objects.RegisterObject(CONF_STATIC_SERVICEFACTORY, createStaticServicesFactory, nil)
}

func createStaticServicesFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &StaticServiceFactory{}, nil
}

//The services start serving when this method is called
func (ds *StaticServiceFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	ds.conf = conf
	return nil
}

//Create the services configured for factory.
func (sf *StaticServiceFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	log.Logger.Trace(ctx, "Creating service for static factory", "name", name, "method", method)
	switch method {
	/*** Provides service for serving any files in a directory*****/
	case CONF_STATICSVC_DIRECTORY:
		{
			return &staticFiles{name: name}, nil

		}
	/*** Provides service for serving files whose path has been specified*****/
	case CONF_STATICSVC_FILE:
		{
			return &FileService{name: name}, nil
		}
	case CONF_STATICSVC_FILEBUNDLE:
		{
			return &BundledFileService{name: name}, nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *StaticServiceFactory) Start(ctx core.ServerContext) error {
	return nil
}
