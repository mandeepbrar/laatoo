package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type Employee_Ref struct {
  Id    string
  Title string
}*/

type Employee struct {
	data.Storable `json:",inline" bson:",inline" laatoo:"auditable, softdelete"`
  
	FirstName	string `json:"FirstName" bson:"FirstName" datastore:"FirstName"`
	LastName	string `json:"LastName" bson:"LastName" datastore:"LastName"`
	EmployeeID	[]string `json:"EmployeeID" bson:"EmployeeID" datastore:"EmployeeID"`
	Job	data.StorableRef `json:"Job" bson:"Job" datastore:"Job"`
}

func (ent *Employee) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Title",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Employee",
		Cacheable:       false,
	}
}



func (ent *Employee) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "FirstName", &ent.FirstName); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "LastName", &ent.LastName); err != nil {
      return err
    }
    
    if err = rdr.ReadArray(c, cdc, "EmployeeID", &ent.EmployeeID); err != nil {
      return err
    }
    
          {
            ent.Job = data.StorableRef{}
            if err = rdr.ReadObject(c, cdc, "Job", &ent.Job); err != nil {
              return err
            }
          }
          

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *Employee) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "FirstName", &ent.FirstName); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "LastName", &ent.LastName); err != nil {
      return err
    }
    
    if err = wtr.WriteArray(c, cdc, "EmployeeID", &ent.EmployeeID); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "Job", &ent.Job); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
