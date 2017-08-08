package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

func newServiceImpl() *serviceImpl {
	info := &serviceInfo{
		request:  &requestInfo{params: make(map[string]core.Param)},
		response: &responseInfo{}}
	return &serviceImpl{svcInfo: info, state: Created, configurableObject: newConfigurableObject()}
}

type serviceImpl struct {
	*configurableObject
	svcInfo *serviceInfo
	state   State
}

func (impl *serviceImpl) Info() core.ServiceInfo {
	return impl.svcInfo
}

func (impl *serviceImpl) Initialize(ctx core.ServerContext) error {
	return nil
}

func (impl *serviceImpl) Start(ctx core.ServerContext) error {
	return nil
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
	impl.svcInfo.svcsToInject = services
}

func (impl *serviceImpl) AddParams(ctx core.ServerContext, params map[string]string) {
	for name, typ := range params {
		impl.svcInfo.request.params[name] = &param{name, typ, false, nil}
	}
}

func (impl *serviceImpl) AddParam(ctx core.ServerContext, name string, datatype string, collection bool) {
	impl.svcInfo.request.params[name] = &param{name, datatype, collection, nil}
}

func (impl *serviceImpl) AddCollectionParams(ctx core.ServerContext, params map[string]string) {
	for name, typ := range params {
		impl.svcInfo.request.params[name] = &param{name, typ, true, nil}
	}
}

func (impl *serviceImpl) AddStringParams(ctx core.ServerContext, params []string, defaultValues []string) {
	for index, name := range params {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
		}
		impl.svcInfo.request.params[name] = &param{name, "", true, defaultValue}
	}
}

func (impl *serviceImpl) AddStringParam(ctx core.ServerContext, name string) {
	impl.AddParam(ctx, name, config.CONF_OBJECT_STRING, false)
}

func (impl *serviceImpl) SetRequestType(ctx core.ServerContext, datatype string, collection bool, stream bool) {
	impl.svcInfo.request.dataType = datatype
	impl.svcInfo.request.isCollection = collection
	impl.svcInfo.request.streaming = stream
}

func (impl *serviceImpl) SetResponseType(ctx core.ServerContext, stream bool) {
	impl.svcInfo.response.streaming = stream
}

func (impl *serviceImpl) SetComponent(ctx core.ServerContext, component bool) {
	impl.svcInfo.component = component
}

func (impl *serviceImpl) SetDescription(ctx core.ServerContext, description string) {
	impl.svcInfo.description = description
}
