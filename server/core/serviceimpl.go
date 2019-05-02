package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

func newServiceImpl(name string) *serviceImpl {
	return &serviceImpl{serviceInfo: newServiceInfo(name, "", nil, nil, nil), state: Created}
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

func (impl *serviceImpl) Describe(ctx core.ServerContext) error {
	return nil
}

func (impl *serviceImpl) Invoke(core.RequestContext) error {
	return nil
}

func (impl *serviceImpl) Stop(ctx core.ServerContext) error {
	return nil
}
func (impl *serviceImpl) Unload(ctx core.ServerContext) error {
	return nil
}

func (impl *serviceImpl) ConfigureService(ctx core.ServerContext, params []string, config []string, description string) {
	//impl.SetRequestType(ctx, requestType, collection, stream)
	impl.AddStringParams(ctx, params, nil)
	impl.AddStringConfigurations(ctx, config, nil)
	impl.SetDescription(ctx, description)
}

func (impl *serviceImpl) InjectServices(ctx core.ServerContext, services map[string]string) {
	impl.serviceInfo.svcsToInject = services
}

func (impl *serviceImpl) AddParams(ctx core.ServerContext, params map[string]string, required bool) error {
	for name, typ := range params {
		err := impl.AddParam(ctx, name, typ, false, required, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (impl *serviceImpl) AddParam(ctx core.ServerContext, name string, datatype string, collection, required, stream bool) error {
	p, err := newParam(ctx, name, datatype, collection, stream, required)
	if err != nil {
		return err
	} else {
		impl.serviceInfo.request.params[name] = p
		return nil
	}
}

func (impl *serviceImpl) AddParamWithType(ctx core.ServerContext, name string, datatype string) error {
	return impl.AddParam(ctx, name, datatype, false, false, true)
}

func (impl *serviceImpl) AddOptionalParamWithType(ctx core.ServerContext, name string, datatype string) error {
	return impl.AddParam(ctx, name, datatype, false, false, false)
}

func (impl *serviceImpl) AddCollectionParams(ctx core.ServerContext, params map[string]string) error {
	for name, typ := range params {
		p, err := newParam(ctx, name, typ, true, false, true)
		if err != nil {
			return err
		} else {
			impl.serviceInfo.request.params[name] = p
		}
	}
	return nil
}

func (impl *serviceImpl) AddStringParams(ctx core.ServerContext, params []string, defaultValues []string) {
	required := true
	for index, name := range params {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
			required = false
		}
		p, _ := newParam(ctx, name, "", false, false, required)
		p.value = defaultValue
		impl.serviceInfo.request.params[name] = p
	}
}

func (impl *serviceImpl) AddStringParam(ctx core.ServerContext, name string) {
	impl.AddParam(ctx, name, config.OBJECTTYPE_STRING, true, false, false)
}

/*
func (impl *serviceImpl) SetRequestType(ctx core.ServerContext, datatype string, collection bool, stream bool) {
	impl.serviceInfo.request.dataType = datatype
	impl.serviceInfo.request.isCollection = collection
	impl.serviceInfo.request.streaming = stream
}*/

/*
func (impl *serviceImpl) SetResponseType(ctx core.ServerContext, stream bool) {
	impl.serviceInfo.response.streaming = stream
}*/

func (impl *serviceImpl) SetComponent(ctx core.ServerContext, component bool) {
	impl.serviceInfo.component = component
}

func (impl *serviceImpl) SetDescription(ctx core.ServerContext, description string) {
	impl.serviceInfo.description = description
}
