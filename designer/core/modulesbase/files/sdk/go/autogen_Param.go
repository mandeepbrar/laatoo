package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Param_Ref struct {
  Id    string
}

type Param struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Required	bool `json:"Required" bson:"Required" datastore:"Required"`
	Default	string `json:"Default" bson:"Default" datastore:"Default"`
}

func (ent *Param) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Param",
	}
}


func (ent *Param) GetName()string {
	return ent.Name
}
func (ent *Param) SetName(val string) {
	ent.Name=val
}
func (ent *Param) GetType()string {
	return ent.Type
}
func (ent *Param) SetType(val string) {
	ent.Type=val
}
func (ent *Param) GetDescription()string {
	return ent.Description
}
func (ent *Param) SetDescription(val string) {
	ent.Description=val
}
func (ent *Param) GetRequired()bool {
	return ent.Required
}
func (ent *Param) SetRequired(val bool) {
	ent.Required=val
}
func (ent *Param) GetDefault()string {
	return ent.Default
}
func (ent *Param) SetDefault(val string) {
	ent.Default=val
}
