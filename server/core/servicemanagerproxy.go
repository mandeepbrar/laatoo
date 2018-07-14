package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
)

type serviceManagerProxy struct {
	manager *serviceManager
}

func (proxy *serviceManagerProxy) GetService(ctx core.ServerContext, serviceName string) (elements.Service, error) {
	return proxy.manager.getService(ctx, serviceName)
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
