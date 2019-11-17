package main

import (
  
	"bytes" 
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type OrgNode_Ref struct {
  Id    string
  Title string
}*/

type OrgNode struct {
	data.Storable `json:",inline" bson:",inline" laatoo:"auditable, softdelete"`
  
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
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "OrgNode",
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

  
      var buf bytes.Buffer
      
    if err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Description", &ent.Description); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Code", &ent.Code); err != nil {
      return err
    }
    
      {
        objRdr, _, err := rdr.ReadObject(c, cdc, "Parent", nil)
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
        ent.Parent = obj.(data.StorableRef)
      }
      
    if err = rdr.ReadString(c, cdc, "Level", &ent.Level); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Attr1", &ent.Attr1); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Attr2", &ent.Attr2); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Attr3", &ent.Attr3); err != nil {
      return err
    }
    
      {
        objRdr, _, err := rdr.ReadObject(c, cdc, "Data1", nil)
        if err != nil {
          return err
        }
        obj, err := c.(core.RequestContext).CreateObject("")
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
        ent.Data1 = obj.(data.StorableRef)
      }
      
      {
        objRdr, _, err := rdr.ReadObject(c, cdc, "Data2", nil)
        if err != nil {
          return err
        }
        obj, err := c.(core.RequestContext).CreateObject("")
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
        ent.Data2 = obj.(data.StorableRef)
      }
      
      {
        objRdr, _, err := rdr.ReadObject(c, cdc, "Data3", nil)
        if err != nil {
          return err
        }
        obj, err := c.(core.RequestContext).CreateObject("")
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
        ent.Data3 = obj.(data.StorableRef)
      }
      
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *OrgNode) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
  var ok bool
	var err error

  
      var buf bytes.Buffer
      
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
    
    _, ok = props["Code"]
    if ok {
      err = rdr.ReadString(c, cdc, "Code", &ent.Code)
      if err != nil {
        return err
      }
    }
    
        {
          objRdr, objMap, err := rdr.ReadObject(c, cdc, "Parent", props)
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
          ent.Parent = obj.(data.StorableRef)
        }  
      
    _, ok = props["Level"]
    if ok {
      err = rdr.ReadString(c, cdc, "Level", &ent.Level)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["Attr1"]
    if ok {
      err = rdr.ReadString(c, cdc, "Attr1", &ent.Attr1)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["Attr2"]
    if ok {
      err = rdr.ReadString(c, cdc, "Attr2", &ent.Attr2)
      if err != nil {
        return err
      }
    }
    
    _, ok = props["Attr3"]
    if ok {
      err = rdr.ReadString(c, cdc, "Attr3", &ent.Attr3)
      if err != nil {
        return err
      }
    }
    
        {
          objRdr, objMap, err := rdr.ReadObject(c, cdc, "Data1", props)
          if err != nil {
            return err
          }
          obj, err := c.(core.RequestContext).CreateObject("")
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
          ent.Data1 = obj.(data.StorableRef)
        }  
      
        {
          objRdr, objMap, err := rdr.ReadObject(c, cdc, "Data2", props)
          if err != nil {
            return err
          }
          obj, err := c.(core.RequestContext).CreateObject("")
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
          ent.Data2 = obj.(data.StorableRef)
        }  
      
        {
          objRdr, objMap, err := rdr.ReadObject(c, cdc, "Data3", props)
          if err != nil {
            return err
          }
          obj, err := c.(core.RequestContext).CreateObject("")
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
          ent.Data3 = obj.(data.StorableRef)
        }  
      
    
  
	return ent.Storable.ReadProps(c, cdc, rdr, props)
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
    
        {
          objWtr, _, err := wtr.WriteObject(c, cdc, "Parent", nil)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Parent
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
    
        {
          objWtr, _, err := wtr.WriteObject(c, cdc, "Data1", nil)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Data1
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
  
      
        {
          objWtr, _, err := wtr.WriteObject(c, cdc, "Data2", nil)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Data2
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
  
      
        {
          objWtr, _, err := wtr.WriteObject(c, cdc, "Data3", nil)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Data3
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

func (ent *OrgNode) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
  var ok bool
	var err error

  
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
      
    _, ok = props["Code"]
    if ok {
      err = wtr.WriteString(c, cdc, "Code", &ent.Code)
      if err != nil {
        return err
      }
    }
      
        {
          objWtr, objMap, err := wtr.WriteObject(c, cdc, "Parent", props)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Parent
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
  
      
    _, ok = props["Level"]
    if ok {
      err = wtr.WriteString(c, cdc, "Level", &ent.Level)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["Attr1"]
    if ok {
      err = wtr.WriteString(c, cdc, "Attr1", &ent.Attr1)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["Attr2"]
    if ok {
      err = wtr.WriteString(c, cdc, "Attr2", &ent.Attr2)
      if err != nil {
        return err
      }
    }
      
    _, ok = props["Attr3"]
    if ok {
      err = wtr.WriteString(c, cdc, "Attr3", &ent.Attr3)
      if err != nil {
        return err
      }
    }
      
        {
          objWtr, objMap, err := wtr.WriteObject(c, cdc, "Data1", props)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Data1
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
  
      
        {
          objWtr, objMap, err := wtr.WriteObject(c, cdc, "Data2", props)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Data2
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
  
      
        {
          objWtr, objMap, err := wtr.WriteObject(c, cdc, "Data3", props)
          if err != nil {
            return err
          }
          var objToWrite interface{}
          objToWrite = ent.Data3
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