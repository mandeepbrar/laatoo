package laatoo

import (
	"github.com/spf13/viper"
)

//Environment hosting an application
type Application struct {
	Environments  map[string]*Environment
	Config        *viper.Viper
	ServicesStore map[string]*Service
}
