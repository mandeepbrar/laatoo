package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type objectInfo struct {
	description string
	objtype     string
}

func newObjectInfo(description, objtype string) *objectInfo {
	return &objectInfo{description, objtype}
}

const (
	OBJECT_TYPE = "type"
	DESCRIPTION = "description"
)

func buildObjectInfo(ctx core.ServerContext, conf config.Config) *objectInfo {
	desc, _ := conf.GetString(ctx, DESCRIPTION)
	objtype, _ := conf.GetString(ctx, OBJECT_TYPE)
	return newObjectInfo(desc, objtype)
}

func (inf *objectInfo) clone() *objectInfo {
	return newObjectInfo(inf.description, inf.objtype)
}

func (inf *objectInfo) GetDescription() string {
	return inf.description
}
func (inf *objectInfo) GetType() string {
	return inf.objtype
}
func (inf *objectInfo) setDescription(desc string) {
	inf.description = desc
}

type metadataProvider struct {
}

func (provider *metadataProvider) CreateServiceInfo(description string, reqInfo core.RequestInfo, streamedResponse bool, configurations []core.Configuration) core.ServiceInfo {
	return newServiceInfo(description, reqInfo, streamedResponse, configurations)
}

func (provider *metadataProvider) CreateFactoryInfo(description string, configurations []core.Configuration) core.ServiceFactoryInfo {
	return newFactoryInfo(description, configurations)
}
func (provider *metadataProvider) CreateModuleInfo(description string, configurations []core.Configuration) core.ModuleInfo {
	return newModuleInfo(description, configurations)

}
func (provider *metadataProvider) CreateRequestInfo(requesttype string, collection bool, stream bool, params []core.Param) core.RequestInfo {
	return newRequestInfo(requesttype, collection, stream, params)

}
func (provider *metadataProvider) CreateConfiguration(name, conftype string, required bool, defaultValue interface{}) core.Configuration {
	return newConfiguration(name, conftype, required, defaultValue)
}

func (provider *metadataProvider) CreateParam(name, paramtype string, collection, required bool) core.Param {
	return newParam(name, paramtype, collection, required)
}
