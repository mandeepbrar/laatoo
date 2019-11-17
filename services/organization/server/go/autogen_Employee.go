package main

import (
  
	"bytes" 
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
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "Employee",
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

  
      var buf bytes.Buffer
      
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
        objRdr, _, err := rdr.ReadObject(c, cdc, "Job", nil)
        if err != nil {
          return err
        }
        obj, err := c.(core.RequestContext).CreateObject("Job")
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
        ent.Job = obj.(data.StorableRef)
      }
      
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *Employee) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
  var ok bool
	var err error

  
      var buf bytes.Buffer
      
    _, ok = props["FirstName"]
    if ok {
      err = rdr.ReadString(c, cdc, "FirstName", &ent.FirstName)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["LastName"]
    if ok {
      err = rdr.ReadString(c, cdc, "LastName", &ent.LastName)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["EmployeeID"]
    if ok {
      err = rdr.ReadArray(c, cdc, "EmployeeID", &ent.EmployeeID)
      if err != nil {
        return err
      }
    }
    
        {
          objRdr, objMap, err := rdr.ReadObject(c, cdc, "Job", props)
          if err != nil {
            return err
          }
          obj, err := c.(core.RequestContext).CreateObject("Job")
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
          ent.Job = obj.(data.StorableRef)
        }  
      
    
  
	return ent.Storable.ReadProps(c, cdc, rdr, props)
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
    
        {
          objWtr, _, err := wtr.WriteObject(c, cdc, "Job", nil)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Job
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

func (ent *Employee) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
  var ok bool
	var err error

  
    _, ok = props["FirstName"]
    if ok {
      err = wtr.WriteString(c, cdc, "FirstName", &ent.FirstName)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["LastName"]
    if ok {
      err = wtr.WriteString(c, cdc, "LastName", &ent.LastName)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["EmployeeID"]
    if ok {
      err = wtr.WriteArray(c, cdc, "EmployeeID", &ent.EmployeeID)
      if err != nil {
        return err
      }
    }
      
        {
          objWtr, objMap, err := wtr.WriteObject(c, cdc, "Job", props)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Job
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