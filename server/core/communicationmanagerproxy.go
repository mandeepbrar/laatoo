package core

import (
	"laatoo/sdk/server/core"
)

type communicatorProxy struct {
	manager *communicator
}

func (comm *communicatorProxy) SendCommunication(ctx core.RequestContext, communication map[interface{}]interface{}) error {
	return comm.manager.sendCommunication(ctx, communication)
}

func (proxy *communicatorProxy) Reference() core.ServerElement {
	return &communicatorProxy{manager: proxy.manager}
}
func (proxy *communicatorProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *communicatorProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *communicatorProxy) GetType() core.ServerElementType {
	return core.ServerElementCommunicator
}
func (proxy *communicatorProxy) GetContext() core.ServerContext {
	return proxy.manager.svrContext
}
