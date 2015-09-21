// +build !appengine

package main

import (
	"laatoocore"
	"log"
	"os"
)

func main() {
	//arg 1 if config is provided as argument
	var configName string
	if len(os.Args) > 1 {
		configName = os.Args[1]
	} else {
		configName = "server"
	}
	log.Println("Main Server Init")
	//create a server with config name
	start(configName, laatoocore.CONF_SERVERTYPE_STANDALONE)
}
