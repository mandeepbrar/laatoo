package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type ObjectDefinition_Ref struct {
  Id    string
}

type ObjectDefinition struct {
	*data.SerializableBase `initialize:"SerializableBase"`
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	RequestType	string `json:"RequestType" bson:"RequestType" datastore:"RequestType"`
	RequestParams	[]Param `json:"RequestParams" bson:"RequestParams" datastore:"RequestParams"`
	Configurations	[]Param `json:"Configurations" bson:"Configurations" datastore:"Configurations"`
}

func (ent *ObjectDefinition) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ObjectDefinition",
	}
}


func (ent *ObjectDefinition) GetName()string {
	return ent.Name
}
func (ent *ObjectDefinition) SetName(val string) {
	ent.Name=val
}
func (ent *ObjectDefinition) GetDescription()string {
	return ent.Description
}
func (ent *ObjectDefinition) SetDescription(val string) {
	ent.Description=val
}
func (ent *ObjectDefinition) GetType()string {
	return ent.Type
}
func (ent *ObjectDefinition) SetType(val string) {
	ent.Type=val
}
func (ent *ObjectDefinition) GetRequestType()string {
	return ent.RequestType
}
func (ent *ObjectDefinition) SetRequestType(val string) {
	ent.RequestType=val
}
func (ent *ObjectDefinition) GetRequestParams()[]Param {
	return ent.RequestParams
}
func (ent *ObjectDefinition) SetRequestParams(val []Param) {
	ent.RequestParams=val
}
func (ent *ObjectDefinition) GetConfigurations()[]Param {
	return ent.Configurations
}
func (ent *ObjectDefinition) SetConfigurations(val []Param) {
	ent.Configurations=val
}
