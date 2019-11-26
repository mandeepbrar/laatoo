package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type EmpClassification_Ref struct {
  Id    string
  Name string
}*/

type EmpClassification struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	AvgSalary	float32 `json:"AvgSalary" bson:"AvgSalary" datastore:"AvgSalary"`
	OnCost	float32 `json:"OnCost" bson:"OnCost" datastore:"OnCost"`
}

func (ent *EmpClassification) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Name",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "EmpClassification",
		Cacheable:       false,
	}
}



func (ent *EmpClassification) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "Name", &ent.Name); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = rdr.ReadFloat32(c, cdc, "AvgSalary", &ent.AvgSalary); err != nil {
      return err
    }
    
    if err = rdr.ReadFloat32(c, cdc, "OnCost", &ent.OnCost); err != nil {
      return err
    }
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *EmpClassification) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "Name", &ent.Name); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = wtr.WriteFloat32(c, cdc, "AvgSalary", &ent.AvgSalary); err != nil {
      return err
    }
    
    if err = wtr.WriteFloat32(c, cdc, "OnCost", &ent.OnCost); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
