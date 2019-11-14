package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type Dependency_Ref struct {
  Id    string
}*/

type Dependency struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Version	string `json:"Version" bson:"Version" datastore:"Version"`
	Comparison	string `json:"Comparison" bson:"Comparison" datastore:"Comparison"`
}

func (ent *Dependency) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Dependency",
	}
}


func (ent *Dependency) GetName()string {
	return ent.Name
}
func (ent *Dependency) SetName(val string) {
	ent.Name=val
}
func (ent *Dependency) GetDescription()string {
	return ent.Description
}
func (ent *Dependency) SetDescription(val string) {
	ent.Description=val
}
func (ent *Dependency) GetVersion()string {
	return ent.Version
}
func (ent *Dependency) SetVersion(val string) {
	ent.Version=val
}
func (ent *Dependency) GetComparison()string {
	return ent.Comparison
}
func (ent *Dependency) SetComparison(val string) {
	ent.Comparison=val
}
