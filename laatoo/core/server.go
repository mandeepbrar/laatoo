package laatoo

import (
	"github.com/spf13/viper"
)

//Environment hosting an application
type Server struct {
	Name         string
	Applications []*Environment
	Config       *viper.Viper
}
