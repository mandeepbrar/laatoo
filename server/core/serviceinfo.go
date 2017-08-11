package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

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
		component:    false,
		request:      newRequestInfo("", false, false, nil),
		response:     newResponseInfo(streamedResponse),
		svcsToInject: make(map[string]string)}
}

func buildServiceInfo(conf config.Config) *serviceInfo {
	comp, _ := conf.GetBool(SVCCOMP)
	return &serviceInfo{configurableObject: buildConfigurableObject(conf),
		component:    comp,
		request:      buildRequestInfo(conf),
		response:     buildResponseInfo(conf),
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

const (
	SVCCOMP            = "component"
	SVCREQ             = "request"
	SVCRES             = "response"
	SVCCOLLECTION      = "collection"
	SVCDATATYPE        = "type"
	SVCPARAMS          = "params"
	SVCSTREAM          = "stream"
	SVCPARAMNAME       = "name"
	SVCPARAMTYPE       = "type"
	SVCPARAMCOLLECTION = "collection"
)

func buildRequestInfo(conf config.Config) *requestInfo {

	req, ok := conf.GetSubConfig(SVCREQ)
	if ok {
		requesttype, _ := req.GetString(SVCDATATYPE)
		collection, _ := req.GetBool(SVCCOLLECTION)
		stream, _ := req.GetBool(SVCSTREAM)
		paramsConf, ok := req.GetSubConfig(SVCPARAMS)
		var params []core.Param
		if ok {
			paramNames := paramsConf.AllConfigurations()
			params = make([]core.Param, len(paramNames))
			for ind, paramName := range paramNames {
				paramDesc, _ := paramsConf.GetSubConfig(paramName)
				collection, _ := paramDesc.GetBool(SVCPARAMCOLLECTION)
				paramtype, _ := paramDesc.GetString(SVCPARAMTYPE)
				params[ind] = newParam(paramName, paramtype, collection)
			}
		}
		return newRequestInfo(requesttype, collection, stream, params)
	}
	return nil
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

func buildResponseInfo(conf config.Config) *responseInfo {
	res, ok := conf.GetSubConfig(SVCRES)
	if ok {
		stream, _ := res.GetBool(SVCSTREAM)
		return newResponseInfo(stream)
	}
	return nil
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
