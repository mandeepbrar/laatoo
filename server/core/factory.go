package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type serviceFactory struct {
	name       string
	factory    core.ServiceFactory
	conf       config.Config
	owner      *factoryManager
	svrContext *serverContext
}
