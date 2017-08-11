package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

func newServiceImpl() *serviceImpl {
	return &serviceImpl{serviceInfo: newServiceInfo("", nil, false, nil), state: Created}
}

type serviceImpl struct {
	*serviceInfo
	state State
}

func (impl *serviceImpl) setServiceInfo(si *serviceInfo) {
	impl.serviceInfo = si
}

func (impl *serviceImpl) info() core.ServiceInfo {
	return impl.serviceInfo
}

func (impl *serviceImpl) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (impl *serviceImpl) Start(ctx core.ServerContext) error {
	return nil
}

func (impl *serviceImpl) Describe(ctx core.ServerContext) {
}

func (impl *serviceImpl) Invoke(core.RequestContext) error {
	return nil
}

func (impl *serviceImpl) ConfigureService(ctx core.ServerContext, requestType string, collection bool, stream bool, params []string, config []string, description string) {
	impl.SetRequestType(ctx, requestType, collection, stream)
	impl.AddStringParams(ctx, params, nil)
	impl.AddStringConfigurations(ctx, config, nil)
	impl.SetDescription(ctx, description)
}

func (impl *serviceImpl) InjectServices(ctx core.ServerContext, services map[string]string) {
	impl.serviceInfo.svcsToInject = services
}

func (impl *serviceImpl) AddParams(ctx core.ServerContext, params map[string]string) {
	for name, typ := range params {
		impl.serviceInfo.request.params[name] = &param{name, typ, false, nil}
	}
}

func (impl *serviceImpl) AddParam(ctx core.ServerContext, name string, datatype string, collection bool) {
	impl.serviceInfo.request.params[name] = &param{name, datatype, collection, nil}
}

func (impl *serviceImpl) AddCollectionParams(ctx core.ServerContext, params map[string]string) {
	for name, typ := range params {
		impl.serviceInfo.request.params[name] = &param{name, typ, true, nil}
	}
}

func (impl *serviceImpl) AddStringParams(ctx core.ServerContext, params []string, defaultValues []string) {
	for index, name := range params {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
		}
		impl.serviceInfo.request.params[name] = &param{name, "", true, defaultValue}
	}
}

func (impl *serviceImpl) AddStringParam(ctx core.ServerContext, name string) {
	impl.AddParam(ctx, name, config.CONF_OBJECT_STRING, false)
}

func (impl *serviceImpl) SetRequestType(ctx core.ServerContext, datatype string, collection bool, stream bool) {
	impl.serviceInfo.request.dataType = datatype
	impl.serviceInfo.request.isCollection = collection
	impl.serviceInfo.request.streaming = stream
}

func (impl *serviceImpl) SetResponseType(ctx core.ServerContext, stream bool) {
	impl.serviceInfo.response.streaming = stream
}

func (impl *serviceImpl) SetComponent(ctx core.ServerContext, component bool) {
	impl.serviceInfo.component = component
}

func (impl *serviceImpl) SetDescription(ctx core.ServerContext, description string) {
	impl.serviceInfo.description = description
}
