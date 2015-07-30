package laatoocodegen

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	CONF_CODEGEN_SERVICENAME = "codegen_service"
)

//Environment hosting an application
type CodegenService struct {
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_CODEGEN_SERVICENAME, CodegenServiceFactory)
}

//factory method returns the service object to the environment
func CodegenServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating static service")
	svc := &CodegenService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(CODEGEN_ERROR_MISSING_ROUTER)
	}
	router := routerInt.(*echo.Group)
	log.Logger.Infof("Router %s", router)

	return svc, nil
}

//Provides the name of the service
func (svc *CodegenService) GetName() string {
	return CONF_CODEGEN_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *CodegenService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *CodegenService) Serve() error {
	return nil
}

//Type of service
func (svc *CodegenService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}
