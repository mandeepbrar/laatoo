package core

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

func (modMgr *moduleManager) createInstances(ctx core.ServerContext) (map[string]*serverModule, error) {
	pendingModuleInstances := make(map[string]config.Config)
	createdInstances := make(map[string]*serverModule)
	//loop through module instances
	for instance, instanceConf := range modMgr.moduleInstancesConfig {
		_, ok := modMgr.moduleInstances[instance]
		if ok {
			continue
		}
		log.Error(ctx, "Loading module instance", "Name", instance)

		modins, loaded, err := modMgr.createModuleInstanceFromConf(ctx, instance, instanceConf, pendingModuleInstances)
		if err != nil {
			return nil, err
		}
		createdInstances[instance] = modins
		if !loaded {
			pendingModuleInstances[instance] = instanceConf
		}

	}

	//load pending modules
	if len(pendingModuleInstances) > 0 {
		instancesCreated, err := modMgr.iterateAndLoadPendingModuleInstances(ctx, pendingModuleInstances)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		for k, v := range instancesCreated {
			createdInstances[k] = v
		}
	}
	return createdInstances, nil
}

func (modMgr *moduleManager) createModuleInstanceFromConf(ctx core.ServerContext, instance string, instanceConf config.Config, pendingModuleInstances map[string]config.Config) (*serverModule, bool, error) {
	ctx = ctx.SubContext("Process Instance Conf " + instance)

	//get module to be used
	moduleName, ok := instanceConf.GetString(ctx, constants.CONF_MODULE)
	if !ok {
		moduleName = instance
	}

	_, moduleInstalled := modMgr.installedModules[moduleName]
	modDir, modAvailable := modMgr.availableModules[moduleName]
	if !moduleInstalled {
		return nil, false, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Name", moduleName, "Module Installed", moduleInstalled, "Module Available", modAvailable)
	}

	ctx, fnd, err := modMgr.setupInstanceContext(ctx, instance, instanceConf, modDir)
	if err != nil {
		return nil, false, errors.WrapError(ctx, err)
	}
	if !fnd {
		return nil, false, nil
	}

	disabled, _ := instanceConf.GetBool(ctx, constants.CONF_MODULE_DISABLED)
	if !disabled {
		log.Debug(ctx, "Creating instance", "instance", instance)
		return modMgr.createModuleInstance(ctx, instance, moduleName, modDir, instanceConf, pendingModuleInstances)
	} else {
		log.Debug(ctx, "Instance has been disabled")
	}
	return nil, true, nil
}

func (modMgr *moduleManager) createModuleInstance(ctx core.ServerContext, moduleInstance, moduleName, dirPath string, instanceConf config.Config, pendingModuleInstances map[string]config.Config) (*serverModule, bool, error) {
	ctx = ctx.SubContext("Create Instance " + moduleInstance)

	//get the environment in which module should operate
	modenv, ok := instanceConf.GetSubConfig(ctx, constants.CONF_MODULE_ENV)
	if ok {
		ctx.SetVals(modenv.(common.GenericConfig))
	}

	modConf, confFound := modMgr.moduleConf[moduleName]

	var modSettings config.Config

	if confFound {
		modSettings, _ = modConf.GetSubConfig(ctx, constants.CONF_MODULE_SETTINGS)
	}

	insSettings, settingsFound := instanceConf.GetSubConfig(ctx, constants.CONF_MODULE_SETTINGS)

	if settingsFound {
		//merge modSettings and insSettings
		modSettings = common.Merge(ctx, modSettings, insSettings)
	}

	log.Info(ctx, "Creating module instance", "name", moduleInstance, "Conf", modConf, "Settings", modSettings)
	if confFound && (modSettings != nil) {
		//		ctx.SetVals(modSettings.(common.GenericConfig))
		moduleparams, _ := modConf.GetSubConfig(ctx, constants.CONF_MODULE_PARAMS)
		log.Info(ctx, "Creating module instance ", "Params", moduleparams)
		if moduleparams != nil {
			paramNames := moduleparams.AllConfigurations(ctx)
			for _, paramName := range paramNames {
				val, ok := modSettings.Get(ctx, paramName)
				if ok {
					ctx.Set(paramName, val)
				}
			}
		}
	}

	/*if env != nil {
		envvars := env.AllConfigurations(ctx)
		for _, varname := range envvars {
			varvalue, _ := env.GetString(ctx, varname)
			ctx.Set(varname, varvalue)
		}
	}*/
	/*
		if modSettings != nil {
			moduleparams, _ := modConf.GetSubConfig(ctx, constants.CONF_MODULE_PARAMS)
			if moduleparams != nil {
				paramNames := moduleparams.AllConfigurations(ctx)
				for _, paramName := range paramNames {
					val, ok := modSettings.Get(ctx, paramName)
					if ok {
						ctx.Set(paramName, val)
					}
				}
			}
		}
	*/

	//create a new mod conf with all variables set in the context so that they can be used by yml templates
	modConf, err := modMgr.getModuleConf(ctx, dirPath)
	if err != nil {
		return nil, false, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName)
	}

	log.Info(ctx, "New conf for module instance", "Conf", modConf)

	modMgr.addModuleSubInstances(ctx, moduleInstance, modConf, pendingModuleInstances)

	modu := newServerModule(ctx, moduleInstance, moduleName, dirPath, modConf, modMgr)
	ctx.(*serverContext).setElements(core.ContextMap{core.ServerElementModule: &moduleProxy{mod: modu}})

	objName, ok := modConf.GetString(ctx, constants.CONF_MODULE_OBJ)
	if ok {
		obj, err := ctx.CreateObject(objName)
		if err != nil {
			return nil, false, errors.WrapError(ctx, err)
		}

		usermod, ok := obj.(core.Module)
		if !ok {
			return nil, false, errors.TypeMismatch(ctx, "Module", moduleInstance, "Object", objName)
		}
		modu.userModule = usermod
		modu.objectName = objName
	}

	if err := modu.loadMetaData(ctx); err != nil {
		return nil, false, errors.WrapError(ctx, err)
	}

	initCtx := ctx.SubContext("Initialize Module Instance")
	err = modu.initialize(initCtx, modSettings)
	if err != nil {
		return nil, false, errors.WrapError(initCtx, err)
	}

	modMgr.moduleInstances[moduleInstance] = modu

	return modu, true, nil
}

func (modMgr *moduleManager) setupInstanceContext(ctx core.ServerContext, instance string, instanceConf config.Config, modDir string) (core.ServerContext, bool, error) {
	parentModuleName, pok := modMgr.parentModules[instance]
	var newCtx core.ServerContext
	if pok {
		parentModuleInstance, ok := modMgr.moduleInstances[parentModuleName]
		if !ok {
			return nil, false, nil
		} else {
			newCtx = parentModuleInstance.svrContext.newContext("Module: " + instance)
		}
	} else {
		newCtx = modMgr.svrref.svrContext.newContext("Module: " + instance)
	}

	newCtx.Set(config.MODULEDIR, modDir)

	newCtx.Set(constants.CONF_MODULE, instance)

	if err := processLogging(newCtx.(*serverContext), instanceConf, instance); err != nil {
		return nil, false, errors.WrapError(ctx, err)
	}
	return newCtx, true, nil
}

const (
	OBJECTS = "objects"
)

//adds any modules that need to be instantiated as  a part of another module instance

func (modMgr *moduleManager) addModuleSubInstances(ctx core.ServerContext, instance string, instanceConf config.Config, pendingModuleInstances map[string]config.Config) {
	//retInstances := make(map[string]config.Config)
	modInstances, ok := instanceConf.GetSubConfig(ctx, constants.CONF_MODULES)
	if ok {
		instanceNames := modInstances.AllConfigurations(ctx)
		for _, subinstanceName := range instanceNames {
			subInstanceConf, _ := modInstances.GetSubConfig(ctx, subinstanceName)
			newInstanceName := fmt.Sprintf("%s->%s", instance, subinstanceName)
			modMgr.parentModules[newInstanceName] = instance
			log.Info(ctx, "Sub module added to the load list", "Instance name", newInstanceName, "Conf", subInstanceConf)
			pendingModuleInstances[newInstanceName] = subInstanceConf
		}
	}
}
