package core

import (
	"laatoo/sdk/server/core"
)

type secretsManagerProxy struct {
	manager *secretsManager
}

func (proxy *secretsManagerProxy) Get(ctx core.ServerContext, key string) ([]byte, bool) {
	return proxy.manager.get(ctx, key)
}
func (proxy *secretsManagerProxy) Put(ctx core.ServerContext, key string, val []byte) error {
	return proxy.manager.put(ctx, key, val)
}

func (proxy *secretsManagerProxy) Reference() core.ServerElement {
	return &secretsManagerProxy{manager: proxy.manager}
}
func (proxy *secretsManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *secretsManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *secretsManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementSecretsManager
}
func (proxy *secretsManagerProxy) GetContext() core.ServerContext {
	return proxy.manager.svrContext
}
