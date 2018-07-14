package core

/*
import (
	"fmt"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"os"
	"path"
	"strings"

	"github.com/blang/semver"
	"github.com/mholt/archiver"
)

func (modMgr *moduleManager) loadAvailableModules(ctx core.ServerContext, modulesRepo string) error {

	files, err := ioutil.ReadDir(modulesRepo)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tar.gz") {
			modName := strings.TrimSuffix(file.Name(), ".tar.gz")
			modver, moduleAlreadyLoaded := modMgr.availableModules[modName]
			if moduleAlreadyLoaded {
				log.Info(ctx, "Module already loaded", "Name", modName, "Version", modver)
				continue
			} else {
				ver, err := modMgr.loadModule(ctx, modulesRepo, modName)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				modMgr.availableModules[modName] = ver
			}
		}
	}
	return nil
}

func (modMgr *moduleManager) loadModule(ctx core.ServerContext, modulesDir string, moduleName string) (*semver.Version, error) {
	log.Info(ctx, "Loading module", "moduleName", moduleName)
	module_ver, err := modMgr.installModule(ctx, modulesDir, moduleName)
	if err != nil {
		return nil, errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, err, "Module", moduleName)
	}

	moduleDir := modMgr.getModuleDir(ctx, modulesDir, moduleName)



	objsPath := path.Join(moduleDir, constants.CONF_OBJECTLDR_OBJECTS)
	err = modMgr.objLoader.loadPluginsFolderIfExists(ctx, objsPath)
	if err != nil {
		return nil, errors.WrapError(ctx, err, "Module", moduleName)
	}

	modConf, _ := modMgr.moduleConf[moduleName]

	if err = modMgr.processModuleConfMetadata(ctx, modConf); err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return module_ver, nil
}

func (modMgr *moduleManager) installModule(ctx core.ServerContext, modulesDir string, moduleName string) (*semver.Version, error) {
	//check if module directory exists
	modDir := modMgr.getModuleDir(ctx, modulesDir, moduleName)
	modDirExists, modDirInf, _ := utils.FileExists(modDir)
	modFile := path.Join(modMgr.modulesRepo, fmt.Sprint(moduleName, ".tar.gz"))
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

func (modMgr *moduleManager) getDependentModules(ctx core.ServerContext, moduleName string) []string {
	modConf, ok := modMgr.moduleConf[moduleName]
	if ok {
		deps, ok := modConf.GetSubConfig(ctx, constants.CONF_MODULE_DEP)
		if ok {
			return deps.AllConfigurations(ctx)
		}
	}
	return nil
}

func (modMgr *moduleManager) processDependentModules(ctx core.ServerContext, modulesDir, modDir string, moduleName string) (*semver.Version, config.Config, error) {

	modconf, err := modMgr.getModuleConf(ctx, modDir)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName)
	}

	ver, ok := modconf.GetString(ctx, constants.CONF_MODULE_VER)
	if !ok {
		return nil, nil, errors.MissingConf(ctx, constants.CONF_MODULE_VER, "Module", moduleName)
	}
	module_ver, err := semver.Parse(ver)
	if err != nil {
		return nil, nil, errors.BadConf(ctx, constants.CONF_MODULE_VER, "Module", moduleName)
	}

	deps, ok := modconf.GetSubConfig(ctx, constants.CONF_MODULE_DEP)
	if ok {
		log.Debug(ctx, "Processing module conf", "deps", deps, "moduleName", moduleName)
		mods := deps.AllConfigurations(ctx)
		for _, mod := range mods {
			ra, ok := deps.GetString(ctx, mod)
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
*/
