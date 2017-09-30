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

func (svc *UI) processItemDir(ctx core.ServerContext, dirPath string, itemType string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		itemName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		if filepath.Ext(path) == ".json" {
			cont, err := ioutil.ReadFile(path)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			svc.addRegItem(ctx, itemType, itemName, string(cont))
		}
		if filepath.Ext(path) == ".xml" {
			cont, err := ioutil.ReadFile(path)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			buf := bytes.NewBuffer(cont)
			dec := xml.NewDecoder(buf)
			var n Node
			err = dec.Decode(&n)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			err = svc.createDisplay(ctx, itemType, itemName, n)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
		return nil
	})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
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
				items, err := config.NewConfigFromFile(ctx, path)
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
