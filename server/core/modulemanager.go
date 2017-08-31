package core

import (
	"fmt"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"

	"github.com/blang/semver"
)

type moduleManager struct {
	name             string
	svrref           *abstractserver
	parent           core.ServerElement
	proxy            server.ModuleManager
	modules          map[string]server.Module
	installedModules map[string]*semver.Version
	moduleConf       map[string]config.Config
	loadedModules    map[string]*semver.Version
	parentModules    map[string]string
	objLoader        *objectLoader
}

func (modMgr *moduleManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Module manager initialized ")
	modulesConfig, err, _ := common.ConfigFileAdapter(ctx, conf, constants.CONF_MODULES)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	modMgr.objLoader = modMgr.svrref.objectLoaderHandle.(*objectLoader)

	baseDir, _ := ctx.GetString(config.BASEDIR)
	modulesDir := path.Join(baseDir, constants.CONF_MODULES)
	ok, fi, _ := utils.FileExists(modulesDir)
	if ok && fi.IsDir() {

		pendingModules := make(map[string]config.Config)
		instances := modulesConfig.AllConfigurations()

		//loop through module instances
		for _, instance := range instances {
			instanceConf, _ := modulesConfig.GetSubConfig(instance)
			log.Info(ctx, "Loading module instance", "Name", instance)

			loaded, err := modMgr.processModuleInstanceConf(ctx, instance, instanceConf, modulesDir, pendingModules)
			if err != nil {
				return err
			}

			if !loaded {
				pendingModules[instance] = instanceConf
			}

		}

		//load pending modules
		if len(pendingModules) > 0 {
			err = modMgr.iterateAndLoadPendingModules(ctx, pendingModules, modulesDir)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
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

func (modMgr *moduleManager) iterateAndLoadPendingModules(ctx core.ServerContext, mods map[string]config.Config, modulesDir string) error {
	//create pending modules from this iteration
	pendingModules := make(map[string]config.Config)

	//loop through provided modules
	for instance, instanceConf := range mods {
		loaded, err := modMgr.processModuleInstanceConf(ctx, instance, instanceConf, modulesDir, pendingModules)
		if err != nil {
			return err
		}
		if !loaded {
			pendingModules[instance] = instanceConf
		}
	}

	//recurse through pending modules of this iteration
	if len(pendingModules) == 0 {
		return nil
	}

	continueRecursion := false
	for k, _ := range mods {
		_, ok := pendingModules[k]
		if !ok {
			//check if atleast one of the initial modules is no longer pending
			continueRecursion = true
			break
		}
	}

	//if no new modules have been loaded in this iteration
	// and there are still modules pending... error out
	if continueRecursion {
		if err := modMgr.iterateAndLoadPendingModules(ctx, pendingModules, modulesDir); err != nil {
			return err
		}
	} else {
		return errors.DepNotMet(ctx, "Multiple Modules", "Modules", pendingModules)
	}
	return nil
}

//adds any modules that need to be instantiated as  a part of another module instance

func (modMgr *moduleManager) addModuleSubInstances(ctx core.ServerContext, instance string, instanceConf config.Config, pendingModules map[string]config.Config) {
	//retInstances := make(map[string]config.Config)
	modInstances, ok := instanceConf.GetSubConfig(constants.CONF_MODULES)
	if ok {
		instanceNames := modInstances.AllConfigurations()
		for _, subinstanceName := range instanceNames {
			subInstanceConf, _ := modInstances.GetSubConfig(subinstanceName)
			newInstanceName := fmt.Sprintf("%s->%s", instance, subinstanceName)
			modMgr.parentModules[newInstanceName] = instance
			log.Trace(ctx, "Sub module added to the load list", "Instance name", newInstanceName, "Conf", subInstanceConf)
			pendingModules[newInstanceName] = subInstanceConf
		}
	}
}

func (modMgr *moduleManager) Start(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Module manager start")
	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		startCtx := mod.svrContext.SubContext("Module Start")
		err := mod.start(startCtx)
		if err != nil {
			return errors.WrapError(startCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) getModuleDir(ctx core.ServerContext, modulesDir string, moduleName string) string {
	return path.Join(modulesDir, moduleName)
}

func (modMgr *moduleManager) getModuleConf(ctx core.ServerContext, modDir string) (config.Config, error) {
	return common.NewConfigFromFile(path.Join(modDir, constants.CONF_CONFIG_FILE))
}

func (modMgr *moduleManager) loadServices(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		svcCtx := mod.svrContext.SubContext("Load Services")
		log.Trace(svcCtx, "Services to process", "Services", mod.services)
		if err := common.ProcessObjects(svcCtx, mod.services, processor); err != nil {
			return errors.WrapError(svcCtx, err)
		}
	}
	return nil
}

/*
processor func(core.ServerContext, config.Config, string) error
*/

func (modMgr *moduleManager) loadExtensions(ctx core.ServerContext) error {

	svcMgr := modMgr.svrref.serviceManagerHandle.(*serviceManager)

	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		modPlugins := mod.plugins(ctx)
		for svcName, pluginConf := range modPlugins {
			pluginObj, err := svcMgr.getService(ctx, svcName)
			if err != nil {
				return errors.BadConf(ctx, constants.MODULEMGR_PLUGIN, "Module Plugin", svcName, "pluginConf", pluginConf)
			}
			pluginSvc := pluginObj.(*serviceProxy)

			plugin, ok := pluginSvc.svc.service.(components.ModuleManagerPlugin)
			if !ok {
				return errors.BadConf(ctx, constants.MODULEMGR_PLUGIN, "Module Plugin", svcName, "pluginConf", pluginConf)
			}

			for passedModName, passedModProxy := range modMgr.modules {
				log.Info(ctx, "Processing module with module manager plugin", "Module", passedModName, "Service name", svcName)
				passedMod := passedModProxy.(*moduleProxy).mod
				passedModCtx := passedMod.svrContext.SubContext("Process module plugin: " + passedModName)
				err := plugin.Load(passedModCtx, passedModName, passedMod.dir, passedMod.userModule, passedMod.modConf, passedMod.modSettings)
				if err != nil {
					return errors.WrapError(passedModCtx, err)
				}
			}
			err = plugin.Loaded(pluginSvc.svc.svrContext)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}

	/*
		for name, pluginConf := range modMgr.modulePlugins {
			svcName, ok := pluginConf.GetString(constants.MODULEPLUGIN_SVC)
			if !ok {
				return errors.MissingConf(ctx, constants.MODULEPLUGIN_SVC, "Module Plugin", name)
			}

		}*/
	return nil
}

func (modMgr *moduleManager) loadFactories(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		facCtx := mod.svrContext.SubContext("Load Factories")
		if err := common.ProcessObjects(facCtx, mod.factories, processor); err != nil {
			return errors.WrapError(facCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadChannels(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		chanCtx := mod.svrContext.SubContext("Load Channels")
		log.Trace(chanCtx, "Channels to process", "channels", mod.channels, "name", mod.name)
		if err := common.ProcessObjects(chanCtx, mod.channels, processor); err != nil {
			return errors.WrapError(chanCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadRules(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		ruleCtx := mod.svrContext.SubContext("Load Rules")
		if err := common.ProcessObjects(ruleCtx, mod.rules, processor); err != nil {
			return errors.WrapError(ruleCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadTasks(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.modules {
		mod := modProxy.(*moduleProxy).mod
		taskCtx := mod.svrContext.SubContext("Load Tasks")
		if err := common.ProcessObjects(taskCtx, mod.tasks, processor); err != nil {
			return errors.WrapError(taskCtx, err)
		}
	}
	return nil
}
