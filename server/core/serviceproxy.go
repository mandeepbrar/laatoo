package core

import (
	"io"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type serviceProxy struct {
	svc *serverService
}

func (svc *serviceProxy) Service() core.Service {
	return svc.svc.service
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

func (svc *serviceProxy) HandleStreamedRequest(ctx core.RequestContext, vals map[string]interface{}, body io.ReadCloser) (*core.Response, error) {
	return svc.svc.handleStreamedRequest(ctx, vals, body)
}

func (svc *serviceProxy) HandleEncodedRequest(ctx core.RequestContext, vals map[string]interface{}, body []byte) (*core.Response, error) {
	return svc.svc.handleEncodedRequest(ctx, vals, body)
}

func (svc *serviceProxy) HandleRequest(ctx core.RequestContext, vals map[string]interface{}, body interface{}) (*core.Response, error) {
	return svc.svc.handleRequest(ctx, vals, body)
}
