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

type application struct {
	name string

	objectLoader       server.ObjectLoader
	objectLoaderHandle server.ServerElementHandle

	channelMgr       server.ChannelManager
	channelMgrHandle server.ServerElementHandle

	factoryManager       server.FactoryManager
	factoryManagerHandle server.ServerElementHandle

	serviceManager       server.ServiceManager
	serviceManagerHandle server.ServerElementHandle

	securityHandler server.SecurityHandler

	env *environment

	//all applets deployed on this server
	applet server.Applet
	proxy  server.Application
}

func newApplication(svrCtx core.ServerContext, name string, filterConf config.Config) (*application, *applicationProxy) {
	envProxy := svrCtx.GetServerElement(core.ServerElementEnvironment).(*environmentProxy)
	env := envProxy.env
	app := &application{name: name, env: env}
	//create application context as a child of environment
	appCtx := envProxy.NewCtx(name)
	proxy := &applicationProxy{Context: appCtx.(*common.Context), app: app}
	app.proxy = proxy
	loader := env.objectLoader
	factoryManager := env.factoryManager
	serviceManager := env.serviceManager
	channelMgr := env.channelMgr
	appLoaderHandle, appLoader := objects.ChildLoader(svrCtx, name, loader)
	app.objectLoaderHandle = appLoaderHandle
	app.objectLoader = appLoader.(server.ObjectLoader)
	factoryManagerHandle, appfactoryManager := factory.ChildFactoryManager(svrCtx, name, factoryManager, proxy)
	app.factoryManagerHandle = factoryManagerHandle
	app.factoryManager = appfactoryManager.(server.FactoryManager)
	serviceManagerHandle, appserviceManager := service.ChildServiceManager(svrCtx, name, serviceManager, proxy)
	app.serviceManagerHandle = serviceManagerHandle
	app.serviceManager = appserviceManager.(server.ServiceManager)
	chanMgrHandle, appChannelMgr := childChannelManager(svrCtx, name, channelMgr, proxy)
	app.channelMgrHandle = chanMgrHandle
	app.channelMgr = appChannelMgr.(server.ChannelManager)
	log.Logger.Debug(svrCtx, "Created application", "Name", name)
	return app, proxy
}

//initialize application with object loader, factory manager, service manager
func (app *application) Initialize(ctx core.ServerContext, conf config.Config) error {
	appInitCtx := app.createContext(ctx, "InitializeApplication: "+app.name)

	objinit := appInitCtx.SubContext("Initialize object loader")
	err := initializeObjectLoader(objinit, conf, app.objectLoaderHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(objinit, "Initialized application object loader")

	secinit := appInitCtx.SubContext("Initialize security handler")
	err = app.initializeSecurityHandler(secinit, conf)
	if err != nil {
		return err
	}
	log.Logger.Trace(secinit, "Initialized application security handler")

	chaninit := appInitCtx.SubContextWithElement("Initialize channel manager", core.ServerElementChannelManager)
	err = initializeChannelManager(chaninit, conf, app.channelMgrHandle)
	if err != nil {
		return err
	}
	log.Logger.Debug(chaninit, "Initialized application channel manager")

	facinit := appInitCtx.SubContext("Initialize factory manager")
	err = initializeFactoryManager(facinit, conf, app.factoryManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(facinit, "Initialized application factory manager")

	svcinit := appInitCtx.SubContext("Initialize service manager")
	err = initializeServiceManager(svcinit, conf, app.serviceManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(appInitCtx, "Initialized application")

	return nil
}

//start application with object loader, factory manager, service manager
func (app *application) Start(ctx core.ServerContext) error {
	applicationStartCtx := app.createContext(ctx, "Start Application: "+app.name)

	appobjldrCtx := applicationStartCtx.SubContextWithElement("Start ObjectLoader", core.ServerElementLoader)
	err := app.objectLoaderHandle.Start(appobjldrCtx)
	if err != nil {
		return errors.WrapError(appobjldrCtx, err)
	}
	log.Logger.Trace(appobjldrCtx, "Started Object Loader")

	chanstart := applicationStartCtx.SubContextWithElement("Start Channel manager", core.ServerElementChannelManager)
	err = app.channelMgrHandle.Start(chanstart)
	if err != nil {
		return errors.WrapError(chanstart, err)
	}

	appfmCtx := applicationStartCtx.SubContextWithElement("Start Factory Manager", core.ServerElementFactoryManager)
	err = app.factoryManagerHandle.Start(appfmCtx)
	if err != nil {
		return errors.WrapError(appfmCtx, err)
	}
	log.Logger.Trace(appfmCtx, "Started factory manager")

	appsmCtx := applicationStartCtx.SubContextWithElement("Start Service Manager", core.ServerElementServiceManager)
	err = app.serviceManagerHandle.Start(appsmCtx)
	if err != nil {
		return errors.WrapError(appsmCtx, err)
	}
	log.Logger.Trace(appsmCtx, "Started service manager")
	log.Logger.Debug(applicationStartCtx, "Started application")
	return nil
}
func (app *application) initializeSecurityHandler(ctx core.ServerContext, conf config.Config) error {
	secConf, ok := conf.GetSubConfig(config.CONF_SECURITY)
	if ok {
		shElem, sh := newSecurityHandler(ctx, "Application:"+app.name, app.proxy)
		err := shElem.Initialize(ctx, secConf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		app.securityHandler = sh.(server.SecurityHandler)
	} else {
		app.securityHandler = app.env.securityHandler
	}
	return nil
}

//create applets
func (app *application) createApplet(ctx core.ServerContext, name string, appletConf config.Config) error {
	appletCreateCtx := app.createContext(ctx, "Creating applet: "+name)
	applprovider, ok := appletConf.GetString(config.CONF_APPL_OBJECT)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Applet Name", name, "Missing Config", config.CONF_APPL_OBJECT)
	}

	log.Logger.Debug(appletCreateCtx, "Creating applet")
	obj, err := appletCreateCtx.CreateObject(applprovider, nil)
	if err != nil {
		return errors.RethrowError(appletCreateCtx, errors.CORE_ERROR_BAD_CONF, err)
	}
	applet, ok := obj.(server.Applet)
	if !ok {
		return errors.ThrowError(appletCreateCtx, errors.CORE_ERROR_BAD_CONF, "Not an applet", applprovider)
	}

	appletCtx := appletCreateCtx.NewContextWithElements(name, core.ContextMap{core.ServerElementApplet: applet}, core.ServerElementApplet)
	log.Logger.Trace(ctx, "Initializing applet")
	err = applet.Initialize(appletCtx, appletConf)
	if err != nil {
		return errors.WrapError(appletCtx, err)
	}

	log.Logger.Trace(appletCtx, "Starting applet")
	err = applet.Start(appletCtx)
	if err != nil {
		return errors.WrapError(appletCtx, err)
	}
	app.applet = applet
	log.Logger.Debug(appletCtx, "Created applet")
	return nil
}

//creates a context specific to environment
func (app *application) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementApplication: app.proxy,
			core.ServerElementLoader:          app.objectLoader,
			core.ServerElementSecurityHandler: app.securityHandler,
			core.ServerElementChannelManager:  app.channelMgr,
			core.ServerElementFactoryManager:  app.factoryManager,
			core.ServerElementServiceManager:  app.serviceManager}, core.ServerElementApplication)
}
