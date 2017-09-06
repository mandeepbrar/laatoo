package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
)

type environment struct {
	*abstractserver

	applications map[string]server.Application

	server *serverObject
}

func newEnvironment(svrCtx *serverContext, name string, svr *serverObject, baseDir string) (*environment, *environmentProxy, error) {

	env := &environment{server: svr, applications: make(map[string]server.Application, 5)}
	proxy := &environmentProxy{env: env}
	abstractserver, err := newAbstractServer(svrCtx, name, svr.abstractserver, proxy, baseDir)
	if err != nil {
		return nil, nil, err
	}
	env.abstractserver = abstractserver
	svrCtx.Set(constants.RELATIVE_DIR, constants.CONF_ENVIRONMENTS)
	svrCtx.Set(constants.CONF_APP_ENVIRONMENT, name)
	env.proxy = proxy
	log.Debug(svrCtx, "Created environment", "Name", name)
	return env, proxy, nil
}

func (env *environment) Initialize(ctx core.ServerContext, conf config.Config) error {
	envInitCtx := ctx.(*serverContext)
	if err := env.initialize(envInitCtx, conf); err != nil {
		return errors.WrapError(envInitCtx, err)
	}
	log.Trace(envInitCtx, "Initialized environment "+env.name)
	return nil
}

func (env *environment) Start(ctx core.ServerContext) error {
	envStartCtx := ctx.(*serverContext)
	if err := env.start(envStartCtx); err != nil {
		return errors.WrapError(envStartCtx, err)
	}
	log.Debug(envStartCtx, "Started environment "+env.name)
	return nil
}

func (env *environment) createApplications(ctx core.ServerContext, baseDir string, name string, applicationConf config.Config) error {
	appCreateCtx := ctx.SubContext("Create").(*serverContext)

	if applicationConf == nil {
		applicationConf = make(config.GenericConfig, 0)
	}

	log.Trace(appCreateCtx, "Creating Application", "Base Directory", baseDir)
	//create an application
	applHandle, applElem, err := newApplication(appCreateCtx, name, env, baseDir)
	if err != nil {
		return err
	}
	log.Debug(appCreateCtx, "Created")

	appInitCtx := ctx.SubContext("Initialize")
	err = applHandle.Initialize(appInitCtx, applicationConf)
	if err != nil {
		return errors.WrapError(appInitCtx, err)
	}

	appStartCtx := ctx.SubContext("Start")
	err = applHandle.Start(appStartCtx)
	if err != nil {
		return errors.WrapError(appStartCtx, err)
	}

	env.applications[name] = applElem
	return nil
}
