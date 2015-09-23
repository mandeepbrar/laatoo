package laatoocore

import (
	"github.com/labstack/echo"
	"laatoosdk/config"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	CONF_SERVERTYPE_STANDALONE = "STANDALONE"
	CONF_SERVERTYPE_GOOGLEAPP  = "GOOGLE_APP"
	CONF_SERVERTYPE_HOSTNAME   = "hostname"
	CONF_ENVIRONMENTS          = "environments"
	CONF_ENVNAME               = "name"
	CONF_ENVCONF               = "conf"
	CONF_ENVPATH               = "path"
	CONF_ENV_SERVERTYPE        = "servertype"
)

//Server hosting a number of environments
type Server struct {
	//name of the server
	Name string
	//if server is standalone or google app
	ServerType string
	//all environments deployed on this server
	Applications map[string]*Environment
	//config file for the server
	Config config.Config
}

//Initialize Server
func (server *Server) InitServer(configName string, router *echo.Echo) {
	server.Applications = make(map[string]*Environment)
	log.Logger.Info("core.server", "Initializing server", "config", configName)
	//read config for standalone
	server.Config = config.NewConfigFromFile(configName)
	//config logger
	debug := log.ConfigLogger(server.Config)
	router.SetDebug(debug)

	log.Logger.Trace("core.server", "Getting environments")
	//read config
	envs := server.Config.GetArray(CONF_ENVIRONMENTS)
	log.Logger.Debug("core.server", " Environments to be initialized", "Number of environments", len(envs))
	for _, val := range envs {
		env := val.(map[string]interface{})
		envName := env[CONF_ENVNAME].(string)
		envConf := env[CONF_ENVCONF].(string)
		envPath := env[CONF_ENVPATH].(string)
		envServerType, ok := env[CONF_ENV_SERVERTYPE]
		if ok && (envServerType.(string) != server.ServerType) {
			log.Logger.Info("core.server", "Skipping environment", "Environment", envName)
			continue
		}
		log.Logger.Debug("core.server", "Creating environment", "Environment", envName)
		environment, err := newEnvironment(envName, envConf, router.Group(envPath), server.ServerType)
		if err != nil {
			errors.RethrowError(CORE_ENVIRONMENT_NOT_CREATED, err, envName)
		}
		server.Applications[envName] = environment
	}
	//Initializes application environments to be hosted on this server
	for envName, app := range server.Applications {
		err := app.InitializeEnvironment()
		if err != nil {
			errors.RethrowError(CORE_ENVIRONMENT_NOT_INITIALIZED, err, "Environment", envName)
		}
	}
}

//start the server
func (server *Server) Start() {
	//Initializes application environments to be hosted on this server
	for _, app := range server.Applications {
		app.StartEnvironment()
	}
}
