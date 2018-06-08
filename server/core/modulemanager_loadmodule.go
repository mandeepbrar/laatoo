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

func (modMgr *moduleManager) loadAvailableModules(ctx core.ServerContext, modulesRepo, modulesDir string, modulesToInstall []string) error {

	err := modMgr.extractArchives(ctx, modulesRepo, modulesDir, modulesToInstall)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	for _, moduleName := range modulesToInstall {
		modDir := modMgr.getModuleDir(ctx, modulesDir, moduleName)
		modDirExists, modDirInf, _ := utils.FileExists(modDir)
		if !modDirExists || !modDirInf.IsDir() {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Module", moduleName)
		}
		installedVer, installed := modMgr.installedModules[moduleName]
		if installed && installedVer != nil {
			log.Info(ctx, "Module already installed", "Name", moduleName, "Version", installedVer)
		} else {
			//path, moduleAlreadyAvailable := modMgr.availableModules[moduleName]
			modMgr.availableModules[moduleName] = modDir
			modMgr.installedModules[moduleName] = nil
		}
	}

	err = modMgr.installModules(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (modMgr *moduleManager) extractArchives(ctx core.ServerContext, modulesRepo, modulesDir string, modulesToInstall []string) error {
	ctx = ctx.SubContext("Extracting archives ")
	for _, modToExtract := range modulesToInstall {
		modFile := path.Join(modMgr.modulesRepo, fmt.Sprint(modToExtract, ".tar.gz"))
		archiveExists, modFileInf, _ := utils.FileExists(modFile)
		if archiveExists {
			modDir := modMgr.getModuleDir(ctx, modulesDir, modToExtract)
			modDirExists, modDirInf, _ := utils.FileExists(modDir)
			log.Debug(ctx, "Extracting archive", "Module", modToExtract, "Dir exists", modDirExists)
			//ensure latest module directory is present
			// if zip file is newer than module dir
			// delete the directory and extract latest zip file
			extractFile := true
			if modDirExists {
				tim := modDirInf.ModTime().Sub(modFileInf.ModTime())
				if tim < 0 {
					err := os.RemoveAll(modDir)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
					log.Debug(ctx, "Deleted old version of module", "Module", modToExtract)
				} else {
					extractFile = false
				}
			}
			if extractFile {
				if err := archiver.TarGz.Open(modFile, modulesDir); err != nil {
					return errors.WrapError(ctx, err)
				}
				log.Info(ctx, "Extracted module ", "Module", modToExtract, "Module file", modFile, "Repo", modulesRepo, "Destination", modulesDir, "Module directory", modDir)
			}
		} else {
			log.Error(ctx, "Archive not found for module", "Module", modToExtract)
		}
	}

	return nil
}

func (modMgr *moduleManager) installModules(ctx core.ServerContext) error {
	for mod, ver := range modMgr.installedModules {
		if ver == nil {
			modDir := modMgr.availableModules[mod]
			_, _, err := modMgr.installModule(ctx, mod, modDir)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

func (modMgr *moduleManager) installModule(ctx core.ServerContext, moduleName, modDir string) (*semver.Version, config.Config, error) {
	ctx = ctx.SubContext("Installing module " + moduleName)

	modconf, err := modMgr.getModuleConf(ctx, modDir)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err, "Info", "Error in opening config", "Module", moduleName, "dir", modDir)
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
		log.Debug(ctx, "Processing module dependencies", "deps", deps, "moduleName", moduleName)
		mods := deps.AllConfigurations(ctx)
		for _, dep := range mods {
			depra, ok := deps.GetString(ctx, dep)
			if !ok {
				return nil, nil, errors.MissingConf(ctx, constants.CONF_MODULE_DEP, "Module", moduleName)
			}
			depr, err := semver.ParseRange(depra)
			if err != nil {
				return nil, nil, errors.WrapError(ctx, err, "Module", moduleName)
			}

			depVer, _ := modMgr.installedModules[dep]

			if depVer != nil {
				if !depr(*depVer) {
					return nil, nil, errors.DepNotMet(ctx, moduleName, "Dependency", dep, "Version required", depra, "Version installed", depVer)
				}
			} else {
				depDir, depAvailable := modMgr.availableModules[dep]
				if !depAvailable {
					return nil, nil, errors.DepNotMet(ctx, moduleName, "Dependency", dep)
				}
				installedDepVer, _, err := modMgr.installModule(ctx, dep, depDir)
				if err != nil {
					return nil, nil, errors.WrapError(ctx, err, "Info", "Could not install dependency", "Dependency", dep, "Dir", depDir)
				}
				if !depr(*installedDepVer) {
					return nil, nil, errors.DepNotMet(ctx, moduleName, "Dependency", dep, "Version required", depra, "Version installed", installedDepVer)
				}
			}
		}
	}

	log.Debug(ctx, "Module dependencies resolved. Loading Objects", "moduleName", moduleName)

	objsPath := path.Join(modDir, constants.CONF_OBJECTLDR_OBJECTS)
	err = modMgr.objLoader.loadPluginsFolderIfExists(ctx, objsPath)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err, "Module", moduleName)
	}

	log.Debug(ctx, "Objects Loaded. Building object info", "moduleName", moduleName)

	if err = modMgr.buildObjectInfo(ctx, modconf); err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}

	modMgr.installedModules[moduleName] = &module_ver
	modMgr.moduleConf[moduleName] = modconf

	log.Info(ctx, "Module installed", "moduleName", moduleName)
	return &module_ver, modconf, nil
}

func (modMgr *moduleManager) buildObjectInfo(ctx core.ServerContext, conf config.Config) error {
	objsconf, ok := conf.GetSubConfig(ctx, OBJECTS)
	if ok {
		objs := objsconf.AllConfigurations(ctx)
		for _, objname := range objs {
			objconf, _ := objsconf.GetSubConfig(ctx, objname)
			objtyp, _ := objconf.GetString(ctx, OBJECT_TYPE)
			var inf core.Info
			var err error
			switch objtyp {
			case "service":
				inf, err = buildServiceInfo(ctx, objconf)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			case "module":
				inf = buildModuleInfo(ctx, objconf)
			case "factory":
				inf = buildFactoryInfo(ctx, objconf)
			}
			if inf != nil {
				ldr := ctx.GetServerElement(core.ServerElementLoader).(*objectLoaderProxy).loader
				ldr.setObjectInfo(ctx, objname, inf)
			}
		}
	}
	return nil
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
