package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	"laatoo/server/constants"
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
func newServer(rootctx *serverContext, baseDir string) (*serverObject, core.ServerElement, *serverContext) {
	//set a server type from the standalone/appengine file
	svr := &serverObject{serverType: SERVER_TYPE}
	svr.environments = make(map[string]server.Environment, 5)
	//create a proxy for the server
	svrElem := &serverProxy{server: svr}

	svrContext := rootctx.newContext("Server")
	svrContext.setElements(core.ContextMap{core.ServerElementServer: svrElem}, core.ServerElementServer)

	svr.abstractserver = newAbstractServer(svrContext, "Server", nil, svrElem, baseDir, nil)

	cmap := svr.contextMap(svrContext)
	cmap[core.ServerElementServer] = svr.proxy
	svrContext.setElements(cmap, core.ServerElementServer)

	log.Info(svrContext, "Created server")

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
		log.Trace(initctx, "Created server messaging manager")
	}*/

	if err := svr.initialize(initctx, conf); err != nil {
		return errors.WrapError(initctx, err)
	}
	log.Trace(initctx, "Initialized server")

	cmap := svr.contextMap(svrCtx)
	cmap[core.ServerElementServer] = svr.proxy
	svrCtx.setElements(cmap, core.ServerElementServer)

	return nil
}

func (svr *serverObject) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Starting Server").(*serverContext)

	log.Trace(startCtx, "Starting server")
	if err := svr.start(startCtx); err != nil {
		return errors.WrapError(startCtx, err)
	}

	log.Info(ctx, "Started server")
	return nil
}

func (svr *serverObject) createEnvironment(ctx core.ServerContext, baseDir string, name string, envConf config.Config) error {
	envCreate := ctx.SubContext("Creating Environment: " + name).(*serverContext)

	if envConf == nil {
		envConf = make(common.GenericConfig, 0)
	}

	log.Trace(envCreate, "Creating Environment")
	filterConf, _ := envConf.GetSubConfig(constants.CONF_FILTERS)

	envHandle, envElem := newEnvironment(envCreate, name, svr, baseDir, filterConf)
	log.Debug(envCreate, "Created environment")

	err := envHandle.Initialize(envCreate, envConf)
	if err != nil {
		return errors.WrapError(envCreate, err)
	}
	log.Debug(envCreate, "Initialized environment")

	err = envHandle.Start(envCreate)
	if err != nil {
		return errors.WrapError(envCreate, err)
	}

	log.Debug(envCreate, "Registered environment")
	svr.environments[name] = envElem
	return nil
}
