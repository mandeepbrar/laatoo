package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

func newServiceImpl() *serviceImpl {
	info := &serviceInfo{
		request:        &requestInfo{params: make(map[string]core.Param)},
		response:       &responseInfo{},
		configurations: make(map[string]interface{})}
	return &serviceImpl{svcInfo: info}
}

type serviceImpl struct {
	svcInfo *serviceInfo
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

func (impl *serviceImpl) Invoke(core.RequestContext, core.Request) (*core.Response, error) {
	return core.StatusSuccessResponse, nil
}

func (impl *serviceImpl) ConfigureService(requestType string, collection bool, stream bool, params []string, config []string, description string) {
	impl.SetRequestType(requestType, collection, stream)
	impl.AddStringParams(params, nil)
	impl.AddStringConfigurations(config, nil)
	impl.SetDescription(description)
}

func (impl *serviceImpl) InjectServices(services map[string]string) {
	impl.svcInfo.svcsToInject = services
}

func (impl *serviceImpl) AddParams(params map[string]string) {
	for name, typ := range params {
		impl.svcInfo.request.params[name] = &param{name, typ, false, nil}
	}
}

func (impl *serviceImpl) AddParam(name string, datatype string, collection bool) {
	impl.svcInfo.request.params[name] = &param{name, datatype, collection, nil}
}

func (impl *serviceImpl) AddCollectionParams(params map[string]string) {
	for name, typ := range params {
		impl.svcInfo.request.params[name] = &param{name, typ, true, nil}
	}
}

func (impl *serviceImpl) AddStringParams(params []string, defaultValues []string) {
	for index, name := range params {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
		}
		impl.svcInfo.request.params[name] = &param{name, "", true, defaultValue}
	}
}

func (impl *serviceImpl) AddStringParam(name string) {
	impl.AddParam(name, config.CONF_OBJECT_STRING, false)
}

func (impl *serviceImpl) AddStringConfigurations(configs []string, defaultValues []string) {
	for index, name := range configs {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
		}
		required := true
		if defaultValues != nil {
			required = false
		}
		impl.svcInfo.configurations[name] = &configuration{name, config.CONF_OBJECT_STRING, required, "", defaultValue}
	}
}

func (impl *serviceImpl) AddStringConfiguration(name string) {
	impl.AddConfigurations(map[string]string{name: config.CONF_OBJECT_STRING})
}

func (impl *serviceImpl) AddConfigurations(configs map[string]string) {
	for name, typ := range configs {
		impl.svcInfo.configurations[name] = &configuration{name, typ, true, nil, nil}
	}
}

func (impl *serviceImpl) AddOptionalConfigurations(configs map[string]string, defaultValues map[string]interface{}) {
	for name, typ := range configs {
		var defaultValue interface{}
		if defaultValues != nil {
			defaultValue = defaultValues[name]
		}
		impl.svcInfo.configurations[name] = &configuration{name, typ, false, nil, defaultValue}
	}
}

func (impl *serviceImpl) SetRequestType(datatype string, collection bool, stream bool) {
	impl.svcInfo.request.dataType = datatype
	impl.svcInfo.request.isCollection = collection
	impl.svcInfo.request.streaming = stream
}

func (impl *serviceImpl) SetResponseType(stream bool) {
	impl.svcInfo.response.streaming = stream
}

func (impl *serviceImpl) GetConfiguration(name string) (interface{}, bool) {
	c, ok := impl.svcInfo.configurations[name]
	if !ok {
		return nil, false
	}
	conf := c.(*configuration)
	if conf.value != nil {
		return conf.value, true
	}
	return conf.defaultValue, false
}

func (impl *serviceImpl) GetStringConfiguration(name string) (string, bool) {
	c, ok := impl.GetConfiguration(name)
	if !ok {
		if c != nil {
			return c.(string), ok
		} else {
			return "", ok
		}
	}
	return c.(string), ok
}

func (impl *serviceImpl) GetBoolConfiguration(name string) (bool, bool) {
	c, ok := impl.GetConfiguration(name)
	if !ok {
		if c != nil {
			return c.(bool), ok
		} else {
			return false, ok
		}
	}
	return c.(bool), ok
}

func (impl *serviceImpl) GetMapConfiguration(name string) (config.Config, bool) {
	c, ok := impl.GetConfiguration(name)
	if !ok {
		if c != nil {
			return c.(config.Config), ok
		} else {
			return nil, ok
		}
	}
	return c.(config.Config), ok
}

func (impl *serviceImpl) SetComponent(component bool) {
	impl.svcInfo.component = component
}

func (impl *serviceImpl) SetDescription(description string) {
	impl.svcInfo.description = description
}

/*
type ServiceParams map[string]*Param
type ServiceConfigurations map[string]*Configuration

func (paramsMap ServiceParams) AddParam(name string, val interface{}, typ string, collection bool) {
	paramsMap[name] = &Param{name, val, typ, collection}
}

func (configs ServiceConfigurations) AddConfiguration(name string, val interface{}, typ string, required bool) {
	configs[name] = &Configuration{name, val, typ, required}
}

func BuildParams(names []string) ServiceParams {
	smap := make(ServiceParams)
	if names != nil {
		for _, name := range names {
			smap.AddParam(name, "", "", false)
		}
	}
	return smap
}

func BuildConfigurations(configs []string) ServiceConfigurations {
	cmap := make(ServiceConfigurations)
	if configs != nil {
		for _, name := range configs {
			cmap.AddConfiguration(name, "", "", true)
		}
	}
	return cmap
}

func BuildRequestInfo(bodytype string, params []string) *RequestInfo {
	return &RequestInfo{DataType: bodytype, Params: BuildParams(params)}
}

func BuildServiceInfo(bodytype string, params []string, configs []string) *ServiceInfo {
	return &ServiceInfo{Request: BuildRequestInfo(bodytype, params), Configurations: BuildConfigurations(configs)}
}

type Request interface {
	GetBody() interface{}
	SetBody(interface{})
	GetParams() ServiceParams
	SetParams(ServiceParams)
	GetParam(string) (*Param, bool)
	AddParam(name string, val string)
	AddObjectParam(name string, val interface{}, typ string, collection bool)
}
*/
