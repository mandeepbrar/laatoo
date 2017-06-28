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

type application struct {
	*abstractserver

	env *environment

	//all applets deployed on this server
	applets map[string]server.Applet
}

func newApplication(svrCtx *serverContext, name string, env *environment, baseDir string, filterConf config.Config) (*application, *applicationProxy) {
	app := &application{env: env, applets: make(map[string]server.Applet, 1)}
	proxy := &applicationProxy{app: app}
	app.abstractserver = newAbstractServer(svrCtx, name, env.abstractserver, proxy, baseDir, filterConf)
	app.proxy = proxy
	log.Debug(svrCtx, "Created application", "Name", name)
	return app, proxy
}

//initialize application with object loader, factory manager, service manager
func (app *application) Initialize(ctx core.ServerContext, conf config.Config) error {
	appInitCtx := app.createContext(ctx, "InitializeApplication: "+app.name)
	if err := app.initialize(appInitCtx, conf); err != nil {
		return errors.WrapError(appInitCtx, err)
	}

	if err := app.createApplets(appInitCtx, conf); err != nil {
		return errors.WrapError(appInitCtx, err)
	}
	log.Debug(appInitCtx, "Initialized application "+app.name)
	return nil
}

//start application with object loader, factory manager, service manager
func (app *application) Start(ctx core.ServerContext) error {
	applicationStartCtx := app.createContext(ctx, "Start Application: "+app.name)
	if err := app.start(applicationStartCtx); err != nil {
		return errors.WrapError(applicationStartCtx, err)
	}

	for name, applet := range app.applets {
		log.Trace(applicationStartCtx, "Starting applet:"+name)
		err := applet.Start(applicationStartCtx)
		if err != nil {
			return errors.WrapError(applicationStartCtx, err)
		}
	}
	log.Debug(applicationStartCtx, "Started application"+app.name)
	return nil
}

//create applets
func (app *application) createApplets(ctx core.ServerContext, conf config.Config) error {
	appletsConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_APPLETS)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	appletNames := appletsConf.AllConfigurations()
	for _, name := range appletNames {
		appletConf, err, _ := common.ConfigFileAdapter(ctx, appletsConf, name)
		if err != nil {
			return err
		}
		appletCreateCtx := app.createContext(ctx, "Creating applet: ")
		applprovider, ok := appletConf.GetString(constants.CONF_APPL_OBJECT)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Applet Name", name, "Missing Config", constants.CONF_APPL_OBJECT)
		}

		log.Debug(appletCreateCtx, "Creating applet")
		obj, err := appletCreateCtx.CreateObject(applprovider)
		if err != nil {
			return errors.RethrowError(appletCreateCtx, errors.CORE_ERROR_BAD_CONF, err)
		}

		applet, ok := obj.(server.Applet)
		if !ok {
			return errors.ThrowError(appletCreateCtx, errors.CORE_ERROR_BAD_CONF, "Not an applet", applprovider)
		}

		appletCtx := appletCreateCtx.NewContext(name)
		log.Trace(ctx, "Initializing applet")
		err = applet.Initialize(appletCtx, appletConf)
		if err != nil {
			return errors.WrapError(appletCtx, err)
		}

		app.applets[name] = applet
		log.Debug(appletCtx, "Created applet")

	}

	return nil
}

//creates a context specific to environment
func (app *application) createContext(ctx core.ServerContext, name string) *serverContext {
	cmap := app.contextMap(ctx)
	cmap[core.ServerElementApplication] = app.proxy
	return ctx.(*serverContext).newContextWithElements(name, cmap, core.ServerElementApplication)
}
