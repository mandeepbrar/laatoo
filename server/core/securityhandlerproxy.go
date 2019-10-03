package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type securityHandlerProxy struct {
	secHandler *securityHandler
}

func (proxy *securityHandlerProxy) Reference() core.ServerElement {
	return &securityHandlerProxy{secHandler: proxy.secHandler}
}

func (proxy *securityHandlerProxy) GetProperty(name string) interface{} {
	switch name {
	case config.REALM:
		return proxy.secHandler.realm
	case config.ADMINROLE:
		return proxy.secHandler.adminRole
	case config.ROLE:
		return proxy.secHandler.roleObject
	case config.USER:
		return proxy.secHandler.userObject
	case config.AUTHHEADER:
		return proxy.secHandler.authHeader
	}
	return nil
}
func (proxy *securityHandlerProxy) GetName() string {
	return proxy.secHandler.name
}
func (proxy *securityHandlerProxy) GetType() core.ServerElementType {
	return core.ServerElementSecurityHandler
}
func (proxy *securityHandlerProxy) GetContext() core.ServerContext {
	return proxy.secHandler.svrContext
}

func (sh *securityHandlerProxy) AuthenticateRequest(ctx core.RequestContext, loadFresh bool) (string, error) {
	return sh.secHandler.authenticateRequest(ctx, loadFresh)
}

func (sh *securityHandlerProxy) HasPermission(ctx core.RequestContext, perm string) bool {
	return sh.secHandler.hasPermission(ctx, perm)
}

func (sh *securityHandlerProxy) AllPermissions(ctx core.RequestContext) []string {
	return sh.secHandler.allPermissions(ctx)
}
