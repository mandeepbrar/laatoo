package factory

import (
	"laatoo/server/common"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type factoryManagerProxy struct {
	*common.Context
	manager *factoryManager
}

func (fm *factoryManagerProxy) GetFactory(ctx core.ServerContext, factoryName string) (server.Factory, error) {
	elem, ok := fm.manager.serviceFactoryStore[factoryName]
	if ok {
		return elem, nil
	}
	return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Factory Name", factoryName)
}
