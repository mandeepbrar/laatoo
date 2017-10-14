package core

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/common"
	"laatoo/server/constants"
	"os"
	"path"
)

func main(rootctx *serverContext, configDir string) error {
	log.Info(rootctx, "Config directory for server: "+configDir)

	configFile := path.Join(configDir, constants.CONF_CONFIG_FILE)
	//read the config file
	conf, err := common.NewConfigFromFile(rootctx, configFile, nil)
	if err != nil {
		log.Info(rootctx, "Config file not found in dir: "+configDir)
		return err
	}

	//create the server
	//object loader and engines are created
	ctx := rootctx.SubContext("Creating server")
	serverHandle, err := newServer(ctx.(*serverContext), configDir)
	if err != nil {
		return err
	}

	//initialize server
	//factory and service manager are configured
	ctx = rootctx.SubContext("Initializing server")
	err = serverHandle.Initialize(ctx, conf)
	if err != nil {
		return err
	}

	//start the server
	ctx = rootctx.SubContext("Starting server")
	err = serverHandle.Start(ctx)
	if err != nil {
		return err
	}

	//create environments on a running server
	ctx = rootctx.SubContext("Creating environments")
	envs, err := createEnvironments(ctx, configDir, serverHandle)
	if err != nil {
		return err
	}

	//create applications on environments
	//each application is hosted on an environment
	ctx = rootctx.SubContext("Creating applications")
	err = createApplications(ctx, envs, conf, serverHandle)
	if err != nil {
		return err
	}

	ctx = rootctx.SubContext("Listening")
	err = startListening(ctx, conf)
	if err != nil {
		return err
	}
	return nil
}

// create environments in the config on a running server
func createEnvironments(ctx core.ServerContext, confDir string, svr *serverObject) (map[string]string, error) {
	envDir := path.Join(confDir, constants.CONF_ENVIRONMENTS)
	envs := make(map[string]string)
	if _, err := os.Stat(envDir); err == nil {
		svrCtx := ctx.(*serverContext)

		files, err := ioutil.ReadDir(envDir)
		if err != nil {
			return envs, errors.WrapError(ctx, err, "Environment directory", envDir)
		}

		for _, info := range files {
			if info.IsDir() {
				envName := info.Name()
				baseEnvDir := path.Join(envDir, envName)
				envCtx := svrCtx.newContext("Environment: " + envName)
				var envConfig config.Config
				configFile := path.Join(baseEnvDir, constants.CONF_CONFIG_FILE)
				if _, err := os.Stat(configFile); err == nil {
					//read the config file
					envConfig, err = common.NewConfigFromFile(envCtx, configFile, nil)
					if err != nil {
						return envs, errors.WrapError(envCtx, err, "Environment config file", configFile)
					}
					name, ok := envConfig.GetString(envCtx, constants.CONF_OBJECT_NAME)
					if ok {
						envName = name
					}
				}

				//create named environment from a config
				err = svr.createEnvironment(envCtx, baseEnvDir, envName, envConfig)
				if err != nil {
					return envs, errors.WrapError(envCtx, err, "Environment", envName, "Base directory", baseEnvDir)
				}
				envs[envName] = baseEnvDir
			}
		}

		/*
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
			}*/
	} else {
		log.Info(ctx, "No environments were found", "Conf Directory", confDir)
	}

	return envs, nil
}

//create applications on named environments
func createApplications(ctx core.ServerContext, envs map[string]string, conf config.Config, svr *serverObject) error {

	for envName, baseDir := range envs {
		//get the environment from the server
		envElem, ok := svr.environments[envName]
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", constants.CONF_APP_ENVIRONMENT)
		}
		//ask the environment to create application using the config
		envProxy := envElem.(*environmentProxy)
		appDir := path.Join(baseDir, constants.CONF_APPLICATIONS)
		if _, err := os.Stat(appDir); err == nil {
			files, err := ioutil.ReadDir(appDir)
			if err != nil {
				return errors.WrapError(ctx, err, "Application directory", appDir)
			}
			for _, info := range files {
				if info.IsDir() {
					appName := info.Name()
					baseAppDir := path.Join(appDir, appName)
					appCtx := envProxy.env.svrContext.newContext("Application: " + appName)
					var appConfig config.Config
					configFile := path.Join(baseAppDir, constants.CONF_CONFIG_FILE)
					if _, err := os.Stat(configFile); err == nil {
						//read the config file
						appConfig, err = common.NewConfigFromFile(appCtx, configFile, nil)
						if err != nil {
							return errors.WrapError(appCtx, err, "Application config file", configFile)
						}
						name, ok := appConfig.GetString(appCtx, constants.CONF_OBJECT_NAME)
						if ok {
							appName = name
						}
					}

					err = envProxy.env.createApplications(appCtx, baseAppDir, appName, appConfig)
					if err != nil {
						return errors.WrapError(appCtx, err, "Environment", envName, "Application", appName, "Base directory", baseAppDir)
					}
				}
			}
		} else {
			log.Info(ctx, "No applications were found", "Conf Directory", baseDir)
		}
	}
	return nil
}
