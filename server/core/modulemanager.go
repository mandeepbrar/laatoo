package core

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"path"
)

type module struct {
}

type moduleManager struct {
	name       string
	svrref     *abstractserver
	parent     core.ServerElement
	proxy      server.ModuleManager
	modules    map[string]*module
	modulesDir string
	objLoader  *objectLoader
	svcMgr     *serviceManager
	facMgr     *factoryManager
	chanMgr    *channelManager
	rulMgr     *rulesManager
	tskMgr     *taskManager
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
		}
	}
	return nil
}

func (modMgr *moduleManager) Start(ctx core.ServerContext) error {
	return nil
}

func (modMgr *moduleManager) loadModule(ctx core.ServerContext, dirName, dirPath string) error {
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
