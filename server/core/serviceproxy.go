package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type serviceProxy struct {
	svc *serverService
}

func (svc *serviceProxy) Service() core.Service {
	return svc.svc.service
}

func (svc *serviceProxy) ServiceContext() core.ServerContext {
	return svc.svc.svrContext
}

func (proxy *serviceProxy) Reference() core.ServerElement {
	return &serviceProxy{svc: proxy.svc}
}

/*func (proxy *serviceProxy) ParamsConfig() config.Config {
	return proxy.svc.paramsConf
}*/
func (proxy *serviceProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *serviceProxy) GetName() string {
	return proxy.svc.name
}
func (proxy *serviceProxy) GetType() core.ServerElementType {
	return core.ServerElementService
}
func (proxy *serviceProxy) GetConfiguration() config.Config {
	return proxy.svc.conf
}
func (proxy *serviceProxy) GetContext() core.ServerContext {
	return proxy.svc.svrContext
}

func (svc *serviceProxy) HandleRequest(ctx core.RequestContext, vals map[string]interface{}) (*core.Response, error) {
	return svc.svc.handleRequest(ctx.(*requestContext), vals)
}
