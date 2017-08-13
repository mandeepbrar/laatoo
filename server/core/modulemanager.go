package core

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"os"
	"path"

	"github.com/blang/semver"
	"github.com/mholt/archiver"
)

type moduleManager struct {
	name          string
	svrref        *abstractserver
	parent        core.ServerElement
	proxy         server.ModuleManager
	modules       map[string]*serverModule
	loadedModules map[string]semver.Version
	objLoader     *objectLoader
}

func (modMgr *moduleManager) Initialize(ctx core.ServerContext, conf config.Config) error {

	modulesConfig, err, _ := common.ConfigFileAdapter(ctx, conf, constants.CONF_MODULES)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	modMgr.objLoader = modMgr.svrref.objectLoaderHandle.(*objectLoader)

	baseDir, _ := ctx.GetString(constants.CONF_BASE_DIR)
	modulesDir := path.Join(baseDir, constants.CONF_MODULES)
	ok, fi, _ := utils.FileExists(modulesDir)
	if ok && fi.IsDir() {

		pendingModules := make(map[string]config.Config)
		instances := modulesConfig.AllConfigurations()

		//loop through module instances
		for _, instance := range instances {
			instanceConf, _ := modulesConfig.GetSubConfig(instance)
			modMgr.addModuleSubInstances(ctx, instanceConf, pendingModules)
			loaded, err := modMgr.processModuleInstanceConf(ctx, instance, instanceConf, modulesDir)
			if err != nil {
				return err
			}

			if !loaded {
				pendingModules[instance] = instanceConf
			}

		}

		//load pending modules
		if len(pendingModules) > 0 {
			return modMgr.iterateAndLoadPendingModules(ctx, pendingModules, modulesDir)
		}
	}
	return nil
}

func (modMgr *moduleManager) iterateAndLoadPendingModules(ctx core.ServerContext, mods map[string]config.Config, modulesDir string) error {
	//create pending modules from this iteration
	pendingModules := make(map[string]config.Config)

	//loop through provided modules
	for instance, instanceConf := range mods {
		modMgr.addModuleSubInstances(ctx, instanceConf, pendingModules)
		loaded, err := modMgr.processModuleInstanceConf(ctx, instance, instanceConf, modulesDir)
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

	//if no new modules have been loaded in this iteration
	// and there are still modules pending... error out
	if len(pendingModules) < len(mods) {
		if err := modMgr.iterateAndLoadPendingModules(ctx, pendingModules, modulesDir); err != nil {
			return err
		}
	} else {
		return errors.DepNotMet(ctx, "Multiple Modules", "Modules", pendingModules)
	}
	return nil
}

//adds any modules that need to be instantiated as  a part of another module instance

func (modMgr *moduleManager) addModuleSubInstances(ctx core.ServerContext, instanceConf config.Config, pendingModules map[string]config.Config) {
	//retInstances := make(map[string]config.Config)
	modInstances, ok := instanceConf.GetSubConfig(constants.CONF_MODULES)
	if ok {
		instanceNames := modInstances.AllConfigurations()
		for _, instanceName := range instanceNames {
			subInstanceConf, _ := modInstances.GetSubConfig(instanceName)
			pendingModules[instanceName] = subInstanceConf
		}
	}
}

func (modMgr *moduleManager) Start(ctx core.ServerContext) error {
	for _, mod := range modMgr.modules {
		startCtx := mod.svrContext.SubContext("Module Start")
		err := mod.start(startCtx)
		if err != nil {
			return errors.WrapError(startCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) processModuleInstanceConf(ctx core.ServerContext, instance string, instanceConf config.Config, modulesDir string) (bool, error) {
	//get module to be used
	mod, ok := instanceConf.GetString(constants.CONF_MODULE)
	if !ok {
		mod = instance
	}

	ctx = modMgr.svrref.svrContext.newContext("Module: " + instance)
	ctx.SetVariable(constants.CONF_MODULE, instance)

	if err := processLogging(ctx.(*serverContext), instanceConf, instance); err != nil {
		return false, errors.WrapError(ctx, err)
	}

	disabled, _ := instanceConf.GetBool(constants.CONF_MODULE_DISABLED)
	if !disabled {
		//check if module directory exists
		modDir := path.Join(modulesDir, mod)
		modDirExists, modDirInf, _ := utils.FileExists(modDir)
		modFile := path.Join(modulesDir, fmt.Sprint(mod, ".tar.gz"))
		modFileExist, modFileInf, _ := utils.FileExists(modFile)
		log.Debug(ctx, "Processing module conf", "Module", mod, "Dir exists", modDirExists, "File exists", modFileExist)
		if modFileExist {

			//ensure latest module directory is present
			// if zip file is newer than module dir
			// delete the directory and extract latest zip file

			extractFile := true
			if modDirExists {
				tim := modDirInf.ModTime().Sub(modFileInf.ModTime())
				if tim < 0 {
					err := os.RemoveAll(modDir)
					if err != nil {
						return false, errors.WrapError(ctx, err)
					}
					log.Debug(ctx, "Deleted old version of module", "Module", mod)
				} else {
					extractFile = false
				}
			}
			if extractFile {
				if err := archiver.TarGz.Open(modFile, modulesDir); err != nil {
					return false, errors.WrapError(ctx, err)
				}
				log.Info(ctx, "Extracted module ", "Module", mod, "Module file", modFile, "Destination", modulesDir, "Module directory", modDir)
			}
			//create a new module instance with provided settings
			return modMgr.createModuleInstance(ctx, instance, mod, modDir, instanceConf)
		} else {
			if modDirExists {
				//create a new module instance with provided settings
				return modMgr.createModuleInstance(ctx, instance, mod, modDir, instanceConf)
			} else {
				return false, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Module", mod)
			}
		}
	}
	return true, nil
}

func (modMgr *moduleManager) createModuleInstance(ctx core.ServerContext, moduleInstance, moduleName, dirPath string, instanceConf config.Config) (bool, error) {
	ctx = ctx.SubContext("Load Module " + moduleInstance)

	modSettings, _ := instanceConf.GetSubConfig(constants.CONF_MODULE_SETTINGS)

	_, moduleAlreadLoaded := modMgr.loadedModules[moduleName]

	conf, err := common.NewConfigFromFile(path.Join(dirPath, constants.CONF_CONFIG_FILE))
	if err != nil {
		return false, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleInstance)
	}

	modu := newServerModule(ctx, moduleInstance, dirPath, instanceConf)
	ctx.(*serverContext).setElements(core.ContextMap{core.ServerElementModule: &moduleProxy{mod: modu}})

	if !moduleAlreadLoaded {
		objsPath := path.Join(dirPath, constants.CONF_OBJECTLDR_OBJECTS)
		err := modMgr.objLoader.loadPluginsFolderIfExists(ctx, objsPath)
		if err != nil {
			return false, errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}

	if err = modMgr.processModuleConfMetadata(ctx, conf); err != nil {
		return false, errors.WrapError(ctx, err)
	}

	objName, ok := conf.GetString(constants.CONF_MODULE_OBJ)
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

	ver, ok := conf.GetString(constants.CONF_MODULE_VER)
	if !ok {
		return false, errors.MissingConf(ctx, constants.CONF_MODULE_VER, "Module", moduleInstance)
	}
	semver_Ver, err := semver.Parse(ver)
	if err != nil {
		return false, errors.BadConf(ctx, constants.CONF_MODULE_VER, "Module", moduleInstance)
	}
	modu.version = semver_Ver

	dependencies := make(map[string]semver.Range)

	deps, ok := conf.GetSubConfig(constants.CONF_MODULE_DEP)
	if ok {
		mods := deps.AllConfigurations()
		for _, mod := range mods {
			ra, ok := deps.GetString(mod)
			if !ok {
				return false, errors.MissingConf(ctx, constants.CONF_MODULE_DEP, "Module", moduleInstance)
			}
			r, err := semver.ParseRange(ra)
			if err != nil {
				return false, errors.WrapError(ctx, err, "Module", moduleInstance)
			}
			dependencies[moduleInstance] = r

			currentVer, deploaded := modMgr.loadedModules[mod]
			if !deploaded || !r(currentVer) {
				return false, nil
			}
		}
	}

	modu.dependencies = dependencies

	moduleparams, _ := conf.GetSubConfig(constants.CONF_MODULE_PARAMS)
	if moduleparams != nil {
		paramNames := moduleparams.AllConfigurations()
		for _, paramName := range paramNames {
			val, ok := modSettings.Get(paramName)
			if ok {
				ctx.Set(paramName, val)
			}
		}
	}

	//get the environment in which module should operate
	modenv, _ := instanceConf.GetSubConfig(constants.CONF_MODULE_ENV)

	initCtx := ctx.SubContext("Initialize Module")
	err = modu.initialize(initCtx, modSettings, modenv)
	if err != nil {
		return false, errors.WrapError(initCtx, err)
	}

	if !moduleAlreadLoaded {
		modMgr.loadedModules[moduleName] = semver_Ver
	}

	modMgr.modules[moduleInstance] = modu

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

func (modMgr *moduleManager) loadServices(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, mod := range modMgr.modules {
		svcCtx := mod.svrContext.SubContext("Load Services")
		log.Trace(svcCtx, "Services to process", "Services", mod.services)
		if err := common.ProcessObjects(svcCtx, mod.services, processor); err != nil {
			return errors.WrapError(svcCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadFactories(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, mod := range modMgr.modules {
		facCtx := mod.svrContext.SubContext("Load Factories")
		if err := common.ProcessObjects(facCtx, mod.factories, processor); err != nil {
			return errors.WrapError(facCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadChannels(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, mod := range modMgr.modules {
		chanCtx := mod.svrContext.SubContext("Load Channels")
		log.Trace(chanCtx, "Channels to process", "channels", mod.channels, "name", mod.name)
		if err := common.ProcessObjects(chanCtx, mod.channels, processor); err != nil {
			return errors.WrapError(chanCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadRules(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, mod := range modMgr.modules {
		ruleCtx := mod.svrContext.SubContext("Load Rules")
		if err := common.ProcessObjects(ruleCtx, mod.rules, processor); err != nil {
			return errors.WrapError(ruleCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadTasks(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, mod := range modMgr.modules {
		taskCtx := mod.svrContext.SubContext("Load Tasks")
		if err := common.ProcessObjects(taskCtx, mod.tasks, processor); err != nil {
			return errors.WrapError(taskCtx, err)
		}
	}
	return nil
}
