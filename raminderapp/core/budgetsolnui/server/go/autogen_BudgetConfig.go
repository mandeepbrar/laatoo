package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type BudgetConfig_Ref struct {
  Id    string
  Title string
}*/

type BudgetConfig struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	GLAccounts	[]*GLAccount `json:"GLAccounts" bson:"GLAccounts" datastore:"GLAccounts"`
}

func (ent *BudgetConfig) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Title",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "BudgetConfig",
		Cacheable:       false,
	}
}



func (ent *BudgetConfig) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = rdr.ReadArray(c, cdc, "GLAccounts", &ent.GLAccounts); err != nil {
      return err
    }
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *BudgetConfig) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = wtr.WriteArray(c, cdc, "GLAccounts", &ent.GLAccounts); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
