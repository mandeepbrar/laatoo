package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
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
		uiFileCont.Write([]byte(fmt.Sprintln("console.log('file written ", name, "');")))
	}

	//reqTemplate := "_rm('%s');"

	/*for name, _ := range filesWritten {
		_, err := uiFileCont.WriteString(fmt.Sprintf(reqTemplate, name))
		if err != nil {
			return err
		}
	}*/

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

	/*wasmMods := make([]string, 0)
	for modName, _ := range svc.wasmFiles {
		wasmMods = append(wasmMods, modName)
	}
	wasmModsStr := ""
	if len(wasmMods) > 0 {
		log.Error(ctx, "Joinging wasm mods", "mods", wasmMods)
		wasmModsbyt, err := json.Marshal(wasmMods)
		wasmModsStr = string(wasmModsbyt)
		log.Error(ctx, "Joinging wasm mods", "wasmModsStr", wasmModsStr)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}*/

	/*settingsStr, err := json.Marshal(svc.insSettings)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/
	loadingComplete := fmt.Sprintf("var insSettings=[%s];console.log('insSettings', insSettings);appLoadingComplete('%s','%s',insSettings, {'%s':'%s'});", modsList.String(), svc.application, "/properties/default."+svc.application+".json", svc.application, "/app."+svc.application+".wasm")
	_, err = uiFileCont.WriteString(loadingComplete)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	uifile := path.Join(baseDir, FILES_DIR, SCRIPTS_DIR, svc.mergeduifile)
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
	initFunc := fmt.Sprintf("console.log('initializing application js');Window.InitializeApplication=function(){var app=require('%s'); app.StartApplication();};", svc.application)
	descFileCont.WriteString(initFunc)

	descfile := path.Join(baseDir, FILES_DIR, SCRIPTS_DIR, svc.mergeduidescriptor)
	err := ioutil.WriteFile(descfile, descFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (svc *UI) writeVendorFile(ctx core.ServerContext, baseDir string) error {
	vendorFileCont := new(bytes.Buffer)
	_, err := vendorFileCont.WriteString(fmt.Sprintf("console.log('loading vendor file');document.InitConfig={Name: '%s', Services:{}, Actions:{}};", svc.application))
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
	nl := []byte(fmt.Sprintln(""))
	for _, cont := range svc.vendorFiles {
		vendorFileCont.Write(nl)
		_, err = vendorFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		vendorFileCont.Write(nl)
	}
	for mod, _ := range svc.vendorFiles {
		initStr := "var k=require('%s'); console.log(\"module found %s\", k); if(k && k.Initialize) {k.Initialize('%s',{},define);};"
		_, err = vendorFileCont.WriteString(fmt.Sprintf(initStr, mod, mod, svc.application))
	}
	vendorFileCont.WriteString("console.log('loaded vendor file');")
	vendorfile := path.Join(baseDir, FILES_DIR, SCRIPTS_DIR, svc.mergedvendorfile)
	log.Info(ctx, "Writing vendor file", "file", vendorfile)
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
	nl := []byte(fmt.Sprintln(""))

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
	buf.Write(nl)
	_, err := buf.Write(cont)
	buf.Write(nl)
	if err == nil {
		writtenMods[name] = true
	}
	return err
}

/*
/*Application.LoadWasm = function(mod, str) {
  try {
    var decodedMod = Application.Base64Decoder(str);
    WebAssembly.instantiate(decodedMod, Application.Modules).then(wasmModule => {
      Application.Modules[mod] = wasmModule.instance.exports;
      console.log("Application after loading mod", mod, Application);
      var testsock = Application.Modules["testsocket"];
      if(testsock) {
        console.log("testsocket", testsock);
        if(testsock.my_func) {
          alert(testsock.my_func());
        }
      }
    });
  }catch(ex) {
    console.log("exception in instantiating wasm", mod, ex);
  }
}*/
