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
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: UI_SVC, Object: UI{}}}
}

const (
	UI_SVC               = "ui"
	FILES_DIR            = "files"
	CONF_FILE_UI         = "uifile"
	CONF_FILE_DESC       = "descriptor"
	CONF_PROPS_EXTENSION = "properties_ext"
	CONF_APPLICATION     = "application"
	DEPENDENCIES         = "dependencies"
	PROPERTIES_DIR       = "properties"
	MERGED_SVCS_FILE     = "mergeduidescriptor"
	MERGED_CSS_FILE      = "mergedcssfile"
	MERGED_VENDOR_FILE   = "mergedvendorfile"
	MERGED_UI_FILE       = "mergeduifile"
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
	propsExt           string
	uiFiles            map[string][]byte
	vendorFiles        map[string][]byte
	cssFiles           map[string][]byte
	modDeps            map[string][]string
	insMods            map[string]string
	insSettings        map[string]config.Config
	descriptorFiles    map[string][]byte
	requiredUIPkgs     utils.StringSet
	propertyFiles      map[string]interface{}
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
	svc.propsExt, _ = svc.GetStringConfiguration(ctx, CONF_PROPS_EXTENSION)
	svc.uiFiles = make(map[string][]byte)
	svc.vendorFiles = make(map[string][]byte)
	svc.cssFiles = make(map[string][]byte)
	svc.insSettings = make(map[string]config.Config)
	svc.insMods = make(map[string]string)
	svc.descriptorFiles = make(map[string][]byte)
	svc.modDeps = make(map[string][]string)
	svc.propertyFiles = make(map[string]interface{})
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

		err := svc.readPropertyFiles(ctx, modName, path.Join(dir, FILES_DIR, PROPERTIES_DIR))
		if err != nil {
			return errors.WrapError(ctx, err)
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

func (svc *UI) readPropertyFiles(ctx core.ServerContext, modName string, propsDir string) error {
	modprops := make(map[string]interface{})
	ok, _, _ := utils.FileExists(propsDir)
	if !ok {
		return nil
	}
	err := filepath.Walk(propsDir, func(path string, info os.FileInfo, err error) error {
		log.Error(ctx, "Reading properties file", "path", path)
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			cont, err := ioutil.ReadFile(path)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			log.Error(ctx, "Reading properties file", "cont", string(cont))
			obj := make(map[string]interface{})
			err = json.Unmarshal(cont, &obj)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			fileName := info.Name()
			locale := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			modprops[locale] = obj
		}
		return nil
	})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if len(modprops) > 0 {
		svc.propertyFiles[modName] = modprops
	}
	return nil
}

func (svc *UI) Loaded(ctx core.ServerContext) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)

	err := svc.writeVendorFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svc.writeCssFile(ctx, baseDir)
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

func (svc *UI) writePropertyFiles(ctx core.ServerContext, baseDir string) error {
	log.Error(ctx, "Writing property files", "map", svc.propertyFiles)
	propsToWrite := make(map[string]interface{})
	for mod, val := range svc.propertyFiles {
		localeProps, _ := val.(map[string]interface{})
		for locale, props := range localeProps {
			val, ok := propsToWrite[locale]
			var modProps map[string]interface{}
			if !ok {
				modProps = make(map[string]interface{})
				propsToWrite[locale] = modProps
			} else {
				modProps = val.(map[string]interface{})
			}
			modProps[mod] = props
			log.Error(ctx, "Writing property files", "mod", mod, "props", props, "locale", locale)
		}
	}
	if len(propsToWrite) > 0 {
		err := os.MkdirAll(path.Join(baseDir, FILES_DIR, "properties"), 0755)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	for locale, props := range propsToWrite {
		data, err := json.Marshal(props)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		localefile := path.Join(baseDir, FILES_DIR, "properties", locale+svc.propsExt)
		err = ioutil.WriteFile(localefile, data, 0755)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *UI) writeAppFile(ctx core.ServerContext, baseDir string) error {
	uiFileCont := new(bytes.Buffer)

	laatoocorejsfile := path.Join(baseDir, FILES_DIR, "laatoocore.js")
	laatoocorejs, err := os.Open(laatoocorejsfile)
	if err != nil {
		return errors.WrapError(ctx, err, "basedir", baseDir)
	}
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	err = m.Minify("text/javascript", uiFileCont, laatoocorejs)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	filesWritten := make(map[string]bool)
	for name, _ := range svc.uiFiles {
		err := svc.appendContent(ctx, name, uiFileCont, filesWritten)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	reqTemplate := "var l=require('%s');"

	for name, _ := range filesWritten {
		_, err := uiFileCont.WriteString(fmt.Sprintf(reqTemplate, name))
		if err != nil {
			return err
		}
	}
	/*
		initMods := new(bytes.Buffer)
		modTemplate := "define('%s', ['%s'], function (m) { var conf=%s; if(m.Initialize) { m.Initialize('%s', conf, define, require); } return m; });" + reqTemplate

		for insName, settings := range svc.insSettings {
			settingsStr, err := json.Marshal(settings)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			initMods.WriteString(fmt.Sprintf(modTemplate, insName, svc.insMods[insName], string(settingsStr), svc.application, insName))
		}

		initapp := "if (window.InitializeApplication) {console.log('Initializing app');window.InitializeApplication();}"
		propsLdr := fmt.Sprintf("var propsurl = window.location.origin+'/properties/default.%s.json'; fetch(propsurl).then(function(resp){resp.json().then(function(data){document.InitConfig.Properties=data;%s;%s;});});", svc.application, initMods.String(), initapp)
		_, err = uiFileCont.WriteString(propsLdr)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	*/

	modsList := new(bytes.Buffer)
	for insName, settings := range svc.insSettings {
		settingsStr, err := json.Marshal(settings)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		_, err = modsList.WriteString(fmt.Sprintf("['%s','%s',%s],", insName, svc.insMods[insName], string(settingsStr)))
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	/*settingsStr, err := json.Marshal(svc.insSettings)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/
	loadingComplete := fmt.Sprintf("var insSettings=[%s];appLoadingComplete('%s','%s',insSettings);", modsList.String(), svc.application, "/properties/default."+svc.application+".json")
	_, err = uiFileCont.WriteString(loadingComplete)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	uifile := path.Join(baseDir, FILES_DIR, svc.mergeduifile)
	err = ioutil.WriteFile(uifile, uiFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (svc *UI) writeDescriptorFile(ctx core.ServerContext, baseDir string) error {
	descFileCont := new(bytes.Buffer)
	for _, cont := range svc.descriptorFiles {
		_, err := descFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	initFunc := fmt.Sprintf("Window.InitializeApplication=function(){var app=require('%s'); app.StartApplication();};", svc.application)
	descFileCont.WriteString(initFunc)

	descfile := path.Join(baseDir, FILES_DIR, svc.mergeduidescriptor)
	err := ioutil.WriteFile(descfile, descFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (svc *UI) writeVendorFile(ctx core.ServerContext, baseDir string) error {
	vendorFileCont := new(bytes.Buffer)
	_, err := vendorFileCont.WriteString(fmt.Sprintf("document.InitConfig={Name: '%s', Services:{}, Actions:{}};", svc.application))
	if err != nil {
		return errors.WrapError(ctx, err)
	}

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

	vendorfile := path.Join(baseDir, FILES_DIR, svc.mergedvendorfile)
	err = ioutil.WriteFile(vendorfile, vendorFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) writeCssFile(ctx core.ServerContext, baseDir string) error {
	cssFileCont := new(bytes.Buffer)
	for _, cont := range svc.cssFiles {
		_, err := cssFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	cssfile := path.Join(baseDir, FILES_DIR, svc.mergedcssfile)
	err := ioutil.WriteFile(cssfile, cssFileCont.Bytes(), 0755)
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
