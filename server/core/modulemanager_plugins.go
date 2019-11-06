package core

import (
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
)

func (modMgr *moduleManager) loadPlugins(ctx core.ServerContext) error {

	/*svcMgrPrxy := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
	svcMgr := svcMgrPrxy.(*serviceManagerProxy).manager
	for svcName, svcProxy := range svcMgr.servicesStore {
		svcConf := svcProxy.svc.conf
		isPlugin, _ := svcConf.GetBool(ctx, constants.MODULEMGR_PLUGIN)
		if isPlugin {
			pluginObj, err := svcMgr.getService(ctx, svcName)
			if err != nil {
				return errors.BadConf(ctx, constants.MODULEMGR_PLUGIN, "Module Plugin", svcName, "pluginConf", svcConf)
			}
			pluginSvc := pluginObj.(*serviceProxy)

			plugin, ok := pluginSvc.svc.service.(components.ModuleManagerPlugin)
			if !ok {
				return errors.BadConf(ctx, constants.MODULEMGR_PLUGIN, "Module Plugin", svcName, "pluginConf", svcConf)
			}
			modMgr.modulePlugins[svcName] = plugin
		}
	}*/

	log.Error(ctx, "Going to load mod mgr plugins")

	for insName, modIns := range modMgr.moduleInstances {
		log.Error(ctx, "Testing mods for plugins", "insname", insName, "user module", modIns.userModule)
		if modIns.userModuleObj != nil {
			log.Error(ctx, "Testing mods for plugins not nil")
			plugin, ok := modIns.userModuleObj.(components.ModuleManagerPlugin)
			if ok {
				log.Info(ctx, "Module manager plugin registered ", "plugin", plugin)
				modMgr.modulePlugins[insName] = plugin
			} else {
				log.Error(ctx, "Not module manager plugin", "type", reflect.TypeOf(modIns.userModuleObj))
			}
		}
	}
	log.Error(ctx, "Loaded mod mgr plugins", "plugins", modMgr.modulePlugins)

	return nil
}

func (modMgr *moduleManager) loadInstancesToPluginsforload(ctx core.ServerContext, instances map[string]*serverModule) error {
	for insName, plugin := range modMgr.modulePlugins {
		for _, modIns := range instances {
			err := modMgr.loadPluginWithMod(ctx, modIns, insName, plugin)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}

	for insName, plugin := range modMgr.modulePlugins {
		pluginMod := modMgr.moduleInstances[insName]
		err := plugin.Loaded(pluginMod.svrContext)
		if err != nil {
			return errors.WrapError(ctx, err)
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

/*func (modMgr *moduleManager) processPlugins(ctx core.ServerContext, mod *serverModule, svcMgr *serviceManager, processedMods map[string]bool) error {
	_, ok := processedMods[mod.name]
	if ok {
		return nil
	}

	//dependent modules of the module being processed for plugins
	deps := modMgr.getDependentModules(ctx, mod.name)
	if deps != nil {
		for _, depName := range deps {
			modProxy, ok := modMgr.moduleInstances[depName]
			if ok {
				depmod := modProxy.(*moduleProxy).mod
				modMgr.processPlugins(ctx, depmod, svcMgr, processedMods)
			}
		}
	}
	err := modMgr.updatePlugins(ctx, mod, svcMgr)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	processedMods[mod.name] = true
	return nil
}

func (modMgr *moduleManager) processModulePlugins(ctx core.ServerContext, mod *serverModule, svcMgr *serviceManager) error {
	modPlugins := mod.plugins(ctx)
	for svcName, plugin := range modPlugins {
		log.Info(ctx, "process plugin ", "svc name", svcName)
		modMgr.modulePlugins[svcName] = plugin

		for passedModName, passedModProxy := range modMgr.moduleInstances {
			err := modMgr.loadPluginWithMod(ctx, passedModProxy.(*moduleProxy).mod, svcName, plugin)
			if err != nil {
				return errors.WrapError(passedModCtx, err)
			}
		}
		err = plugin.Loaded(pluginSvc.svc.svrContext)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}*/

func (modMgr *moduleManager) loadPluginWithMod(ctx core.ServerContext, modIns *serverModule, pluginName string, plugin components.ModuleManagerPlugin) error {
	log.Info(ctx, "Processing module with module manager plugin", "Module", modIns.name, "Service name", pluginName)
	modInsCtx := modIns.svrContext.SubContext("Process module plugin: " + pluginName)
	log.Debug(modInsCtx, "Loading plugin with module", "Instance", modIns.name, "Module name", modIns.moduleName, "Settings", modIns.modSettings)

	parentName := ""
	if modIns.parentInstance != nil {
		parentName = modIns.parentInstance.name
	}
	_, hot := modMgr.hotModules[modIns.moduleName]
	modInfo := &components.ModInfo{
		InstanceName:    modIns.name,
		ModName:         modIns.moduleName,
		ModDir:          modIns.dir,
		ParentModName:   parentName,
		Mod:             modIns.impl,
		UserObj:         modIns.userModule,
		ModConf:         modIns.modConf,
		ModSettings:     modIns.modSettings,
		Configurations:  modIns.impl.GetConfigurations(),
		ModProps:        modIns.properties,
		IsExtended:      modIns.isExtended,
		ExtendedModName: modIns.extendedMod,
		ExtendedModConf: modIns.extendedModConf,
		ExtendedModDir:  modIns.extendedModDir,
		Hot:             hot,
	}

	err := plugin.Load(modInsCtx, modInfo)
	if err != nil {
		return errors.WrapError(modInsCtx, err)
	}
	return nil
}
