package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"path"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: UI_SVC, Object: UI{}}}
}

const (
	UI_SVC             = "ui"
	FILES_DIR          = "files"
	CONF_FILE_UI       = "uifile"
	CONF_FILE_DESC     = "descriptor"
	CONF_APPLICATION   = "application"
	DEPENDENCIES       = "dependencies"
	MERGED_SVCS_FILE   = "mergeduidescriptor"
	MERGED_CSS_FILE    = "mergedcssfile"
	MERGED_VENDOR_FILE = "mergedvendorfile"
	MERGED_UI_FILE     = "mergeduifile"
)

type UI struct {
	core.Service
	uifile             string
	descfile           string
	mergeduifile       string
	mergedvendorfile   string
	mergeduidescriptor string
	mergedcssfile      string
	application        string
	uiFiles            map[string][]byte
	vendorFiles        map[string][]byte
	cssFiles           map[string][]byte
	modDeps            map[string][]string
	insMods            map[string]string
	insSettings        map[string]config.Config
	descriptorFiles    map[string][]byte
	requiredUIPkgs     utils.StringSet
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
	svc.uifile, _ = svc.GetStringConfiguration(ctx, CONF_FILE_UI)
	svc.descfile, _ = svc.GetStringConfiguration(ctx, CONF_FILE_DESC)
	svc.mergeduifile, _ = svc.GetStringConfiguration(ctx, MERGED_UI_FILE)
	svc.mergedvendorfile, _ = svc.GetStringConfiguration(ctx, MERGED_VENDOR_FILE)
	svc.mergeduidescriptor, _ = svc.GetStringConfiguration(ctx, MERGED_SVCS_FILE)
	svc.mergedcssfile, _ = svc.GetStringConfiguration(ctx, MERGED_CSS_FILE)
	svc.application, _ = svc.GetStringConfiguration(ctx, CONF_APPLICATION)
	svc.uiFiles = make(map[string][]byte)
	svc.vendorFiles = make(map[string][]byte)
	svc.cssFiles = make(map[string][]byte)
	svc.insSettings = make(map[string]config.Config)
	svc.insMods = make(map[string]string)
	svc.descriptorFiles = make(map[string][]byte)
	svc.modDeps = make(map[string][]string)
	svc.requiredUIPkgs = utils.NewStringSet([]string{})
	return nil
}

func (svc *UI) Load(ctx core.ServerContext, insName, modName, dir string, mod core.Module, modConf config.Config, settings config.Config) error {
	ui, ok := settings.GetBool(ctx, UI_SVC)
	if ok && !ui {
		return nil
	}

	app, ok := settings.GetString(ctx, CONF_APPLICATION)
	if app != "" && svc.application != app {
		log.Debug(ctx, "Skipping module from ui", "module", modName, "application", svc.application)
		return nil
	}

	_, modRead := svc.uiFiles[modName]

	if !modRead {
		uifile := path.Join(dir, FILES_DIR, svc.uifile)
		ok, _, _ = utils.FileExists(uifile)
		if ok {
			cont, err := ioutil.ReadFile(uifile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}

			modDeps, ok := modConf.GetSubConfig(ctx, DEPENDENCIES)
			if ok {
				svc.modDeps[modName] = modDeps.AllConfigurations(ctx)
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
			modRead = true
		}

		vendorfile := path.Join(dir, FILES_DIR, "vendor.js")
		ok, _, _ = utils.FileExists(vendorfile)
		if ok {
			cont, err := ioutil.ReadFile(vendorfile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			svc.vendorFiles[modName] = cont
		}

		cssfile := path.Join(dir, FILES_DIR, "app.css")
		ok, _, _ = utils.FileExists(cssfile)
		if ok {
			cont, err := ioutil.ReadFile(cssfile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			svc.cssFiles[modName] = cont
		}

		descfile := path.Join(dir, FILES_DIR, svc.descfile)
		ok, _, _ = utils.FileExists(descfile)
		if ok {
			cont, err := ioutil.ReadFile(descfile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			svc.descriptorFiles[modName] = cont
		}

	}

	if modRead && (modName != insName) {
		svc.insSettings[insName] = settings
		svc.insMods[insName] = modName
	}

	return nil
}

func (svc *UI) Loaded(ctx core.ServerContext) error {
	uiFileCont := new(bytes.Buffer)
	cssFileCont := new(bytes.Buffer)
	descFileCont := new(bytes.Buffer)
	vendorFileCont := new(bytes.Buffer)
	_, err := vendorFileCont.WriteString(fmt.Sprintf("document.InitConfig={Name: '%s', Services:{}, Actions:{}};", svc.application))
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	baseDir, _ := ctx.GetString(config.MODULEDIR)

	almondjsfile := path.Join(baseDir, FILES_DIR, "almond.js")
	almondjs, err := ioutil.ReadFile(almondjsfile)
	if err != nil {
		return errors.WrapError(ctx, err, "basedir", baseDir)
	}
	_, err = vendorFileCont.Write(almondjs)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	/*
		err = svc.writeDependenciesSourceFile(ctx, baseDir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	*/
	for mod, cont := range svc.vendorFiles {
		_, err = vendorFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		initStr := "var k=require('%s_vendor'); if(k && k.Initialize) {k.Initialize('%s',{},define);}"
		_, err = vendorFileCont.WriteString(fmt.Sprintf(initStr, mod, svc.application))
	}

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

	filesWritten := make(map[string]bool)
	for name, _ := range svc.uiFiles {
		err := svc.appendContent(ctx, name, uiFileCont, filesWritten)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	reqTemplate := "var l=require('%s');"

	for name, _ := range filesWritten {
		_, err = uiFileCont.WriteString(fmt.Sprintf(reqTemplate, name))
		if err != nil {
			return err
		}
	}

	for _, cont := range svc.cssFiles {
		_, err := cssFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	cssfile := path.Join(baseDir, FILES_DIR, svc.mergedcssfile)
	err = ioutil.WriteFile(cssfile, cssFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	for _, cont := range svc.descriptorFiles {
		_, err := descFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	modTemplate := "define('%s', ['%s'], function (m) { var conf=%s; if(m.Initialize) { m.Initialize('%s', conf, define, require); } return m; });" + reqTemplate

	for insName, settings := range svc.insSettings {
		settingsStr, err := json.Marshal(settings)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		uiFileCont.WriteString(fmt.Sprintf(modTemplate, insName, svc.insMods[insName], string(settingsStr), svc.application, insName))
	}

	initFunc := fmt.Sprintf("function InitializeApplication(){var app=require('%s'); app.StartApplication();};", svc.application)
	descFileCont.WriteString(initFunc)

	vendorfile := path.Join(baseDir, FILES_DIR, svc.mergedvendorfile)
	err = ioutil.WriteFile(vendorfile, vendorFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	uifile := path.Join(baseDir, FILES_DIR, svc.mergeduifile)
	err = ioutil.WriteFile(uifile, uiFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	descfile := path.Join(baseDir, FILES_DIR, svc.mergeduidescriptor)
	err = ioutil.WriteFile(descfile, descFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

/*/
func (svc *UI) writeDependenciesSourceFile(ctx core.ServerContext, baseDir string) error {
	deps := svc.requiredUIPkgs.Values()
	depsStr, err := json.Marshal(deps)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	srcFile := "var deps=%s;var retVal={}; deps.forEach(function(k){var m = require(k); retVal[k]=m;});module.exports=retVal;"

	srcfile := path.Join(baseDir, "src", "index.js")
	srcCont := fmt.Sprintf(srcFile, depsStr)
	err = ioutil.WriteFile(srcfile, []byte(srcCont), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}*/

func (svc *UI) appendContent(ctx core.ServerContext, name string, buf *bytes.Buffer, writtenMods map[string]bool) error {
	log.Info(ctx, "Appending module content ", "Name", name)

	cont, uiprs := svc.uiFiles[name]
	if !uiprs {
		return nil
	}

	deps := svc.modDeps[name]
	for _, dep := range deps {
		_, ok := writtenMods[dep]
		if !ok {
			err := svc.appendContent(ctx, dep, buf, writtenMods)
			if err != nil {
				return err
			}
		}
	}
	/*cont, ok := svc.uiFiles[name]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Module ", name)
	}*/
	_, err := buf.Write(cont)
	if err == nil {
		writtenMods[name] = true
	}
	return err
}
