package laatooview

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	CONF_VIEW_SERVICENAME = "view_service"
	CONF_VIEW_VIEWS       = "views"
	CONF_VIEW_DATASVC     = "data_svc"
	CONF_VIEW_VIEWNAME    = "name"
	CONF_VIEW_VIEWPATH    = "path"
	//CONF_VIEW_OBJECT      = "object"
	//CONF_VIEW_TYPE        = "object"
)

//Environment hosting an application
type ViewService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_VIEW_SERVICENAME, ViewServiceFactory)
}

//factory method returns the service object to the environment
func ViewServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating view service")
	svc := &ViewService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(VIEW_ERROR_MISSING_ROUTER)
	}

	router := routerInt.(*echo.Group)
	log.Logger.Infof("Router %s", router)
	svc.Router = router

	datasvcInt, ok := conf[CONF_VIEW_DATASVC]
	if !ok {
		return nil, errors.ThrowError(VIEW_ERROR_MISSING_DATASVC, CONF_VIEW_SERVICENAME)
	}
	svc.dataSvcName = datasvcInt.(string)

	viewsInt, ok := conf[CONF_VIEW_VIEWS]
	if !ok {
		return nil, errors.ThrowError(VIEW_ERROR_MISSING_VIEWS, CONF_VIEW_SERVICENAME)
	}

	views, ok := viewsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(VIEW_ERROR_MISSING_VIEWS, CONF_VIEW_SERVICENAME)
	}

	for name, val := range views {

		viewConfig, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.ThrowError(VIEW_ERROR_INCORRECT_VIEW_CONF, CONF_VIEW_SERVICENAME, name)
		}

		nameInt, ok := viewConfig[CONF_VIEW_VIEWNAME]
		if !ok {
			return nil, errors.ThrowError(VIEW_ERROR_MISSING_VIEWNAME, CONF_VIEW_SERVICENAME, name)
		} else {
			name := nameInt.(string)
			pathInt, ok := viewConfig[CONF_VIEW_VIEWPATH]
			if !ok {
				return nil, errors.ThrowError(VIEW_ERROR_MISSING_VIEWPATH, CONF_VIEW_SERVICENAME, name)
			}

			path := pathInt.(string)

			viewInt, err := laatoocore.CreateObject(name, nil)
			if err != nil {
				return nil, errors.RethrowError(VIEW_ERROR_MISSING_VIEW, err, CONF_VIEW_SERVICENAME, name)
			}

			view := viewInt.(data.View)

			router.Get(path, func(ctx *echo.Context) error {
				return view.Execute(svc.DataStore, ctx)
			})
		}
	}

	return svc, nil
}

//Provides the name of the service
func (svc *ViewService) GetName() string {
	return CONF_VIEW_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *ViewService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(VIEW_ERROR_MISSING_DATASVC, err, CONF_VIEW_SERVICENAME)
	}

	svc.DataStore = dataSvc.(data.DataService)
	return nil
}

//The service starts serving when this method is called
func (svc *ViewService) Serve() error {
	return nil
}

//Type of service
func (svc *ViewService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}
