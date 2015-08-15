package laatooactions

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	CONF_ACTION_SERVICENAME = "action_service"
	CONF_ACTION_ACTIONS     = "actions"
	CONF_ACTION_ACTIONPATH  = "path"
	CONF_ACTION_VIEWMODE    = "viewmode"
	CONF_ACTION_ACTIONTYPE  = "actiontype"
)

//Environment hosting an application
type ActionService struct {
	Router     *echo.Group
	allActions map[string]*entities.Action
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_ACTION_SERVICENAME, ActionServiceFactory)
}

//factory method returns the service object to the environment
func ActionServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating actions service")
	svc := &ActionService{}

	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ACTION_ERROR_MISSING_ROUTER)
	}

	router := routerInt.(*echo.Group)
	log.Logger.Infof("Router %s", router)
	svc.Router = router

	actionsInt, ok := conf[CONF_ACTION_ACTIONS]
	if !ok {
		return nil, errors.ThrowError(ACTION_ERROR_MISSING_ACTIONS, CONF_ACTION_SERVICENAME)
	}

	actionsConf, ok := actionsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ACTION_ERROR_MISSING_ACTIONS, CONF_ACTION_SERVICENAME)
	}
	svc.allActions = make(map[string]*entities.Action, len(actionsConf))
	for actionName, val := range actionsConf {
		actionConfig, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.ThrowError(ACTION_ERROR_INCORRECT_ACTION_CONF, CONF_ACTION_SERVICENAME, actionName, actionName)
		}
		pathInt, ok := actionConfig[CONF_ACTION_ACTIONPATH]
		if !ok {
			return nil, errors.ThrowError(ACTION_ERROR_MISSING_ACTION_PATH, CONF_ACTION_SERVICENAME, actionName)
		}
		path := pathInt.(string)
		action := &entities.Action{Name: actionName, Path: path}

		actionTypeInt, ok := actionConfig[CONF_ACTION_ACTIONTYPE]
		if !ok {
			return nil, errors.ThrowError(ACTION_ERROR_MISSING_ACTION_TYPE, CONF_ACTION_SERVICENAME, actionName)
		}
		action.ActionType = actionTypeInt.(string)

		viewModeInt, ok := actionConfig[CONF_ACTION_VIEWMODE]
		if ok {
			action.Viewmode = viewModeInt.(string)
		}

		svc.allActions[actionName] = action
		log.Logger.Infof("Path", path)
	}

	router.Get("", func(ctx *echo.Context) error {
		ctx.JSON(http.StatusOK, svc.allActions)
		return nil
	})

	//svc.GetActions
	return svc, nil
}

//Provides the name of the service
func (svc *ActionService) GetName() string {
	return CONF_ACTION_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *ActionService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *ActionService) Serve() error {
	return nil
}

//Type of service
func (svc *ActionService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *ActionService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	returnVal := make(map[string]interface{}, 5)
	if name == "getAllActions" {
		returnVal["actions"] = svc.allActions
		return returnVal, nil
	}
	return nil, nil
}
