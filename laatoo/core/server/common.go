package server

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_SERVERTYPE_HOSTNAME = "hostname"
	CONF_APPLICATIONS        = "applications"
	CONF_APPNAME             = "name"
	CONF_APPCONF             = "conf"
	CONF_APP_SERVERTYPE      = "servertype"
)

var (
	__initContext__ = NewServerContext("__init__", nil, nil)
)

//Initialize Server
func (server *Server) InitServer(ctx *serverContext, configFileName string, appContextProvider core.ApplicationContextProvider) error {
	server.Applications = make(map[string]*Application)
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
	log.Logger.Trace(ctx, "Getting applications")
	//read config
	apps, ok := server.Config.GetConfigArray(CONF_APPLICATIONS)
	if !ok {
		panic("No applications found")
	}
	log.Logger.Debug(ctx, " Applications to be initialized", "Number of applications", len(apps))
	createctx := ctx.subCtx("Create Application", serverConf, nil)
	for _, val := range apps {
		appName, ok := val.GetString(CONF_APPNAME)
		if !ok {
			panic("Application name not provided")
		}
		appConf, err := common.ConfigFileAdapter(val, CONF_APPCONF)
		if err != nil {
			panic("Application conf not provided")
		}
		appServerType, ok := val.GetString(CONF_APP_SERVERTYPE)
		if !ok {
			panic("Server type not provided for application")
		}
		if appServerType != server.ServerType {
			log.Logger.Info(createctx, "Skipping application", "Application", appName)
			continue
		}
		log.Logger.Debug(createctx, "Creating application", "Application", appName)
		appserverctx := createctx.subCtx(appName, val, nil)
		var appCtx core.ApplicationContext
		if appContextProvider != nil {
			appCtx, err = appContextProvider(appName)
			if err != nil {
				return errors.RethrowError(appserverctx, CORE_APPLICATION_NOT_CREATED, err, "Application", appName)
			}
		}
		application, err := newApplication(appserverctx, appName, appCtx, appConf, server.ServerType)
		if err != nil {
			return errors.RethrowError(appserverctx, CORE_APPLICATION_NOT_CREATED, err, "Application", appName)
		}
		server.Applications[appName] = application
	}
	initctx := ctx.subCtx("Initialize Applications", nil, nil)
	//Initializes applications to be hosted on this server
	for appName, app := range server.Applications {
		intappctx := initctx.subCtx(appName, app.Config, app)
		err := app.InitializeApplication(intappctx)
		if err != nil {
			return errors.RethrowError(intappctx, CORE_APPLICATION_NOT_INITIALIZED, err, "Application", appName)
		}
	}
	return nil
}

//start the server
func (server *Server) Start(ctx *serverContext) error {
	startctx := ctx.subCtx("Start Application", nil, nil)
	//starts applications to be hosted on this server
	for appName, app := range server.Applications {
		appstartctx := startctx.subCtx(appName, app.Config, app)
		err := app.StartApplication(appstartctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
