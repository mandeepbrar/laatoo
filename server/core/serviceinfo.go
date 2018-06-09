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

func newServiceInfo(description string, reqInfo core.RequestInfo, resInfo core.ResponseInfo, configurations []core.Configuration) *serviceInfo {
	co := newConfigurableObject(description, "Service")
	co.setConfigurations(configurations)
	si := &serviceInfo{configurableObject: co,
		component:    false,
		svcsToInject: make(map[string]string)}
	if resInfo == nil {
		si.request = newRequestInfo(make(map[string]core.Param))
	} else {
		si.request = reqInfo.(*requestInfo)
	}
	if resInfo == nil {
		si.response = newResponseInfo(make(map[string]core.Param))
	} else {
		si.response = resInfo.(*responseInfo)
	}
	return si
}

func buildServiceInfo(ctx core.ServerContext, conf config.Config) (*serviceInfo, error) {
	comp, _ := conf.GetBool(ctx, SVCCOMP)
	reqInfo, err := buildRequestInfo(ctx, conf)
	if err != nil {
		return nil, err
	}
	resInfo, err := buildResponseInfo(ctx, conf)
	if err != nil {
		return nil, err
	}
	return &serviceInfo{configurableObject: buildConfigurableObject(ctx, conf),
		component:    comp,
		request:      reqInfo,
		response:     resInfo,
		svcsToInject: make(map[string]string)}, nil
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
	params map[string]core.Param
}

func newRequestInfo(params map[string]core.Param) *requestInfo {
	/*params := make(map[string]core.Param)
	if params != nil {
		for _, p := range params {
			reqInfo.params[p.GetName()] = p
		}
	}*/
	return &requestInfo{params}
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

func buildRequestInfo(ctx core.ServerContext, conf config.Config) (*requestInfo, error) {
	req, ok := conf.GetSubConfig(ctx, SVCREQ)
	var params map[string]core.Param
	var err error
	if ok {
		paramsConf, ok := req.GetSubConfig(ctx, SVCPARAMS)
		if ok {
			params, err = readParamsConf(ctx, paramsConf)
			if err != nil {
				return nil, err
			}
		}
	}
	return newRequestInfo(params), nil
}

func (ri *requestInfo) clone() *requestInfo {
	return &requestInfo{cloneParamsMap(ri.params)}
}

func (ri *requestInfo) ParamInfo() map[string]core.Param {
	return ri.params
}

type responseInfo struct {
	params map[string]core.Param
}

func buildResponseInfo(ctx core.ServerContext, conf config.Config) (*responseInfo, error) {
	res, ok := conf.GetSubConfig(ctx, SVCRES)
	var params map[string]core.Param
	var err error
	if ok {
		paramsConf, ok := res.GetSubConfig(ctx, SVCPARAMS)
		if ok {
			params, err = readParamsConf(ctx, paramsConf)
			if err != nil {
				return nil, err
			}
		}
	}
	return newResponseInfo(params), nil
}

func newResponseInfo(params map[string]core.Param) *responseInfo {
	return &responseInfo{params}
}

func (ri *responseInfo) clone() *responseInfo {
	return &responseInfo{cloneParamsMap(ri.params)}
}

func (ri *responseInfo) ParamInfo() map[string]core.Param {
	return ri.params
}
