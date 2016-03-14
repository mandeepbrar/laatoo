package laatoormi

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
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
)

//Environment hosting an application
type RmiService struct {
	dataServiceName string
	DataStore       data.DataService
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_RMI_SERVICENAME, RmiServiceFactory)
}

//factory method returns the service object to the environment
func RmiServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating rmi service")
	svc := &RmiService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(core.Router)
	entitydatasvcInt, ok := conf[CONF_RMI_DATA_SVC]
	if ok {
		svc.dataServiceName = entitydatasvcInt.(string)
	}

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
					router.Put(ctx, path, methodConfig, func(ctx core.Context) error {
						return svc.invokeMethod(ctx, method, methodName, nil)
					})
				case "POST":
					router.Post(ctx, path, methodConfig, func(ctx core.Context) error {
						return svc.invokeMethod(ctx, method, methodName, nil)
					})
				}
			} else {
				router.Post(ctx, path, methodConfig, func(ctx core.Context) error {
					return svc.invokeMethod(ctx, method, methodName, nil)
				})
			}

		}
	}
	return svc, nil
}

func (svc *RmiService) invokeMethod(ctx core.Context, method laatoocore.InvokableMethod, methodName string, params map[string]interface{}) error {
	if params == nil {
		params = map[string]interface{}{}
	}
	if svc.DataStore != nil {
		params[CONF_RMI_DATASTORE] = svc.DataStore
	}
	err := method(ctx, params)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Invoked method", "method", methodName, "err", err)
	return err
}

//Provides the name of the service
func (svc *RmiService) GetName() string {
	return CONF_RMI_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *RmiService) Initialize(ctx core.Context) error {
	if svc.dataServiceName != "" {
		dataSvc, err := ctx.GetService(svc.dataServiceName)
		if err != nil {
			return errors.RethrowError(ctx, RMI_ERROR_MISSING_DATASVC, err)
		}
		svc.DataStore = dataSvc.(data.DataService)
	}
	return nil
}

//The service starts serving when this method is called
func (svc *RmiService) Serve(ctx core.Context) error {
	return nil
}

//Type of service
func (svc *RmiService) GetServiceType() string {
	return core.SERVICE_TYPE_WEB
}

//Execute method
func (svc *RmiService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
