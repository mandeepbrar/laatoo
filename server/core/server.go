package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/adapters"
	"laatoo/server/common"
	"laatoo/server/constants"
	"laatoo/server/security"
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
func newServer(ctx *serverContext, baseDir string) (*serverObject, error) {
	ctx.Set(constants.RELATIVE_DIR, constants.CONF_APP_SERVER)
	//set a server type from the standalone/appengine file
	svr := &serverObject{serverType: SERVER_TYPE}
	svr.environments = make(map[string]server.Environment, 5)
	//create a proxy for the server
	svrElem := &serverProxy{server: svr}

	ctx.setElements(core.ContextMap{core.ServerElementServer: svrElem})

	//	svrContext := ctx.SubContext("Abstract Server")
	abstractserver, err := newAbstractServer(ctx, "Server", nil, svrElem, baseDir)
	if err != nil {
		return nil, err
	}
	svr.abstractserver = abstractserver

	log.Info(ctx, "Created server")

	return svr, nil
}

//initialize the server with the read config
func (svr *serverObject) Initialize(ctx core.ServerContext, conf config.Config) error {
	//svr.objectLoader.Register(ctx, config.DEFAULT_ROLE, security.Role{}, nil)
	//svr.objectLoader.Register(ctx, config.DEFAULT_USER, security.DefaultUser{}, nil)
	svr.objectLoader.Register(ctx, common.CONF_DEFAULTFACTORY_NAME, adapters.DefaultFactory{}, nil)
	svr.objectLoader.Register(ctx, common.CONF_SERVICEAGGREGATOR_NAME, adapters.ServiceAggregator{}, nil)
	svr.objectLoader.Register(ctx, constants.CONST_ALL_PERMISSIONS, security.AllPermissions{}, nil)

	initctx := ctx.SubContext("Initializing Server").(*serverContext)
	svr.conf = conf
	if err := svr.initialize(initctx, conf); err != nil {
		return errors.WrapError(initctx, err)
	}

	log.Trace(initctx, "Initialized server")

	return nil
}

func (svr *serverObject) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Starting Server").(*serverContext)

	if err := svr.start(startCtx); err != nil {
		return errors.WrapError(startCtx, err)
	}

	log.Info(ctx, "Started server")
	return nil
}

func (svr *serverObject) createEnvironment(ctx core.ServerContext, baseDir string, name string, envConf config.Config) error {
	envCreate := ctx.SubContext("Create").(*serverContext)

	if envConf == nil {
		envConf = ctx.CreateConfig()
	}

	log.Trace(envCreate, "Creating Environment")
	envHandle, envElem, err := newEnvironment(envCreate, name, svr, baseDir)
	if err != nil {
		return err
	}
	log.Debug(envCreate, "Created environment")

	envInit := ctx.SubContext("Initialize").(*serverContext)
	err = envHandle.Initialize(envInit, envConf)
	if err != nil {
		return errors.WrapError(envInit, err)
	}
	log.Debug(envInit, "Initialized environment")

	envStart := ctx.SubContext("Start").(*serverContext)
	err = envHandle.Start(envStart)
	if err != nil {
		return errors.WrapError(envStart, err)
	}

	log.Debug(ctx, "Registered environment")
	svr.environments[name] = envElem
	return nil
}
