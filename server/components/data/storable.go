package data

import (
	"fmt"
	"reflect"

	"laatoo.io/sdk/ctx"
	"laatoo.io/sdk/server/core"
	"laatoo.io/sdk/server/log"
	serverutils "laatoo.io/sdk/server/utils"

	"github.com/twinj/uuid"
)

/**
protobuf numbers

id = 51
deleted=52
isnew=53
createdby=54
updatedby=55
createdat=56
updatedat=57
type=59
name=60
tenant=61
AbstractStorable=62
SoftDeleteStorable=63
Entity=64
AbstractStorableMT=65
SoftDeleteStorableMT=66
HardDeleteAuditable=67
SoftDeleteAuditable=68
HardDeleteAuditableMT=69
SoftDeleteAuditableMT=70
SerializableBase=71
collection=72
tenantname=73
*/

type StorableConfig struct {
	LabelField        string
	PartialLoadFields []string
	FullLoadFields    []string
	PreSave           bool
	PostSave          bool
	PostUpdate        bool
	PostLoad          bool
	Trackable         bool
	Collection        string
	Cacheable         bool
	RefOps            bool
	Multitenant       bool
}

// Object stored by data service
type Storable interface {
	Constructor()
	Config() *StorableConfig
	GetId() string
	SetId(string)
	GetLabel(core.RequestContext, interface{}) string
	SetValues(core.RequestContext, interface{}, core.StringMap) error
	PreSave(ctx core.RequestContext) error
	PostSave(ctx core.RequestContext) error
	PostLoad(ctx core.RequestContext) error
	IsMultitenant() bool
	Join(item Storable)
	GetObjectRef() interface{}
}

type StorageInfo struct {
	Id    string      `json:"Id" bson:"Id" protobuf:"bytes,51,opt,name=id,proto3" sql:"type:varchar(100); primary key;" gorm:"primary_key"`
	P_ref interface{} `json:"-" datastore:"-" bson:"-" sql:"-"`
}

func (si *StorageInfo) Constructor() {
	if si.Id == "" {
		si.Id = uuid.NewV4().String()
	}
}

func (si *StorageInfo) GetId() string {
	return si.Id
}
func (si *StorageInfo) SetId(val string) {
	si.Id = val
}

func (si *StorageInfo) GetLabel(ctx core.RequestContext, i interface{}) string {
	stor := i.(Storable)
	c := stor.Config()
	if c != nil && c.LabelField != "" {
		v := reflect.ValueOf(stor).Elem()
		f := v.FieldByName(c.LabelField)
		return f.String()
	}
	return ""
}

func (si *StorageInfo) PreSave(ctx core.RequestContext) error {
	return nil
}
func (si *StorageInfo) PostSave(ctx core.RequestContext) error {
	return nil
}
func (si *StorageInfo) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (si *StorageInfo) SetValues(ctx core.RequestContext, obj interface{}, val core.StringMap) error {
	delete(val, "Id")
	delete(val, "IsNew")
	delete(val, "CreatedBy")
	delete(val, "UpdatedBy")
	delete(val, "CreatedAt")
	delete(val, "UpdatedAt")
	return serverutils.SetObjectFields(ctx.ServerContext(), obj, val, nil, nil)
}

func (si *StorageInfo) IsMultitenant() bool {
	return false
}

func (si *StorageInfo) Join(item Storable) {
}
func (si *StorageInfo) Config() *StorableConfig {
	return nil
}

func (si *StorageInfo) GetObjectRef() interface{} {
	return si.P_ref
}

func (si *StorageInfo) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error
	if err = rdr.ReadString(c, cdc, "Id", &si.Id); err != nil {
		return err
	}
	return nil
}

func (si *StorageInfo) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error
	if err = wtr.WriteString(c, cdc, "Id", &si.Id); err != nil {
		return err
	}
	return nil
}

type StorableRef struct {
	Id     string   `json:"Id" bson:"Id" protobuf:"bytes,51,opt,name=id,proto3" sql:"type:varchar(100);`
	Type   string   `json:"Type" bson:"Type" protobuf:"bytes,59,opt,name=type,proto3" sql:"type:varchar(100);`
	Name   string   `json:"Name" bson:"Name" protobuf:"bytes,60,opt,name=name,proto3" sql:"type:varchar(300);`
	Entity Storable `json:"-" datastore:"-" bson:"-" sql:"-" protobuf:"group,64,opt,name=Entity,proto3"`
}

func StorableArrayToMap(items []Storable) map[string]Storable {
	res := make(map[string]Storable, len(items))
	for _, item := range items {
		res[item.GetId()] = item
	}
	return res
}

//Factory function for creating storable
//type StorableCreator func() interface{}

func CastToStorableCollection(cx ctx.Context, items interface{}) ([]Storable, []string, error) {
	arr := reflect.ValueOf(items)
	if arr.Kind() == reflect.Ptr {
		arr = arr.Elem()
	}
	if arr.Kind() != reflect.Slice {
		return nil, nil, fmt.Errorf("Invalid cast to Storable. Type of Item: %s", arr.Kind())
	}
	length := arr.Len()
	retVal := make([]Storable, length)
	ids := make([]string, length)
	j := 0
	for i := 0; i < length; i++ {
		itemKind := arr.Index(i).Kind()
		var valPtr interface{}
		if itemKind == reflect.Ptr {
			valPtr = arr.Index(i).Interface()
		} else {
			valPtr = arr.Index(i).Addr().Interface()
		}
		if valPtr != nil {
			stor, ok := valPtr.(Storable)
			if !ok {
				return nil, nil, fmt.Errorf("Invalid cast to Storable. Item: %s", valPtr)
			}
			softDeletable, ok := stor.(SoftDeletable)
			if ok && softDeletable.IsDeleted() {
				continue
			}
			ids[j] = stor.GetId()
			retVal[j] = stor
			j++
		} else {
			log.Warn(cx, "Nil object received", "index", i)
		}
	}
	return retVal[0:j], ids, nil
}

func CastToStorableHash(items interface{}) (map[string]Storable, error) {
	arr := reflect.ValueOf(items)
	if arr.Kind() == reflect.Ptr {
		arr = arr.Elem()
	}
	if arr.Kind() != reflect.Slice {
		return nil, fmt.Errorf("Invalid cast to Storable. Type of Item: %s", arr.Kind())
	}
	length := arr.Len()
	retVal := make(map[string]Storable, length)
	for i := 0; i < length; i++ {
		itemKind := arr.Index(i).Kind()
		var valPtr interface{}
		if itemKind == reflect.Ptr {
			valPtr = arr.Index(i).Interface()
		} else {
			valPtr = arr.Index(i).Addr().Interface()
		}
		stor, ok := valPtr.(Storable)
		if !ok {
			return nil, fmt.Errorf("Invalid cast to Storable. Item: %s %s %t", valPtr, arr.Index(i).Kind(), arr.Index(i).IsNil())
		}
		softDeletable, ok := stor.(SoftDeletable)
		if ok && softDeletable.IsDeleted() {
			continue
		}
		retVal[stor.GetId()] = stor
	}
	return retVal, nil
}
