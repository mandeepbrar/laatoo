// +build appengine

package main

import (
	"laatoo/config"
)

//Read Standalone server configurations
func readConfig() config.Config {
	return nil
}

func init() {
	//read config for app
	conf := readConfig()
	//initialize all services with config
	InitializeServices(conf)
}
