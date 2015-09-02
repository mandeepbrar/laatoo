package laatoostatic

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
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
func StaticServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating static service")
	svc := &StaticService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(STATIC_ERROR_MISSING_ROUTER)
	}
	publicdir, ok := conf[CONF_STATIC_PUBLICDIR]
	if !ok {
		return nil, errors.ThrowError(STATIC_ERROR_MISSING_PUBLICDIR)
	}
	router := routerInt.(*echo.Group)

	log.Logger.Infof("Designer service starting with page path %s", publicdir)
	router.Static("/", publicdir.(string))
	return svc, nil
}

//Provides the name of the service
func (svc *StaticService) GetName() string {
	return CONF_STATIC_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *StaticService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *StaticService) Serve() error {
	return nil
}

//Type of service
func (svc *StaticService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *StaticService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
