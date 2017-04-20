package service

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type serviceManagerProxy struct {
	*common.Context
	manager *serviceManager
}

func (sm *serviceManagerProxy) GetService(ctx core.ServerContext, serviceName string) (server.Service, error) {
	elem, ok := sm.manager.servicesStore[serviceName]
	if ok {
		return elem, nil
	}
	return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Service Alias", serviceName)
}
