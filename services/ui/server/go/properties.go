package main

import (
	"encoding/json"
	"io/ioutil"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"os"
	"path"

	"github.com/imdario/mergo"
)

func (svc *UI) appendPropertyFiles(ctx core.ServerContext, modName string, props map[string]interface{}) error {
	modprops := make(map[string]interface{})
	if props != nil {
		for locale, val := range props {
			modprops[locale] = map[string]interface{}{modName: val}
		}
	}
	if len(modprops) > 0 {
		props := svc.propertyFiles
		mergo.Merge(&props, modprops)
		svc.propertyFiles = props
	}
	return nil
}

func (svc *UI) writePropertyFiles(ctx core.ServerContext, baseDir string) error {
	svrProps := ctx.GetServerProperties()

	propsToWrite := make(map[string]interface{})
	mergo.Merge(&propsToWrite, svrProps)
	mergo.Merge(&propsToWrite, svc.propertyFiles)

	/*for mod, val := range svc.propertyFiles {
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
	}*/
	if len(propsToWrite) > 0 {
		err := os.MkdirAll(path.Join(baseDir, FILES_DIR, "properties"), 0755)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	log.Trace(ctx, "Writing properties", "props", propsToWrite)
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
