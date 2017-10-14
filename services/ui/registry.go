package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"os"
	"path/filepath"
	"strings"
)

func (svc *UI) addRegItem(ctx core.ServerContext, itemType string, itemName string, itemStr string) {
	var itemReg map[string]string
	reg, ok := svc.uiRegistry[itemType]
	if ok {
		itemReg = reg
	} else {
		itemReg = make(map[string]string)
		svc.uiRegistry[itemType] = itemReg
	}
	itemReg[itemName] = fmt.Sprintf("_r('%s', '%s', %s);", itemType, itemName, itemStr)
}

/*
func (svc *UI) processFile(ctx core.ServerContext, cont []byte, filetype string, itemName string, itemType string) error {
	ctx = ctx.SubContext(itemName)
	log.Error(ctx, "processing file", "filetype", filetype, "itemName", itemName, " cont", string(cont))
	if (filetype == ".json") || (filetype == ".yml") {
processConfig
		if itemType == "Block" {
			err := svc.createYmlBlock(ctx, itemType, itemName, cont)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			svc.addRegItem(ctx, itemType, itemName, string(cont))
		}
	}
	if filetype == ".xml" {
		buf := bytes.NewBuffer(cont)
		dec := xml.NewDecoder(buf)
		var n Node
		err := dec.Decode(&n)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		err = svc.createXMLBlock(ctx, itemType, itemName, n)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
*/
func (svc *UI) processItemDir(ctx core.ServerContext, dirPath string, itemType string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		itemName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		ext := filepath.Ext(path)
		var cont []byte
		if ext == ".yml" || ext == ".json" {
			conf, err := ctx.ReadConfig(path, nil)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			err = svc.processConfig(ctx, conf, itemName, itemType)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			if itemType == "Block" && ext == ".xml" {
				cont, err = ioutil.ReadFile(path)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				buf := bytes.NewBuffer(cont)
				dec := xml.NewDecoder(buf)
				var n Node
				err := dec.Decode(&n)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				err = svc.createXMLBlock(ctx, itemType, itemName, n)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
		/*err = svc.processFile(ctx, cont, ext, itemName, itemType)
		if err != nil {
			return errors.WrapError(ctx, err)
		}*/
		return nil
	})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

/*
func (svc *UI) processTemplate(ctx core.ServerContext, cont []byte, filetype string, itemName string, itemType string) error {
	contextVar := func(variable string) string {
		val, _ := ctx.GetString(variable)
		return val
	}
	funcMap := template.FuncMap{"var": contextVar}
	temp, err := template.New(itemName).Funcs(funcMap).Parse(string(cont))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	result := new(bytes.Buffer)
	anon := struct{}{}
	err = temp.Execute(result, anon)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return svc.processFile(ctx, result.Bytes(), filetype, itemName, itemType)
}*/

func (svc *UI) readRegistry(ctx core.ServerContext, modConf config.Config, regDir string) error {

	ui, ok := modConf.GetSubConfig(ctx, "ui")
	if ok {
		registry, ok := ui.GetSubConfig(ctx, "registry")
		if ok {
			confs := registry.AllConfigurations(ctx)
			for _, itemType := range confs {
				items, ok := registry.GetSubConfig(ctx, itemType)
				if ok {
					err := svc.processMutipleItems(ctx, items, itemType)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
				}
			}
		}
	}

	ok, _, _ = utils.FileExists(regDir)
	if ok {
		err := filepath.Walk(regDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && regDir != path {
				svc.processItemDir(ctx, path, info.Name())
				return nil
			}
			fileName := info.Name()
			itemType := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			ext := filepath.Ext(path)
			if ext == ".yml" || ext == ".json" {
				items, err := ctx.ReadConfig(path, nil)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				err = svc.processMutipleItems(ctx, items, itemType)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
			return nil
		})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *UI) processMutipleItems(ctx core.ServerContext, conf config.Config, itemType string) error {
	itemNames := conf.AllConfigurations(ctx)
	for _, item := range itemNames {
		itemVal, _ := conf.GetSubConfig(ctx, item)
		err := svc.processConfig(ctx, itemVal, item, itemType)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *UI) processConfig(ctx core.ServerContext, conf config.Config, itemName, itemType string) error {
	if itemType == "Block" {
		err := svc.createConfBlock(ctx, itemType, itemName, conf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		strVal, err := json.Marshal(conf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.addRegItem(ctx, itemType, itemName, string(strVal))
	}
	return nil
}