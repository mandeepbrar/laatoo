package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Field_Ref struct {
  Id    string
}

type Field struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
}

func (ent *Field) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Field",
	}
}


func (ent *Field) GetName()string {
	return ent.Name
}
func (ent *Field) SetName(val string) {
	ent.Name=val
}
func (ent *Field) GetType()string {
	return ent.Type
}
func (ent *Field) SetType(val string) {
	ent.Type=val
}
