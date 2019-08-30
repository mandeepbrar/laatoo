package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Entitlement_Ref struct {
  Id    string
  Title string
}

type Entitlement struct {
	data.SoftDeleteAuditableMT `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Solution	data.StorableRef `json:"Solution" bson:"Solution" datastore:"Solution"`
	Local	bool `json:"Local" bson:"Local" datastore:"Local"`
	Module	data.StorableRef `json:"Module" bson:"Module" datastore:"Module"`
}

func (ent *Entitlement) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "Entitlement",
		SoftDeleteField: "Deleted",
		PreSave:         true,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "Entitlement",
		Cacheable:       false,
	}
}



func (ent *Entitlement) GetName()string {
	return ent.Name
}
func (ent *Entitlement) SetName(val string) {
	ent.Name=val
}
func (ent *Entitlement) GetSolution()data.StorableRef {
	return ent.Solution
}
func (ent *Entitlement) SetSolution(val data.StorableRef) {
	ent.Solution=val
}
func (ent *Entitlement) GetLocal()bool {
	return ent.Local
}
func (ent *Entitlement) SetLocal(val bool) {
	ent.Local=val
}
func (ent *Entitlement) GetModule()data.StorableRef {
	return ent.Module
}
func (ent *Entitlement) SetModule(val data.StorableRef) {
	ent.Module=val
}
