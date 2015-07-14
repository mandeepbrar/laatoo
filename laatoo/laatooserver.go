package main

import (
	"laatoo/config"
	"laatoo/log"
	//"net/http"
)

func InitializeServices(conf config.Config) {
	//config logger
	log.ConfigLogger(conf)

}
