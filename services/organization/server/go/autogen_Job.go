package main

import (
  
	"bytes" 
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type Job_Ref struct {
  Id    string
  Account string
}*/

type Job struct {
	data.Storable `json:",inline" bson:",inline" laatoo:"auditable, softdelete"`
  
	JobID	string `json:"JobID" bson:"JobID" datastore:"JobID"`
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	OrgUnit	data.StorableRef `json:"OrgUnit" bson:"OrgUnit" datastore:"OrgUnit"`
}

func (ent *Job) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Account",
		Type:            "Job",
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

  
      var buf bytes.Buffer
      
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
        objRdr, _, err := rdr.ReadObject(c, cdc, "OrgUnit", nil)
        if err != nil {
          return err
        }
        obj, err := c.(core.RequestContext).CreateObject("OrgNode")
        if err != nil {
          return err
        }
        srl, ok := obj.(core.Serializable)
        if ok {
          err =srl.ReadAll(c, cdc, objRdr)
          if err != nil {
            return err
          }
        } else {
          buf.Reset()
          err = objRdr.ReadBytes(c, cdc, &buf)
          if err != nil {
            return err
          }
          err = cdc.Unmarshal(c, buf.Bytes(), obj)
          if err != nil {
            return err
          }
        }
        ent.OrgUnit = obj.(data.StorableRef)
      }
      
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *Job) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
  var ok bool
	var err error

  
      var buf bytes.Buffer
      
    _, ok = props["JobID"]
    if ok {
      err = rdr.ReadString(c, cdc, "JobID", &ent.JobID)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["Title"]
    if ok {
      err = rdr.ReadString(c, cdc, "Title", &ent.Title)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["Description"]
    if ok {
      err = rdr.ReadString(c, cdc, "Description", &ent.Description)
      if err != nil {
        return err
      }
    }
    
        {
          objRdr, objMap, err := rdr.ReadObject(c, cdc, "OrgUnit", props)
          if err != nil {
            return err
          }
          obj, err := c.(core.RequestContext).CreateObject("OrgNode")
          if err != nil {
            return err
          }
          srl, ok := obj.(core.Serializable)
          if ok {
            err =srl.ReadProps(c, cdc, objRdr, objMap)
            if err != nil {
              return err
            }
          } else {
            buf.Reset()
            err = objRdr.ReadBytes(c, cdc, &buf)
            if err != nil {
              return err
            }
            err = cdc.Unmarshal(c, buf.Bytes(), obj)
            if err != nil {
              return err
            }
          }
          ent.OrgUnit = obj.(data.StorableRef)
        }  
      
    
  
	return ent.Storable.ReadProps(c, cdc, rdr, props)
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
    
        {
          objWtr, _, err := wtr.WriteObject(c, cdc, "OrgUnit", nil)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.OrgUnit
          srl, ok := objToWrite.(core.Serializable)
          if ok {
            err =srl.WriteAll(c, cdc, objWtr)
            if err != nil {
              return err
            }
          } else {
            byts, err := cdc.Marshal(c, objToWrite)
            if err != nil {
              return err
            }
            err = objWtr.WriteBytes(c, cdc, &byts)
            if err != nil {
              return err
            }
          }
        }
  
      

	return ent.Storable.WriteAll(c, cdc, wtr)
}

func (ent *Job) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
  var ok bool
	var err error

  
    _, ok = props["JobID"]
    if ok {
      err = wtr.WriteString(c, cdc, "JobID", &ent.JobID)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["Title"]
    if ok {
      err = wtr.WriteString(c, cdc, "Title", &ent.Title)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["Description"]
    if ok {
      err = wtr.WriteString(c, cdc, "Description", &ent.Description)
      if err != nil {
        return err
      }
    }
      
        {
          objWtr, objMap, err := wtr.WriteObject(c, cdc, "OrgUnit", props)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.OrgUnit
          srl, ok := objToWrite.(core.Serializable)
          if ok {
            err =srl.WriteProps(c, cdc, objWtr, objMap)
            if err != nil {
              return err
            }
          } else {
            byts, err := cdc.Marshal(c, objToWrite)
            if err != nil {
              return err
            }
            err = objWtr.WriteBytes(c, cdc, &byts)
            if err != nil {
              return err
            }
          }
        }
  
      

	return ent.Storable.WriteProps(c, cdc, wtr, props)
}