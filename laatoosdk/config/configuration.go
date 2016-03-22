package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

//Json based viper implementation of Config interface
type ConfigImpl struct {
	*viper.Viper
}

//constructs the json config object
func constructConfig() *ConfigImpl {
	conf := &ConfigImpl{viper.New()}
	//set config type
	conf.SetConfigType("json")
	return conf
}

//creates a new config object from file provided to it
//only works for json configs
func NewConfigFromFile(file string) *ConfigImpl {
	conf := constructConfig()
	//set file name to read
	conf.SetConfigFile(file)
	//read the file
	err := conf.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return conf
}

/*
//creates config from json data provided to it
func NewConfigFromJSON(json string) *ConfigImpl {
	conf := constructConfig()
	//set file name to read
	conf.SetConfigName(file)
	//read the file
	err := conf.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return conf
}
*/
//Get string configuration value
func (conf *ConfigImpl) GetString(configurationName string) string {
	val := conf.Get(configurationName)
	if val != nil {
		return val.(string)
	}
	return ""
}

//Get string configuration value
func (conf *ConfigImpl) GetBool(configurationName string) bool {
	val := conf.Get(configurationName)
	if val != nil {
		config, err := strconv.ParseBool(val.(string))
		if err != nil {
			return false
		}
		return config
	}
	return false
}

func (conf *ConfigImpl) GetArray(configurationName string) []interface{} {
	arrInt := conf.Get(configurationName)
	if arrInt == nil {
		return nil
	}
	return arrInt.([]interface{})
}

func (conf *ConfigImpl) GetMap(configurationName string) map[string]interface{} {
	mapInt := conf.Get(configurationName)
	mapVal := mapInt.(map[string]interface{})
	return mapVal
}

//Set string configuration value
func (conf *ConfigImpl) SetString(configurationName string, configurationValue string) {
	conf.Set(configurationName, configurationValue)
}

//Config Interface used by Laatoo
type Config interface {
	GetString(configurationName string) string
	GetBool(configurationName string) bool
	GetMap(configurationName string) map[string]interface{}
	GetArray(configurationName string) []interface{}
	SetString(configurationName string, configurationValue string)
}
