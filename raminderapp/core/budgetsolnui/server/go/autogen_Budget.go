package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type Budget_Ref struct {
  Id    string
  Title string
}*/

type Budget struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Year	string `json:"Year" bson:"Year" datastore:"Year"`
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Active	bool `json:"Active" bson:"Active" datastore:"Active"`
	Forecast	bool `json:"Forecast" bson:"Forecast" datastore:"Forecast"`
	GLAccounts	[]*GLAccountLineItem `json:"GLAccounts" bson:"GLAccounts" datastore:"GLAccounts"`
}

func (ent *Budget) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Title",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Budget",
		Cacheable:       false,
	}
}



func (ent *Budget) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "Year", &ent.Year); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = rdr.ReadBool(c, cdc, "Active", &ent.Active); err != nil {
      return err
    }
    
    if err = rdr.ReadBool(c, cdc, "Forecast", &ent.Forecast); err != nil {
      return err
    }
    
    if err = rdr.ReadArray(c, cdc, "GLAccounts", &ent.GLAccounts); err != nil {
      return err
    }
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *Budget) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "Year", &ent.Year); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = wtr.WriteBool(c, cdc, "Active", &ent.Active); err != nil {
      return err
    }
    
    if err = wtr.WriteBool(c, cdc, "Forecast", &ent.Forecast); err != nil {
      return err
    }
    
    if err = wtr.WriteArray(c, cdc, "GLAccounts", &ent.GLAccounts); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
