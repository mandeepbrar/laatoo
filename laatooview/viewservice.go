package laatooview

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	LOGGING_CONTEXT          = "view_service"
	CONF_VIEW_SERVICENAME    = "view_service"
	CONF_VIEW_VIEWS          = "views"
	CONF_VIEW_DATASVC        = "data_svc"
	CONF_VIEW_VIEWOBJECTNAME = "objectname"
	CONF_VIEW_VIEWPATH       = "path"
	CONF_VIEW_VIEWCONF       = "conf"
	//CONF_VIEW_OBJECT      = "object"
	//CONF_VIEW_TYPE        = "object"
)

//Environment hosting an application
type ViewService struct {
	DataStore   data.DataService
	dataSvcName string
	Views       map[string]data.View
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_VIEW_SERVICENAME, ViewServiceFactory)
}

//factory method returns the service object to the environment
func ViewServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating view service")
	svc := &ViewService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, VIEW_ERROR_MISSING_ROUTER)
	}

	router := routerInt.(core.Router)

	datasvcInt, ok := conf[CONF_VIEW_DATASVC]
	if !ok {
		return nil, errors.ThrowError(ctx, VIEW_ERROR_MISSING_DATASVC)
	}
	svc.dataSvcName = datasvcInt.(string)

	viewsInt, ok := conf[CONF_VIEW_VIEWS]
	if !ok {
		return nil, errors.ThrowError(ctx, VIEW_ERROR_MISSING_VIEWS)
	}

	views, ok := viewsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ctx, VIEW_ERROR_MISSING_VIEWS)
	}
	svc.Views = make(map[string]data.View, len(views))
	for name, val := range views {

		viewConfig, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.ThrowError(ctx, VIEW_ERROR_INCORRECT_VIEW_CONF, "View Name", name)
		}

		objectnameInt, ok := viewConfig[CONF_VIEW_VIEWOBJECTNAME]
		if !ok {
			return nil, errors.ThrowError(ctx, VIEW_ERROR_MISSING_VIEWOBJECTNAME, "View Name", name)
		} else {
			objectname := objectnameInt.(string)
			pathInt, ok := viewConfig[CONF_VIEW_VIEWPATH]
			if !ok {
				return nil, errors.ThrowError(ctx, VIEW_ERROR_MISSING_VIEWPATH, "View Name", name)
			}

			path := pathInt.(string)

			viewInt, err := laatoocore.CreateObject(ctx, objectname, viewConfig)
			if err != nil {
				return nil, errors.RethrowError(ctx, VIEW_ERROR_MISSING_VIEW, err, "View Name", name)
			}

			view := viewInt.(data.View)

			/*conf := make(map[string]interface{})
			confInt, ok := viewConfig[CONF_VIEW_VIEWCONF]
			if ok {
				conf = confInt.(map[string]interface{})
			}*/

			router.Get(ctx, path, viewConfig, func(ctx core.Context) error {
				return view.Execute(ctx, svc.DataStore)
			})
			log.Logger.Trace(ctx, LOGGING_CONTEXT, "Registering view", "View Name", name)
			svc.Views[name] = view

		}
	}
	//router.Get(ctx, "", svc.GetView)

	return svc, nil
}

func (svc *ViewService) GetView(ctx core.Context) error {
	name := ctx.Query("viewname")
	if name == "" {
		return errors.ThrowError(ctx, VIEW_ERROR_MISSING_VIEW)
	}
	view, ok := svc.Views[name]
	if !ok {
		return errors.ThrowError(ctx, VIEW_ERROR_MISSING_VIEW, "View Name", name)
	}
	return view.Execute(ctx, svc.DataStore)
}

//Provides the name of the service
func (svc *ViewService) GetName() string {
	return CONF_VIEW_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *ViewService) Initialize(ctx core.Context) error {
	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, VIEW_ERROR_MISSING_DATASVC, err)
	}

	svc.DataStore = dataSvc.(data.DataService)
	return nil
}

//The service starts serving when this method is called
func (svc *ViewService) Serve(ctx core.Context) error {
	return nil
}

//Type of service
func (svc *ViewService) GetServiceType() string {
	return core.SERVICE_TYPE_WEB
}

//Execute method
func (svc *ViewService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
