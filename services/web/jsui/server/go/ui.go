package main

import (
	"fmt"
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

	"github.com/fsnotify/fsnotify"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: UI_SVC, Object: UI{}}}
}

const (
	UI_SVC               = "ui"
	UI_SVC_ENABLED       = "ui"
	FILES_DIR            = "files"
	SCRIPTS_DIR          = "scripts"
	WASM_DIR             = "wasm"
	CSS_DIR              = "css"
	CONF_FILE_UI         = "uifile"
	CONF_FILE_DESC       = "descriptor"
	CONF_PROPS_EXTENSION = "properties_ext"
	CONF_APPLICATION     = "uiapplication"
	DEPENDENCIES         = "dependencies"
	UI_DIR               = "ui"
	REG_DIR              = "registry"
	MERGED_SVCS_FILE     = "mergeduidescriptor"
	MERGED_CSS_FILE      = "mergedcssfile"
	MERGED_WASM_FILE     = "mergedwasm"
	MERGED_VENDOR_FILE   = "mergedvendorfile"
	MERGED_UI_FILE       = "mergeduifile"
	BOOT_FILE            = "bootfile"
)

type UI struct {
	core.Module
	svrCtx             core.ServerContext
	uifile             string
	descfile           string
	mergeduifile       string
	mergedvendorfile   string
	mergeduidescriptor string
	mergedwasmfile     string
	mergedcssfile      string
	bootfile           string
	application        string
	propsExt           string
	uiFiles            map[string][]byte
	vendorFiles        map[string][]byte
	cssFiles           map[string][]byte
	wasmFiles          map[string]string
	wasmImportFiles    map[string]string
	wasmImportScript   []byte
	modDeps            map[string][]string
	insMods            map[string]string
	insSettings        map[string]config.Config
	descriptorFiles    map[string][]byte
	requiredUIPkgs     utils.StringSet
	propertyFiles      map[string]interface{}
	uiRegistry         map[string]map[string]string
	uiPlugins          map[string]*components.ModInfo
	watchers           []*fsnotify.Watcher
	wasmModName        string
}

/*
func (svc *StaticFiles) Initialize(ctx core.ServerContext) error {
	svc.SetDescription(ctx, "Static files service")
	svc.AddStringConfigurations(ctx, []string{CONF_FILE_STORAGE}, nil)
	svc.AddStringConfigurations(ctx, []string{CONF_FILE_OPER, CONF_FILE_TRANSFORM_STG, CONF_FILE_DEFAULT, CONF_IMAGE_WIDTH, CONF_IMAGE_HEIGHT}, []string{"", "", "", "0", "0"})
	svc.AddParam(ctx, CONF_STATIC_FILEPARAM, config.OBJECTTYPE_STRING, false)

	return nil
}*/

func (svc *UI) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.svrCtx = ctx
	svc.uifile, _ = svc.GetStringConfiguration(ctx, CONF_FILE_UI)
	svc.descfile, _ = svc.GetStringConfiguration(ctx, CONF_FILE_DESC)
	svc.mergeduifile, _ = svc.GetStringConfiguration(ctx, MERGED_UI_FILE)
	svc.mergedvendorfile, _ = svc.GetStringConfiguration(ctx, MERGED_VENDOR_FILE)
	svc.mergeduidescriptor, _ = svc.GetStringConfiguration(ctx, MERGED_SVCS_FILE)
	svc.mergedcssfile, _ = svc.GetStringConfiguration(ctx, MERGED_CSS_FILE)
	svc.mergedwasmfile, _ = svc.GetStringConfiguration(ctx, MERGED_WASM_FILE)
	svc.bootfile, _ = svc.GetStringConfiguration(ctx, BOOT_FILE)
	svc.application, _ = svc.GetStringConfiguration(ctx, CONF_APPLICATION)
	svc.propsExt, _ = svc.GetStringConfiguration(ctx, CONF_PROPS_EXTENSION)
	svc.watchers = make([]*fsnotify.Watcher, 0)
	svc.uiFiles = make(map[string][]byte)
	svc.vendorFiles = make(map[string][]byte)
	svc.cssFiles = make(map[string][]byte)
	svc.wasmFiles = make(map[string]string)
	svc.wasmImportFiles = make(map[string]string)
	svc.insSettings = make(map[string]config.Config)
	svc.insMods = make(map[string]string)
	svc.descriptorFiles = make(map[string][]byte)
	svc.modDeps = make(map[string][]string)
	svc.propertyFiles = make(map[string]interface{})
	svc.requiredUIPkgs = utils.NewStringSet([]string{})
	svc.uiRegistry = make(map[string]map[string]string)
	svc.wasmModName = fmt.Sprintf("wasm_%s", svc.application)
	svc.uiPlugins = make(map[string]*components.ModInfo)
	log.Error(ctx, "UI conf file", "conf", conf, "svc", svc.uifile, "svc", svc)
	//svc.uiDisplays = make(map[string]string)
	return nil
}

func (svc *UI) Load(ctx core.ServerContext, modInfo *components.ModInfo) error {
	ctx = ctx.SubContext("Load UI module " + modInfo.InstanceName)

	modName := modInfo.ModName
	ui, ok := modInfo.ModSettings.GetBool(ctx, UI_SVC_ENABLED)
	if ok && !ui {
		return nil
	}

	app, ok := modInfo.ModSettings.GetString(ctx, CONF_APPLICATION)
	if app != "" && svc.application != app {
		log.Error(ctx, "Skipping module from ui", "module", modName, "application", svc.application)
		return nil
	}
	log.Info(ctx, "Module read", "mod name", modName)

	modFilesDir := svc.getFilesDir(ctx, modName, modInfo.ModDir, modInfo)

	_, modRead := svc.uiFiles[modName]

	if modInfo.Hot || !modRead {

		if modInfo.IsExtended {
			if err := svc.LoadFiles(ctx, modInfo, modInfo.ExtendedModName, modInfo.ExtendedModConf, modInfo.ExtendedModDir); err != nil {
				return errors.WrapError(ctx, err)
			}
		}

		if err := svc.LoadFiles(ctx, modInfo, modName, modInfo.ModConf, modInfo.ModDir); err != nil {
			return errors.WrapError(ctx, err)
		}
		_, modRead = svc.uiFiles[modName]
	}

	/*
		instance := insName
		if parentIns != "" {
			instance = parentIns
		}


	*/

	_, modExists := svc.uiFiles[modName]
	_, extendedModExists := svc.uiFiles[modInfo.ExtendedModName]

	log.Trace(ctx, "Mod exists ", "application", svc.application, "modname", modName, "extended mod name", modInfo.ExtendedModName, "mod exists", modExists, "extended mod exists", extendedModExists)

	if modExists || extendedModExists {
		log.Error(ctx, "creating instance", "modName", modName, "modInfo.InstanceName", modInfo.InstanceName)
		svc.insSettings[modInfo.InstanceName] = modInfo.ModSettings
		if modInfo.IsExtended {
			svc.insMods[modInfo.InstanceName] = modInfo.ExtendedModName
		} else {
			svc.insMods[modInfo.InstanceName] = modName
		}
	}

	err := svc.readRegistry(ctx, modName, modInfo.Mod, modInfo.ModConf, modInfo.ModDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.appendPropertyFiles(ctx, modName, modInfo.ModProps)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	publicImgs := path.Join(modFilesDir, "images")
	ok, _, _ = utils.FileExists(publicImgs)
	if ok {
		log.Debug(ctx, "Processing images", "dir", publicImgs)
		err = svc.copyImages(ctx, modName, publicImgs)
		if err != nil {
			return errors.WrapError(ctx, err)
		}

	}

	fontsDir := path.Join(modFilesDir, "fonts")
	ok, _, _ = utils.FileExists(fontsDir)
	if ok {
		log.Debug(ctx, "Copying fonts ", "fonts", fontsDir)
		err = svc.copyFonts(ctx, modName, fontsDir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	if modInfo.Mod != nil {
		uiplugin, ok := modInfo.Mod.(UIPlugin)
		if ok {
			svc.uiPlugins[modInfo.InstanceName] = modInfo
			comps := uiplugin.UILoad(ctx)
			reg, _ := comps["registry"]
			if reg != nil {
				registry, _ := reg.(config.Config)
				err := svc.processRegistryConfig(ctx, registry, modInfo.ModDir)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
	}

	return nil
}

func (svc *UI) LoadFiles(ctx core.ServerContext, modInfo *components.ModInfo, modName string, modConf config.Config, modDir string) error {

	modFilesDir := svc.getFilesDir(ctx, modName, modDir, modInfo)

	modDeps, ok := modInfo.ModConf.GetSubConfig(ctx, DEPENDENCIES)
	if ok {
		svc.modDeps[modName] = modDeps.AllConfigurations(ctx)
	}

	uifile := path.Join(modFilesDir, SCRIPTS_DIR, svc.uifile)

	ok, _, _ = utils.FileExists(uifile)
	if ok {
		cont, err := ioutil.ReadFile(uifile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		ui, ok := modConf.GetSubConfig(ctx, "ui")
		if ok {
			pkgs, ok := ui.GetSubConfig(ctx, "packages")
			if ok {
				pkgNames := pkgs.AllConfigurations(ctx)
				svc.requiredUIPkgs.Append(pkgNames)
			}
		}
		log.Debug(ctx, "adding files", "modname", modName)

		svc.uiFiles[modName] = cont
	}

	vendorfile := path.Join(modFilesDir, SCRIPTS_DIR, "vendor.js")
	ok, _, _ = utils.FileExists(vendorfile)
	if ok {
		log.Trace(ctx, "Reading vendor file", "file", vendorfile)
		cont, err := ioutil.ReadFile(vendorfile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.vendorFiles[fmt.Sprintf("%s_vendor", modName)] = cont
	}

	cssfile := path.Join(modFilesDir, CSS_DIR, "app.css")
	ok, _, _ = utils.FileExists(cssfile)
	log.Error(ctx, "CSS file exists", "file", cssfile, "ok", ok)
	if ok {
		cont, err := ioutil.ReadFile(cssfile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.cssFiles[modName] = cont
	}

	wasmfile := path.Join(modFilesDir, WASM_DIR, modName+".wasm")
	ok, _, _ = utils.FileExists(wasmfile)
	if ok {
		svc.wasmFiles[modName] = wasmfile
		/*	cont, err := ioutil.ReadFile(wasmfile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			svc.wasmFiles[modName] = cont*/
		//log.Error(ctx, "Wasm file read", "Mod", modName)
	}

	wasmimportsfile := path.Join(modFilesDir, WASM_DIR, modName+".js")
	ok, _, _ = utils.FileExists(wasmimportsfile)
	if ok {
		svc.wasmImportFiles[modName] = wasmimportsfile
	}
	/*svc.wasmImportFiles[modName] = path.Join(modFilesDir, WASM_DIR,  modName+".js")
	wasmimportsfile := path.Join(modFilesDir, WASM_DIR, modName+".js")
	ok, _, _ = utils.FileExists(wasmimportsfile)
	if ok {
		cont, err := ioutil.ReadFile(wasmimportsfile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.wasmImportFiles[modName] = cont
		//log.Error(ctx, "Wasm file read", "Mod", modName)
	}*/

	descfile := path.Join(modFilesDir, SCRIPTS_DIR, svc.descfile)
	ok, _, _ = utils.FileExists(descfile)
	if ok {
		cont, err := ioutil.ReadFile(descfile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.descriptorFiles[modName] = cont
	}

	return nil
}

func (svc *UI) getFilesDir(ctx core.ServerContext, modName string, modDir string, modInfo *components.ModInfo) string {
	return path.Join(modDir, FILES_DIR)
}

func (svc *UI) Loaded(ctx core.ServerContext) error {
	log.Error(ctx, " UI module loaded", "plugins", svc.uiPlugins)
	for _, modInfo := range svc.uiPlugins {
		log.Error(ctx, " UI module loaded", "mod info", modInfo)
		compsToProcess := modInfo.Mod.(UIPlugin).LoadingComplete(ctx)
		if compsToProcess != nil {
			reg, _ := compsToProcess["registry"]
			if reg != nil {
				registry, _ := reg.(config.Config)
				err := svc.processRegistryConfig(ctx, registry, modInfo.ModDir)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
	}
	log.Error(ctx, "going to write output", "service", svc)
	return svc.writeOutput(ctx)
}

func (svc *UI) writeOutput(ctx core.ServerContext) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)

	scriptsdir := path.Join(baseDir, FILES_DIR, SCRIPTS_DIR)
	err := os.MkdirAll(scriptsdir, 0700)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	cssdir := path.Join(baseDir, FILES_DIR, CSS_DIR)
	err = os.MkdirAll(cssdir, 0700)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	wasmdir := path.Join(baseDir, FILES_DIR, WASM_DIR)
	err = os.MkdirAll(wasmdir, 0700)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeVendorFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeCssFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeWasmFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeDescriptorFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writePropertyFiles(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeAppFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeBootFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
	/*
		deps := svc.requiredUIPkgs.Values()
		log.Info(ctx, "UI Packages required", "pkgs", deps)
		vendorFileCont.WriteString("var k = require('designerdependencies'); console.log(k);")
		for _, pkg := range deps {
			defineStr := "define('%s',[],function(){console.log('defined %s');var l= window['%s'];console.log(l);return l});"
			_, err = vendorFileCont.WriteString(fmt.Sprintf(defineStr, pkg, pkg, pkg))
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}*/

}

func (svc *UI) copyFonts(ctx core.ServerContext, mod, dirPath string) error {
	/*files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/

	basedir, _ := ctx.GetString(config.BASEDIR)
	/*
		err = os.MkdirAll(filepath.Join(basedir, "images", mod), 0700)
		if err != nil {
			return errors.WrapError(ctx, err)
		}*/
	dest := filepath.Join(basedir, "fonts")
	log.Debug(ctx, "Copying fonts", "src", dirPath, "dest", dest)

	err := utils.CopyDir(dirPath, filepath.Join(basedir, "fonts"), "")
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	/*
		for _, info := range files {
			if info.IsDir() {
				continue
			}
			path := filepath.Join(dirPath, info.Name())

			dest := filepath.Join(basedir, "images", info.Name())
			log.Trace(ctx, "Copying file", "src", path, "dest", dest)
			err = utils.CopyFile(path, dest)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}*/
	return nil
}

func (svc *UI) copyImages(ctx core.ServerContext, mod, dirPath string) error {
	/*files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/

	basedir, _ := ctx.GetString(config.BASEDIR)
	/*
		err = os.MkdirAll(filepath.Join(basedir, "images", mod), 0700)
		if err != nil {
			return errors.WrapError(ctx, err)
		}*/
	dest := filepath.Join(basedir, "images")
	log.Error(ctx, "Copying images ", "src", dirPath, "dest", dest)
	err := utils.CopyDir(dirPath, dest, "")
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	/*
		for _, info := range files {
			if info.IsDir() {
				continue
			}
			path := filepath.Join(dirPath, info.Name())

			dest := filepath.Join(basedir, "images", info.Name())
			log.Trace(ctx, "Copying file", "src", path, "dest", dest)
			err = utils.CopyFile(path, dest)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}*/
	return nil
}

func (svc *UI) Unloading(ctx core.ServerContext, insName, modName string) error {
	delete(svc.uiFiles, modName)
	delete(svc.vendorFiles, modName)
	delete(svc.cssFiles, modName)
	delete(svc.wasmFiles, modName)
	delete(svc.wasmImportFiles, modName)
	delete(svc.descriptorFiles, modName)
	delete(svc.propertyFiles, modName)
	delete(svc.uiRegistry, modName)
	return nil
}

func (svc *UI) Unloaded(ctx core.ServerContext, insName, modName string) error {
	if err := svc.writeOutput(ctx); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
