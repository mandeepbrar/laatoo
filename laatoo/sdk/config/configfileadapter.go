package config

import (
	"fmt"
)

func ConfigFileAdapter(conf Config, configName string) (Config, error) {
	var configToRet Config
	var err error
	confFileName, ok := conf.GetString(configName)
	if ok {
		configToRet, err = NewConfigFromFile(confFileName)
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
