package common

import (
	"fmt"
	"laatoo/sdk/config"
)

func ConfigFileAdapter(conf config.Config, configName string) (config.Config, error) {
	var configToRet config.Config
	var err error
	confFileName, ok := conf.GetString(configName)
	if ok {
		configToRet, err = config.NewConfigFromFile(confFileName)
		if err != nil {
			return nil, fmt.Errorf("Could not read from file %s. Error:%s", confFileName, err)
		}
	} else {
		configToRet, ok = conf.GetSubConfig(configName)
		if !ok {
			return nil, fmt.Errorf("Could not find config: %s", configName)
		}
	}
	return configToRet, nil
}
