package main

import (
	"encoding/json"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/modulesbase"
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

func processArchive(ctx core.RequestContext, mod, file, tmppath string, repositoryFiles components.StorageComponent) (*modulesbase.ModuleDefinition, error) {
	err := extractArchive(ctx, mod, file, tmppath, repositoryFiles)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return readMod(ctx, mod, file, tmppath)
}

func extractArchive(ctx core.RequestContext, mod, file, tmppath string, repositoryFiles components.StorageComponent) error {
	modTmpDir := path.Join(tmppath, mod)
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
	localmodFile := path.Join(tmppath, mod+".tar.gz")
	fil, err := os.Create(localmodFile)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Archive exists", "file", file)
	err = repositoryFiles.CopyFile(ctx, file, fil)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if err := archiver.Unarchive(localmodFile, tmppath); err != nil {
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

func readMod(ctx core.RequestContext, modName, file, tmppath string) (*modulesbase.ModuleDefinition, error) {
	confDir := path.Join(tmppath, modName, "config")
	confPath := path.Join(confDir, "config.yml")
	conf, err := ctx.ServerContext().ReadConfig(confPath, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod := &modulesbase.ModuleDefinition{Name: modName, File: file}
	mod.SetId(modName)
	err = readConf(ctx, mod, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	svcs, err := readServices(ctx, confDir, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Services = svcs
	channels, err := readChannels(ctx, confDir, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Services for mod", "services", svcs)
	mod.Channels = channels
	factories, err := readFactories(ctx, confDir, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Factories = factories
	engines, err := readEngines(ctx, confDir, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Engines = engines
	rules, err := readRules(ctx, confDir, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Rules = rules
	tasks, err := readTasks(ctx, confDir, conf)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	mod.Tasks = tasks
	err = readSubModules(ctx, mod, conf, tmppath)
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

func readConf(ctx core.RequestContext, mod *modulesbase.ModuleDefinition, conf config.Config) error {
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

func readDependencies(ctx core.RequestContext, mod *modulesbase.ModuleDefinition, conf config.Config) error {
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

func readParams(ctx core.RequestContext, mod *modulesbase.ModuleDefinition, conf config.Config) error {
	params := make([]modulesbase.Param, 0)
	log.Error(ctx, "Reading params", "conf", conf)
	paramsConf, ok := conf.GetSubConfig(ctx, "params")
	if ok {
		paramNames := paramsConf.AllConfigurations(ctx)
		for _, paramName := range paramNames {
			paramConf, _ := paramsConf.GetSubConfig(ctx, paramName)
			ptype, _ := paramConf.GetString(ctx, "type")
			desc, _ := paramConf.GetString(ctx, "description")
			modParam := modulesbase.Param{Name: paramName, Type: ptype, Description: desc}
			params = append(params, modParam)
		}
	}
	log.Error(ctx, "Reading params", "params", params)
	mod.Params = params
	return nil
}

func readObjects(ctx core.RequestContext, mod *modulesbase.ModuleDefinition, conf config.Config) error {
	objects := make([]modulesbase.ObjectDefinition, 0)
	objsConf, ok := conf.GetSubConfig(ctx, "objects")
	if ok {
		objNames := objsConf.AllConfigurations(ctx)
		for _, objName := range objNames {
			objDef := modulesbase.ObjectDefinition{Name: objName}
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
				objConfDefs := make([]modulesbase.Param, 0)
				allConfs := objConfs.AllConfigurations(ctx)
				for _, confName := range allConfs {
					confDef := modulesbase.Param{Name: confName}
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

func readModObj(ctx core.RequestContext, objName string, mod *modulesbase.ModuleDefinition, conf config.Config) error {
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
			svcsConf, ok := svcs.([]config.Config)
			if ok {
				svcsArr := make([]modulesbase.Service, len(svcsConf))
				for i, conf := range svcsConf {
					name, _ := conf.GetString(ctx, "name")
					svc, err := readServiceConf(ctx, name, conf)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
					svcsArr[i] = svc
				}
				mod.Services = svcsArr
			}
			facs := inf["factories"]
			facsConf, ok := facs.([]config.Config)
			if ok {
				facsArr := make([]modulesbase.Factory, len(facsConf))
				for i, conf := range facsConf {
					name, _ := conf.GetString(ctx, "name")
					fac, err := readFactoryConf(ctx, name, conf)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
					facsArr[i] = fac
				}
				mod.Factories = facsArr
			}
			chns := inf["channels"]
			chansConf, ok := chns.([]config.Config)
			if ok {
				chnsArr := make([]modulesbase.Channel, len(chansConf))
				for i, conf := range chansConf {
					name, _ := conf.GetString(ctx, "name")
					svc, err := readChannelConf(ctx, name, conf)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
					chnsArr[i] = svc
				}
				mod.Channels = chnsArr
			}
		}
	}
	return nil
}

func readElementConfs(ctx core.RequestContext, modDir, element string) ([]config.Config, []string, error) {
	elems := make([]config.Config, 0)
	elemNames := make([]string, 0)
	elementDir := path.Join(modDir, element)
	ok, fi, _ := utils.FileExists(elementDir)
	log.Trace(ctx, "Mod dir exists", "dir", elementDir, "ok", ok, "modDir", modDir, "element", element)
	if ok && fi.IsDir() {
		files, err := ioutil.ReadDir(elementDir)
		if err != nil {
			return nil, nil, err
		}
		for _, info := range files {
			elemfileName := info.Name()
			file := path.Join(elementDir, elemfileName)
			if !info.IsDir() {
				extension := filepath.Ext(elemfileName)
				elemName := elemfileName[0 : len(elemfileName)-len(extension)]
				elemConf, err := ctx.ServerContext().ReadConfig(file, nil)
				if err != nil {
					return nil, nil, errors.WrapError(ctx, err)
				}
				name, ok := elemConf.GetString(ctx, "name")
				if ok {
					elemName = name
				}
				elems = append(elems, elemConf)
				elemNames = append(elemNames, elemName)
			}
		}
	}
	return elems, elemNames, nil
}

func readServices(ctx core.RequestContext, modDir string, conf config.Config) ([]modulesbase.Service, error) {

	elemConfs, elemNames, err := readElementConfs(ctx, modDir, "services")
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	elems := make([]modulesbase.Service, len(elemConfs))

	for i, ec := range elemConfs {
		svcName := elemNames[i]
		svc, err := readServiceConf(ctx, svcName, ec)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		elems[i] = svc
	}

	return elems, nil
}

func readFactories(ctx core.RequestContext, modDir string, conf config.Config) ([]modulesbase.Factory, error) {

	elemConfs, elemNames, err := readElementConfs(ctx, modDir, "factories")
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	elems := make([]modulesbase.Factory, len(elemConfs))

	for i, ec := range elemConfs {
		facName := elemNames[i]
		fac, err := readFactoryConf(ctx, facName, ec)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		elems[i] = fac
	}

	return elems, nil
}

func readChannels(ctx core.RequestContext, modDir string, conf config.Config) ([]modulesbase.Channel, error) {

	elemConfs, elemNames, err := readElementConfs(ctx, modDir, "channels")
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	elems := make([]modulesbase.Channel, len(elemConfs))

	for i, ec := range elemConfs {
		chnName := elemNames[i]
		chn, err := readChannelConf(ctx, chnName, ec)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		elems[i] = chn
	}

	return elems, nil
}

func readEngines(ctx core.RequestContext, modDir string, conf config.Config) ([]modulesbase.Engine, error) {
	elemConfs, elemNames, err := readElementConfs(ctx, modDir, "engines")
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	elems := make([]modulesbase.Engine, len(elemConfs))

	for i, ec := range elemConfs {
		engName := elemNames[i]
		eng, err := readEngineConf(ctx, engName, ec)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		elems[i] = eng
	}

	return elems, nil
}

func readRules(ctx core.RequestContext, modDir string, conf config.Config) ([]modulesbase.Rule, error) {
	elemConfs, elemNames, err := readElementConfs(ctx, modDir, "rules")
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	elems := make([]modulesbase.Rule, len(elemConfs))

	for i, ec := range elemConfs {
		ruleName := elemNames[i]
		rul, err := readRuleConf(ctx, ruleName, ec)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		elems[i] = rul
	}

	return elems, nil
}

func readTasks(ctx core.RequestContext, modDir string, conf config.Config) ([]modulesbase.Task, error) {
	elemConfs, elemNames, err := readElementConfs(ctx, modDir, "tasks")
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	elems := make([]modulesbase.Task, len(elemConfs))

	for i, ec := range elemConfs {
		taskName := elemNames[i]
		tsk, err := readTaskConf(ctx, taskName, ec)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		elems[i] = tsk
	}

	return elems, nil
}

func readSubModules(ctx core.RequestContext, mod *modulesbase.ModuleDefinition, conf config.Config, tmppath string) error {
	/*modsConf, ok := conf.GetSubConfig(ctx, "modules")
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
			subModDef, err := readMod(ctx, subModName, tmppath)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			log.Error(ctx, "SubMod def", "subModDef", subModDef)
		}
	}*/
	return nil
}

func readServiceConf(ctx core.RequestContext, name string, conf config.Config) (modulesbase.Service, error) {
	return modulesbase.Service{}, nil
}

func readFactoryConf(ctx core.RequestContext, name string, conf config.Config) (modulesbase.Factory, error) {
	return modulesbase.Factory{}, nil
}

func readChannelConf(ctx core.RequestContext, name string, conf config.Config) (modulesbase.Channel, error) {
	return modulesbase.Channel{}, nil
}

func readRuleConf(ctx core.RequestContext, name string, conf config.Config) (modulesbase.Rule, error) {
	return modulesbase.Rule{}, nil
}

func readTaskConf(ctx core.RequestContext, name string, conf config.Config) (modulesbase.Task, error) {
	return modulesbase.Task{}, nil
}

func readEngineConf(ctx core.RequestContext, name string, conf config.Config) (modulesbase.Engine, error) {
	return modulesbase.Engine{}, nil
}

func writeConfigForm(ctx core.RequestContext, mod *modulesbase.ModuleDefinition) ([]byte, error) {
	formConf := make(map[string]interface{})
	formInfo := make(map[string]interface{})
	//configFormInfo.Set(ctx, "entity", entity.object)
	formInfo["className"] = " configform " + strings.ToLower(mod.Name)
	formConf["info"] = formInfo

	formFields := make(map[string]interface{})

	for _, param := range mod.Params {
		pname := param.Name
		formParam := make(map[string]interface{})
		formParam["name"] = pname
		formParam["type"] = param.Type
		formParam["className"] = " configformfield " + pname
		formFields[pname] = formParam
	}
	formConf["fields"] = formFields

	return json.Marshal(formConf)
}
