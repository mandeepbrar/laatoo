package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

/*type objectInfo struct {
	description string
	objtype     string
	properties  map[string]interface{}
}

func newObjectInfo(description, objtype string) *objectInfo {
	return &objectInfo{description, objtype, make(map[string]interface{})}
}
*/
const (
	OBJECT_TYPE = "type"
	DESCRIPTION = "description"
)

func newObjectInfo(description, objtype string) core.Info {
	return core.NewInfo(description, objtype, make(map[string]interface{}))
}

func buildObjectInfo(ctx core.ServerContext, conf config.Config) core.Info {
	desc, _ := conf.GetString(ctx, DESCRIPTION)
	objtype, _ := conf.GetString(ctx, OBJECT_TYPE)
	return newObjectInfo(desc, objtype)
}

/*func (inf *objectInfo) clone() *objectInfo {
	return newObjectInfo(inf.description, inf.objtype)
}

func (inf *objectInfo) GetDescription() string {
	return inf.description
}
func (inf *objectInfo) GetType() string {
	return inf.objtype
}

func (inf *objectInfo) GetProperty(prop string) interface{} {
	return inf.properties[prop]
}
func (inf *objectInfo) setProperty(prop string, val interface{}) {
	inf.properties[prop] = val
}
func (inf *objectInfo) setDescription(desc string) {
	inf.description = desc
}*/

type metadataProvider struct {
}

func (provider *metadataProvider) CreateServiceInfo(description string, reqInfo core.RequestInfo, resInfo core.ResponseInfo, configurations []core.Configuration) core.ServiceInfo {
	return newServiceInfo(description, reqInfo, resInfo, configurations)
}

func (provider *metadataProvider) CreateFactoryInfo(description string, configurations []core.Configuration) core.ServiceFactoryInfo {
	return newFactoryInfo(description, configurations)
}

func (provider *metadataProvider) CreateModuleInfo(description string, configurations []core.Configuration) core.ModuleInfo {
	return newModuleInfo(description, configurations)

}
func (provider *metadataProvider) CreateRequestInfo(params map[string]core.Param) core.RequestInfo {
	return newRequestInfo(params)
}

func (provider *metadataProvider) CreateResponseInfo(params map[string]core.Param) core.ResponseInfo {
	return newResponseInfo(params)
}

func (provider *metadataProvider) CreateConfiguration(name, conftype string, required bool, defaultValue interface{}) core.Configuration {
	return newConfiguration(name, conftype, required, defaultValue)
}

func (provider *metadataProvider) CreateParam(ctx core.ServerContext, name, paramtype string, collection, isStream bool, required bool) (core.Param, error) {
	return newParam(ctx, name, paramtype, collection, isStream, required)
}
