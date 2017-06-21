package common

import (
	"fmt"
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"path"
)

const (
	CONF_DEFAULTFACTORY_NAME       = "__defaultfactory__"
	CONF_DEFAULTMETHODFACTORY_NAME = "__defaultmethodfactory__"
)

func ConfigFileAdapter(ctx core.ServerContext, conf config.Config, configName string) (config.Config, error, bool) {
	confFileName, ok := conf.GetString(configName)
	if ok {
		log.Debug(ctx, "Reading config from file "+confFileName)
		return FileAdapter(conf, configName)
	} else {
		return nil, nil, false
	}
}

func ProcessDirectoryFiles(ctx core.ServerContext, subdir string, processor func(core.ServerContext, config.Config, string) error) error {
	baseDir, _ := ctx.GetString(constants.CONF_BASE_DIR)
	subDir := path.Join(baseDir, subdir)
	ok, _, _ := utils.FileExists(subDir)
	if ok {
		files, err := ioutil.ReadDir(subDir)
		if err != nil {
			return errors.WrapError(ctx, err, "Subdirectory", subDir)
		}
		for _, info := range files {
			if !info.IsDir() {
				elemName := info.Name()
				confFile := path.Join(subDir, elemName)
				elemConf, err := NewConfigFromFile(confFile)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				name, ok := elemConf.GetString(constants.CONF_OBJECT_NAME)
				if ok {
					elemName = name
				}
				elemCtx := ctx.SubContext(elemName)
				if err := processor(elemCtx, elemConf, elemName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func FileAdapter(conf config.Config, configName string) (config.Config, error, bool) {
	var configToRet config.Config
	var err error
	confFileName, ok := conf.GetString(configName)
	if ok {
		configToRet, err = NewConfigFromFile(confFileName)
		if err != nil {
			return nil, fmt.Errorf("Could not read from file %s. Error:%s", confFileName, err), true
		}
	} else {
		configToRet, ok = conf.GetSubConfig(configName)
		if !ok {
			return nil, nil, false
		}
	}
	return configToRet, nil, true
}

func Cast(conf interface{}) (config.Config, bool) {
	var gc GenericConfig
	cf, ok := conf.(map[string]interface{})
	if ok {
		gc = cf
		return gc, true
	}
	return nil, false
}

func Merge(conf1 config.Config, conf2 config.Config) config.Config {
	mergedConf := make(GenericConfig)
	copyConfs := func(conf config.Config) {
		if conf == nil {
			return
		}
		confNames := conf.AllConfigurations()
		for _, confName := range confNames {
			val, _ := conf.Get(confName)
			subConf, ok := val.(config.Config)
			if ok {
				existingVal, eok := mergedConf[confName]
				if eok {
					existingConf, cok := existingVal.(config.Config)
					if cok {
						mergedSubConf := Merge(existingConf, subConf)
						mergedConf[confName] = mergedSubConf
					} else {
						mergedConf[confName] = val
					}
				} else {
					mergedConf[confName] = val
				}
			} else {
				mergedConf[confName] = val
			}
		}
	}
	copyConfs(conf1)
	copyConfs(conf2)
	return mergedConf
}
