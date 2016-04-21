package service

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type service struct {
	*common.Context
	name    string
	service core.Service
	factory server.Factory
	conf    config.Config
	owner   *serviceManager
}

func (svc *service) Service() core.Service {
	return svc.service
}
