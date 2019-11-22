package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type OrgNode_Ref struct {
  Id    string
  Title string
}*/

type OrgNode struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Code	string `json:"Code" bson:"Code" datastore:"Code"`
	Parent	data.StorableRef `json:"Parent" bson:"Parent" datastore:"Parent"`
	Level	string `json:"Level" bson:"Level" datastore:"Level"`
	Attr1	string `json:"Attr1" bson:"Attr1" datastore:"Attr1"`
	Attr2	string `json:"Attr2" bson:"Attr2" datastore:"Attr2"`
	Attr3	string `json:"Attr3" bson:"Attr3" datastore:"Attr3"`
	Data1	data.StorableRef `json:"Data1" bson:"Data1" datastore:"Data1"`
	Data2	data.StorableRef `json:"Data2" bson:"Data2" datastore:"Data2"`
	Data3	data.StorableRef `json:"Data3" bson:"Data3" datastore:"Data3"`
}

func (ent *OrgNode) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Title",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "OrgNode",
		Cacheable:       false,
	}
}



func (ent *OrgNode) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if _, err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if _, err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if _, err = rdr.ReadString(c, cdc, "Code", &ent.Code); err != nil {
      return err
    }
    
          {
            hasKey, err := rdr.ReadObject(c, cdc, "Parent", &ent.Parent)
            if err != nil || !hasKey {
              return err
            }
          }
          
    if _, err = rdr.ReadString(c, cdc, "Level", &ent.Level); err != nil {
      return err
    }
    
    if _, err = rdr.ReadString(c, cdc, "Attr1", &ent.Attr1); err != nil {
      return err
    }
    
    if _, err = rdr.ReadString(c, cdc, "Attr2", &ent.Attr2); err != nil {
      return err
    }
    
    if _, err = rdr.ReadString(c, cdc, "Attr3", &ent.Attr3); err != nil {
      return err
    }
    
          {
            hasKey, err := rdr.ReadObject(c, cdc, "Data1", &ent.Data1)
            if err != nil || !hasKey {
              return err
            }
          }
          
          {
            hasKey, err := rdr.ReadObject(c, cdc, "Data2", &ent.Data2)
            if err != nil || !hasKey {
              return err
            }
          }
          
          {
            hasKey, err := rdr.ReadObject(c, cdc, "Data3", &ent.Data3)
            if err != nil || !hasKey {
              return err
            }
          }
          

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *OrgNode) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
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
    
    if err = wtr.WriteObject(c, cdc, "Parent", &ent.Parent); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Level", &ent.Level); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Attr1", &ent.Attr1); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Attr2", &ent.Attr2); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Attr3", &ent.Attr3); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "Data1", &ent.Data1); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "Data2", &ent.Data2); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "Data3", &ent.Data3); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
