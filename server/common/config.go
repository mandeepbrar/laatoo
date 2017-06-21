package common

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"path"
)

const (
	CONF_DEFAULTFACTORY_NAME       = "__defaultfactory__"
	CONF_DEFAULTMETHODFACTORY_NAME = "__defaultmethodfactory__"
)

func ConfigFileAdapter(ctx core.ServerContext, conf config.Config, configName string) (config.Config, error, bool) {
	confFileName, ok := conf.GetString(configName)
	if ok {
		log.Logger.Debug(ctx, "Reading config from file "+confFileName)
		return config.FileAdapter(conf, configName)
	} else {
		return nil, nil, false
	}
}

func ProcessDirectoryFiles(ctx core.ServerContext, subdir string, processor func(core.ServerContext, config.Config, string) error) error {
	baseDir, _ := ctx.GetString(config.CONF_BASE_DIR)
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
				elemConf, err := config.NewConfigFromFile(confFile)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				name, ok := elemConf.GetString(config.CONF_OBJECT_NAME)
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
