package main

import (
	"fmt"
	"io/ioutil"
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
		modDef, err := svc.processArchive(ctx, mod, fileName)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		err = svc.dataStore.Put(ctx, mod, modDef)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		return modDef, nil
	}
	return nil, nil
}

func (svc *RepositoryUpdate) processArchive(ctx core.RequestContext, mod, file string) (*ModuleDefinition, error) {
	err := svc.extractArchive(ctx, mod, file)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return svc.readMod(ctx, mod)
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
		return errors.WrapError(ctx, err)
	}
	if err := archiver.TarGz.Read(str, TMPPATH); err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Extracted module ", "Module", mod, "Module file", file)
	return nil
}

func (svc *RepositoryUpdate) readMod(ctx core.RequestContext, modName string) (*ModuleDefinition, error) {
	confPath := path.Join(TMPPATH, modName, "config.yml")
	conf, err := ctx.ServerContext().ReadConfig(confPath, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod := &ModuleDefinition{Name: modName}
	err = svc.readConf(ctx, mod, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	modDir := path.Join(TMPPATH, modName)
	svcs, err := svc.readElementNames(ctx, modDir, "services", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Services = svcs
	channels, err := svc.readElementNames(ctx, modDir, "channels", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Services for mod", "services", svcs)
	mod.Channels = channels
	factories, err := svc.readElementNames(ctx, modDir, "factories", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Factories = factories
	engines, err := svc.readElementNames(ctx, modDir, "engines", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Engines = engines
	rules, err := svc.readElementNames(ctx, modDir, "rules", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Rules = rules
	tasks, err := svc.readElementNames(ctx, modDir, "tasks", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Tasks = tasks
	err = svc.readSubModules(ctx, mod, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	err = svc.writeParamsForm(ctx, mod)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Module Conf", "mod", modName, "conf", conf, "mod", mod)
	return mod, nil
}

func (svc *RepositoryUpdate) readConf(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	desc, ok := conf.GetString(ctx, "description")
	if ok {
		mod.Description = desc
	}
	ver, ok := conf.GetString(ctx, "version")
	if ok {
		mod.Version = ver
	}
	err := svc.readDependencies(ctx, mod, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = svc.readParams(ctx, mod, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = svc.readObjects(ctx, mod, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *RepositoryUpdate) readDependencies(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	depVals := make(map[string]string)
	deps, ok := conf.GetSubConfig(ctx, "dependencies")
	if ok {
		depNames := deps.AllConfigurations(ctx)
		for _, dep := range depNames {
			ver, _ := deps.GetString(ctx, dep)
			depVals[dep] = ver
		}
	}
	mod.Dependencies = depVals
	return nil
}

func (svc *RepositoryUpdate) readParams(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	params := make(map[string]ModuleParam)
	log.Error(ctx, "Reading params", "conf", conf)
	paramsConf, ok := conf.GetSubConfig(ctx, "params")
	if ok {
		paramNames := paramsConf.AllConfigurations(ctx)
		for _, paramName := range paramNames {
			paramConf, _ := paramsConf.GetSubConfig(ctx, paramName)
			ptype, _ := paramConf.GetString(ctx, "type")
			desc, _ := paramConf.GetString(ctx, "description")
			modParam := ModuleParam{Name: paramName, Type: ptype, Description: desc}
			params[paramName] = modParam
		}
	}
	log.Error(ctx, "Reading params", "params", params)
	mod.Params = params
	return nil
}

func (svc *RepositoryUpdate) readObjects(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	objects := make([]ObjectDefinition, 0)
	objsConf, ok := conf.GetSubConfig(ctx, "objects")
	if ok {
		objNames := objsConf.AllConfigurations(ctx)
		for _, objName := range objNames {
			objDef := ObjectDefinition{Name: objName}
			objConf, _ := objsConf.GetSubConfig(ctx, objName)
			objType, _ := objConf.GetString(ctx, "type")
			if objType == "module" {
				err := svc.readModObj(ctx, objName, mod, conf)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
			objDef.Type = objType
			objConfs, ok := objConf.GetSubConfig(ctx, "configurations")
			if ok {
				objConfDefs := make([]ConfigurationDefinition, 0)
				allConfs := objConfs.AllConfigurations(ctx)
				for _, confName := range allConfs {
					confDef := ConfigurationDefinition{Name: confName}
					confConf, _ := objConfs.GetSubConfig(ctx, confName)
					confType, ok := confConf.GetString(ctx, "type")
					if ok {
						confDef.Type = confType
					}
					confDesc, ok := confConf.GetString(ctx, "description")
					if ok {
						confDef.Description = confDesc
					}
					confDef.Required, _ = confConf.GetBool(ctx, "required")
					objConfDefs = append(objConfDefs, confDef)
				}
				objDef.Configurations = objConfDefs
			}
			log.Error(ctx, "Objects encountered", "objs", objConf)
			objects = append(objects, objDef)
		}
	}
	log.Error(ctx, "Objects encountered", "objs", objects)
	mod.Objects = objects
	return nil
}

func (svc *RepositoryUpdate) readModObj(ctx core.RequestContext, objName string, mod *ModuleDefinition, conf config.Config) error {
	log.Error(ctx, "Creating obj", "objName", objName)
	obj, err := ctx.ServerContext().CreateObject(objName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	modObj := obj.(core.Module)
	log.Error(ctx, "module obj", "modObj", modObj)
	if modObj != nil {
		inf := modObj.MetaInfo(ctx.ServerContext())
		log.Error(ctx, "Meta Info", "inf", inf)
		if inf != nil {
			svcs := inf["services"]
			svcsArr, ok := svcs.([]string)
			if ok {
				mod.Services = svcsArr
			}
			facs := inf["factories"]
			facsArr, ok := facs.([]string)
			if ok {
				mod.Factories = facsArr
			}
			chns := inf["channels"]
			chnsArr, ok := chns.([]string)
			if ok {
				mod.Channels = chnsArr
			}
		}
	}
	return nil
}

func (svc *RepositoryUpdate) readElementNames(ctx core.RequestContext, modDir, element string, conf config.Config) ([]string, error) {
	elems := make([]string, 0)
	elementDir := path.Join(modDir, element)
	ok, fi, _ := utils.FileExists(elementDir)
	log.Trace(ctx, "Mod dir exists", "dir", elementDir, "ok", ok, "modDir", modDir, "element", element)
	if ok && fi.IsDir() {
		files, err := ioutil.ReadDir(elementDir)
		if err != nil {
			return elems, err
		}
		for _, info := range files {
			elemfileName := info.Name()
			file := path.Join(elementDir, elemfileName)
			if !info.IsDir() {
				extension := filepath.Ext(elemfileName)
				elemName := elemfileName[0 : len(elemfileName)-len(extension)]
				elemConf, err := ctx.ServerContext().ReadConfig(file, nil)
				if err != nil {
					return elems, errors.WrapError(ctx, err)
				}
				name, ok := elemConf.GetString(ctx, "name")
				if ok {
					elemName = name
				}
				elems = append(elems, elemName)
			}
		}
	}
	return elems, nil
}

func (svc *RepositoryUpdate) readSubModules(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	modsConf, ok := conf.GetSubConfig(ctx, "modules")
	if ok {
		allSubMods := modsConf.AllConfigurations(ctx)
		for _, subMod := range allSubMods {
			subModConf, _ := modsConf.GetSubConfig(ctx, subMod)
			subModName := subMod
			name, ok := subModConf.GetString(ctx, "module")
			if ok {
				subModName = name
			}
			log.Error(ctx, "SubMod found", "submod", subModName, "conf", subModConf)
			subModDef, err := svc.processModule(ctx, subModName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			log.Error(ctx, "SubMod def", "subModDef", subModDef)
		}
	}
	return nil
}

func (svc *RepositoryUpdate) writeParamsForm(ctx core.RequestContext, mod *ModuleDefinition) error {
	formConf := make(map[string]interface{})
	formFields := make(map[string]interface{})

	for pname, pConf := range mod.Params {
		formParam := make(map[string]interface{})
		formParam["name"] = pname
		formParam["type"] = pConf.Type
		formFields[pname] = formParam
	}
	formConf["fields"] = formFields

	mod.ParamsForm = formConf
	return nil
}
