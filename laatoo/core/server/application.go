package server

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type application struct {
	*abstractserver

	env *environment

	//all applets deployed on this server
	applet server.Applet
}

func newApplication(svrCtx *serverContext, name string, env *environment, filterConf config.Config) (*application, *applicationProxy) {
	app := &application{env: env}
	appCtx := env.proxy.NewCtx(name)
	proxy := &applicationProxy{Context: appCtx.(*common.Context), app: app}
	app.abstractserver = newAbstractServer(svrCtx, name, env.abstractserver, proxy, filterConf)
	app.proxy = proxy
	log.Logger.Debug(svrCtx, "Created application", "Name", name)
	return app, proxy
}

//initialize application with object loader, factory manager, service manager
func (app *application) Initialize(ctx core.ServerContext, conf config.Config) error {
	appInitCtx := app.createContext(ctx, "InitializeApplication: "+app.name)
	if err := app.initialize(appInitCtx, conf); err != nil {
		return errors.WrapError(appInitCtx, err)
	}
	log.Logger.Debug(appInitCtx, "Initialized application "+app.name)

	return nil
}

//start application with object loader, factory manager, service manager
func (app *application) Start(ctx core.ServerContext) error {
	applicationStartCtx := app.createContext(ctx, "Start Application: "+app.name)
	if err := app.start(applicationStartCtx); err != nil {
		return errors.WrapError(applicationStartCtx, err)
	}
	log.Logger.Debug(applicationStartCtx, "Started application"+app.name)
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
func (app *application) createContext(ctx core.ServerContext, name string) *serverContext {
	cmap := app.contextMap(ctx, name)
	cmap[core.ServerElementApplication] = app.proxy
	return ctx.(*serverContext).newContextWithElements(name, cmap, core.ServerElementApplication)
}
