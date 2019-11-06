package modulesrepository

import (
  
  "laatoo/sdk/server/components/data"
)

type Entitlement_Ref struct {
  Id    string
  Title string
}

type Entitlement struct {
	*data.SoftDeleteAuditableMT `initialize:"SoftDeleteAuditableMT"`
  
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
