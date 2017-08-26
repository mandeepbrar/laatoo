package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/server/constants"
)

func (modMgr *moduleManager) processModuleInstanceConf(ctx core.ServerContext, instance string, instanceConf config.Config, modulesDir string,
	pendingModules map[string]config.Config) (bool, error) {

	//get module to be used
	moduleName, ok := instanceConf.GetString(constants.CONF_MODULE)
	if !ok {
		moduleName = instance
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

	ctx.Set(config.MODULEDIR, modMgr.getModuleDir(ctx, modulesDir, moduleName))

	ctx.SetVariable(constants.CONF_MODULE, instance)

	if err := processLogging(ctx.(*serverContext), instanceConf, instance); err != nil {
		return false, errors.WrapError(ctx, err)
	}

	disabled, _ := instanceConf.GetBool(constants.CONF_MODULE_DISABLED)
	if !disabled {
		_, err := modMgr.loadModule(ctx, modulesDir, moduleName)
		if err != nil {
			return false, errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, err, "Module", moduleName)
		}
		return modMgr.createModuleInstance(ctx, instance, moduleName, modMgr.getModuleDir(ctx, modulesDir, moduleName), instanceConf, pendingModules)
	}
	return true, nil
}

func (modMgr *moduleManager) createModuleInstance(ctx core.ServerContext, moduleInstance, moduleName, dirPath string, instanceConf config.Config, pendingModules map[string]config.Config) (bool, error) {
	ctx = ctx.SubContext("Load Module " + moduleInstance)

	modSettings, _ := instanceConf.GetSubConfig(constants.CONF_MODULE_SETTINGS)
	modConf, _ := modMgr.moduleConf[moduleName]

	modMgr.addModuleSubInstances(ctx, moduleInstance, modConf, pendingModules)

	modu := newServerModule(ctx, moduleInstance, dirPath, instanceConf)
	ctx.(*serverContext).setElements(core.ContextMap{core.ServerElementModule: &moduleProxy{mod: modu}})

	objName, ok := modConf.GetString(constants.CONF_MODULE_OBJ)
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

	if modSettings != nil {
		moduleparams, _ := modConf.GetSubConfig(constants.CONF_MODULE_PARAMS)
		if moduleparams != nil {
			paramNames := moduleparams.AllConfigurations()
			for _, paramName := range paramNames {
				val, ok := modSettings.Get(paramName)
				if ok {
					ctx.Set(paramName, val)
				}
			}
		}
	}

	//get the environment in which module should operate
	modenv, _ := instanceConf.GetSubConfig(constants.CONF_MODULE_ENV)

	initCtx := ctx.SubContext("Initialize Module")
	err := modu.initialize(initCtx, modSettings, modenv)
	if err != nil {
		return false, errors.WrapError(initCtx, err)
	}

	modMgr.modules[moduleInstance] = &moduleProxy{mod: modu}

	return true, nil
}

const (
	OBJECTS = "objects"
)

func (modMgr *moduleManager) processModuleConfMetadata(ctx core.ServerContext, conf config.Config) error {
	objsconf, ok := conf.GetSubConfig(OBJECTS)
	if ok {
		objs := objsconf.AllConfigurations()
		for _, objname := range objs {
			objconf, _ := objsconf.GetSubConfig(objname)
			objtyp, _ := objconf.GetString(OBJECT_TYPE)
			var inf core.Info
			switch objtyp {
			case "service":
				inf = buildServiceInfo(objconf)
			case "module":
				inf = buildModuleInfo(objconf)
			case "factory":
				inf = buildFactoryInfo(objconf)
			}
			if inf != nil {
				ldr := ctx.GetServerElement(core.ServerElementLoader).(*objectLoaderProxy).loader
				ldr.setObjectInfo(ctx, objname, inf)
			}
		}
	}
	return nil
}
