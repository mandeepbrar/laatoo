package main

import (
	"fmt"
	"laatoo/sdk/components"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"os"
	"path"
	"path/filepath"

	"github.com/mholt/archiver"
)

//"path/filepath"

type RepositoryUpdate struct {
	core.Service
	dataStore       data.DataComponent
	repositoryFiles components.StorageComponent
}

const (
	CONF_DATASTORE = "datastore"
	PARAM_MOD      = "module"
	TMPPATH        = "/tmp"
)

func (svc *RepositoryUpdate) Describe(ctx core.ServerContext) {
	svc.AddStringParam(ctx, PARAM_MOD)
}

func (svc *RepositoryUpdate) Start(ctx core.ServerContext) error {
	dataSvcName := "repository.modules.database"
	dataSvc, err := ctx.GetService(dataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	svc.dataStore = dataSvc.(data.DataComponent)
	repositorySvc := "repository.modules.directory"
	filesSvc, err := ctx.GetService(repositorySvc)
	if err != nil {
		return errors.MissingService(ctx, repositorySvc)
	}
	svc.repositoryFiles = filesSvc.(components.StorageComponent)
	return nil
}

func (svc *RepositoryUpdate) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("RepositoryUpdate")
	mod, ok := ctx.GetStringParam(PARAM_MOD)
	log.Error(ctx, "Update Module", "Module", mod)
	if ok {
		files, err := svc.repositoryFiles.ListFiles(ctx, fmt.Sprintf("%s*.tar.gz", mod))
		if err != nil {
			return err
		}
		log.Error(ctx, "Update Module", "files", files)
		for _, file := range files {
			_, fileName := filepath.Split(file)
			log.Error(ctx, "Process archive", "fileName", fileName)
			err = svc.processArchive(ctx, mod, fileName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (svc *RepositoryUpdate) processArchive(ctx core.RequestContext, mod, file string) error {
	err := svc.extractArchive(ctx, mod, file)
	if err != nil {
		return err
	}
	err = svc.readMod(ctx, mod)
	if err != nil {
		return err
	}
	return nil
}

func (svc *RepositoryUpdate) extractArchive(ctx core.RequestContext, mod, file string) error {
	modTmpDir := path.Join(TMPPATH, mod)
	modDirExists, _, _ := utils.FileExists(modTmpDir)
	log.Error(ctx, "Extracting archive", "Module", modTmpDir, "Dir exists", modDirExists)
	//ensure latest module directory is present
	// if zip file is newer than module dir
	// delete the directory and extract latest zip file
	if modDirExists {
		err := os.RemoveAll(modTmpDir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	str, err := svc.repositoryFiles.Open(ctx, file)
	if err != nil {
		return err
	}
	if err := archiver.TarGz.Read(str, TMPPATH); err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Extracted module ", "Module", mod, "Module file", file)
	return nil
}

func (svc *RepositoryUpdate) readMod(ctx core.RequestContext, modName string) error {
	confPath := path.Join(TMPPATH, modName, "config.yml")
	conf, err := ctx.ServerContext().ReadConfig(confPath, nil)
	if err != nil {
		return err
	}
	log.Error(ctx, "Module Conf", "mod", modName, "conf", conf)
	mod := &ModuleDefinition{Name: modName}
	err = svc.readConf(ctx, mod, conf)
	if err != nil {
		return err
	}
	return nil
}

func (svc *RepositoryUpdate) readConf(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	return nil
}

/*

func (modMgr *moduleManager) buildObjectInfo(ctx core.ServerContext, conf config.Config) error {
	objsconf, ok := conf.GetSubConfig(ctx, OBJECTS)
	if ok {
		objs := objsconf.AllConfigurations(ctx)
		for _, objname := range objs {
			objconf, _ := objsconf.GetSubConfig(ctx, objname)
			objtyp, _ := objconf.GetString(ctx, OBJECT_TYPE)
			var inf core.Info
			switch objtyp {
			case "service":
				inf = buildServiceInfo(ctx, objconf)
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
}*/
