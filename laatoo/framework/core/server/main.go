package server

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

func main(rootctx *serverContext, configFile string) error {
	//read the config file
	conf, err := config.NewConfigFromFile(configFile)
	if err != nil {
		return err
	}

	//config logger
	debug := log.ConfigLogger(conf)
	if !debug {
		errors.ShowStack = false
	}

	//create the server
	//object loader and engines are created
	serverHandle, _, ctx := newServer(rootctx)

	//initialize server
	//factory and service manager are configured
	err = serverHandle.Initialize(ctx, conf)
	if err != nil {
		return err
	}

	//start the server
	err = serverHandle.Start(ctx)
	if err != nil {
		return err
	}

	//create environments on a running server
	err = createEnvironments(ctx, conf)
	if err != nil {
		return err
	}

	//create applications on environments
	//each application is hosted on an environment
	err = createApplications(ctx, conf)
	if err != nil {
		return err
	}

	err = startListening(ctx, conf)
	if err != nil {
		return err
	}
	return nil
}

// create environments in the config on a running server
func createEnvironments(ctx core.ServerContext, conf config.Config) error {
	svrCtx := ctx.(*serverContext)
	svrProx := svrCtx.server.(*serverProxy)
	//read all configs
	envs, ok := conf.GetSubConfig(config.CONF_ENVIRONMENTS)
	if ok {
		envNames := envs.AllConfigurations()
		for _, envName := range envNames {
			envConfig, err, _ := common.ConfigFileAdapter(ctx, envs, envName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			//create named environment from a config
			err = svrProx.server.createEnvironment(ctx, envName, envConfig)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//create applications on named environments
func createApplications(ctx core.ServerContext, conf config.Config) error {
	svrCtx := ctx.(*serverContext)
	svrProx := svrCtx.server.(*serverProxy)
	apps, ok := conf.GetSubConfig(config.CONF_APPLICATIONS)
	if ok {
		appNames := apps.AllConfigurations()
		//iterate applications
		for _, appName := range appNames {
			appConfig, err, _ := common.ConfigFileAdapter(ctx, apps, appName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			//find the environment on which the application is to be hosted
			envName, ok := appConfig.GetString(config.CONF_APP_ENVIRONMENT)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_APP_ENVIRONMENT)
			}
			//get the environment from the server
			envElem, ok := svrProx.server.environments[envName]
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_APP_ENVIRONMENT)
			}
			//ask the environment to create application using the config
			envProxy := envElem.(*environmentProxy)
			err = envProxy.env.createApplications(ctx, appName, appConfig)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
