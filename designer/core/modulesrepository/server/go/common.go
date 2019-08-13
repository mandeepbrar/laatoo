package main

import (
	"encoding/json"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"
)

func processArchive(ctx core.RequestContext, mod, file string, repositoryFiles components.StorageComponent) (*ModuleDefinition, error) {
	err := extractArchive(ctx, mod, file, repositoryFiles)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return readMod(ctx, mod)
}

func extractArchive(ctx core.RequestContext, mod, file string, repositoryFiles components.StorageComponent) error {
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
	//	os.MkdirAll(modTmpDir, os.ModeTemporary)
	archiveExists := repositoryFiles.Exists(ctx, file)
	if !archiveExists {
		log.Info(ctx, "Archive does not exist", "file", file)
		ctx.SetResponse(core.StatusNotFoundResponse)
		return nil
	}
	localmodFile := path.Join(TMPPATH, mod+".tar.gz")
	fil, err := os.Create(localmodFile)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Archive exists", "file", file)
	err = repositoryFiles.CopyFile(ctx, file, fil)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if err := archiver.Unarchive(localmodFile, TMPPATH); err != nil {
		return errors.WrapError(ctx, err)
	}
	/*	str, err := repositoryFiles.Open(ctx, file)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		targz := archiver.NewTarGz()
		err = targz.Open(str, -1)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		if err := targz.Extract(mod, mod, TMPPATH); err != nil {
			return errors.WrapError(ctx, err)
		}
	*/
	utils.PrintDirContents(modTmpDir)
	log.Error(ctx, "Extracted module ", "Module", mod, "Module file", file, "modTmpDir", modTmpDir)
	return nil
}

func readMod(ctx core.RequestContext, modName string) (*ModuleDefinition, error) {
	confDir := path.Join(TMPPATH, modName, "config")
	confPath := path.Join(confDir, "config.yml")
	conf, err := ctx.ServerContext().ReadConfig(confPath, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod := &ModuleDefinition{Name: modName}
	mod.SetId(modName)
	err = readConf(ctx, mod, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	svcs, err := readElementNames(ctx, confDir, "services", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Services = svcs
	channels, err := readElementNames(ctx, confDir, "channels", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Services for mod", "services", svcs)
	mod.Channels = channels
	factories, err := readElementNames(ctx, confDir, "factories", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Factories = factories
	engines, err := readElementNames(ctx, confDir, "engines", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Engines = engines
	rules, err := readElementNames(ctx, confDir, "rules", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Rules = rules
	tasks, err := readElementNames(ctx, confDir, "tasks", conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Tasks = tasks
	err = readSubModules(ctx, mod, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	/*	err = writeParamsForm(ctx, mod)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}*/
	log.Error(ctx, "Module Conf", "mod", modName, "conf", conf, "mod", mod)
	return mod, nil
}

func readConf(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	desc, ok := conf.GetString(ctx, "description")
	if ok {
		mod.Description = desc
	}
	ver, ok := conf.GetString(ctx, "version")
	if ok {
		mod.Version = ver
	}
	err := readDependencies(ctx, mod, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = readParams(ctx, mod, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = readObjects(ctx, mod, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func readDependencies(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
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

func readParams(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
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

func readObjects(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
	objects := make([]ObjectDefinition, 0)
	objsConf, ok := conf.GetSubConfig(ctx, "objects")
	if ok {
		objNames := objsConf.AllConfigurations(ctx)
		for _, objName := range objNames {
			objDef := ObjectDefinition{Name: objName}
			objConf, _ := objsConf.GetSubConfig(ctx, objName)
			objType, _ := objConf.GetString(ctx, "type")
			if objType == "module" {
				err := readModObj(ctx, objName, mod, conf)
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

func readModObj(ctx core.RequestContext, objName string, mod *ModuleDefinition, conf config.Config) error {
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

func readElementNames(ctx core.RequestContext, modDir, element string, conf config.Config) ([]string, error) {
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

func readSubModules(ctx core.RequestContext, mod *ModuleDefinition, conf config.Config) error {
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
			subModDef, err := readMod(ctx, subModName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			log.Error(ctx, "SubMod def", "subModDef", subModDef)
		}
	}
	return nil
}

func writeParamsForm(ctx core.RequestContext, mod *ModuleDefinition) ([]byte, error) {
	formConf := make(map[string]interface{})
	formInfo := make(map[string]interface{})
	//configFormInfo.Set(ctx, "entity", entity.object)
	formInfo["className"] = " configform " + strings.ToLower(mod.Name)
	formConf["info"] = formInfo

	formFields := make(map[string]interface{})

	for pname, pConf := range mod.Params {
		formParam := make(map[string]interface{})
		formParam["name"] = pname
		formParam["type"] = pConf.Type
		formParam["className"] = " configformfield " + pname
		formFields[pname] = formParam
	}
	formConf["fields"] = formFields

	return json.Marshal(formConf)
}
