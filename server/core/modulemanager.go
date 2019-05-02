package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"

	//"github.com/fsnotify/fsnotify"
	"github.com/blang/semver"
	"github.com/radovskyb/watcher"
)

type moduleManager struct {
	name                  string
	svrref                *abstractserver
	parent                core.ServerElement
	proxy                 elements.ModuleManager
	modulesRepo           string
	hotModulesRepo        string
	availableModules      map[string]string
	moduleInstances       map[string]elements.Module
	installedModules      map[string]*semver.Version
	moduleConf            map[string]config.Config
	moduleInstancesConfig map[string]config.Config
	loadedModules         map[string]*semver.Version
	modulePlugins         map[string]components.ModuleManagerPlugin
	parentModules         map[string]string
	hotModules            map[string]string
	objLoader             *objectLoader
	watchers              []*watcher.Watcher
}

func (modMgr *moduleManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Module manager initialized ")
	moduleInstancesConfig, err, _ := common.ConfigFileAdapter(ctx, conf, constants.CONF_MODULE_INSTANCES)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	modMgr.moduleInstancesConfig = make(map[string]config.Config)
	if moduleInstancesConfig != nil {
		instances := moduleInstancesConfig.AllConfigurations(ctx)
		//loop through module instances
		for _, instance := range instances {
			instanceConf, _ := moduleInstancesConfig.GetSubConfig(ctx, instance)
			modMgr.moduleInstancesConfig[instance] = instanceConf
		}
	}

	availableModules, _ := conf.GetSubConfig(ctx, constants.CONF_MODULES)

	modMgr.hotModules = make(map[string]string)
	modMgr.watchers = make([]*watcher.Watcher, 0)

	modMgr.objLoader = modMgr.svrref.objectLoaderHandle.(*objectLoader)

	baseDir, _ := ctx.GetString(config.BASEDIR)
	modulesDir := path.Join(baseDir, constants.CONF_MODULES)
	modulesRepository, ok := conf.GetString(ctx, constants.CONF_MODULES_REPO)
	if !ok {
		modulesRepository = modulesDir
	}
	modMgr.modulesRepo = modulesRepository

	modulesDevRepo, ok := conf.GetString(ctx, constants.CONF_MODULES_DEVREPO)
	if ok {
		modMgr.hotModulesRepo = modulesDevRepo
	} else {
		modMgr.hotModulesRepo = "/devmodulesrepo"
	}

	repoExists, _, _ := utils.FileExists(modulesRepository)
	if repoExists && (availableModules != nil) {
		err = modMgr.loadAvailableModules(ctx, modulesRepository, modulesDir, availableModules)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	ok, fi, _ := utils.FileExists(modulesDir)
	if ok && fi.IsDir() {
		if _, err = modMgr.createInstances(ctx); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	/*
		pluginsConf, ok := conf.GetSubConfig(constants.MODULEPLUGINS)
		if ok {
			plugins := pluginsConf.AllConfigurations()
			for _, pluginName := range plugins {
				pluginConf, _ := pluginsConf.GetSubConfig(pluginName)
				modMgr.modulePlugins[pluginName] = pluginConf
			}
		}
	*/
	return nil
}

func (modMgr *moduleManager) Start(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Module manager start")
	for _, modProxy := range modMgr.moduleInstances {
		err := modMgr.startInstance(ctx, modProxy.(*moduleProxy))
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	err := modMgr.startPlugins(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (modMgr *moduleManager) iterateAndLoadPendingModuleInstances(ctx core.ServerContext, mods map[string]config.Config) ([]string, error) {
	//create pending modules from this iteration
	pendingModuleInstances := make(map[string]config.Config)
	createdInstances := []string{}

	//loop through provided modules
	for instance, instanceConf := range mods {
		loaded, err := modMgr.createModuleInstanceFromConf(ctx, instance, instanceConf, pendingModuleInstances)
		if err != nil {
			return nil, err
		}
		createdInstances = append(createdInstances, instance)
		if !loaded {
			pendingModuleInstances[instance] = instanceConf
		}
	}

	//recurse through pending modules of this iteration
	if len(pendingModuleInstances) == 0 {
		return createdInstances, nil
	}

	continueRecursion := false
	for k, _ := range mods {
		_, ok := pendingModuleInstances[k]
		if !ok {
			//check if atleast one of the initial modules is no longer pending
			continueRecursion = true
			break
		}
	}

	//if no new modules have been loaded in this iteration
	// and there are still modules pending... error out
	if continueRecursion {
		insCreated, err := modMgr.iterateAndLoadPendingModuleInstances(ctx, pendingModuleInstances)
		if err != nil {
			return nil, err
		}
		createdInstances = append(createdInstances, insCreated...)
	} else {
		return nil, errors.DepNotMet(ctx, "Multiple Modules", "Modules", pendingModuleInstances)
	}
	return createdInstances, nil
}

func (modMgr *moduleManager) startInstance(ctx core.ServerContext, ins *moduleProxy) error {
	mod := ins.mod
	startCtx := mod.svrContext.SubContext("Module Start")
	err := mod.start(startCtx)
	if err != nil {
		return errors.WrapError(startCtx, err)
	}
	return nil
}

func (modMgr *moduleManager) getModuleDir(ctx core.ServerContext, modulesDir string, moduleName string) string {
	return path.Join(modulesDir, moduleName)
}

func (modMgr *moduleManager) getModuleConf(ctx core.ServerContext, modDir string) (config.Config, error) {
	return common.NewConfigFromFile(ctx, path.Join(modDir, constants.CONF_CONFIG_DIR, constants.CONF_CONFIG_FILE), nil)
}
