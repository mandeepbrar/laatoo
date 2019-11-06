package modulesbase

import (
  
  "laatoo/sdk/server/components/data"
)

type ConfigForm_Ref struct {
  Id    string
  Title string
}

type ConfigForm struct {
	*data.SoftDeleteAuditable `initialize:"SoftDeleteAuditable"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Form	string `json:"Form" bson:"Form" datastore:"Form"`
}

func (ent *ConfigForm) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ConfigForm",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ConfigForm",
		Cacheable:       false,
	}
}
