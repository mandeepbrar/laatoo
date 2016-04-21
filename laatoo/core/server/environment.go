package server

import (
	"laatoo/core/common"
	"laatoo/core/factory"
	"laatoo/core/objects"
	"laatoo/core/service"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type environment struct {
	name string

	objectLoader       server.ObjectLoader
	objectLoaderHandle server.ServerElementHandle

	channelMgr       server.ChannelManager
	channelMgrHandle server.ServerElementHandle

	factoryManager       server.FactoryManager
	factoryManagerHandle server.ServerElementHandle

	serviceManager       server.ServiceManager
	serviceManagerHandle server.ServerElementHandle

	securityHandler       server.SecurityHandler
	securityHandlerHandle server.ServerElementHandle

	messagingManager       server.MessagingManager
	messagingManagerHandle server.ServerElementHandle

	//all applets deployed on this server
	applications map[string]server.Application
	proxy        server.Environment
	server       *serverObject
}

func newEnvironment(svrCtx core.ServerContext, name string, filterConf config.Config) (*environment, core.ServerElement) {
	svr := svrCtx.GetServerElement(core.ServerElementServer).(*serverProxy).server
	env := &environment{name: name, applications: make(map[string]server.Application, 5), server: svr}
	envCtx := svrCtx.NewCtx(name)
	proxy := &environmentProxy{Context: envCtx.(*common.Context), env: env}
	env.proxy = proxy
	loader := svr.objectLoader
	factoryManager := svr.factoryManager
	serviceManager := svr.serviceManager
	channelMgr := svr.channelMgr
	msgMgr := svr.messagingManager
	envLoaderHandle, envLoader := objects.ChildLoader(svrCtx, name, loader)
	env.objectLoaderHandle = envLoaderHandle
	env.objectLoader = envLoader.(server.ObjectLoader)
	factoryManagerHandle, envfactoryManager := factory.ChildFactoryManager(svrCtx, name, factoryManager, proxy)
	env.factoryManagerHandle = factoryManagerHandle
	env.factoryManager = envfactoryManager.(server.FactoryManager)
	serviceManagerHandle, envserviceManager := service.ChildServiceManager(svrCtx, name, serviceManager, proxy)
	env.serviceManagerHandle = serviceManagerHandle
	env.serviceManager = envserviceManager.(server.ServiceManager)
	chanMgrHandle, envChannelMgr := childChannelManager(svrCtx, name, channelMgr, proxy)
	env.channelMgrHandle = chanMgrHandle
	env.channelMgr = envChannelMgr.(server.ChannelManager)
	msgMgrHandle, envmsgMgr := childMessagingManager(svrCtx, name, msgMgr, proxy)
	env.messagingManagerHandle = msgMgrHandle
	env.messagingManager = envmsgMgr.(server.MessagingManager)
	log.Logger.Debug(svrCtx, "Created environment", "Name", name)
	return env, proxy
}

func (env *environment) Initialize(ctx core.ServerContext, conf config.Config) error {
	envInitCtx := env.createContext(ctx, "InitializeEnvironment")

	msginit := envInitCtx.SubContextWithElement("Initialize messaging manager", core.ServerElementMessagingManager)
	err := initializeMessagingManager(msginit, conf, env.messagingManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(msginit, "Initialized environment messaging manager")

	secinit := envInitCtx.SubContext("Initialize security handler")
	err = env.initializeSecurityHandler(secinit, conf)
	if err != nil {
		return err
	}
	log.Logger.Trace(secinit, "Initialized environment security handler")

	objinit := envInitCtx.SubContextWithElement("Initialize object loader", core.ServerElementLoader)
	err = initializeObjectLoader(objinit, conf, env.objectLoaderHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(objinit, "Initialized environment object loader")

	chaninit := envInitCtx.SubContextWithElement("Initialize channel manager", core.ServerElementChannelManager)
	err = initializeChannelManager(chaninit, conf, env.channelMgrHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(chaninit, "Initialized environment channel manager")

	facinit := envInitCtx.SubContextWithElement("Initialize factory manager", core.ServerElementFactoryManager)
	err = initializeFactoryManager(facinit, conf, env.factoryManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(facinit, "Initialized environment factory manager")

	svcinit := envInitCtx.SubContextWithElement("Initialize service manager", core.ServerElementServiceManager)
	err = initializeServiceManager(svcinit, conf, env.serviceManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(svcinit, "Initialized environment service manager")

	return nil
}

func (env *environment) Start(ctx core.ServerContext) error {
	envStartCtx := env.createContext(ctx, "StartEnvironment")

	msgstart := envStartCtx.SubContextWithElement("Start messaging manager", core.ServerElementMessagingManager)
	err := env.messagingManagerHandle.Start(msgstart)
	if err != nil {
		return errors.WrapError(msgstart, err)
	}

	err = env.startSecurityHandler(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	objstart := envStartCtx.SubContextWithElement("Start object loader", core.ServerElementLoader)
	err = env.objectLoaderHandle.Start(objstart)
	if err != nil {
		return errors.WrapError(objstart, err)
	}

	chanstart := envStartCtx.SubContextWithElement("Start channel manager", core.ServerElementChannelManager)
	err = env.channelMgrHandle.Start(chanstart)
	if err != nil {
		return errors.WrapError(chanstart, err)
	}

	facstart := envStartCtx.SubContextWithElement("Start factory manager", core.ServerElementFactoryManager)
	err = env.factoryManagerHandle.Start(facstart)
	if err != nil {
		return errors.WrapError(facstart, err)
	}

	svcstart := envStartCtx.SubContextWithElement("Start service manager", core.ServerElementServiceManager)
	err = env.serviceManagerHandle.Start(svcstart)
	if err != nil {
		return errors.WrapError(svcstart, err)
	}

	log.Logger.Debug(ctx, "Started environment", "Name", env.name)
	return nil
}

func (env *environment) initializeSecurityHandler(ctx core.ServerContext, conf config.Config) error {
	envSecCtx := env.createContext(ctx, "Security Handler")
	secConf, ok := conf.GetSubConfig(config.CONF_SECURITY)
	if ok {
		shElem, sh := newSecurityHandler(envSecCtx, "Environment:"+env.name, env.proxy)
		err := shElem.Initialize(envSecCtx, secConf)
		if err != nil {
			return errors.WrapError(envSecCtx, err)
		}
		env.securityHandlerHandle = shElem
		env.securityHandler = sh.(server.SecurityHandler)
	} else {
		env.securityHandler = env.server.securityHandler
	}
	return nil
}

func (env *environment) startSecurityHandler(ctx core.ServerContext) error {
	if env.securityHandler != env.server.securityHandler {
		envSecCtx := ctx.SubContextWithElement("Start Security Handler", core.ServerElementSecurityHandler)
		log.Logger.Trace(envSecCtx, "Starting Security Handler")
		return env.securityHandlerHandle.Start(envSecCtx)
	}
	return nil
}

func (env *environment) createApplications(ctx core.ServerContext, name string, applicationConf config.Config) error {
	appCreateCtx := env.createContext(ctx, "CreateApplication: "+name)

	log.Logger.Trace(appCreateCtx, "Creating Application")
	filterConf, _ := applicationConf.GetSubConfig(config.CONF_FILTERS)

	//create an application
	applHandle, applElem := newApplication(appCreateCtx, name, filterConf)
	log.Logger.Debug(appCreateCtx, "Created")

	err := applHandle.Initialize(appCreateCtx, applicationConf)
	if err != nil {
		return errors.WrapError(appCreateCtx, err)
	}

	log.Logger.Debug(appCreateCtx, "Initialized")

	err = applHandle.Start(appCreateCtx)
	if err != nil {
		return errors.WrapError(appCreateCtx, err)
	}

	log.Logger.Debug(appCreateCtx, "Started")
	env.applications[name] = applElem
	return nil
}

//creates a context specific to environment
func (env *environment) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementEnvironment: env.proxy,
			core.ServerElementLoader:           env.objectLoader,
			core.ServerElementChannelManager:   env.channelMgr,
			core.ServerElementMessagingManager: env.messagingManager,
			core.ServerElementSecurityHandler:  env.securityHandler,
			core.ServerElementFactoryManager:   env.factoryManager,
			core.ServerElementServiceManager:   env.serviceManager}, core.ServerElementEnvironment)
}
