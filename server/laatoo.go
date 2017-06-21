package main

import (
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"laatoo/server/core"
	"log"
	"os"
)

func main() {
	//arg 1 if config is provided as argument
	configDir := os.Getenv("LAATOO_CONFIGDIR")
	if len(configDir) == 0 {
		exists, _, _ := utils.FileExists(constants.CONF_CONFIG_FILE)
		if exists {
			configDir = "."
		} else {
			configDir = "/etc/laatoo"
		}
	}
	log.Println("Main Server Init from config directory ", configDir)
	//create a server with config name
	err := core.Main(configDir)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
