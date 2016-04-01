package static

import (
	"fmt"
	"laatoo/core/registry"
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_STATIC_SERVICEFACTORY = "static_files"
	CONF_STATICSVC_DIRECTORY   = "directory"
	CONF_STATIC_CACHE          = "cache"
	CONF_STATIC_DIR            = "directory"
)

//Environment hosting an application
type StaticServiceFactory struct {
	Conf config.Config
}

//Initialize service, register provider with laatoo
func init() {
	registry.RegisterServiceFactoryProvider(CONF_STATIC_SERVICEFACTORY, NewStaticServiceFactory)
}

//factory method returns the service object to the environment
func NewStaticServiceFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Info(ctx, "Creating static service")
	svc := &StaticServiceFactory{conf}
	return svc, nil
}

//Create the services configured for factory.
func (sf *StaticServiceFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	sf.Conf = conf
	log.Logger.Info(ctx, "creating service files", "name", name)
	switch name {
	/*** Provides service for serving any files in a directory*****/
	case CONF_STATICSVC_DIRECTORY:
		{
			svcFunc, err := CreateDirectorySvc(ctx, name, conf)
			if err != nil {
				return nil, err
			}
			return services.NewService(ctx, svcFunc, conf), nil
		}
	/*** Provides service for serving files whose path has been specified*****/
	case CONF_STATICSVC_FILE:
		{
			return &FileService{conf: conf}, nil
		}
	case CONF_STATICSVC_FILEBUNDLE:
		{
			return &BundledFileService{conf: conf}, nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *StaticServiceFactory) StartServices(ctx core.ServerContext) error {
	return nil
}

func CreateDirectorySvc(ctx core.ServerContext, name string, conf config.Config) (core.ServiceFunc, error) {
	dir, ok := conf.GetString(CONF_STATICSVC_DIRECTORY)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATICSVC_DIRECTORY)
	}
	return func(ctx core.RequestContext) error {
		filename, ok := ctx.GetString(CONF_STATIC_FILEPARAM)
		if ok {
			ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, fmt.Sprintf("%s/%s", dir, filename), nil))
		} else {
			ctx.SetResponse(core.StatusNotFoundResponse)
		}
		return nil
	}, nil
}
