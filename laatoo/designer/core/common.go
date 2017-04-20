package core

import (
	"laatoo/sdk"
	_ "laatoo/sdk/includes/cache"
	_ "laatoo/sdk/includes/data"
	_ "laatoo/sdk/includes/middleware"
	_ "laatoo/sdk/includes/pubsub"
	//	"laatoo/sdk/log"
	_ "laatoo/designer/core/entities"
	_ "laatoo/sdk/includes/security"
	_ "laatoo/sdk/includes/static"
	_ "laatoo/sdk/includes/storage"
	"log"
	//_ "laatoo/designer/core/services"
)

func Start(configName string) {
	//create a server with config name
	err := sdk.Server("server.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
