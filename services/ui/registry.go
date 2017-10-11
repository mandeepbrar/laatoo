package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
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

func (svc *UI) processFile(ctx core.ServerContext, cont []byte, filetype string, itemName string, itemType string) error {
	if filetype == ".json" {
		if itemType == "Block" {
			err := svc.createJsonBlock(ctx, itemType, itemName, cont)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			svc.addRegItem(ctx, itemType, itemName, string(cont))
		}
	}
	if itemType == ".xml" {
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

func (svc *UI) processItemDir(ctx core.ServerContext, dirPath string, itemType string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		itemName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		cont, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		if filepath.Ext(path) == ".json" {
			err = svc.processFile(ctx, cont, ".json", itemName, itemType)
		}
		if filepath.Ext(path) == ".tpl" {
			fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
			itemName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			err = svc.processTemplate(ctx, cont, filepath.Ext(fileName), itemName, itemType)
		}
		if filepath.Ext(path) == ".xml" {
			err = svc.processFile(ctx, cont, ".xml", itemName, itemType)
		}
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		return nil
	})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

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
}

func (svc *UI) readRegistry(ctx core.ServerContext, modConf config.Config, regDir string) error {

	ui, ok := modConf.GetSubConfig(ctx, "ui")
	if ok {
		registry, ok := ui.GetSubConfig(ctx, "registry")
		if ok {
			confs := registry.AllConfigurations(ctx)
			for _, itemType := range confs {
				items, ok := registry.GetSubConfig(ctx, itemType)
				if ok {
					err := svc.processConfig(ctx, items, itemType)
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
			if filepath.Ext(path) == ".json" {
				items, err := ctx.ReadConfig(path)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				err = svc.processConfig(ctx, items, itemType)
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

func (svc *UI) processConfig(ctx core.ServerContext, conf config.Config, itemType string) error {
	itemNames := conf.AllConfigurations(ctx)
	for _, item := range itemNames {
		itemVal, _ := conf.Get(ctx, item)
		strVal, err := json.Marshal(itemVal)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		svc.addRegItem(ctx, itemType, item, string(strVal))
	}
	return nil
}
