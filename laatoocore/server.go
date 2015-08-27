package laatoocore

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"laatoosdk/config"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net/http"
)

const (
	CONF_SERVERTYPE_STANDALONE = "STANDALONE"
	CONF_SERVERTYPE_GOOGLEAPP  = "GOOGLE_APP"
	CONF_SERVERTYPE_HOSTNAME   = "hostname"
	CONF_ENVIRONMENTS          = "environments"
	CONF_ENVNAME               = "name"
	CONF_ENVCONF               = "conf"
	CONF_ENVPATH               = "path"
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
func (server *Server) InitServer(configName string) {
	server.Applications = make(map[string]*Environment)
	log.Logger.Infof("Initializing server with config %s", configName)
	//read config for standalone
	server.Config = config.NewConfigFromFile(configName)
	//config logger
	debug := log.ConfigLogger(server.Config)
	//initialize router
	router := echo.New()
	// Middleware
	router.Use(mw.Logger())
	router.Use(mw.Recover())
	router.SetDebug(debug)

	log.Logger.Debugf("Getting environments")
	//read config
	envs := server.Config.GetArray(CONF_ENVIRONMENTS)
	log.Logger.Debugf("%s environments to be initialized", len(envs))
	for _, val := range envs {
		env := val.(map[string]interface{})
		envName := env[CONF_ENVNAME].(string)
		envConf := env[CONF_ENVCONF].(string)
		envPath := env[CONF_ENVPATH].(string)
		log.Logger.Debugf("Creating environment", envName)
		environment, err := newEnvironment(envName, envConf, router.Group(envPath))
		if err != nil {
			errors.RethrowError(CORE_ENVIRONMENT_NOT_CREATED, err, envName)
		}
		server.Applications[envName] = environment
	}
	//Initializes application environments to be hosted on this server
	for envName, app := range server.Applications {
		err := app.InitializeEnvironment()
		if err != nil {
			errors.RethrowError(CORE_ENVIRONMENT_NOT_INITIALIZED, err, envName)
		}
	}
	log.Logger.Debugf("Router %v", router)

	http.Handle("/", router)
}

//start the server
func (server *Server) Start() {
	//Initializes application environments to be hosted on this server
	for _, app := range server.Applications {
		app.StartEnvironment()
	}
}

//Create a new server
func NewServer(configName string, serverType string) (*Server, error) {
	server := &Server{ServerType: serverType}
	server.InitServer(configName)
	server.Start()
	//listen if server type is standalone
	if serverType == CONF_SERVERTYPE_STANDALONE {
		//find the address to bind from the server
		address := server.Config.GetString(CONF_SERVERTYPE_HOSTNAME)
		if address == "" {
			return nil, errors.ThrowError(CORE_SERVERADD_NOT_FOUND)
		}
		//start listening
		http.ListenAndServe(address, nil)
	}
	return server, nil
}
