package laatoormi

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	LOGGING_CONTEXT         = "rmi_service"
	CONF_RMI_SERVICENAME    = "rmi_service"
	CONF_RMI_DATA_SVC       = "data_svc"
	CONF_RMI_DATASTORE      = "datastore"
	CONF_RMI_METHODS        = "methods"
	CONF_RMI_PATH           = "path"
	CONF_RMI_METHODNAME     = "method"
	CONF_RMI_HTTPMETHODNAME = "httpmethod"
	CONF_RMI_METHODCONFIG   = "methodconfig"
)

//Environment hosting an application
type RmiService struct {
	serviceEnv      service.Environment
	dataServiceName string
	DataStore       data.DataService
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_RMI_SERVICENAME, RmiServiceFactory)
}

//factory method returns the service object to the environment
func RmiServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating rmi service")
	serviceEnv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)
	svc := &RmiService{serviceEnv: serviceEnv}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(*echo.Group)

	entitydatasvcInt, ok := conf[CONF_RMI_DATA_SVC]
	svc.dataServiceName = entitydatasvcInt.(string)

	rmimethodsInt, ok := conf[CONF_RMI_METHODS]
	if !ok {
		return nil, errors.ThrowError(ctx, RMI_ERROR_MISSING_METHODS)
	}

	rmimethods, ok := rmimethodsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ctx, RMI_ERROR_MISSING_METHODS)
	}

	for name, val := range rmimethods {

		methodConfig, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.ThrowError(ctx, RMI_ERROR_INCORRECT_METHOD_CONF, "Method", name)
		}

		pathInt, ok := methodConfig[CONF_RMI_PATH]
		if !ok {
			return nil, errors.ThrowError(ctx, RMI_ERROR_MISSING_PATH, "Method", name)
		} else {
			methodInt, ok := methodConfig[CONF_RMI_METHODNAME]
			if !ok {
				return nil, errors.ThrowError(ctx, RMI_ERROR_MISSING_METHODNAME, "Method", name)
			}
			path := pathInt.(string)
			methodName := methodInt.(string)
			method, err := laatoocore.GetMethod(ctx, methodName)
			if err != nil {
				return nil, err
			}
			httpmethodInt, ok := methodConfig[CONF_RMI_HTTPMETHODNAME]
			if ok {
				switch httpmethodInt.(string) {
				case "PUT":
					router.Put(path, func(ctx *echo.Context) error {
						return svc.invokeMethod(ctx, method, methodConfig)
					})
				case "POST":
					router.Post(path, func(ctx *echo.Context) error {
						return svc.invokeMethod(ctx, method, methodConfig)
					})
				}
			} else {
				router.Post(path, func(ctx *echo.Context) error {
					return svc.invokeMethod(ctx, method, methodConfig)
				})
			}

		}
	}
	return svc, nil
}

func (svc *RmiService) invokeMethod(ctx *echo.Context, method laatoocore.InvokableMethod, methodConfig map[string]interface{}) error {
	ctx.Set(CONF_RMI_DATASTORE, svc.DataStore)
	ctx.Set(CONF_RMI_METHODCONFIG, methodConfig)
	err := method(ctx)
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Error in invoking method", "method", method, "err", err)
	return err
}

//Provides the name of the service
func (svc *RmiService) GetName() string {
	return CONF_RMI_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *RmiService) Initialize(ctx *echo.Context) error {
	svcenv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)
	dataSvc, err := svcenv.GetService(ctx, svc.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, RMI_ERROR_MISSING_DATASVC, err)
	}

	svc.DataStore = dataSvc.(data.DataService)
	return nil
}

//The service starts serving when this method is called
func (svc *RmiService) Serve(ctx *echo.Context) error {
	return nil
}

//Type of service
func (svc *RmiService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *RmiService) Execute(ctx *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
