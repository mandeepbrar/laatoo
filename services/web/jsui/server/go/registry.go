package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"path"
	"path/filepath"
	"strings"
)

const (
	ITEM_CONFIG      = "config"
	ITEM_SKIP        = "skip"
	BLOCK_REG        = "Blocks"
	FORM_REG         = "Forms"
	DEPENDENCIES_REG = "Dependencies"
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
	itemStrVal, _ := processJSRegex(ctx, itemStr)
	itemReg[itemName] = fmt.Sprintf("_r('%s', '%s', %s);", itemType, itemName, itemStrVal)
}

func (svc *UI) processItemDir(ctx core.ServerContext, modName string, dirPath string, itemType string, modDir string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, info := range files {
		if info.IsDir() {
			continue
		}
		path := filepath.Join(dirPath, info.Name())
		err = svc.processRegItem(ctx, path, itemType, modDir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *UI) processRegItem(ctx core.ServerContext, path string, itemType string, modDir string) error {
	ctx = ctx.SubContext("Process reg item " + path)
	ext := filepath.Ext(path)
	fileName := filepath.Base(path)
	itemName := strings.TrimSuffix(fileName, ext)
	if ext == ".yml" || ext == ".json" {
		conf, err := ctx.ReadConfig(path, nil)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		err = svc.processConfig(ctx, conf, itemName, itemType, modDir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		if ((itemType == BLOCK_REG) || (itemType == FORM_REG)) && ext == ".xml" {
			cont, err := ioutil.ReadFile(path)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			proccont,err := utils.ProcessTemplate(ctx, cont, nil)
			if err != nil {
				log.Error(ctx, "Template not correct", "Content", string(cont), "Path",path) 
				return errors.WrapError(ctx, err)
			}

			buf := bytes.NewBuffer(proccont)
			dec := xml.NewDecoder(buf)
			var n Node
			err = dec.Decode(&n)
			if err != nil {
				log.Error(ctx, "Xml Error ", "err", err, "xml", string(cont))
				return errors.WrapError(ctx, err)
			}
			if itemType == BLOCK_REG {
				err = svc.createXMLBlock(ctx, itemType, itemName, n)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			} else {
				err = svc.createXMLForm(ctx, itemType, itemName, n)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
	}
	return nil
}

func (svc *UI) readRegistry(ctx core.ServerContext, modName string, mod core.Module, modConf config.Config, moddir string) error {

	processRegistryConfig := func(ctx core.ServerContext, registry config.Config) error {
		confs := registry.AllConfigurations(ctx)
		for _, itemType := range confs {
			items, ok := registry.GetSubConfig(ctx, itemType)
			if ok {
				err := svc.processMutipleItems(ctx, items, itemType, moddir)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
		return nil
	}

	ui, ok := modConf.GetSubConfig(ctx, "ui")
	if ok {
		registry, ok := ui.GetSubConfig(ctx, "registry")
		if ok {
			err := processRegistryConfig(ctx, registry)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}

	if mod != nil {
		uiplugin, ok := mod.(UIPlugin)
		if ok {
			registry := uiplugin.GetRegistry(ctx)
			if registry != nil {
				err := processRegistryConfig(ctx, registry)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
	}

	uiRegDir := path.Join(moddir, UI_DIR, REG_DIR)
	ok, _, _ = utils.FileExists(uiRegDir)
	log.Error(ctx, "Reading registry ", "regDir", uiRegDir)
	if ok {
		files, err := ioutil.ReadDir(uiRegDir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		for _, info := range files {
			path := filepath.Join(uiRegDir, info.Name())
			if info.IsDir() {
				err = svc.processItemDir(ctx, modName, path, info.Name(), moddir)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				continue
			}
			fileName := info.Name()
			ext := filepath.Ext(fileName)
			itemType := strings.TrimSuffix(fileName, ext)
			if ext == ".yml" || ext == ".json" {
				items, err := ctx.ReadConfig(path, nil)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				err = svc.processMutipleItems(ctx, items, itemType, moddir)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
	}
	return nil
}

func (svc *UI) processMutipleItems(ctx core.ServerContext, conf config.Config, itemType string, modDir string) error {
	itemNames := conf.AllConfigurations(ctx)
	for _, item := range itemNames {
		itemVal, ok := conf.GetSubConfig(ctx, item)
		if ok {
			err := svc.processConfig(ctx, itemVal, item, itemType, modDir)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			log.Error(ctx, "Invalid registry item", "item", item)
		}
	}
	return nil
}

func (svc *UI) getRegistryItemName(ctx core.ServerContext, conf config.Config, itemName string) string {
	if conf != nil {
		name, ok := conf.GetString(ctx, "name")
		if ok {
			return name
		}
		name, ok = conf.GetString(ctx, "Name")
		if ok {
			return name
		}
	}
	return itemName
}

func (svc *UI) processConfig(ctx core.ServerContext, conf config.Config, itemName, itemType, modDir string) error {
	ctx = ctx.SubContext("Registry Item: " + itemName)
	itemApp, ok := conf.GetString(ctx, CONF_APPLICATION)
	if ok {
		if itemApp != svc.application {
			return nil
		}
	}
	itemCfg, ok := conf.GetSubConfig(ctx, ITEM_CONFIG)
	if ok {
		skipItem, _ := itemCfg.GetBool(ctx, ITEM_SKIP)
		if skipItem {
			return nil
		}
	}
	log.Trace(ctx, "Processing config", "itemType", itemType)
	switch itemType {
	case BLOCK_REG:
		{
			err := svc.createConfBlock(ctx, itemType, itemName, conf)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	case FORM_REG:
		{
			err := svc.createForm(ctx, itemType, itemName, conf)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	case DEPENDENCIES_REG:
		{
			err := svc.processDependencies(ctx, itemType, itemName, conf, modDir)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	default:
		{
			strVal, err := json.Marshal(conf)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			svc.addRegItem(ctx, itemType, svc.getRegistryItemName(ctx, conf, itemName), string(strVal))
		}
	}
	return nil
}
