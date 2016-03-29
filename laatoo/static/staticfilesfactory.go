package laatoostatic

import (
	"laatoocore"
	"laatoocore/services"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	CONF_STATIC_SERVICEFACTORY = "static_files"
	CONF_STATIC_PUBLICDIR      = "publicdir"
)

//Environment hosting an application
type StaticServiceFactory struct {
	*services.DefaultFactory
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterServiceFactoryProvider(CONF_ENTITY_SERVICES, NewStaticServiceFactory)
}

//factory method returns the service object to the environment
func NewStaticServiceFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Info(ctx, "Creating static service")
	svc := &StaticServiceFactory{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, STATIC_ERROR_MISSING_ROUTER)
	}
	publicdir, ok := conf[CONF_STATIC_PUBLICDIR]
	if !ok {
		return nil, errors.ThrowError(ctx, STATIC_ERROR_MISSING_PUBLICDIR)
	}
	router := routerInt.(core.Router)

	log.Logger.Info(ctx, LOGGING_CONTEXT, "Image service starting", " Path", publicdir)
	router.Static(ctx, "/", conf, publicdir.(string))
	return svc, nil
}
