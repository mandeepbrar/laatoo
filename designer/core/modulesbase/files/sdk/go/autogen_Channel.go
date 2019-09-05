package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Channel_Ref struct {
  Id    string
}

type Channel struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Factory	string `json:"Factory" bson:"Factory" datastore:"Factory"`
	ServiceMethod	string `json:"ServiceMethod" bson:"ServiceMethod" datastore:"ServiceMethod"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *Channel) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Channel",
	}
}


func (ent *Channel) GetName()string {
	return ent.Name
}
func (ent *Channel) SetName(val string) {
	ent.Name=val
}
func (ent *Channel) GetDescription()string {
	return ent.Description
}
func (ent *Channel) SetDescription(val string) {
	ent.Description=val
}
func (ent *Channel) GetFactory()string {
	return ent.Factory
}
func (ent *Channel) SetFactory(val string) {
	ent.Factory=val
}
func (ent *Channel) GetServiceMethod()string {
	return ent.ServiceMethod
}
func (ent *Channel) SetServiceMethod(val string) {
	ent.ServiceMethod=val
}
func (ent *Channel) GetSettings()map[string]interface{} {
	return ent.Settings
}
func (ent *Channel) SetSettings(val map[string]interface{}) {
	ent.Settings=val
}
