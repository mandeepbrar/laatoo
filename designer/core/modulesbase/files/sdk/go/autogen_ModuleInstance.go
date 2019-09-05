package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleInstance_Ref struct {
  Id    string
}

type ModuleInstance struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Module	data.StorableRef `json:"Module" bson:"Module" datastore:"Module"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *ModuleInstance) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleInstance",
	}
}


func (ent *ModuleInstance) GetName()string {
	return ent.Name
}
func (ent *ModuleInstance) SetName(val string) {
	ent.Name=val
}
func (ent *ModuleInstance) GetModule()data.StorableRef {
	return ent.Module
}
func (ent *ModuleInstance) SetModule(val data.StorableRef) {
	ent.Module=val
}
func (ent *ModuleInstance) GetLoggingLevel()string {
	return ent.LoggingLevel
}
func (ent *ModuleInstance) SetLoggingLevel(val string) {
	ent.LoggingLevel=val
}
func (ent *ModuleInstance) GetLoggingFormat()string {
	return ent.LoggingFormat
}
func (ent *ModuleInstance) SetLoggingFormat(val string) {
	ent.LoggingFormat=val
}
func (ent *ModuleInstance) GetDescription()string {
	return ent.Description
}
func (ent *ModuleInstance) SetDescription(val string) {
	ent.Description=val
}
func (ent *ModuleInstance) GetSettings()map[string]interface{} {
	return ent.Settings
}
func (ent *ModuleInstance) SetSettings(val map[string]interface{}) {
	ent.Settings=val
}
