package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Entity_Ref struct {
  Id    string
}

type Entity struct {
	*data.SerializableBase `initialize:"SerializableBase"`
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Fields	[]Field `json:"Fields" bson:"Fields" datastore:"Fields"`
}

func (ent *Entity) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Entity",
	}
}


func (ent *Entity) GetName()string {
	return ent.Name
}
func (ent *Entity) SetName(val string) {
	ent.Name=val
}
func (ent *Entity) GetDescription()string {
	return ent.Description
}
func (ent *Entity) SetDescription(val string) {
	ent.Description=val
}
func (ent *Entity) GetFields()[]Field {
	return ent.Fields
}
func (ent *Entity) SetFields(val []Field) {
	ent.Fields=val
}
