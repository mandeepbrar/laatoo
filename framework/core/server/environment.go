package server

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type environment struct {
	*abstractserver

	applications map[string]server.Application

	server *serverObject
}

func newEnvironment(svrCtx *serverContext, name string, svr *serverObject, filterConf config.Config) (*environment, *environmentProxy) {

	env := &environment{server: svr, applications: make(map[string]server.Application, 5)}
	envCtx := svr.proxy.NewCtx(name)
	proxy := &environmentProxy{Context: envCtx.(*common.Context), env: env}
	env.abstractserver = newAbstractServer(svrCtx, name, svr.abstractserver, proxy, filterConf)
	env.proxy = proxy
	log.Logger.Debug(svrCtx, "Created environment", "Name", name)
	return env, proxy
}

func (env *environment) Initialize(ctx core.ServerContext, conf config.Config) error {
	envInitCtx := env.createContext(ctx, "InitializeEnvironment:"+env.name)
	if err := env.initialize(envInitCtx, conf); err != nil {
		return errors.WrapError(envInitCtx, err)
	}
	log.Logger.Trace(envInitCtx, "Initialized environment "+env.name)
	return nil
}

func (env *environment) Start(ctx core.ServerContext) error {
	envStartCtx := env.createContext(ctx, "StartEnvironment:"+env.name)
	if err := env.start(envStartCtx); err != nil {
		return errors.WrapError(envStartCtx, err)
	}
	log.Logger.Debug(envStartCtx, "Started environment "+env.name)
	return nil
}

func (env *environment) createApplications(ctx core.ServerContext, name string, applicationConf config.Config) error {
	appCreateCtx := env.createContext(ctx, "CreateApplication: "+name)
	log.Logger.Trace(appCreateCtx, "Creating Application")
	filterConf, _ := applicationConf.GetSubConfig(config.CONF_FILTERS)
	//create an application
	applHandle, applElem := newApplication(appCreateCtx, name, env, filterConf)
	log.Logger.Debug(appCreateCtx, "Created")

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
	cmap := env.contextMap(ctx, name)
	cmap[core.ServerElementEnvironment] = env.proxy
	return ctx.(*serverContext).newContextWithElements(name, cmap, core.ServerElementEnvironment)
}
