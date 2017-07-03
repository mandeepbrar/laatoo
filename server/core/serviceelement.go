package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type serviceProxy struct {
	svc *service
}

func (svc *serviceProxy) Service() core.Service {
	return svc.svc.service
}

func (proxy *serviceProxy) Reference() core.ServerElement {
	return &serviceProxy{svc: proxy.svc}
}
func (proxy *serviceProxy) ParamsConfig() config.Config {
	return proxy.svc.paramsConf
}
func (proxy *serviceProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *serviceProxy) GetName() string {
	return proxy.svc.name
}
func (proxy *serviceProxy) GetType() core.ServerElementType {
	return core.ServerElementService
}

func (svc *serviceProxy) Invoke(ctx core.RequestContext) error {
	return svc.svc.invoke(ctx)
}
