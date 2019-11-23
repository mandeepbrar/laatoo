package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type Job_Ref struct {
  Id    string
  Account string
}*/

type Job struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	JobID	string `json:"JobID" bson:"JobID" datastore:"JobID"`
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	OrgUnit	data.StorableRef `json:"OrgUnit" bson:"OrgUnit" datastore:"OrgUnit"`
}

func (ent *Job) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "Account",
		PreSave:         true,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Job",
		Cacheable:       false,
	}
}



func (ent *Job) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "JobID", &ent.JobID); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
          {
            err := rdr.ReadObject(c, cdc, "OrgUnit", &ent.OrgUnit)
            if err != nil {
              return err
            }
          }
          

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *Job) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "JobID", &ent.JobID); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = wtr.WriteObject(c, cdc, "OrgUnit", &ent.OrgUnit); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
