package core

import (
	"laatoo/sdk/core"
)

type request struct {
	Body   interface{}
	Params map[string]core.Param
}

func (req *request) GetBody() interface{} {
	return req.Body
}

func (req *request) setBody(body interface{}) {
	req.Body = body
}

func (req *request) GetParams() map[string]core.Param {
	return req.Params
}

func (req *request) setParams(params map[string]core.Param) {
	req.Params = params
}

func (req *request) GetParam(name string) (core.Param, bool) {
	val, ok := req.Params[name]
	return val, ok
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
