package main

import (
  
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type GLNormalLineItem_Ref struct {
  Id    string
  LineItem string
}*/

type GLNormalLineItem struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	LineItem	string `json:"LineItem" bson:"LineItem" datastore:"LineItem"`
	Budget	string `json:"Budget" bson:"Budget" datastore:"Budget"`
	Jan	int64 `json:"Jan" bson:"Jan" datastore:"Jan"`
	Feb	int64 `json:"Feb" bson:"Feb" datastore:"Feb"`
	Mar	int64 `json:"Mar" bson:"Mar" datastore:"Mar"`
	Apr	int64 `json:"Apr" bson:"Apr" datastore:"Apr"`
	May	int64 `json:"May" bson:"May" datastore:"May"`
	Jun	int64 `json:"Jun" bson:"Jun" datastore:"Jun"`
	Jul	int64 `json:"Jul" bson:"Jul" datastore:"Jul"`
	Aug	int64 `json:"Aug" bson:"Aug" datastore:"Aug"`
	Sep	int64 `json:"Sep" bson:"Sep" datastore:"Sep"`
	Oct	int64 `json:"Oct" bson:"Oct" datastore:"Oct"`
	Nov	int64 `json:"Nov" bson:"Nov" datastore:"Nov"`
	Dec	int64 `json:"Dec" bson:"Dec" datastore:"Dec"`
}

func (ent *GLNormalLineItem) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "LineItem",
		PreSave:         true,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "GLNormalLineItem",
		Cacheable:       false,
	}
}



func (ent *GLNormalLineItem) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  
    if err = rdr.ReadString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "LineItem", &ent.LineItem); err != nil {
      return err
    }
    
    if err = rdr.ReadString(c, cdc, "Budget", &ent.Budget); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Jan", &ent.Jan); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Feb", &ent.Feb); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Mar", &ent.Mar); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Apr", &ent.Apr); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "May", &ent.May); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Jun", &ent.Jun); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Jul", &ent.Jul); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Aug", &ent.Aug); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Sep", &ent.Sep); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Oct", &ent.Oct); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Nov", &ent.Nov); err != nil {
      return err
    }
    
    if err = rdr.ReadInt64(c, cdc, "Dec", &ent.Dec); err != nil {
      return err
    }
    

	return ent.Storable.ReadAll(c, cdc, rdr)
}


func (ent *GLNormalLineItem) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  
    if err = wtr.WriteString(c, cdc, "Title", &ent.Title); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "LineItem", &ent.LineItem); err != nil {
      return err
    }
    
    if err = wtr.WriteString(c, cdc, "Budget", &ent.Budget); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Jan", &ent.Jan); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Feb", &ent.Feb); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Mar", &ent.Mar); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Apr", &ent.Apr); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "May", &ent.May); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Jun", &ent.Jun); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Jul", &ent.Jul); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Aug", &ent.Aug); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Sep", &ent.Sep); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Oct", &ent.Oct); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Nov", &ent.Nov); err != nil {
      return err
    }
    
    if err = wtr.WriteInt64(c, cdc, "Dec", &ent.Dec); err != nil {
      return err
    }
    

	return ent.Storable.WriteAll(c, cdc, wtr)
}
