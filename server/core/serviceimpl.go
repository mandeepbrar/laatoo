package core

import "laatoo/sdk/core"

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
	impl.AddStringParams(params)
	impl.AddStringConfigurations(config)
	impl.SetDescription(description)
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

func (impl *serviceImpl) AddStringParams(params []string) {
	for _, name := range params {
		impl.svcInfo.request.params[name] = &param{name, "", true, nil}
	}
}

func (impl *serviceImpl) AddStringConfigurations(configs []string) {
	for _, name := range configs {
		impl.svcInfo.configurations[name] = &configuration{name, "", true, nil}
	}
}

func (impl *serviceImpl) AddConfigurations(configs map[string]string) {
	for name, typ := range configs {
		impl.svcInfo.configurations[name] = &configuration{name, typ, true, nil}
	}
}
func (impl *serviceImpl) AddOptionalConfigurations(configs map[string]string) {
	for name, typ := range configs {
		impl.svcInfo.configurations[name] = &configuration{name, typ, false, nil}
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

func (impl *serviceImpl) GetConfiguration(name string) interface{} {
	conf, ok := impl.svcInfo.configurations[name]
	if !ok {
		return nil
	}
	return conf.(*configuration).value
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
