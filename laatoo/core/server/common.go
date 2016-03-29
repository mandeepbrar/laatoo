package server

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_SERVERTYPE_HOSTNAME = "hostname"
	CONF_ENVIRONMENTS        = "environments"
	CONF_ENVNAME             = "name"
	CONF_ENVCONF             = "conf"
	CONF_ENV_SERVERTYPE      = "servertype"
)

var (
	__initContext__ = NewServerContext("__init__", nil, nil)
)

//Initialize Server
func (server *Server) InitServer(ctx *serverContext, configFileName string) error {
	server.Applications = make(map[string]*Environment)
	log.Logger.Info(ctx, "Initializing server", "config", configFileName)
	//read config for standalone
	serverConf, err := config.NewConfigFromFile(configFileName)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err)
	}
	server.Config = serverConf
	//config logger
	debug := log.ConfigLogger(server.Config)
	log.Logger.Info(ctx, "Configuration file read", "config", configFileName, "debug", debug)
	if !debug {
		errors.ShowStack = false
	}
	log.Logger.Trace(ctx, "Getting environments")
	//read config
	envs, ok := server.Config.GetConfigArray(CONF_ENVIRONMENTS)
	if !ok {
		panic("No environments found")
	}
	log.Logger.Debug(ctx, " Environments to be initialized", "Number of environments", len(envs))
	createctx := ctx.subCtx("CreateEnvironment", serverConf, nil)
	for _, val := range envs {
		envName, ok := val.GetString(CONF_ENVNAME)
		if !ok {
			panic("Environment name not provided")
		}
		envConf, ok := val.GetString(CONF_ENVCONF)
		if !ok {
			panic("Environment conf not provided")
		}
		envServerType, ok := val.GetString(CONF_ENV_SERVERTYPE)
		if !ok {
			panic("Server type not provided for environment")
		}
		if envServerType != server.ServerType {
			log.Logger.Info(createctx, "Skipping environment", "Environment", envName)
			continue
		}
		log.Logger.Debug(createctx, "Creating environment", "Environment", envName)
		envctx := createctx.subCtx(envName, val, nil)
		envConfFile := fmt.Sprintf("%s/%s", envName, envConf)
		environment, err := newEnvironment(envctx, envName, envConfFile, server.ServerType)
		if err != nil {
			return errors.RethrowError(envctx, CORE_ENVIRONMENT_NOT_CREATED, err, "Environment", envName)
		}
		server.Applications[envName] = environment
	}
	initctx := ctx.subCtx("InitializeEnvironment", nil, nil)
	//Initializes application environments to be hosted on this server
	for envName, app := range server.Applications {
		envctx := initctx.subCtx(envName, app.Config, app)
		envctx.environment = app
		err := app.InitializeEnvironment(envctx)
		if err != nil {
			return errors.RethrowError(envctx, CORE_ENVIRONMENT_NOT_INITIALIZED, err, "Environment", envName)
		}
	}
	return nil
}

//start the server
func (server *Server) Start(ctx *serverContext) error {
	startctx := ctx.subCtx("StartEnvironment", nil, nil)
	//Initializes application environments to be hosted on this server
	for envName, app := range server.Applications {
		envctx := startctx.subCtx(envName, app.Config, app)
		err := app.StartEnvironment(envctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
