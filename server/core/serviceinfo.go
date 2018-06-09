package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
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
		si.request = nil
	} else {
		si.request = reqInfo.(*requestInfo)
	}
	if resInfo == nil {
		si.response = nil
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

type param struct {
	name                        string
	ptype                       string
	oDataType                   objectDataType
	dataObjectCreator           core.ObjectCreator
	dataObjectCollectionCreator core.ObjectCollectionCreator
	collection                  bool
	isStream                    bool
	required                    bool
	value                       interface{}
}

func newParam(ctx core.ServerContext, name, ptype string, collection, stream, required bool) (*param, error) {
	dataObjectType := convertDataType(ptype)
	p := &param{name, ptype, dataObjectType, nil, nil, collection, stream, required, nil}
	if dataObjectType == __custom {
		if collection {
			dataObjectCollectionCreator, err := ctx.GetObjectCollectionCreator(ptype)
			if err != nil {
				return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", ptype)
			}
			p.dataObjectCollectionCreator = dataObjectCollectionCreator
		} else {
			dataObjectCreator, err := ctx.GetObjectCreator(ptype)
			if err != nil {
				return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", ptype)
			}
			p.dataObjectCreator = dataObjectCreator
		}
	}

	return p, nil
}

func (p *param) clone() *param {
	return &param{p.name, p.ptype, p.oDataType, p.dataObjectCreator, p.dataObjectCollectionCreator, p.collection, p.isStream, p.required, p.value}
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

func (p *param) IsStream() bool {
	return p.isStream
}

func (p *param) GetValue() interface{} {
	return p.value
}

func cloneParamsMap(params map[string]core.Param) map[string]core.Param {
	cparams := make(map[string]core.Param, len(params))
	for k, v := range params {
		cparams[k] = v.(*param).clone()
	}
	return cparams
}

func readParamsConf(ctx core.ServerContext, paramsConf config.Config) (map[string]core.Param, error) {
	params := make(map[string]core.Param)
	paramNames := paramsConf.AllConfigurations(ctx)
	for _, paramName := range paramNames {
		paramDesc, _ := paramsConf.GetSubConfig(ctx, paramName)
		collection, _ := paramDesc.GetBool(ctx, SVCPARAMCOLLECTION)
		paramtype, _ := paramDesc.GetString(ctx, SVCPARAMTYPE)
		stream, _ := paramDesc.GetBool(ctx, SVCSTREAM)
		paramreqd, _ := paramDesc.GetBool(ctx, SVCPARAMREQD)
		p, err := newParam(ctx, paramName, paramtype, collection, stream, paramreqd)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		} else {
			params[paramName] = p
		}
	}
	return params, nil
}

type objectDataType int

const (
	__stringmap objectDataType = iota
	__stringsmap
	__bytes
	__files
	__inttype
	__stringtype
	__stringarr
	__booltype
	__custom
	__none
)

func convertDataType(dtype string) objectDataType {
	switch dtype {
	case "":
		return __none
	case config.OBJECTTYPE_STRINGMAP:
		return __stringmap
	case config.OBJECTTYPE_STRINGSMAP:
		return __stringsmap
	case config.OBJECTTYPE_BYTES:
		return __bytes
	case config.OBJECTTYPE_STRING:
		return __stringtype
	case config.OBJECTTYPE_STRINGARR:
		return __stringarr
	case config.OBJECTTYPE_BOOL:
		return __booltype
	case config.OBJECTTYPE_FILES:
		return __files
	default:
		return __custom
	}
}
