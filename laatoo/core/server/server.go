package server

import (
	"laatoo/core/common"
	"laatoo/core/engine/http"
	"laatoo/core/factory"
	"laatoo/core/objects"
	"laatoo/core/service"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

//Server hosting a number of applications
type serverObject struct {
	//name of the server
	name string
	//if server is standalone or google app
	serverType string

	//root Object Loader of the server
	objectLoader       server.ObjectLoader
	objectLoaderHandle server.ServerElementHandle

	//root factory manager for the server
	factoryManager       server.FactoryManager
	factoryManagerHandle server.ServerElementHandle

	//root service manager for the server
	serviceManager       server.ServiceManager
	serviceManagerHandle server.ServerElementHandle

	//all environments deployed on this server
	environments map[string]server.Environment

	//engines configured on the server
	engineHandles map[string]server.ServerElementHandle
	engines       map[string]server.Engine

	channelMgr       server.ChannelManager
	channelMgrHandle server.ServerElementHandle

	securityHandler server.SecurityHandler

	//config for the server
	conf  config.Config
	proxy server.Server
}

//Create a new server
func newServer(rootctx *serverContext) (*serverObject, core.ServerElement, core.ServerContext) {
	//set a server type from the standalone/appengine file
	svr := &serverObject{serverType: SERVER_TYPE}
	//create a proxy for the server
	svrElem := &serverProxy{server: svr, Context: rootctx.NewCtx("Server").(*common.Context)}
	svr.proxy = svrElem
	svrctx := rootctx.NewContextWithElements("Server", core.ContextMap{core.ServerElementServer: svrElem}, core.ServerElementServer)

	objCreateCtx := svrctx.SubContext("Create Object Loader")
	//create root object loader for the server
	objectLoaderHandle, objectLoader := objects.NewObjectLoader(objCreateCtx, "Server Object Loader", svrElem)
	//assign the object loader handle to the context and server
	//the object loader for context may change depending upon the context
	svr.objectLoaderHandle = objectLoaderHandle
	svr.objectLoader = objectLoader.(server.ObjectLoader)

	fmCreateCtx := svrctx.SubContext("Create Factory Manager")
	//create root factory manager for the server
	//assign it to the server and proxy
	factoryManagerHandle, factoryManager := factory.NewFactoryManager(fmCreateCtx, "Server Factory Manager", svrElem)
	svr.factoryManagerHandle = factoryManagerHandle
	svr.factoryManager = factoryManager.(server.FactoryManager)

	smCreateCtx := svrctx.SubContext("Create Service Manager")
	//create root service manager for the server
	//assign it to the server and proxy
	serviceManagerHandle, serviceManager := service.NewServiceManager(smCreateCtx, "Server Service Manager", svrElem)
	svr.serviceManagerHandle = serviceManagerHandle
	svr.serviceManager = serviceManager.(server.ServiceManager)

	svr.engineHandles = make(map[string]server.ServerElementHandle, 5)
	svr.engines = make(map[string]server.Engine, 5)
	svr.environments = make(map[string]server.Environment, 5)

	cmCreateCtx := svrctx.SubContext("Create Channel Manager")
	channelMgrHandle, channelMgr := newChannelManager(cmCreateCtx, "Server Channel Manager", svrElem)
	svr.channelMgr = channelMgr
	svr.channelMgrHandle = channelMgrHandle

	log.Logger.Info(svrctx, "Created server")
	return svr, svrElem, rootctx.NewContextWithElements("Server", core.ContextMap{core.ServerElementServer: svrElem, core.ServerElementLoader: svr.objectLoader,
		core.ServerElementFactoryManager: svr.factoryManager, core.ServerElementServiceManager: svr.serviceManager}, core.ServerElementServer)
}

//initialize the server with the read config
func (svr *serverObject) Initialize(ctx core.ServerContext, conf config.Config) error {
	initctx := svr.createContext(ctx, "InitializingServer")
	log.Logger.Trace(initctx, "Initializing")
	svr.conf = conf

	objinit := initctx.SubContextWithElement("Initialize object loader", core.ServerElementLoader)
	err := initializeObjectLoader(objinit, conf, svr.objectLoaderHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(objinit, "Initialized server object loader")

	secinit := initctx.SubContext("Initialize security handler")
	err = svr.initializeSecurityHandler(secinit, conf)
	if err != nil {
		return err
	}
	log.Logger.Trace(secinit, "Initialized server security handler")

	log.Logger.Trace(ctx, "Initializing engines")
	engines, ok := conf.GetSubConfig(config.CONF_ENGINES)
	if ok {
		engineNames := engines.AllConfigurations()
		for _, engName := range engineNames {
			engConf, _ := config.ConfigFileAdapter(engines, engName)
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

	//initializine the factory manager. The context already has object loader form server creation
	facinit := initctx.SubContext("Initialize factory manager")
	log.Logger.Trace(facinit, "Initializing server factory manager")
	err = initializeFactoryManager(facinit, conf, svr.factoryManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(facinit, "Initialized server factory manager")

	//initialize the service manager, The context already has service manager from server creation
	//set the server as parent of all services created by service manager
	svcinit := initctx.SubContext("Initialize service manager")
	err = initializeServiceManager(svcinit, conf, svr.serviceManagerHandle)
	if err != nil {
		return err
	}

	chminit := initctx.SubContext("Initialize channel manager")
	err = initializeChannelManager(chminit, conf, svr.channelMgrHandle)
	if err != nil {
		return err
	}

	log.Logger.Debug(svcinit, "Initialized server service manager")

	return nil
}

func (svr *serverObject) Start(ctx core.ServerContext) error {
	startCtx := svr.createContext(ctx, "StartingServer")
	log.Logger.Trace(startCtx, "Starting server")
	objStart := startCtx.SubContextWithElement("Start object loader", core.ServerElementLoader)
	log.Logger.Trace(objStart, "Started server object loader")
	err := svr.objectLoaderHandle.Start(objStart)
	if err != nil {
		return errors.WrapError(objStart, err)
	}

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

	facStart := startCtx.SubContext("Start server factory manager")
	err = svr.factoryManagerHandle.Start(facStart)
	if err != nil {
		return errors.WrapError(facStart, err)
	}
	log.Logger.Trace(facStart, "Started server factory manager")

	smStart := startCtx.SubContextWithElement("Start server service manager", core.ServerElementServiceManager)
	err = svr.serviceManagerHandle.Start(smStart)
	if err != nil {
		return errors.WrapError(smStart, err)
	}
	log.Logger.Trace(smStart, "Started server service manager")

	cmStart := startCtx.SubContextWithElement("Start server channel manager", core.ServerElementChannelManager)
	err = svr.channelMgrHandle.Start(cmStart)
	if err != nil {
		return errors.WrapError(cmStart, err)
	}
	log.Logger.Trace(cmStart, "Started server channel manager")

	log.Logger.Info(ctx, "Started server")
	return nil
}

func (svr *serverObject) initializeSecurityHandler(ctx core.ServerContext, conf config.Config) error {
	secConf, ok := conf.GetSubConfig(config.CONF_SECURITY)
	if ok {
		shElem, sh := newSecurityHandler(ctx, "Server", svr.proxy)
		err := shElem.Initialize(ctx, secConf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svr.securityHandler = sh.(server.SecurityHandler)
	}
	return nil
}

func (svr *serverObject) createEnvironment(ctx core.ServerContext, name string, envConf config.Config) error {
	envCreate := svr.createContext(ctx, "Creating Environment: "+name)
	log.Logger.Trace(envCreate, "Creating Environment")
	filterConf, _ := envConf.GetSubConfig(config.CONF_FILTERS)

	envHandle, envElem := newEnvironment(envCreate, name, filterConf)
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
func (svr *serverObject) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementServer: svr.proxy,
			core.ServerElementLoader:          svr.objectLoader,
			core.ServerElementSecurityHandler: svr.securityHandler,
			core.ServerElementChannelManager:  svr.channelMgr,
			core.ServerElementFactoryManager:  svr.factoryManager,
			core.ServerElementServiceManager:  svr.serviceManager}, core.ServerElementServer)
}
