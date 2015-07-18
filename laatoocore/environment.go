package laatoocore

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"laatoosdk/config"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/utils"
)

const (
	CONF_ENV_SERVICES    = "services"
	CONF_ENV_SERVICENAME = "servicename"
	CONF_ENV_ROUTER      = "router"
)

//Environment hosting an application
type Environment struct {
	Router        *gin.RouterGroup
	Config        config.Config
	ServicesStore *utils.MemoryStorer
}

//creates a new environment
func newEnvironment(envName string, conf string, router *gin.RouterGroup) (*Environment, error) {
	env := &Environment{Router: router}
	env.ServicesStore = utils.NewMemoryStorer()
	//read config for standalone
	env.Config = config.NewConfigFromFile(conf)
	//create all services in the environment
	err := env.createServices()
	if err != nil {
		return nil, fmt.Errorf("Could not create environment %s Error: %s", envName, err)
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
			err := fmt.Errorf("Service name of alias %s is wrong", alias)
			return err
		}
		//router to be passed in the configuration
		serviceConfig[CONF_ENV_ROUTER] = env.Router
		//get the service with a given name alias and config
		svcInt, err := GetService(svcName, alias, serviceConfig)
		if err != nil {
			err := fmt.Errorf("Service %s could not be created. Error: %s", alias, err)
			return err
		}
		//put the created service in the store
		svc := svcInt.(service.Service)
		env.ServicesStore.PutObject(alias, svc)
	}
	return nil
}

//Initialize an environment
func (env *Environment) InitializeEnvironment() {
	//go through list of all the services
	svcs := env.ServicesStore.GetList()
	//iterate through all the services
	for _, svcInt := range svcs {
		svc := svcInt.(service.Service)
		//initialize service
		svc.Initialize(env)
	}
}

//start services
func (env *Environment) StartEnvironment() {
	//go through list of all the services
	svcs := env.ServicesStore.GetList()
	//iterate through all the services
	for _, svcInt := range svcs {
		svc := svcInt.(service.Service)
		//start service
		svc.Serve()
	}
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
