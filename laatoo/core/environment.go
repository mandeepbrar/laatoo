package laatoo

import (
	"github.com/spf13/viper"
)

//Environment hosting an application
type Environment struct {
	Application   *Application
	Config        *viper.Viper
	ServicesStore map[string]*Service
}
