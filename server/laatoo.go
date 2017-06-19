package main

import (
	"laatoo/server/core"
	"log"
	"os"
)

func main() {
	//arg 1 if config is provided as argument
	configName := os.Getenv("LAATOO_CONFIG")
	if len(configName) == 0 {
		configName = "/etc/laatoo/server.json"
	}
	log.Println("Main Server Init")
	//create a server with config name
	err := core.Main(configName)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
