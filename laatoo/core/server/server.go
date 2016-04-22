package server

import (
	"laatoo/core/common"
	"laatoo/core/engine/http"

	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

//Server hosting a number of applications
type serverObject struct {
	*abstractserver
	//if server is standalone or google app
	serverType string

	//all environments deployed on this server
	environments map[string]server.Environment

	//engines configured on the server
	engineHandles map[string]server.ServerElementHandle
	engines       map[string]server.Engine

	//config for the server
	conf config.Config
}

//Create a new server
func newServer(rootctx *serverContext) (*serverObject, core.ServerElement, core.ServerContext) {
	//set a server type from the standalone/appengine file
	svr := &serverObject{serverType: SERVER_TYPE}
	svr.engineHandles = make(map[string]server.ServerElementHandle, 5)
	svr.engines = make(map[string]server.Engine, 5)
	svr.environments = make(map[string]server.Environment, 5)

	//create a proxy for the server
	svrElem := &serverProxy{server: svr, Context: rootctx.NewCtx("Server").(*common.Context)}
	svrctx := svr.createContext(rootctx, "Server Elements Creation")
	svr.abstractserver = newAbstractServer(svrctx, "Server", nil, svrElem, nil)
	svr.proxy = svrElem

	log.Logger.Info(svrctx, "Created server")
	return svr, svrElem, svr.createContext(rootctx, "Server")
}

//initialize the server with the read config
func (svr *serverObject) Initialize(ctx core.ServerContext, conf config.Config) error {

	initctx := svr.createContext(ctx, "InitializingServer")
	svr.conf = conf

	svrMsgCtx := initctx.SubContext("Create Messaging Manager")
	msgSvcName, ok := conf.GetString(config.CONF_MESSAGING_SVC)
	if ok {
		msgHandle, msgElem := newMessagingManager(svrMsgCtx, "Server", svr.proxy, msgSvcName)
		svr.messagingManager = msgElem
		svr.messagingManagerHandle = msgHandle
		log.Logger.Trace(initctx, "Created server messaging manager")
	}

	log.Logger.Trace(ctx, "Initializing engines")
	engines, ok := conf.GetSubConfig(config.CONF_ENGINES)
	if ok {
		engineNames := engines.AllConfigurations()
		for _, engName := range engineNames {
			engConf, err, _ := config.ConfigFileAdapter(engines, engName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			engCreateCtx := initctx.SubContext("Create Engine: " + engName)
			log.Logger.Trace(engCreateCtx, "Creating")
			engHandle, eng, err := svr.createEngine(engCreateCtx, engName, engConf)
			if err != nil {
				return errors.WrapError(engCreateCtx, err)
			}

			engInitCtx := initctx.SubContext("Initialize Engine: " + engName)
			log.Logger.Trace(engInitCtx, "Initializing engine")
			err = engHandle.Initialize(engInitCtx, engConf)
			if err != nil {
				return errors.WrapError(engInitCtx, err)
			}

			//get a root channel and assign it to server channel manager
			svr.channelMgr.(*channelManagerProxy).manager.channelStore[engName] = eng.GetRootChannel(engInitCtx)
			log.Logger.Info(engInitCtx, "Registered root channel", "Name", engName)

			svr.engines[engName] = eng
			svr.engineHandles[engName] = engHandle
			log.Logger.Trace(initctx, "Registered engine", "Name", engName)
		}
	}
	log.Logger.Debug(initctx, "Initialized Engines")

	if err := svr.initialize(initctx, conf); err != nil {
		return errors.WrapError(initctx, err)
	}
	log.Logger.Trace(initctx, "Initialized server")

	return nil
}

func (svr *serverObject) Start(ctx core.ServerContext) error {
	startCtx := svr.createContext(ctx, "Starting Server")

	engStartCtx := startCtx.SubContext("Start Engines")
	for _, engineHandle := range svr.engineHandles {
		errorsChan := make(chan error)
		go func() {
			log.Logger.Info(engStartCtx, "Starting engine")
			errorsChan <- engineHandle.Start(engStartCtx)
			err := <-errorsChan
			if err != nil {
				panic(err.Error())
			}
		}()
	}
	log.Logger.Trace(startCtx, "Started engines")

	log.Logger.Trace(startCtx, "Starting server")
	if err := svr.start(startCtx); err != nil {
		return errors.WrapError(startCtx, err)
	}

	log.Logger.Info(ctx, "Started server")
	return nil
}

func (svr *serverObject) createEnvironment(ctx core.ServerContext, name string, envConf config.Config) error {
	envCreate := svr.createContext(ctx, "Creating Environment: "+name)
	log.Logger.Trace(envCreate, "Creating Environment")
	filterConf, _ := envConf.GetSubConfig(config.CONF_FILTERS)

	envHandle, envElem := newEnvironment(envCreate, name, svr, filterConf)
	log.Logger.Debug(envCreate, "Created environment")

	err := envHandle.Initialize(envCreate, envConf)
	if err != nil {
		return errors.WrapError(envCreate, err)
	}
	log.Logger.Debug(envCreate, "Initialized environment")

	err = envHandle.Start(envCreate)
	if err != nil {
		return errors.WrapError(envCreate, err)
	}

	log.Logger.Debug(envCreate, "Registered environment")
	svr.environments[name] = envElem
	return nil
}

func (svr *serverObject) createEngine(ctx core.ServerContext, name string, engConf config.Config) (server.ServerElementHandle, server.Engine, error) {
	enginetype, ok := engConf.GetString(config.CONF_ENGINE_TYPE)
	if !ok {
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", config.CONF_ENGINE_TYPE)
	}
	var engineHandle server.ServerElementHandle
	var engine server.Engine
	switch enginetype {
	case config.CONF_ENGINETYPE_HTTP:
		engineHandle, engine = http.NewEngine(ctx, name)
	case core.CONF_ENGINE_TCP:
	default:
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Config Name", config.CONF_ENGINE_TYPE)
	}
	return engineHandle, engine, nil
}

//creates a context specific to server
func (svr *serverObject) createContext(ctx core.ServerContext, name string) *serverContext {
	cmap := svr.contextMap(ctx, name)
	cmap[core.ServerElementServer] = svr.proxy
	return ctx.(*serverContext).newContextWithElements(name, cmap, core.ServerElementServer)
}
