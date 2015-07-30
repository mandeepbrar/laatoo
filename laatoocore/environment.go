package laatoocore

import (
	"github.com/labstack/echo"
	"laatoosdk/config"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/utils"
)

const (
	CONF_ENV_SERVICES     = "services"
	CONF_ENV_SERVICENAME  = "servicename"
	CONF_ENV_ROUTER       = "router"
	CONF_ENV_CONTEXT      = "context"
	CONF_SERVICE_BINDPATH = "path"
)

//Environment hosting an application
type Environment struct {
	Router        *echo.Group
	Config        config.Config
	ServicesStore *utils.MemoryStorer
}

//creates a new environment
func newEnvironment(envName string, conf string, router *echo.Group) (*Environment, error) {
	env := &Environment{Router: router}
	env.ServicesStore = utils.NewMemoryStorer()
	//read config for standalone
	env.Config = config.NewConfigFromFile(conf)
	//create all services in the environment

	if err := env.createServices(); err != nil {
		return nil, errors.RethrowError(CORE_ENVIRONMENT_NOT_CREATED, err, envName)
	}
	return env, nil
}

func (env *Environment) createServices() error {
	//get a map of all the services
	svcs := env.Config.GetMap(CONF_ENV_SERVICES)
	for alias, val := range svcs {
		//get the config for the service with given alias
		serviceConfig := val.(map[string]interface{})
		//get the service name to be created for the alias
		log.Logger.Info("Creating service %s", alias)

		svcName, ok := serviceConfig[CONF_ENV_SERVICENAME].(string)
		if !ok {
			return errors.ThrowError(CORE_ERROR_MISSING_SERVICE_NAME, alias)
		}

		svcBindPath, ok := serviceConfig[CONF_SERVICE_BINDPATH]
		if ok {
			//router to be passed in the configuration
			serviceConfig[CONF_ENV_ROUTER] = env.Router.Group(svcBindPath.(string))
		}

		serviceConfig[CONF_ENV_CONTEXT] = env

		//get the service with a given name alias and config
		svcInt, err := CreateObject(svcName, serviceConfig)
		if err != nil {
			return errors.RethrowError(CORE_ERROR_SERVICE_CREATION, err, alias)
		}

		//put the created service in the store
		svc := svcInt.(service.Service)
		env.ServicesStore.PutObject(alias, svc)
	}
	return nil
}

//Initialize an environment
func (env *Environment) InitializeEnvironment() error {
	//go through list of all the services
	svcs := env.ServicesStore.GetList()
	//iterate through all the services
	for _, svcInt := range svcs {
		svc := svcInt.(service.Service)
		//initialize service
		err := svc.Initialize(env)
		if err != nil {
			return errors.RethrowError(CORE_ERROR_SERVICE_INITIALIZATION, err, svc.GetName())
		}
	}
	return nil
}

//Provides the service reference by alias
func (env *Environment) GetService(alias string) (service.Service, error) {
	svcInt, err := env.ServicesStore.GetObject(alias)
	if err != nil {
		return nil, err
	}
	svc, ok := svcInt.(service.Service)
	if !ok {
		return nil, errors.RethrowError(CORE_ERROR_SERVICE_NOT_FOUND, err, alias)
	}
	return svc, nil
}

//creates a named object if the factory has been registered with environment
func (env *Environment) CreateObject(objName string, confData map[string]interface{}) (interface{}, error) {
	return CreateObject(objName, confData)
}

func (env *Environment) CreateEmptyObject(objName string) (interface{}, error) {
	return CreateEmptyObject(objName)
}

func (env *Environment) CreateCollection(objName string) (interface{}, error) {
	return CreateCollection(objName)
}

//start services
func (env *Environment) StartEnvironment() error {
	//go through list of all the services
	svcs := env.ServicesStore.GetList()
	//iterate through all the services
	for _, svcInt := range svcs {
		svc := svcInt.(service.Service)
		//start service
		if err := svc.Serve(); err != nil {
			return errors.RethrowError(CORE_ERROR_SERVICE_NOT_STARTED, err, svc.GetName())
		}
	}
	return nil
}

// Initialize routes of designer
/*func (env *Environment) InitializeRoutes(conf config.Config) {
	//initialize routes for api
	apiRouter := router.Group("/api")
	log.Logger.Debug(apiRouter)
	//initialize routes for designer
	router.Static("/", "designerhtml")
}
*/
