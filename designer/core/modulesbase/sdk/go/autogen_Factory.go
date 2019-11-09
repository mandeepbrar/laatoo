package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Factory_Ref struct {
  Id    string
}

type Factory struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Object	string `json:"Object" bson:"Object" datastore:"Object"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *Factory) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Factory",
	}
}


func (ent *Factory) GetName()string {
	return ent.Name
}
func (ent *Factory) SetName(val string) {
	ent.Name=val
}
func (ent *Factory) GetDescription()string {
	return ent.Description
}
func (ent *Factory) SetDescription(val string) {
	ent.Description=val
}
func (ent *Factory) GetObject()string {
	return ent.Object
}
func (ent *Factory) SetObject(val string) {
	ent.Object=val
}
func (ent *Factory) GetSettings()map[string]interface{} {
	return ent.Settings
}
func (ent *Factory) SetSettings(val map[string]interface{}) {
	ent.Settings=val
}
