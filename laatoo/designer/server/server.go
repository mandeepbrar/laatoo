// +build !appengine

package main

import (
	"laatoo/designer/core"
	"log"
	"os"
)

func main() {
	//arg 1 if config is provided as argument
	var configName string
	if len(os.Args) > 1 {
		configName = os.Args[1]
	} else {
		configName = "server.json"
	}
	log.Println("Main Server Init")
	//create a server with config name
	core.Start(configName)
}
