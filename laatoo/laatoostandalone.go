// +build !appengine

package main

import (
	"laatoo/config"
	//"laatoo/log"
	"net/http"
)

//Read Standalone server configurations
func readConfig() config.Config {
	return nil
}

func main() {
	//read config for standalone
	conf := readConfig()
	address := conf.GetConfig(config.HOSTNAME)
	//initialize all services with config
	InitializeServices(conf)
	//start listening standalone
	http.ListenAndServe(address, nil)
}
