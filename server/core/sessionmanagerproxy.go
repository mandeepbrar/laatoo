package core

import (
	"laatoo/sdk/server/core"
)

type sessionManagerProxy struct {
	manager *sessionManager
}

func (proxy *sessionManagerProxy) GetSession(ctx core.ServerContext, sessionId string) (core.Session, error) {
	return proxy.manager.getSession(ctx, sessionId)
}

func (proxy *sessionManagerProxy) GetUserSession(ctx core.ServerContext, userId string) (core.Session, error) {
	return proxy.manager.getUserSession(ctx, userId)
}

func (proxy *sessionManagerProxy) Broadcast(ctx core.ServerContext, messageFunc func(core.ServerContext, core.Session) error) error {
	return proxy.manager.broadcast(ctx, messageFunc)
}

func (proxy *sessionManagerProxy) Reference() core.ServerElement {
	return &sessionManagerProxy{manager: proxy.manager}
}
func (proxy *sessionManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *sessionManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *sessionManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementSessionManager
}
func (proxy *sessionManagerProxy) GetContext() core.ServerContext {
	return proxy.manager.svrContext
}
