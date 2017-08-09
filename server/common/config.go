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
	"os"
	"path"
	"path/filepath"
)

const (
	CONF_DEFAULTFACTORY_NAME     = "__defaultfactory__"
	CONF_SERVICEAGGREGATOR_NAME  = "__serviceaggregator__"
	CONF_TRANSFORMERSERVICE_NAME = "__transformerservice__"
)

func ConfigFileAdapter(ctx core.ServerContext, conf config.Config, configName string) (config.Config, error, bool) {
	retconf, ok := conf.GetSubConfig(configName)
	if ok {
		return retconf, nil, ok
	}
	confFileName, ok := conf.GetString(configName)
	if ok {
		log.Debug(ctx, "Reading config from file "+confFileName)
		return FileAdapter(conf, configName)
	} else {
		return nil, nil, false
	}
}

func ProcessObjects(ctx core.ServerContext, objs map[string]config.Config, processor func(core.ServerContext, config.Config, string) error) error {
	for elemName, elemConf := range objs {
		elemCtx := ctx.SubContext(elemName)
		if err := processor(elemCtx, elemConf, elemName); err != nil {
			return err
		}
	}
	return nil
}

func processDirectoryFiles(ctx core.ServerContext, subDir string, objs map[string]config.Config, recurse bool) error {
	ok, fi, _ := utils.FileExists(subDir)
	if ok && fi.IsDir() {
		files, err := ioutil.ReadDir(subDir)
		if err != nil {
			return errors.WrapError(ctx, err, "Subdirectory", subDir)
		}

		for _, info := range files {
			elemfileName := info.Name()
			file := path.Join(subDir, elemfileName)
			if !info.IsDir() {
				extension := filepath.Ext(elemfileName)
				elemName := elemfileName[0 : len(elemfileName)-len(extension)]
				elemConf, err := NewConfigFromFile(file)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				if !CheckContextCondition(ctx, elemConf) {
					continue
				}
				name, ok := elemConf.GetString(constants.CONF_OBJECT_NAME)
				if ok {
					elemName = name
				}
				objs[elemName] = elemConf
				/*elemCtx := ctx.SubContext(elemName)
				if err := processor(elemCtx, elemConf, elemName); err != nil {
					return err
				}*/
				if (info.Mode() & os.ModeSymlink) != 0 {
					s, err := os.Readlink(file)
					if err == nil && recurse {
						err = processDirectoryFiles(ctx, s, objs, recurse)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
					}
				}
			} else {
				if recurse {
					err = processDirectoryFiles(ctx, file, objs, recurse)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
				}
			}
		}
	}
	return nil
}

func ProcessDirectoryFiles(ctx core.ServerContext, baseDir string, object string, recurse bool) (map[string]config.Config, error) {
	objs := make(map[string]config.Config)
	subDir := path.Join(baseDir, object)
	err := processDirectoryFiles(ctx, subDir, objs, recurse)
	if err != nil {
		return nil, err
	}
	return objs, nil
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
	var gc config.GenericConfig
	cf, ok := conf.(map[string]interface{})
	if ok {
		gc = cf
		return gc, true
	}
	return nil, false
}

func MergeConfigMaps(conf1 map[string]config.Config, conf2 map[string]config.Config) map[string]config.Config {
	res := make(map[string]config.Config)
	for k, v := range conf1 {
		res[k] = v
	}
	for k, v := range conf2 {
		res[k] = v
	}
	return res
}

func Merge(conf1 config.Config, conf2 config.Config) config.Config {
	mergedConf := make(config.GenericConfig)
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
