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

type EntitlementCreationService struct {
	core.Service
	dataStore        data.DataComponent
	repositoryFiles  components.StorageComponent
	modulesdataStore data.DataComponent
	formDataStore    data.DataComponent
}

func (svc *EntitlementCreationService) Start(ctx core.ServerContext) error {

	dataSvcName := "repository.entitlement.database"
	dataSvc, err := ctx.GetService(dataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	svc.dataStore = dataSvc.(data.DataComponent)
	repositorySvc := "repository.solution.module.import"
	filesSvc, err := ctx.GetService(repositorySvc)
	if err != nil {
		return errors.MissingService(ctx, repositorySvc)
	}
	svc.repositoryFiles = filesSvc.(components.StorageComponent)

	modulesDataSvcName := "repository.modules.database"
	modulesDataSvc, err := ctx.GetService(modulesDataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	svc.modulesdataStore = modulesDataSvc.(data.DataComponent)
	formDataSvcName := "repository.moduleform.database"
	dataSvc, err = ctx.GetService(formDataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	svc.formDataStore = dataSvc.(data.DataComponent)

	return nil
}

func (svc *EntitlementCreationService) Invoke(ctx core.RequestContext) error {
	solution, ok := ctx.GetStringParam("solutionId")
	if ok {
		log.Error(ctx, "Entitlement creation ", "solution", solution)
		modID, ok := ctx.GetStringParam("moduleId")
		if ok {
			stor, err := svc.modulesdataStore.GetById(ctx, modID)
			log.Error(ctx, "Entitlement creation. Looking for mod by id ", "modID", modID, "stor", stor, "err", err)
			if err == nil {
				modDef, _ := stor.(*ModuleDefinition)
				log.Error(ctx, "Entitlement creation ", "solution", solution, "module", modDef)
				modRef := data.StorableRef{Id: modID, Name: modDef.Name, Type: "ModuleDefinition"}

				ent := &Entitlement{Name: modDef.Name, Solution: data.StorableRef{Id: solution, Type: "Solution"}, Module: modRef, Local: false}
				err = svc.dataStore.Save(ctx, ent)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				ctx.SetResponse(core.SuccessResponse(nil))
				return nil
			} else {
				return err
			}
		} else {
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
						modRef := data.StorableRef{Id: modID, Name: modDef.Name, Type: "ModuleDefinition"}
						ent := &Entitlement{Name: modName, Solution: data.StorableRef{Id: solution, Type: "Solution"}, Module: modRef, Local: true}
						err = svc.dataStore.Save(ctx, ent)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
						log.Error(ctx, "Imported Module", "Module", modName, "modDef", modDef)
						conf, err := writeParamsForm(ctx, modDef)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
						modForm := &ModuleForm{Name: modName, Form: string(conf)}
						err = svc.formDataStore.Put(ctx, modName, modForm)
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
		}
	}
	ctx.SetResponse(core.StatusBadRequestResponse)
	return nil
}
