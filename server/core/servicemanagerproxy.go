package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type serviceManagerProxy struct {
	manager *serviceManager
}

func (sm *serviceManagerProxy) GetService(ctx core.ServerContext, serviceName string) (server.Service, error) {
	elem, ok := sm.manager.servicesStore[serviceName]
	if ok {
		return elem, nil
	}
	return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Service Alias", serviceName)
}

func (proxy *serviceManagerProxy) Reference() core.ServerElement {
	return &serviceManagerProxy{manager: proxy.manager}
}
func (proxy *serviceManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *serviceManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *serviceManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementServiceManager
}
