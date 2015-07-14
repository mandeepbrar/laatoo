package laatoo

import (
	"github.com/spf13/viper"
)

//Environment hosting an application
type Service struct {
	Name   string
	Config *viper.Viper
}
