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
	name      string
	svrref    *abstractserver
	parent    core.ServerElement
	proxy     server.ModuleManager
	modules   map[string]*module
	objLoader *objectLoader
	svcMgr    *serviceManager
	facMgr    *factoryManager
	chanMgr   *channelManager
	rulMgr    *rulesManager
	tskMgr    *taskManager
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
			mods := conf.AllConfigurations()
			for _, mod := range mods {
				params, _ := conf.GetSubConfig(mod)
				disabled, _ := params.GetBool(constants.CONF_MODULE_DISABLED)
				if !disabled {
					modDir := path.Join(modulesDir, mod)
					ok, _, _ := utils.FileExists(modDir)
					if ok {
						if err = modMgr.loadModule(ctx, mod, modDir, params); err != nil {
							return err
						}
					} else {
						modFile := path.Join(modulesDir, fmt.Sprint(mod, ".tar.gz"))
						ok, _, _ := utils.FileExists(modFile)
						if ok {
							if err = archiver.TarGz.Open(modFile, modulesDir); err != nil {
								return errors.WrapError(ctx, err)
							}
							if err = modMgr.loadModule(ctx, mod, modDir, params); err != nil {
								return err
							}
						} else {
							return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Module", mod)
						}
					}
				}
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

func (modMgr *moduleManager) loadModule(ctx core.ServerContext, moduleName, dirPath string, params config.Config) error {
	ctx = ctx.SubContext("Load Module " + moduleName)

	conf, err := common.NewConfigFromFile(path.Join(dirPath, constants.CONF_CONFIG_FILE))
	if err != nil {
		return errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName)
	}

	//moduleparams, _ := conf.GetSubConfig(constants.CONF_MODULE_PARAMS)

	ver, ok := conf.GetString(constants.CONF_MODULE_VER)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_MODULE_VER, "Module", moduleName)
	}
	semver_Ver, err := semver.Parse(ver)
	if err != nil {
		return errors.BadConf(ctx, constants.CONF_MODULE_VER, "Module", moduleName)
	}
	modu := &module{name: moduleName, version: semver_Ver}
	dependencies := make(map[string]semver.Range)

	deps, ok := conf.GetSubConfig(constants.CONF_MODULE_DEP)
	if ok {
		mods := deps.AllConfigurations()
		for _, mod := range mods {
			ra, ok := deps.GetString(mod)
			if !ok {
				return errors.MissingConf(ctx, constants.CONF_MODULE_DEP, "Module", mod)
			}
			r, err := semver.ParseRange(ra)
			if err != nil {
				return errors.WrapError(ctx, err, "Module", mod)
			}
			dependencies[mod] = r
		}
	}
	modu.dependencies = dependencies
	modCtx := modMgr.svrref.svrContext.newContext("Module: " + moduleName)
	modCtx.setElements(core.ContextMap{core.ServerElementModule: modu})

	if err := modMgr.loadModuleDirs(modCtx, moduleName, dirPath); err != nil {
		return err
	}

	if err := modMgr.loadModuleFromObj(modCtx, moduleName, dirPath, conf, params); err != nil {
		return err
	}
	modu.svrContext = modCtx
	modMgr.modules[moduleName] = modu

	return nil
}

func (modMgr *moduleManager) loadModuleDirs(ctx core.ServerContext, dirName, dirPath string) error {
	err := modMgr.objLoader.loadPluginsFolderIfExists(ctx, dirPath)
	if err != nil {
		return errors.WrapError(ctx, err, "Module", dirName)
	}
	err = modMgr.facMgr.loadFactoriesFromFolder(ctx, dirPath)
	if err != nil {
		return errors.WrapError(ctx, err, "Module", dirName)
	}
	err = modMgr.svcMgr.loadServicesFromFolder(ctx, dirPath)
	if err != nil {
		return errors.WrapError(ctx, err, "Module", dirName)
	}
	err = modMgr.chanMgr.loadChannelsFromFolder(ctx, dirPath)
	if err != nil {
		return errors.WrapError(ctx, err, "Module", dirName)
	}

	err = modMgr.rulMgr.loadRulesFromDirectory(ctx, dirPath)
	if err != nil {
		return errors.WrapError(ctx, err, "Module", dirName)
	}

	err = modMgr.tskMgr.loadTasksFromDirectory(ctx, dirPath)
	if err != nil {
		return errors.WrapError(ctx, err, "Module", dirName)
	}
	return nil
}

func (modMgr *moduleManager) loadModuleFromObj(ctx core.ServerContext, dirName, dirPath string, conf config.Config, params config.Config) error {
	objName, ok := conf.GetString(constants.CONF_MODULE_OBJ)
	if ok {
		obj, err := ctx.CreateObject(objName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		mod, ok := obj.(core.Module)
		if !ok {
			return errors.TypeMismatch(ctx, err, "Module", dirName, "Object", objName)
		}
		err = mod.Initialize(params)
		if err != nil {
			return errors.WrapError(ctx, err)
		}

		facs := mod.Factories()
		facCtx := ctx.SubContext("Module factories " + dirName)
		for alias, conf := range facs {
			err = modMgr.facMgr.createServiceFactory(facCtx, conf, alias)
			if err != nil {
				return errors.WrapError(facCtx, err)
			}
		}

		svcs := mod.Services()
		svcCtx := ctx.SubContext("Module services " + dirName)
		for alias, conf := range svcs {
			err = modMgr.svcMgr.createService(svcCtx, conf, alias)
			if err != nil {
				return errors.WrapError(svcCtx, err)
			}
		}

		channels := mod.Channels()
		chanCtx := ctx.SubContext("Module channels " + dirName)
		for channel, conf := range channels {
			err = modMgr.chanMgr.createChannel(chanCtx, conf, channel)
			if err != nil {
				return errors.WrapError(chanCtx, err)
			}
		}

		rules := mod.Rules()
		rulCtx := ctx.SubContext("Module rules " + dirName)
		for rul, conf := range rules {
			err = modMgr.rulMgr.processRuleConf(rulCtx, conf, rul)
			if err != nil {
				return errors.WrapError(rulCtx, err)
			}
		}

		tasks := mod.Tasks()
		tskCtx := ctx.SubContext("Module tasks " + dirName)
		for tsk, conf := range tasks {
			err = modMgr.tskMgr.processTaskConf(tskCtx, conf, tsk)
			if err != nil {
				return errors.WrapError(tskCtx, err)
			}
		}
	}
	return nil
}
