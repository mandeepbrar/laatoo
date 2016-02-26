package laatoostatic

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	LOGGING_CONTEXT         = "static_file_service"
	CONF_STATIC_SERVICENAME = "static_file_service"
	CONF_STATIC_PUBLICDIR   = "publicdir"
)

//Environment hosting an application
type StaticService struct {
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_STATIC_SERVICENAME, StaticServiceFactory)
}

//factory method returns the service object to the environment
func StaticServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating static service")
	svc := &StaticService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, STATIC_ERROR_MISSING_ROUTER)
	}
	publicdir, ok := conf[CONF_STATIC_PUBLICDIR]
	if !ok {
		return nil, errors.ThrowError(ctx, STATIC_ERROR_MISSING_PUBLICDIR)
	}
	router := routerInt.(*echo.Group)

	log.Logger.Info(ctx, LOGGING_CONTEXT, "Image service starting", " Path", publicdir)
	router.Static("/", publicdir.(string))
	return svc, nil
}

//Provides the name of the service
func (svc *StaticService) GetName() string {
	return CONF_STATIC_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *StaticService) Initialize(ctx *echo.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *StaticService) Serve(ctx *echo.Context) error {
	return nil
}

//Type of service
func (svc *StaticService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *StaticService) Execute(ctx *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
