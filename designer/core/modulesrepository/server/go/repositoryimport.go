package main

import (
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"strings"
)

//"path/filepath"

type RepositoryImport struct {
	core.Service
	dataStore       data.DataComponent
	repositoryFiles components.StorageComponent
}

func (svc *RepositoryImport) Start(ctx core.ServerContext) error {

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

func (svc *RepositoryImport) Invoke(ctx core.RequestContext) error {
	uploadedFiles, ok := ctx.GetStringsMapParam("Data")
	if ok {
		uploadedModules := map[string]string{}
		for fileName, archiveName := range uploadedFiles {
			modName := strings.TrimSuffix(fileName, ".tar.gz")
			log.Error(ctx, "Import Module", "Module", modName, "Archive Name", archiveName)
			if ok {
				modDef, err := processArchive(ctx, modName, archiveName, svc.repositoryFiles)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				err = svc.dataStore.Put(ctx, modName, modDef)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				/*	_, err := svc.processModule(ctx, mod)
					return err*/
			}
			uploadedModules[fileName] = modName
		}
		ctx.SetResponse(core.SuccessResponse(uploadedModules))
		return nil
	}
	ctx.SetResponse(core.StatusBadRequestResponse)
	return nil
}
