package core

import "laatoo/sdk/core"

type serviceInfo struct {
	*configurableObject
	request      *requestInfo
	response     *responseInfo
	description  string
	component    bool
	svcsToInject map[string]string
}

func newServiceInfo(description string, reqInfo core.RequestInfo, streamedResponse bool, configurations []core.Configuration) *serviceInfo {
	co := newConfigurableObject(description, "Service")
	co.setConfigurations(configurations)
	return &serviceInfo{configurableObject: co,
		request:      newRequestInfo("", false, false, nil),
		response:     newResponseInfo(streamedResponse),
		svcsToInject: make(map[string]string)}
}

func (svcinfo *serviceInfo) GetRequestInfo() core.RequestInfo {
	return svcinfo.request
}

func (svcinfo *serviceInfo) GetResponseInfo() core.ResponseInfo {
	return svcinfo.response
}

func (svcinfo *serviceInfo) GetDescription() string {
	return svcinfo.description
}

func (svcinfo *serviceInfo) IsComponent() bool {
	return svcinfo.component
}

func (svcinfo *serviceInfo) GetRequiredServices() map[string]string {
	return svcinfo.svcsToInject
}

type requestInfo struct {
	dataType     string
	isCollection bool
	streaming    bool
	params       map[string]core.Param
}

func newRequestInfo(requesttype string, collection bool, stream bool, params []core.Param) *requestInfo {
	reqInfo := &requestInfo{requesttype, collection, stream, nil}
	reqInfo.params = make(map[string]core.Param)
	if params != nil {
		for _, p := range params {
			reqInfo.params[p.GetName()] = p
		}
	}
	return reqInfo
}

func (ri *requestInfo) GetDataType() string {
	return ri.dataType
}
func (ri *requestInfo) IsCollection() bool {
	return ri.isCollection
}
func (ri *requestInfo) IsStream() bool {
	return ri.streaming
}
func (ri *requestInfo) GetParams() map[string]core.Param {
	return ri.params
}

type responseInfo struct {
	streaming bool
}

func newResponseInfo(streaming bool) *responseInfo {
	return &responseInfo{streaming}
}
func (ri *responseInfo) IsStream() bool {
	return ri.streaming
}

type param struct {
	name       string
	ptype      string
	collection bool
	value      interface{}
}

func newParam(name, ptype string, collection bool) *param {
	return &param{name, ptype, collection, nil}
}

func (p *param) GetName() string {
	return p.name
}
func (p *param) IsCollection() bool {
	return p.collection
}
func (p *param) GetDataType() string {
	return p.ptype
}
func (p *param) GetValue() interface{} {
	return p.value
}
