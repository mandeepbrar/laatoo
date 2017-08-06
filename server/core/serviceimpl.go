package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/server/common"
	"strings"
)

type ServiceState int

const (
	Created ServiceState = iota
	Initialized
	Started
)

func newServiceImpl() *serviceImpl {
	info := &serviceInfo{
		request:        &requestInfo{params: make(map[string]core.Param)},
		response:       &responseInfo{},
		configurations: make(map[string]interface{})}
	return &serviceImpl{svcInfo: info, state: Created}
}

type serviceImpl struct {
	svcInfo *serviceInfo
	state   ServiceState
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

func (impl *serviceImpl) AddStringConfigurations(ctx core.ServerContext, configs []string, defaultValues []string) {
	required := true
	if defaultValues != nil {
		required = false
	}
	for index, name := range configs {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
		}
		impl.svcInfo.configurations[name] = &configuration{name, config.CONF_OBJECT_STRING, required, nil, defaultValue}
	}
}

func (impl *serviceImpl) AddStringConfiguration(ctx core.ServerContext, name string) {
	impl.AddConfigurations(ctx, map[string]string{name: config.CONF_OBJECT_STRING})
}

func (impl *serviceImpl) AddConfigurations(ctx core.ServerContext, configs map[string]string) {
	for name, typ := range configs {
		impl.svcInfo.configurations[name] = &configuration{name, typ, true, nil, nil}
	}
}

func (impl *serviceImpl) AddOptionalConfigurations(ctx core.ServerContext, configs map[string]string, defaultValues map[string]interface{}) {
	for name, typ := range configs {
		var defaultValue interface{}
		if defaultValues != nil {
			defaultValue = defaultValues[name]
		}
		impl.svcInfo.configurations[name] = &configuration{name, typ, false, nil, defaultValue}
	}
}

func (impl *serviceImpl) SetRequestType(ctx core.ServerContext, datatype string, collection bool, stream bool) {
	impl.svcInfo.request.dataType = datatype
	impl.svcInfo.request.isCollection = collection
	impl.svcInfo.request.streaming = stream
}

func (impl *serviceImpl) SetResponseType(ctx core.ServerContext, stream bool) {
	impl.svcInfo.response.streaming = stream
}

func (impl *serviceImpl) lookupContext(ctx core.ServerContext, name string) (interface{}, bool) {
	val, found := ctx.Get(name)
	if !found {
		val, found = ctx.GetVariable(name)
		if found {
			return val, found
		} else {
			return nil, false
		}
	} else {
		return val, found
	}
}

func (impl *serviceImpl) GetConfiguration(ctx core.ServerContext, name string) (interface{}, bool) {
	var val interface{}
	c, found := impl.svcInfo.configurations[name]
	if !found {
		val, found = impl.lookupContext(ctx, name)
	} else {
		conf := c.(*configuration)
		if conf.value != nil {
			val = conf.value
			valStr, ok := val.(string)
			if ok && strings.HasPrefix(valStr, ":") {
				val, found = impl.lookupContext(ctx, valStr[1:])
			}
		} else {
			val = conf.defaultValue
			found = false
		}
	}
	return val, found
}

func (impl *serviceImpl) GetStringConfiguration(ctx core.ServerContext, name string) (string, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return "", ok
	}
	return c.(string), ok
}

func (impl *serviceImpl) GetBoolConfiguration(ctx core.ServerContext, name string) (bool, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return false, ok
	}
	return c.(bool), ok
}

func (impl *serviceImpl) GetMapConfiguration(ctx core.ServerContext, name string) (config.Config, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return nil, ok
	}
	m, mok := c.(map[string]interface{})
	if mok {
		return common.GenericConfig(m), true
	} else {
		return nil, false
	}
}

func (impl *serviceImpl) SetComponent(ctx core.ServerContext, component bool) {
	impl.svcInfo.component = component
}

func (impl *serviceImpl) SetDescription(ctx core.ServerContext, description string) {
	impl.svcInfo.description = description
}
