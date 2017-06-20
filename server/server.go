package main

import (
	"laatoo/sdk"
	_ "laatoo/sdk/includes/security"
	"log"
)

func main() {
	//arg 1 if config is provided as argument
	configName := "/etc/laatoo/server.json"
	log.Println("Main Server Init")
	//create a server with config name
	err := sdk.Server(configName)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
