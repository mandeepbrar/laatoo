package config

import (
	"github.com/spf13/viper"
)

type ConfigImpl struct {
	*viper.Viper
}

type Config interface {
	GetConfig(configurationName string) string
	SetConfig(configurationName string, configurationValue string)
}
