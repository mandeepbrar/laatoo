package core

import "laatoo/sdk/server/core"

type serviceFactoryProxy struct {
	fac *serviceFactory
}

func (proxy *serviceFactoryProxy) Factory() core.ServiceFactory {
	return proxy.fac.factory
}

func (proxy *serviceFactoryProxy) Reference() core.ServerElement {
	return &serviceFactoryProxy{fac: proxy.fac}
}
func (proxy *serviceFactoryProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *serviceFactoryProxy) GetName() string {
	return proxy.fac.name
}
func (proxy *serviceFactoryProxy) GetType() core.ServerElementType {
	return core.ServerElementServiceFactory
}
func (proxy *serviceFactoryProxy) GetContext() core.ServerContext {
	return proxy.fac.svrContext
}
