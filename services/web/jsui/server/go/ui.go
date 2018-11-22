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
	CONF_APPLICATION     = "application"
	CONF_HOT_MODULES     = "hotmodules"
	DEPENDENCIES         = "dependencies"
	UI_DIR               = "ui"
	MERGED_SVCS_FILE     = "mergeduidescriptor"
	MERGED_CSS_FILE      = "mergedcssfile"
	MERGED_WASM_FILE     = "mergedwasm"
	MERGED_VENDOR_FILE   = "mergedvendorfile"
	MERGED_UI_FILE       = "mergeduifile"
	HOT_MODULES_REPO     = "hotmodulesrepo"
)

type UI struct {
	core.Service
	svrCtx             core.ServerContext
	uifile             string
	descfile           string
	mergeduifile       string
	mergedvendorfile   string
	mergeduidescriptor string
	mergedwasmfile     string
	mergedcssfile      string
	application        string
	propsExt           string
	hotModulesRepo     string
	uiFiles            map[string][]byte
	vendorFiles        map[string][]byte
	cssFiles           map[string][]byte
	wasmFiles          map[string]string
	wasmImportFiles    map[string]string
	wasmImportScript   []byte
	modDeps            map[string][]string
	insMods            map[string]string
	hotloadMods        map[string]string
	insSettings        map[string]config.Config
	descriptorFiles    map[string][]byte
	requiredUIPkgs     utils.StringSet
	propertyFiles      map[string]interface{}
	uiRegistry         map[string]map[string]string
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
	svc.application, _ = svc.GetStringConfiguration(ctx, CONF_APPLICATION)
	svc.propsExt, _ = svc.GetStringConfiguration(ctx, CONF_PROPS_EXTENSION)
	svc.hotloadMods, _ = svc.GetStringsMapConfiguration(ctx, CONF_HOT_MODULES)
	svc.watchers = make([]*fsnotify.Watcher, 0)
	log.Error(ctx, "*************hot modules directory being used**********", "hotloadMods", svc.hotloadMods)
	svc.hotModulesRepo, _ = ctx.GetString(HOT_MODULES_REPO)
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
	log.Error(ctx, "Loading module", "modInfo", modInfo, "application", svc.application)

	app, ok := modInfo.ModSettings.GetString(ctx, CONF_APPLICATION)
	if app != "" && svc.application != app {
		log.Error(ctx, "Skipping module from ui", "module", modName, "application", svc.application)
		return nil
	}
	log.Error(ctx, "Module read", "mod name", modName)

	modDevDir, hot := svc.hotloadMods[modName]
	modFilesDir := ""
	if hot {
		modFilesDir = path.Join(svc.hotModulesRepo, modDevDir, FILES_DIR)
	} else {
		modFilesDir = path.Join(modInfo.ModDir, FILES_DIR)
	}

	_, modRead := svc.uiFiles[modName]

	if !modRead {
		if modInfo.IsExtended {
			if err := svc.LoadFiles(ctx, modInfo, modInfo.ExtendedModName, modInfo.ExtendedModConf, modInfo.ExtendedModDir, modFilesDir, hot); err != nil {
				return errors.WrapError(ctx, err)
			}
		}

		if err := svc.LoadFiles(ctx, modInfo, modName, modInfo.ModConf, modInfo.ModDir, modFilesDir, hot); err != nil {
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
	if modRead {
		log.Error(ctx, "creating instance", "modName", modName, "modInfo.InstanceName", modInfo.InstanceName)
		svc.insSettings[modInfo.InstanceName] = modInfo.ModSettings
		svc.insMods[modInfo.InstanceName] = modName
	}

	uiRegDir := path.Join(modInfo.ModDir, UI_DIR)
	err := svc.readRegistry(ctx, modInfo.Mod, modInfo.ModConf, modInfo.ModDir, uiRegDir)
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
	return nil
}

func (svc *UI) LoadFiles(ctx core.ServerContext, modInfo *components.ModInfo, modName string, modConf config.Config, modDir string, modFilesDir string,
	hot bool) error {
	modDeps, ok := modInfo.ModConf.GetSubConfig(ctx, DEPENDENCIES)
	if ok {
		svc.modDeps[modName] = modDeps.AllConfigurations(ctx)
	}

	uifile := path.Join(modFilesDir, SCRIPTS_DIR, svc.uifile)

	if hot {
		log.Info(ctx, "*************hot modules directory being used**********", "modFilesDir", modFilesDir)
		svc.addWatch(ctx, modName, uifile, modFilesDir, svc.reloadAppFile)
	}

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

		svc.uiFiles[modName] = cont
	}

	vendorfile := path.Join(modFilesDir, SCRIPTS_DIR, "vendor.js")
	ok, _, _ = utils.FileExists(vendorfile)
	if ok {
		if hot {
			svc.addWatch(ctx, modName, vendorfile, modFilesDir, svc.reloadVendorFile)
		}
		log.Trace(ctx, "Reading vendor file", "file", vendorfile)
		cont, err := ioutil.ReadFile(vendorfile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.vendorFiles[fmt.Sprintf("%s_vendor", modName)] = cont
	}

	cssfile := path.Join(modFilesDir, CSS_DIR, "app.css")
	ok, _, _ = utils.FileExists(cssfile)
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

func (svc *UI) Loaded(ctx core.ServerContext) error {
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
	log.Debug(ctx, "Copying images ", "src", dirPath, "dest", dest)
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
