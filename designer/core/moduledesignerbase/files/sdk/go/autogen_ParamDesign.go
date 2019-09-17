package moduledesignerbase


import (
  
  "laatoo/sdk/server/components/data"
)

type ParamDesign_Ref struct {
  Id    string
}

type ParamDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Required	bool `json:"Required" bson:"Required" datastore:"Required"`
	Default	string `json:"Default" bson:"Default" datastore:"Default"`
}

func (ent *ParamDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ParamDesign",
	}
}


func (ent *ParamDesign) GetName()string {
	return ent.Name
}
func (ent *ParamDesign) SetName(val string) {
	ent.Name=val
}
func (ent *ParamDesign) GetType()string {
	return ent.Type
}
func (ent *ParamDesign) SetType(val string) {
	ent.Type=val
}
func (ent *ParamDesign) GetDescription()string {
	return ent.Description
}
func (ent *ParamDesign) SetDescription(val string) {
	ent.Description=val
}
func (ent *ParamDesign) GetRequired()bool {
	return ent.Required
}
func (ent *ParamDesign) SetRequired(val bool) {
	ent.Required=val
}
func (ent *ParamDesign) GetDefault()string {
	return ent.Default
}
func (ent *ParamDesign) SetDefault(val string) {
	ent.Default=val
}
