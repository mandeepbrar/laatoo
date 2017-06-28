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

type environment struct {
	*abstractserver

	applications map[string]server.Application

	server *serverObject
}

func newEnvironment(svrCtx *serverContext, name string, svr *serverObject, baseDir string, filterConf config.Config) (*environment, *environmentProxy) {

	env := &environment{server: svr, applications: make(map[string]server.Application, 5)}
	proxy := &environmentProxy{env: env}
	env.abstractserver = newAbstractServer(svrCtx, name, svr.abstractserver, proxy, baseDir, filterConf)
	env.proxy = proxy
	log.Debug(svrCtx, "Created environment", "Name", name)
	return env, proxy
}

func (env *environment) Initialize(ctx core.ServerContext, conf config.Config) error {
	envInitCtx := env.createContext(ctx, "InitializeEnvironment:"+env.name)
	if err := env.initialize(envInitCtx, conf); err != nil {
		return errors.WrapError(envInitCtx, err)
	}
	log.Trace(envInitCtx, "Initialized environment "+env.name)
	return nil
}

func (env *environment) Start(ctx core.ServerContext) error {
	envStartCtx := env.createContext(ctx, "StartEnvironment:"+env.name)
	if err := env.start(envStartCtx); err != nil {
		return errors.WrapError(envStartCtx, err)
	}
	log.Debug(envStartCtx, "Started environment "+env.name)
	return nil
}

func (env *environment) createApplications(ctx core.ServerContext, baseDir string, name string, applicationConf config.Config) error {
	appCreateCtx := env.createContext(ctx, "CreateApplication: "+name)

	if applicationConf == nil {
		applicationConf = make(common.GenericConfig, 0)
	}

	log.Trace(appCreateCtx, "Creating Application", "Base Directory", baseDir)
	filterConf, _ := applicationConf.GetSubConfig(constants.CONF_FILTERS)
	//create an application
	applHandle, applElem := newApplication(appCreateCtx, name, env, baseDir, filterConf)
	log.Debug(appCreateCtx, "Created")

	appInitCtx := env.createContext(ctx, "InitializeApplication: "+name)
	err := applHandle.Initialize(appInitCtx, applicationConf)
	if err != nil {
		return errors.WrapError(appInitCtx, err)
	}

	err = applHandle.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	env.applications[name] = applElem
	return nil
}

//creates a context specific to environment
func (env *environment) createContext(ctx core.ServerContext, name string) *serverContext {
	cmap := env.contextMap(ctx)
	cmap[core.ServerElementEnvironment] = env.proxy
	return ctx.(*serverContext).newContextWithElements(name, cmap, core.ServerElementEnvironment)
}
