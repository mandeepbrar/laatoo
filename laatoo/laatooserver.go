package main

import (
	"laatoocore"
	_ "laatoodesigner"
	"laatoosdk/log"
	"os"
)

const (
	CONF_SERVERTYPE_HOSTNAME = "STANDALONE"
)

func main() {
	//arg 1 if config is provided as argument
	var configName string
	if len(os.Args) > 1 {
		configName = os.Args[1]
	} else {
		configName = "server"
	}
	//create a server with config name
	_, err := laatoocore.NewServer(configName, laatoocore.CONF_SERVERTYPE_STANDALONE)
	if err != nil {
		log.Logger.Error("Error in server %s", err)
	}
}
