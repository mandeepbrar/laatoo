package core

import (
	"laatoo.io/sdk/config"
)

type Service interface {
	ConfigurableObject
	Describe(ServerContext) error
	Initialize(ctx ServerContext, conf config.Config) error
	Start(ctx ServerContext) error
	Stop(ctx ServerContext) error
	Unload(ctx ServerContext) error
	AddParams(ServerContext, map[string]DataType, bool) error
	AddStringParams(ctx ServerContext, names []string, defaultValues []string)
	AddStringParam(ctx ServerContext, name string)
	AddCustomObjectParam(ctx ServerContext, name string, customObjectType string, collection, required, stream bool) error
	AddParam(ctx ServerContext, name string, datatype DataType, collection, required, stream bool) error
	AddParamWithType(ctx ServerContext, name string, datatype DataType) error
	AddOptionalParamWithType(ctx ServerContext, name string, datatype DataType) error
	AddCollectionParams(ctx ServerContext, params map[string]DataType) error
	//	SetRequestType(ctx ServerContext, datatype string, collection bool, stream bool)
	//	SetResponseType(ctx ServerContext, stream bool)
	SetDescription(ServerContext, string)
	SetComponent(ServerContext, bool)
	//ConfigureService(ctx ServerContext, requestType string, collection bool, stream bool, params []string, config []string, description string)
	ConfigureService(ctx ServerContext, params []string, config []string, description string)
}

type UserInvokableService interface {
	Service
	Invoke(RequestContext) error
}

type Param interface {
	GetName() string
	IsCollection() bool
	IsStream() bool
	IsRequired() bool
	GetDataType() DataType
	GetValue() interface{}
}

type ServiceFunc func(ctx RequestContext) error

type Request interface {
	//GetBody() interface{}
	GetParam(RequestContext, string) (Param, bool)
	GetParams(RequestContext) map[string]Param
	GetParamValue(RequestContext, string) (interface{}, bool)
	GetIntParam(RequestContext, string) (int, bool)
	GetStringParam(RequestContext, string) (string, bool)
	GetStringMapParam(RequestContext, string) (map[string]interface{}, bool)
	GetStringsMapParam(RequestContext, string) (map[string]string, bool)
}

type Response struct {
	Status   int
	Data     interface{}
	MetaInfo map[string]interface{}
	Error    error
	Return   bool
}
