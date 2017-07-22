package core

import "laatoo/sdk/core"

type serviceInfo struct {
	request        *requestInfo
	response       *responseInfo
	description    string
	component      bool
	configurations map[string]interface{}
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
func (svcinfo *serviceInfo) GetConfigurations() map[string]interface{} {
	return svcinfo.configurations
}

type requestInfo struct {
	dataType     string
	isCollection bool
	streaming    bool
	params       map[string]core.Param
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

func (ri *responseInfo) IsStream() bool {
	return ri.streaming
}

type configuration struct {
	name     string
	conftype string
	required bool
	value    interface{}
}

type param struct {
	name       string
	ptype      string
	collection bool
	value      interface{}
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
