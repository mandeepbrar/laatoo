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

func buildServiceInfo(ctx core.ServerContext, conf config.Config) *serviceInfo {
	comp, _ := conf.GetBool(ctx, SVCCOMP)
	return &serviceInfo{configurableObject: buildConfigurableObject(ctx, conf),
		component:    comp,
		request:      buildRequestInfo(ctx, conf),
		response:     buildResponseInfo(ctx, conf),
		svcsToInject: make(map[string]string)}
}

func (svcinfo *serviceInfo) clone() *serviceInfo {
	inf := &serviceInfo{configurableObject: svcinfo.configurableObject.clone(),
		component: svcinfo.component,
		request:   svcinfo.request.clone(),
		response:  svcinfo.response.clone()}
	inf.svcsToInject = make(map[string]string, len(svcinfo.svcsToInject))
	for k, v := range svcinfo.svcsToInject {
		inf.svcsToInject[k] = v
	}
	return inf
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
	SVCPARAMREQD       = "required"
)

func buildRequestInfo(ctx core.ServerContext, conf config.Config) *requestInfo {
	req, ok := conf.GetSubConfig(ctx, SVCREQ)
	if ok {
		requesttype, _ := req.GetString(ctx, SVCDATATYPE)
		collection, _ := req.GetBool(ctx, SVCCOLLECTION)
		stream, _ := req.GetBool(ctx, SVCSTREAM)
		paramsConf, ok := req.GetSubConfig(ctx, SVCPARAMS)
		var params []core.Param
		if ok {
			paramNames := paramsConf.AllConfigurations(ctx)
			params = make([]core.Param, len(paramNames))
			for ind, paramName := range paramNames {
				paramDesc, _ := paramsConf.GetSubConfig(ctx, paramName)
				collection, _ := paramDesc.GetBool(ctx, SVCPARAMCOLLECTION)
				paramtype, _ := paramDesc.GetString(ctx, SVCPARAMTYPE)
				paramreqd, _ := paramDesc.GetBool(ctx, SVCPARAMREQD)
				params[ind] = newParam(paramName, paramtype, collection, paramreqd)
			}
		}
		return newRequestInfo(requesttype, collection, stream, params)
	}
	return newRequestInfo("", false, false, nil)
}

func (ri *requestInfo) clone() *requestInfo {
	params := make(map[string]core.Param, len(ri.params))
	for k, v := range ri.params {
		params[k] = v.(*param).clone()
	}
	return &requestInfo{ri.dataType, ri.isCollection, ri.streaming, params}
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

func buildResponseInfo(ctx core.ServerContext, conf config.Config) *responseInfo {
	res, ok := conf.GetSubConfig(ctx, SVCRES)
	if ok {
		stream, _ := res.GetBool(ctx, SVCSTREAM)
		return newResponseInfo(stream)
	}
	return newResponseInfo(false)
}

func newResponseInfo(streaming bool) *responseInfo {
	return &responseInfo{streaming}
}
func (ri *responseInfo) IsStream() bool {
	return ri.streaming
}

func (ri *responseInfo) clone() *responseInfo {
	return newResponseInfo(ri.streaming)
}

type param struct {
	name       string
	ptype      string
	collection bool
	required   bool
	value      interface{}
}

func newParam(name, ptype string, collection, required bool) *param {
	return &param{name, ptype, collection, required, nil}
}

func (p *param) clone() *param {
	return &param{p.name, p.ptype, p.collection, p.required, p.value}
}

func (p *param) GetName() string {
	return p.name
}
func (p *param) IsCollection() bool {
	return p.collection
}
func (p *param) IsRequired() bool {
	return p.required
}
func (p *param) GetDataType() string {
	return p.ptype
}
func (p *param) GetValue() interface{} {
	return p.value
}
