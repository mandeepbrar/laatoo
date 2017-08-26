package main

import (
	"bytes"
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
	CONF_FILE_UI     = "uifile"
	CONF_FILE_DESC   = "descriptor"
	MERGED_SVCS_FILE = "services.conf.js"
	MERGED_UI_FILE   = "app.js"
)

type UI struct {
	core.Service
	uifile          string
	descfile        string
	uiFiles         map[string][]byte
	uiDeps          map[string][]string
	descriptorFiles map[string][]byte
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
	svc.uiFiles = make(map[string][]byte)
	svc.descriptorFiles = make(map[string][]byte)
	svc.uiDeps = make(map[string][]string)
	return nil
}

func (svc *UI) Load(ctx core.ServerContext, name, dir string, mod core.Module, modConf config.Config) error {
	uifile := path.Join(dir, UI_SVC, svc.uifile)
	ok, _, _ := utils.FileExists(uifile)
	if ok {
		cont, err := ioutil.ReadFile(uifile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.uiFiles[name] = cont
	}

	descfile := path.Join(dir, UI_SVC, svc.descfile)
	ok, _, _ = utils.FileExists(descfile)
	if ok {
		cont, err := ioutil.ReadFile(descfile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.descriptorFiles[name] = cont
	}
	return nil
}

func (svc *UI) Loaded(ctx core.ServerContext) error {
	uiFileCont := new(bytes.Buffer)

	baseDir, _ := ctx.GetString(config.MODULEDIR)
	log.Error(ctx, "base directory of module", "files", baseDir)
	almondjsfile := path.Join(baseDir, "files", "almond.js")
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

	uifile := path.Join(baseDir, "files", MERGED_UI_FILE)
	err = ioutil.WriteFile(uifile, uiFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	descfile := path.Join(baseDir, "files", MERGED_SVCS_FILE)
	err = ioutil.WriteFile(descfile, descFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) appendContent(ctx core.ServerContext, name string, buf *bytes.Buffer, writtenMods map[string]bool) error {
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
