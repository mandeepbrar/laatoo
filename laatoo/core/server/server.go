package server

import (
	"laatoo/sdk/config"
)

//Server hosting a number of applications
type Server struct {
	//name of the server
	Name string
	//if server is standalone or google app
	ServerType string
	//all applications deployed on this server
	Applications map[string]*Application
	//config for the server
	Config config.Config
}
