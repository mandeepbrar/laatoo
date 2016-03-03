package laatoocodegen

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	LOGGING_CONTEXT          = "codegen"
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
func CodegenServiceFactory(ctx *core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating static service")
	svc := &CodegenService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, CODEGEN_ERROR_MISSING_ROUTER)
	}
	router := routerInt.(core.Router)

	return svc, nil
}

//Provides the name of the service
func (svc *CodegenService) GetName() string {
	return CONF_CODEGEN_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *CodegenService) Initialize(ctx *core.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *CodegenService) Serve(ctx *core.Context) error {
	return nil
}

//Type of service
func (svc *CodegenService) GetServiceType() string {
	return core.SERVICE_TYPE_WEB
}

//Execute method
func (svc *CodegenService) Execute(ctx *core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
