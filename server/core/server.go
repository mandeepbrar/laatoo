package core

import (
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
	svrElem := &serverProxy{server: svr}
	svr.proxy = svrElem

	svrContext := rootctx.newContext("Server")
	svrElem.Context = svrContext.Context
	svrContext.setElements(core.ContextMap{core.ServerElementServer: svrElem}, core.ServerElementServer)

	svr.abstractserver = newAbstractServer(svrContext, "Server", nil, svrElem, nil)

	cmap := svr.contextMap(svrContext)
	cmap[core.ServerElementServer] = svr.proxy
	svrContext.setElements(cmap, core.ServerElementServer)

	log.Logger.Info(rootctx, "Created server")

	return svr, svrElem, svrContext
}

//initialize the server with the read config
func (svr *serverObject) Initialize(ctx core.ServerContext, conf config.Config) error {
	svrCtx := ctx.(*serverContext)
	initctx := ctx.SubContext("Initializing Server").(*serverContext)
	svr.conf = conf

	/*svrMsgCtx := initctx.SubContext("Create Messaging Manager")
	msgSvcName, ok := conf.GetString(config.CONF_MESSAGING_SVC)
	if ok {
		msgHandle, msgElem := newMessagingManager(svrMsgCtx, "Server", svr.proxy, msgSvcName)
		svr.messagingManager = msgElem
		svr.messagingManagerHandle = msgHandle
		log.Logger.Trace(initctx, "Created server messaging manager")
	}*/

	if err := svr.initialize(initctx, conf); err != nil {
		return errors.WrapError(initctx, err)
	}
	log.Logger.Trace(initctx, "Initialized server")

	cmap := svr.contextMap(svrCtx)
	cmap[core.ServerElementServer] = svr.proxy
	svrCtx.setElements(cmap, core.ServerElementServer)

	return nil
}

func (svr *serverObject) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Starting Server").(*serverContext)

	log.Logger.Trace(startCtx, "Starting server")
	if err := svr.start(startCtx); err != nil {
		return errors.WrapError(startCtx, err)
	}

	log.Logger.Info(ctx, "Started server")
	return nil
}

func (svr *serverObject) createEnvironment(ctx core.ServerContext, baseDir string, name string, envConf config.Config) error {
	envCreate := svr.createContext(ctx, "Creating Environment: "+name)
	envCreate.Set(config.CONF_BASE_DIR, baseDir)

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
