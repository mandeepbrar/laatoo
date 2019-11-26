package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type Supplier_Ref struct {
  Id    string
  Name string
}*/

type Supplier struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
}

func (ent *Supplier) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Name",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Supplier",
		Cacheable:       false,
	}
}



func (ent *Supplier) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "Name", &ent.Name); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *Supplier) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "Name", &ent.Name); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
