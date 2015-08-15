package laatooentities

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	CONF_ENTITY_SERVICENAME = "entity_service"
	CONF_ENTITY_ENTITIES    = "entities"
	//CONF_ENTITY_TYPE        = "object"
)

//Environment hosting an application
type EntityService struct {
	Router         *echo.Group
	EntitiesConf   map[string]interface{}
	EntityServices map[string]interface{}
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_ENTITY_SERVICENAME, EntityServiceFactory)
}

//factory method returns the service object to the environment
func EntityServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating entity service")
	svc := &EntityService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_ROUTER)
	}
	svc.Router = routerInt.(*echo.Group)

	//get a map of all the entities
	entitiesInt, ok := conf[CONF_ENTITY_ENTITIES]
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_ENTITIES)
	}

	entitiesConf, ok := entitiesInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_ENTITIES)
	}

	svc.EntityServices = make(map[string]interface{}, 20)

	svc.EntitiesConf = entitiesConf

	return svc, nil
}

//Provides the name of the service
func (svc *EntityService) GetName() string {
	return CONF_ENTITY_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *EntityService) Initialize(ctx service.ServiceContext) error {
	for name, val := range svc.EntitiesConf {
		//get the config for the page with given alias
		entityConf, ok := val.(map[string]interface{})
		if !ok {
			return errors.ThrowError(ENTITY_ERROR_WRONG_ENTITYCONFIG)
		}

		//get the service name to be created for the alias
		log.Logger.Infof("Creating entity %s", name)
		entitynameInt, ok := entityConf[entities.CONF_ENTITY_NAME]
		if !ok {
			return errors.ThrowError(entities.ENTITY_ERROR_MISSING_NAME)
		}

		entityConf[laatoocore.CONF_ENV_CONTEXT] = ctx

		entityConf[laatoocore.CONF_ENV_ROUTER] = svc.Router
		//create page with provided conf
		entInt, err := ctx.CreateObject(entitynameInt.(string), entityConf)
		if err != nil {
			return err
		}
		svc.EntityServices[name] = entInt
	}
	return nil
}

//The service starts serving when this method is called
func (svc *EntityService) Serve() error {
	return nil
}

//Type of service
func (svc *EntityService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *EntityService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
