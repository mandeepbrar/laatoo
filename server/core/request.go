package core

import (
	"laatoo/sdk/server/core"
)

type request struct {
	//	Body   interface{}
	Params map[string]core.Param
}

/*
func (req *request) GetBody() interface{} {
	return req.Body
}

func (req *request) setBody(body interface{}) {
	req.Body = body
}*/

func (req *request) GetParams(ctx core.RequestContext) map[string]core.Param {
	return req.Params
}

func (req *request) setParams(ctx core.RequestContext, params map[string]core.Param) {
	req.Params = params
}

func (req *request) GetParam(ctx core.RequestContext, name string) (core.Param, bool) {
	val, ok := req.Params[name]
	return val, ok
}
func (req *request) GetParamValue(ctx core.RequestContext, name string) (interface{}, bool) {
	val, ok := req.Params[name]
	if ok {
		return val.GetValue(), true
	}
	return nil, ok
}
func (req *request) GetStringParam(ctx core.RequestContext, name string) (string, bool) {
	val, ok := req.Params[name]
	if ok {
		pval, ok := val.GetValue().(string)
		if ok {
			return pval, ok
		}
		return "", ok
	}
	return "", ok
}

func (req *request) GetIntParam(ctx core.RequestContext, name string) (int, bool) {
	val, ok := req.Params[name]
	if ok {
		pval, ok := val.GetValue().(int)
		if ok {
			return pval, ok
		}
		return -1, ok
	}
	return -1, ok
}

func (req *request) GetStringMapParam(ctx core.RequestContext, name string) (map[string]interface{}, bool) {
	val, ok := req.Params[name]
	if ok {
		pval, ok := val.GetValue().(map[string]interface{})
		if ok {
			return pval, ok
		}
		return nil, ok
	}
	return nil, ok
}

func (req *request) GetStringsMapParam(ctx core.RequestContext, name string) (map[string]string, bool) {
	val, ok := req.Params[name]
	if ok {
		storedVal := val.GetValue()
		pval, ok := storedVal.(map[string]string)
		if ok {
			return pval, ok
		}
		return nil, ok
	}
	return nil, ok
}

func (req *request) addParam(name string, val string) {
	req.addObjectParam(name, val, "", false)
}
func (req *request) addObjectParam(name string, val interface{}, typ string, collection bool) {
	if req.Params != nil {
		req.Params[name] = &param{name: name, value: val, ptype: typ, collection: collection}
	}
}

/*
GetBody() interface{}
GetParam(string) (Param, bool)
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
