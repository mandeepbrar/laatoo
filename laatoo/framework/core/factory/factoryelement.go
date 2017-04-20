package factory

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type serviceFactory struct {
	*common.Context
	name    string
	factory core.ServiceFactory
	conf    config.Config
	owner   *factoryManager
}

func (fac *serviceFactory) Factory() core.ServiceFactory {
	return fac.factory
}
