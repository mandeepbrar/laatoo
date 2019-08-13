package main

import (
	"fmt"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"path/filepath"
)

//"path/filepath"

type RepositoryUpdate struct {
	core.Service
	dataStore       data.DataComponent
	repositoryFiles components.StorageComponent
	formDataStore   data.DataComponent
}

const (
	CONF_DATASTORE = "datastore"
	PARAM_MOD      = "module"
	TMPPATH        = "/tmp"
)

func (svc *RepositoryUpdate) Describe(ctx core.ServerContext) error {
	svc.AddStringParam(ctx, PARAM_MOD)
	return nil
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
	formDataSvcName := "repository.moduleform.database"
	dataSvc, err = ctx.GetService(formDataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	svc.formDataStore = dataSvc.(data.DataComponent)
	return nil
}

func (svc *RepositoryUpdate) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("RepositoryUpdate")
	mod, ok := ctx.GetStringParam(PARAM_MOD)
	log.Error(ctx, "Update Module", "Module", mod)
	if ok {
		_, err := svc.processModule(ctx, mod)
		return err
	}
	return nil
}

func (svc *RepositoryUpdate) processModule(ctx core.RequestContext, mod string) (*ModuleDefinition, error) {
	files, err := svc.repositoryFiles.ListFiles(ctx, fmt.Sprintf("%s.tar.gz", mod))
	if err != nil {
		return nil, err
	}
	log.Error(ctx, "Update Module", "files", files)
	if len(files) > 0 {
		_, fileName := filepath.Split(files[0])
		log.Error(ctx, "Process archive", "fileName", fileName)
		modDef, err := processArchive(ctx, mod, fileName, svc.repositoryFiles)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		err = svc.dataStore.Put(ctx, mod, modDef)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		conf, err := writeParamsForm(ctx, modDef)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		modForm := &ModuleForm{Name: modDef.Name, Form: string(conf)}
		log.Error(ctx, "Process archive", "modForm", modForm)
		err = svc.formDataStore.Put(ctx, modDef.Name, modForm)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		return modDef, nil
	}
	return nil, nil
}
