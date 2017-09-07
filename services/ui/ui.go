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
	UI_SVC           = "ui"
	FILES_DIR        = "files"
	CONF_FILE_UI     = "uifile"
	CONF_FILE_DESC   = "descriptor"
	CONF_APPLICATION = "application"
	DEPENDENCIES     = "dependencies"
	MERGED_SVCS_FILE = "mergeduidescriptor"
	MERGED_UI_FILE   = "mergeduifile"
)

type UI struct {
	core.Service
	uifile             string
	descfile           string
	mergeduifile       string
	mergeduidescriptor string
	application        string
	uiFiles            map[string][]byte
	uiDeps             map[string][]string
	insMods            map[string]string
	insSettings        map[string]config.Config
	descriptorFiles    map[string][]byte
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
	svc.mergeduidescriptor, _ = svc.GetStringConfiguration(ctx, MERGED_SVCS_FILE)
	svc.application, _ = svc.GetStringConfiguration(ctx, CONF_APPLICATION)
	svc.uiFiles = make(map[string][]byte)
	svc.insSettings = make(map[string]config.Config)
	svc.insMods = make(map[string]string)
	svc.descriptorFiles = make(map[string][]byte)
	svc.uiDeps = make(map[string][]string)
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

			uiDeps, ok := modConf.GetSubConfig(ctx, DEPENDENCIES)
			if ok {
				svc.uiDeps[modName] = uiDeps.AllConfigurations(ctx)
			}

			svc.uiFiles[modName] = cont
			modRead = true
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

	baseDir, _ := ctx.GetString(config.MODULEDIR)
	almondjsfile := path.Join(baseDir, FILES_DIR, "almond.js")
	almondjs, err := ioutil.ReadFile(almondjsfile)
	if err != nil {
		return errors.WrapError(ctx, err, "basedir", baseDir)
	}
	_, err = uiFileCont.Write(almondjs)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	filesWritten := make(map[string]bool)
	for name, _ := range svc.uiFiles {
		err := svc.appendContent(ctx, name, uiFileCont, filesWritten)
		if err != nil {
			return err
		}
	}

	descFileCont := new(bytes.Buffer)
	for _, cont := range svc.descriptorFiles {
		_, err := descFileCont.Write(cont)
		if err != nil {
			return err
		}
	}

	modTemplate := "define('%s', ['%s'], function (m) { var conf=%s; if(m.Initialize) { console.log('Initializing %s'); m.Initialize(%s, conf); } return m; })"

	for insName, settings := range svc.insSettings {
		settingsStr, err := json.Marshal(settings)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		descFileCont.WriteString(fmt.Sprintf(modTemplate, insName, svc.insMods[insName], string(settingsStr), insName, svc.application))
	}

	initFunc := fmt.Sprintf("function InitializeApplication(){document.InitConfig={Name: '%s', Services:{}, Actions:{}}; var app=require('%s'); console.log(app); app.StartApplication();}", svc.application, svc.application)
	descFileCont.WriteString(initFunc)

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

func (svc *UI) appendContent(ctx core.ServerContext, name string, buf *bytes.Buffer, writtenMods map[string]bool) error {
	log.Info(ctx, "Appending module content ", "Name", name)
	deps := svc.uiDeps[name]
	for _, dep := range deps {
		_, ok := writtenMods[dep]
		if !ok {
			err := svc.appendContent(ctx, dep, buf, writtenMods)
			if err != nil {
				return err
			}
		}
	}
	cont, ok := svc.uiFiles[name]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_MODULE, "Module ", name)
	}
	_, err := buf.Write(cont)
	if err == nil {
		writtenMods[name] = true
	}
	return err
}
