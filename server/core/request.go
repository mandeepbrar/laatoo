package core

import (
	"laatoo/sdk/core"
)

type request struct {
	Body   interface{}
	Params core.ServiceParamsMap
}

func (req *request) GetBody() interface{} {
	return req.Body
}

func (req *request) SetBody(body interface{}) {
	req.Body = body
}

func (req *request) GetParams() core.ServiceParamsMap {
	return req.Params
}

func (req *request) SetParams(params core.ServiceParamsMap) {
	req.Params = params
}

func (req *request) GetParam(name string) (*core.ServiceParam, bool) {
	val, ok := req.Params[name]
	return val, ok
}

func (req *request) AddParam(name string, val interface{}, typ string, collection bool) {
	if req.Params != nil {
		req.Params.AddParam(name, val, typ, collection)
	}
}
