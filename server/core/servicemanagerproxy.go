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

func (proxy *serviceManagerProxy) GetServiceContext(ctx core.ServerContext, serviceName string) (core.ServerContext, error) {
	svc, err := proxy.GetService(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	return svc.ServiceContext(), nil
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
func (proxy *serviceManagerProxy) GetContext() core.ServerContext {
	return proxy.manager.svrContext
}
