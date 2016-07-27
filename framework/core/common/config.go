package common

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

func ConfigFileAdapter(ctx core.ServerContext, conf config.Config, configName string) (config.Config, error, bool) {
	confFileName, ok := conf.GetString(configName)
	if ok {
		log.Logger.Debug(ctx, "Reading config from file "+confFileName)
	}
	return config.FileAdapter(conf, configName)
}
