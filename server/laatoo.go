package main

import (
	"laatoo/server/core"
	"log"
	"os"
)

func main() {
	//arg 1 if config is provided as argument
	configDir := os.Getenv("LAATOO_CONFIGDIR")
	if len(configDir) == 0 {
		configDir = "/etc/laatoo"
	}
	log.Println("Main Server Init")
	//create a server with config name
	err := core.Main(configDir)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
