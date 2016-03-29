package server

import (
	"laatoo/sdk/config"
)

//Server hosting a number of environments
type Server struct {
	//name of the server
	Name string
	//if server is standalone or google app
	ServerType string
	//all environments deployed on this server
	Applications map[string]*Environment
	//config for the server
	Config config.Config
}
