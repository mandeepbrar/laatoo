package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

func (modMgr *moduleManager) processModuleInstanceConf(ctx core.ServerContext, instance string, instanceConf config.Config, pendingModules map[string]config.Config) (bool, error) {
	ctx = ctx.SubContext("Process Instance Conf " + instance)

	//get module to be used
	moduleName, ok := instanceConf.GetString(ctx, constants.CONF_MODULE)
	if !ok {
		moduleName = instance
	}

	_, moduleInstalled := modMgr.installedModules[moduleName]
	modDir, modAvailable := modMgr.availableModules[moduleName]
	if !moduleInstalled {
		return false, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Name", moduleName, "Module Installed", moduleInstalled, "Module Available", modAvailable)
	}

	parentModuleName, pok := modMgr.parentModules[instance]
	if pok {
		parentModule, ok := modMgr.modules[parentModuleName]
		if !ok {
			return false, nil
		} else {
			parentmod := parentModule.(*moduleProxy).mod
			ctx = parentmod.svrContext.newContext("Module: " + instance)
		}
	} else {
		ctx = modMgr.svrref.svrContext.newContext("Module: " + instance)
	}

	ctx.Set(config.MODULEDIR, modDir)

	ctx.Set(constants.CONF_MODULE, instance)

	if err := processLogging(ctx.(*serverContext), instanceConf, instance); err != nil {
		return false, errors.WrapError(ctx, err)
	}

	disabled, _ := instanceConf.GetBool(ctx, constants.CONF_MODULE_DISABLED)
	if !disabled {
		return modMgr.createModuleInstance(ctx, instance, moduleName, modDir, instanceConf, pendingModules)
	} else {
		log.Debug(ctx, "Instance has been disabled")
	}
	return true, nil
}

func (modMgr *moduleManager) createModuleInstance(ctx core.ServerContext, moduleInstance, moduleName, dirPath string, instanceConf config.Config, pendingModules map[string]config.Config) (bool, error) {
	ctx = ctx.SubContext("Create Instance " + moduleInstance)

	//get the environment in which module should operate
	modenv, ok := instanceConf.GetSubConfig(ctx, constants.CONF_MODULE_ENV)
	if ok {
		ctx.SetVals(modenv.(common.GenericConfig))
	}

	modConf, confFound := modMgr.moduleConf[moduleName]
	modSettings, settingsFound := instanceConf.GetSubConfig(ctx, constants.CONF_MODULE_SETTINGS)
	log.Error(ctx, "Creating module instance", "Conf", modConf, "Settings", modSettings)
	if confFound && settingsFound {
		//		ctx.SetVals(modSettings.(common.GenericConfig))
		moduleparams, _ := modConf.GetSubConfig(ctx, constants.CONF_MODULE_PARAMS)
		log.Error(ctx, "Creating module instance ", "Params", moduleparams)
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
		return false, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName)
	}

	modMgr.addModuleSubInstances(ctx, moduleInstance, modConf, pendingModules)

	modu := newServerModule(ctx, moduleInstance, moduleName, dirPath, modConf)
	ctx.(*serverContext).setElements(core.ContextMap{core.ServerElementModule: &moduleProxy{mod: modu}})

	objName, ok := modConf.GetString(ctx, constants.CONF_MODULE_OBJ)
	if ok {
		obj, err := ctx.CreateObject(objName)
		if err != nil {
			return false, errors.WrapError(ctx, err)
		}

		usermod, ok := obj.(core.Module)
		if !ok {
			return false, errors.TypeMismatch(ctx, "Module", moduleInstance, "Object", objName)
		}
		modu.userModule = usermod
		modu.objectName = objName
	}

	if err := modu.loadMetaData(ctx); err != nil {
		return false, errors.WrapError(ctx, err)
	}

	initCtx := ctx.SubContext("Initialize Module")
	err = modu.initialize(initCtx, modSettings)
	if err != nil {
		return false, errors.WrapError(initCtx, err)
	}

	modMgr.modules[moduleInstance] = &moduleProxy{mod: modu}

	return true, nil
}

const (
	OBJECTS = "objects"
)
