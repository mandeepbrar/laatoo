package main

import (
	_ "laatooauthentication"
	"laatoocore"
	_ "laatoodata"
	_ "laatooentities"
	_ "laatoofiles"
	_ "laatoopages"
	_ "laatoopubsub"
	"laatoosdk/log"
	_ "laatoostatic"
	_ "laatooview"
	_ "mehfilpublish/publish_prod/entities"
)

func start(configName string, serverType string) {
	//create a server with config name
	_, err := laatoocore.NewServer(configName, serverType)
	if err != nil {
		log.Logger.Error("Error in server %s", err)
	}
}
