package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
	"laatoo/server/common"
)

type factoryManagerProxy struct {
	manager *factoryManager
}

func (fm *factoryManagerProxy) GetFactory(ctx core.ServerContext, factoryName string) (server.Factory, error) {
	factoryName = common.FillVariables(ctx, factoryName)
	elem, ok := fm.manager.serviceFactoryStore[factoryName]
	if ok {
		return elem, nil
	}
	return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Factory Name", factoryName)
}
func (proxy *factoryManagerProxy) Reference() core.ServerElement {
	return &factoryManagerProxy{manager: proxy.manager}
}
func (proxy *factoryManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *factoryManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *factoryManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementFactoryManager
}
