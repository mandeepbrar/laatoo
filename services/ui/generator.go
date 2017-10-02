package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"os"
	"path"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

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

	reqTemplate := "_rm('%s');"

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

	for _, val := range svc.uiRegistry {
		for _, itemStr := range val {
			//regCont := fmt.Sprintf("_r('%s', '%s', %s);", itemType, item, itemStr)
			_, err := descFileCont.WriteString(itemStr)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	/*
		for _, dispFunc := range svc.uiDisplays {
			_, err := descFileCont.WriteString(dispFunc)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	*/
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
