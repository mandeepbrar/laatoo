package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type GLAccount_Ref struct {
  Id    string
  Title string
}*/

type GLAccount struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Code	string `json:"Code" bson:"Code" datastore:"Code"`
	Rollup	bool `json:"Rollup" bson:"Rollup" datastore:"Rollup"`
	Parent	data.StorableRef `json:"Parent" bson:"Parent" datastore:"Parent"`
	LinkedElement	data.StorableRef `json:"LinkedElement" bson:"LinkedElement" datastore:"LinkedElement"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
}

func (ent *GLAccount) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Title",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "GLAccount",
		Cacheable:       false,
	}
}



func (ent *GLAccount) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Code", &ent.Code); err != nil {
      return err
    }
    
    if err = rdr.ReadBool(c, cdc, "Rollup", &ent.Rollup); err != nil {
      return err
    }
    
          {
            err := rdr.ReadObject(c, cdc, "Parent", &ent.Parent)
            if err != nil {
              return err
            }
          }
          
          {
            err := rdr.ReadObject(c, cdc, "LinkedElement", &ent.LinkedElement)
            if err != nil {
              return err
            }
          }
          
    if err = rdr.ReadString(c, cdc, "Type", &ent.Type); err != nil {
      return err
    }
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *GLAccount) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Code", &ent.Code); err != nil {
      return err
    }
    
    if err = wtr.WriteBool(c, cdc, "Rollup", &ent.Rollup); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "Parent", &ent.Parent); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "LinkedElement", &ent.LinkedElement); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Type", &ent.Type); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
