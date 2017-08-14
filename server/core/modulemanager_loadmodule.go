package core

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"os"
	"path"

	"github.com/blang/semver"
	"github.com/mholt/archiver"
)

func (modMgr *moduleManager) loadModule(ctx core.ServerContext, modulesDir string, moduleName string) (*semver.Version, error) {
	modver, moduleAlreadyLoaded := modMgr.loadedModules[moduleName]
	if moduleAlreadyLoaded {
		return modver, nil
	}
	log.Error(ctx, "Loading module", "moduleName", moduleName)
	module_ver, err := modMgr.installModule(ctx, modulesDir, moduleName)
	if err != nil {
		return nil, errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, err, "Module", moduleName)
	}

	moduleDir := modMgr.getModuleDir(ctx, modulesDir, moduleName)

	/*conf, err := common.NewConfigFromFile(path.Join(moduleDir, constants.CONF_CONFIG_FILE))
	if err != nil {
		return false, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName)
	}*/

	objsPath := path.Join(moduleDir, constants.CONF_OBJECTLDR_OBJECTS)
	err = modMgr.objLoader.loadPluginsFolderIfExists(ctx, objsPath)
	if err != nil {
		return nil, errors.WrapError(ctx, err, "Module", moduleName)
	}

	modConf, _ := modMgr.moduleConf[moduleName]

	if err = modMgr.processModuleConfMetadata(ctx, modConf); err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	modMgr.loadedModules[moduleName] = module_ver
	return module_ver, nil
}

func (modMgr *moduleManager) installModule(ctx core.ServerContext, modulesDir string, moduleName string) (*semver.Version, error) {
	module_ver, moduleAlreadyInstalled := modMgr.installedModules[moduleName]
	if moduleAlreadyInstalled {
		return module_ver, nil
	}

	//check if module directory exists
	modDir := modMgr.getModuleDir(ctx, modulesDir, moduleName)
	modDirExists, modDirInf, _ := utils.FileExists(modDir)
	modFile := path.Join(modulesDir, fmt.Sprint(moduleName, ".tar.gz"))
	modFileExist, modFileInf, _ := utils.FileExists(modFile)
	log.Debug(ctx, "Processing module conf", "Module", moduleName, "Dir exists", modDirExists, "File exists", modFileExist)
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
					return nil, errors.WrapError(ctx, err)
				}
				log.Debug(ctx, "Deleted old version of module", "Module", moduleName)
			} else {
				extractFile = false
			}
		}
		if extractFile {
			if err := archiver.TarGz.Open(modFile, modulesDir); err != nil {
				return nil, errors.WrapError(ctx, err)
			}
			log.Info(ctx, "Extracted module ", "Module", moduleName, "Module file", modFile, "Destination", modulesDir, "Module directory", modDir)
		}
	} else {
		if !modDirExists {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Name", moduleName)
		}
	}

	mod_version, modconf, err := modMgr.processDependentModules(ctx, modulesDir, modDir, moduleName)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	modMgr.moduleConf[moduleName] = modconf
	modMgr.installedModules[moduleName] = mod_version

	return mod_version, nil
}

func (modMgr *moduleManager) processDependentModules(ctx core.ServerContext, modulesDir, modDir string, moduleName string) (*semver.Version, config.Config, error) {

	modconf, err := modMgr.getModuleConf(ctx, modDir)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName)
	}

	ver, ok := modconf.GetString(constants.CONF_MODULE_VER)
	if !ok {
		return nil, nil, errors.MissingConf(ctx, constants.CONF_MODULE_VER, "Module", moduleName)
	}
	module_ver, err := semver.Parse(ver)
	if err != nil {
		return nil, nil, errors.BadConf(ctx, constants.CONF_MODULE_VER, "Module", moduleName)
	}

	deps, ok := modconf.GetSubConfig(constants.CONF_MODULE_DEP)
	if ok {
		log.Debug(ctx, "Processing module conf", "deps", deps, "moduleName", moduleName)
		mods := deps.AllConfigurations()
		for _, mod := range mods {
			ra, ok := deps.GetString(mod)
			if !ok {
				return nil, nil, errors.MissingConf(ctx, constants.CONF_MODULE_DEP, "Module", moduleName)
			}
			r, err := semver.ParseRange(ra)
			if err != nil {
				return nil, nil, errors.WrapError(ctx, err, "Module", moduleName)
			}

			currentVer, depInstalled := modMgr.installedModules[mod]

			if depInstalled {
				if !r(*currentVer) {
					return nil, nil, errors.DepNotMet(ctx, moduleName, "Dependency", mod, "Version required", ra, "Version installed", currentVer)
				}
			} else {
				currentVer, err = modMgr.loadModule(ctx, modulesDir, mod)
				if err != nil {
					return nil, nil, errors.WrapError(ctx, err, "Module", moduleName)
				}
				if !r(*currentVer) {
					return nil, nil, errors.DepNotMet(ctx, moduleName, "Dependency", mod, "Version required", ra, "Version installed", currentVer)
				}
			}
		}
	}
	return &module_ver, modconf, nil
}
