package main

import (
	_ "entities/publish_prod"
	_ "laatooauthentication"
	"laatoocore"
	_ "laatoodata"
	_ "laatooentities"
	_ "laatoofiles"
	_ "laatoopages"
	"laatoosdk/log"
	_ "laatoostatic"
	_ "laatooview"
)

func start(configName string, serverType string) {
	//create a server with config name
	_, err := laatoocore.NewServer(configName, serverType)
	if err != nil {
		log.Logger.Error("Error in server %s", err)
	}
}
