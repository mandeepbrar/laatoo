package core

import (
	"laatoo.io/sdk/config"
	"laatoo.io/sdk/ctx"
	"laatoo.io/sdk/server/auth"
)

type RequestContext interface {
	ctx.Context
	ServerContext() ServerContext
	EngineRequestContext() EngineContext
	EngineRequestParams() StringMap
	SubContext(name string) RequestContext
	GetServerElement(elemType ServerElementType) ServerElement
	//NewContext(name string) RequestContext
	GetRequest() Request
	SetResponse(*Response)
	GetSession() Session
	GetResponse() *Response
	//GetBody() interface{}
	GetParam(string) (Param, bool)
	GetParams() map[string]Param
	GetParamValue(string) (interface{}, bool)
	GetIntParam(string) (int, bool)
	GetStringParam(string) (string, bool)
	GetConfigParam(string) (config.Config, bool)
	GetConfigArrParam(string) ([]config.Config, bool)
	GetStringMapParam(string) (StringMap, bool)
	GetStringsMapParam(string) (StringsMap, bool)
	Invoke(alias string, params StringMap) (*Response, error)
	Forward(string, StringMap) error
	ForwardToService(Service, StringMap) error
	GetUser() auth.User
	GetTenant() auth.TenantInfo
	HasPermission(perm string) bool
	GetObjectFactory(name string) (ObjectFactory, bool)
	CreateCollection(objectName string, length int) (interface{}, error)
	CreateObject(objectName string) (interface{}, error)
	CreateObjectPointersCollection(objectName string, length int) (interface{}, error)
	PublishMessage(topic string, message interface{})
	SendSynchronousMessage(msgType string, data interface{}) error
	PutInCache(bucket string, key string, item interface{}) error
	PutMultiInCache(bucket string, vals StringMap) error
	GetFromCache(bucket string, key string) (interface{}, bool)
	GetMultiFromCache(bucket string, keys []string) StringMap
	GetObjectFromCache(bucket string, key string, objectType string) (interface{}, bool)
	IncrementInCache(bucket string, key string) error
	DecrementInCache(bucket string, key string) error
	GetObjectsFromCache(bucket string, keys []string, objectType string) StringMap
	PushTask(queue string, taskdata interface{}) error
	SubscribeTaskCompletion(queue string, callback func(ctx RequestContext, invocationId string, result interface{})) error
	StartWorkflow(workflowName string, initData StringMap, insconf StringMap) (interface{}, error)
	InvalidateCache(bucket string, key string) error
	GetCodec(encoding string) (Codec, bool)
	SendCommunication(communication interface{}) error
	GetRegName(object interface{}) (string, bool, bool)
	GetExpressionValue(expression Expression, vars StringMap) (interface{}, error)
	InvokeActivity(activity string, params StringMap) (interface{}, error)
	InvokeScript(script string, params StringMap) (interface{}, error)
	ExectueAction(action *Action, params StringMap) (interface{}, error)
	CompleteRequest()
}
