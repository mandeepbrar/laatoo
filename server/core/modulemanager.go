package core

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"

	"github.com/blang/semver"
	"github.com/mholt/archiver"
)

type moduleManager struct {
	name          string
	svrref        *abstractserver
	parent        core.ServerElement
	proxy         server.ModuleManager
	modules       map[string]*module
	loadedModules map[string]semver.Version
	objLoader     *objectLoader
	svcMgr        *serviceManager
	facMgr        *factoryManager
	chanMgr       *channelManager
	rulMgr        *rulesManager
	tskMgr        *taskManager
}

func (modMgr *moduleManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	modMgr.svcMgr = modMgr.svrref.serviceManagerHandle.(*serviceManager)
	modMgr.objLoader = modMgr.svrref.objectLoaderHandle.(*objectLoader)
	modMgr.facMgr = modMgr.svrref.factoryManagerHandle.(*factoryManager)
	modMgr.rulMgr = modMgr.svrref.rulesManagerHandle.(*rulesManager)
	modMgr.tskMgr = modMgr.svrref.taskManagerHandle.(*taskManager)
	modMgr.chanMgr = modMgr.svrref.channelManagerHandle.(*channelManager)

	baseDir, _ := ctx.GetString(constants.CONF_BASE_DIR)
	modulesDir := path.Join(baseDir, constants.CONF_MODULES_DIR)
	ok, fi, _ := utils.FileExists(modulesDir)
	if ok && fi.IsDir() {

		modulesConfig := path.Join(modulesDir, constants.CONF_CONFIG_FILE)
		conf, err := common.NewConfigFromFile(modulesConfig)
		if err == nil {
			pendingModules := make(map[string]config.Config)
			instances := conf.AllConfigurations()
			for _, instance := range instances {
				params, _ := conf.GetSubConfig(instance)
				loaded, err := modMgr.loadInstance(ctx, instance, params, modulesDir)
				if err != nil {
					return err
				}
				if !loaded {
					pendingModules[instance] = params
				}
			}
			if len(pendingModules) > 0 {
				return modMgr.iterateAndLoadPendingModules(ctx, pendingModules, modulesDir)
			}
		}

		/*
			files, err := ioutil.ReadDir(modulesDir)
			if err != nil {
				return errors.WrapError(ctx, err, "Subdirectory", modulesDir)
			}
			for _, info := range files {
				fileName := info.Name()
				if info.IsDir() {
					if err = modMgr.loadModule(ctx, fileName, path.Join(modulesDir, fileName)); err != nil {
						return err
					}
				}
			}*/
	}
	return nil
}

func (modMgr *moduleManager) iterateAndLoadPendingModules(ctx core.ServerContext, mods map[string]config.Config, modulesDir string) error {
	pendingModules := make(map[string]config.Config)
	for instance, params := range mods {
		loaded, err := modMgr.loadInstance(ctx, instance, params, modulesDir)
		if err != nil {
			return err
		}
		if !loaded {
			pendingModules[instance] = params
		}
	}
	if len(pendingModules) > 0 && len(pendingModules) < len(mods) {
		if err := modMgr.iterateAndLoadPendingModules(ctx, pendingModules, modulesDir); err != nil {
			return err
		}
	} else {
		return errors.DepNotMet(ctx, "Multiple Modules", "Modules", pendingModules)
	}
	return nil
}

func (modMgr *moduleManager) Start(ctx core.ServerContext) error {
	for modName, module := range modMgr.modules {
		for dep, semver_range := range module.dependencies {
			dependency, ok := modMgr.modules[dep]
			if !ok {
				return errors.DepNotMet(ctx, dep, "Module Name", modName)
			}
			if !semver_range(dependency.version) {
				return errors.DepNotMet(ctx, dep, "Module Name", modName)
			}
		}
	}
	return nil
}

func (modMgr *moduleManager) loadInstance(ctx core.ServerContext, instance string, params config.Config, modulesDir string) (bool, error) {
	mod, ok := params.GetString(constants.CONF_MODULE)
	if !ok {
		mod = instance
	}
	disabled, _ := params.GetBool(constants.CONF_MODULE_DISABLED)
	if !disabled {
		modDir := path.Join(modulesDir, mod)
		ok, _, _ := utils.FileExists(modDir)
		if ok {
			return modMgr.loadModule(ctx, instance, mod, modDir, params)
		} else {
			modFile := path.Join(modulesDir, fmt.Sprint(mod, ".tar.gz"))
			ok, _, _ := utils.FileExists(modFile)
			if ok {
				if err := archiver.TarGz.Open(modFile, modulesDir); err != nil {
					return false, errors.WrapError(ctx, err)
				}
				return modMgr.loadModule(ctx, instance, mod, modDir, params)
			} else {
				return false, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Module", mod)
			}
		}
	}
	return true, nil
}

func (modMgr *moduleManager) loadModule(ctx core.ServerContext, moduleInstance, moduleName, dirPath string, params config.Config) (bool, error) {
	ctx = ctx.SubContext("Load Module " + moduleInstance)

	conf, err := common.NewConfigFromFile(path.Join(dirPath, constants.CONF_CONFIG_FILE))
	if err != nil {
		return false, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleInstance)
	}

	ver, ok := conf.GetString(constants.CONF_MODULE_VER)
	if !ok {
		return false, errors.MissingConf(ctx, constants.CONF_MODULE_VER, "Module", moduleInstance)
	}
	semver_Ver, err := semver.Parse(ver)
	if err != nil {
		return false, errors.BadConf(ctx, constants.CONF_MODULE_VER, "Module", moduleInstance)
	}
	modu := &module{name: moduleInstance, version: semver_Ver}
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
		}
	}
	modu.dependencies = dependencies
	modCtx := modMgr.svrref.svrContext.newContext("Module: " + moduleInstance)
	modCtx.setElements(core.ContextMap{core.ServerElementModule: modu})
	modCtx.SetVariable(constants.CONF_MODULE, moduleInstance)

	moduleparams, _ := conf.GetSubConfig(constants.CONF_MODULE_PARAMS)

	paramNames := moduleparams.AllConfigurations()
	for _, paramName := range paramNames {
		val, ok := params.Get(paramName)
		if ok {
			modCtx.Set(paramName, val)
		}
	}

	_, moduleAlreadLoaded := modMgr.loadedModules[moduleName]
	if !moduleAlreadLoaded {
		modMgr.loadedModules[moduleName] = semver_Ver
	}

	if err := modMgr.loadModuleDirs(modCtx, moduleInstance, dirPath, params, moduleAlreadLoaded); err != nil {
		return false, err
	}

	if err := modMgr.loadModuleFromObj(modCtx, moduleInstance, dirPath, conf, params, moduleAlreadLoaded); err != nil {
		return false, err
	}
	modu.svrContext = modCtx
	modMgr.modules[moduleInstance] = modu

	return true, nil
}

func (modMgr *moduleManager) loadModuleDirs(ctx core.ServerContext, moduleInstance, dirPath string, params config.Config, moduleAlreadLoaded bool) error {
	var err error
	if !moduleAlreadLoaded {
		err := modMgr.objLoader.loadPluginsFolderIfExists(ctx, dirPath)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}

	factoriesEnabled, ok := params.GetBool(constants.CONF_FACTORIES)

	if !ok || factoriesEnabled {
		err = modMgr.facMgr.loadFactoriesFromFolder(ctx, dirPath)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}

	servicesEnabled, ok := params.GetBool(constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		err = modMgr.svcMgr.loadServicesFromFolder(ctx, dirPath)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}

	channelsEnabled, ok := params.GetBool(constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		err = modMgr.chanMgr.loadChannelsFromFolder(ctx, dirPath)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}

	rulesEnabled, ok := params.GetBool(constants.CONF_RULES)

	if !ok || rulesEnabled {
		err = modMgr.rulMgr.loadRulesFromDirectory(ctx, dirPath)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}

	tasksEnabled, ok := params.GetBool(constants.CONF_TASKS)
	if !ok || tasksEnabled {
		err = modMgr.tskMgr.loadTasksFromDirectory(ctx, dirPath)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", moduleInstance)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadModuleFromObj(ctx core.ServerContext, moduleInstance, dirPath string, conf config.Config, params config.Config, moduleAlreadLoaded bool) error {
	objName, ok := conf.GetString(constants.CONF_MODULE_OBJ)
	if ok {

		if !moduleAlreadLoaded {
			err := modMgr.objLoader.loadPluginsFolderIfExists(ctx, dirPath)
			if err != nil {
				return errors.WrapError(ctx, err, "Module", moduleInstance)
			}
		}

		obj, err := ctx.CreateObject(objName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		mod, ok := obj.(core.Module)
		if !ok {
			return errors.TypeMismatch(ctx, err, "Module", moduleInstance, "Object", objName)
		}
		err = mod.Initialize(params)
		if err != nil {
			return errors.WrapError(ctx, err)
		}

		factoriesEnabled, ok := params.GetBool(constants.CONF_FACTORIES)

		if !ok || factoriesEnabled {
			facs := mod.Factories()
			facCtx := ctx.SubContext("Module factories " + moduleInstance)
			for alias, conf := range facs {
				err = modMgr.facMgr.createServiceFactory(facCtx, conf, alias)
				if err != nil {
					return errors.WrapError(facCtx, err)
				}
			}
		}

		servicesEnabled, ok := params.GetBool(constants.CONF_SERVICES)

		if !ok || servicesEnabled {
			svcs := mod.Services()
			svcCtx := ctx.SubContext("Module services " + moduleInstance)
			for alias, conf := range svcs {
				err = modMgr.svcMgr.createService(svcCtx, conf, alias)
				if err != nil {
					return errors.WrapError(svcCtx, err)
				}
			}
		}

		channelsEnabled, ok := params.GetBool(constants.CONF_CHANNELS)

		if !ok || channelsEnabled {
			channels := mod.Channels()
			chanCtx := ctx.SubContext("Module channels " + moduleInstance)
			for channel, conf := range channels {
				err = modMgr.chanMgr.createChannel(chanCtx, conf, channel)
				if err != nil {
					return errors.WrapError(chanCtx, err)
				}
			}
		}

		rulesEnabled, ok := params.GetBool(constants.CONF_RULES)

		if !ok || rulesEnabled {
			rules := mod.Rules()
			rulCtx := ctx.SubContext("Module rules " + moduleInstance)
			for rul, conf := range rules {
				err = modMgr.rulMgr.processRuleConf(rulCtx, conf, rul)
				if err != nil {
					return errors.WrapError(rulCtx, err)
				}
			}
		}

		tasksEnabled, ok := params.GetBool(constants.CONF_TASKS)
		if !ok || tasksEnabled {
			tasks := mod.Tasks()
			tskCtx := ctx.SubContext("Module tasks " + moduleInstance)
			for tsk, conf := range tasks {
				err = modMgr.tskMgr.processTaskConf(tskCtx, conf, tsk)
				if err != nil {
					return errors.WrapError(tskCtx, err)
				}
			}
		}
	}
	return nil
}
