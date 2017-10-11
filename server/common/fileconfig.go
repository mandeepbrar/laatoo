package common

import (
	"fmt"
	"laatoo/sdk/config"
	context "laatoo/sdk/ctx"
	"laatoo/sdk/log"

	"gopkg.in/yaml.v2"
)

//creates a new config object from file provided to it
//only works for json configs
func NewConfigFromFile(ctx context.Context, file string) (config.Config, error) {
	conf := make(GenericConfig, 50)
	fileData, err := GetTemplateFileContent(ctx, file, nil)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file %s. Error: %s", file, err.Error())
	}
	log.Info(ctx, "Config File", "file", file, "conf", string(fileData))
	if err = yaml.Unmarshal(fileData, &conf); err != nil {
		return nil, fmt.Errorf("Error parsing config file %s. Error: %s", file, err.Error())
	}
	return conf, nil
}
