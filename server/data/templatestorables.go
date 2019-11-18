package data

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/utils"
	"reflect"
	"time"

	"github.com/golang/protobuf/proto"
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
*/

type SerializableBase struct {
}

func (b *SerializableBase) Reset() {
	*b = reflect.New(reflect.TypeOf(b).Elem()).Elem().Interface().(SerializableBase)
}

func (m *SerializableBase) String() string { return proto.CompactTextString(m) }

func (*SerializableBase) ProtoMessage() {}

type AbstractStorable struct {
	Id    string      `json:"Id" bson:"Id" protobuf:"bytes,51,opt,name=id,proto3" sql:"type:varchar(100); primary key; unique;index" gorm:"primary_key"`
	P_ref interface{} `json:"-" bson:"-" sql:"-"`
}

func NewAbstractStorable() *AbstractStorable {
	return &AbstractStorable{Id: uuid.NewV4().String()}
}

func (as *AbstractStorable) Constructor() {
	if as.Id == "" {
		as.Id = uuid.NewV4().String()
	}
}

func (as *AbstractStorable) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (as *AbstractStorable) GetId() string {
	return as.Id
}
func (as *AbstractStorable) SetId(val string) {
	as.Id = val
}

func (as *AbstractStorable) GetLabel(ctx core.RequestContext, i interface{}) string {
	stor := i.(data.Storable)
	c := stor.Config()
	if c != nil && c.LabelField != "" {
		v := reflect.ValueOf(stor).Elem()
		f := v.FieldByName(c.LabelField)
		return f.String()
	}
	return ""
}

func (as *AbstractStorable) PreSave(ctx core.RequestContext) error {
	return nil
}
func (as *AbstractStorable) PostSave(ctx core.RequestContext) error {
	return nil
}
func (as *AbstractStorable) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (as *AbstractStorable) SetValues(ctx core.RequestContext, obj interface{}, val map[string]interface{}) error {
	delete(val, "Id")
	delete(val, "IsNew")
	delete(val, "CreatedBy")
	delete(val, "UpdatedBy")
	delete(val, "CreatedAt")
	delete(val, "UpdatedAt")
	return utils.SetObjectFields(ctx, obj, val, nil, nil)
}

func (as *AbstractStorable) IsMultitenant() bool {
	return false
}

func (as *AbstractStorable) Join(item data.Storable) {
}
func (as *AbstractStorable) Config() *data.StorableConfig {
	return nil
}
func (as *AbstractStorable) GetTenant() string {
	return ""
}

func (as *AbstractStorable) SetTenant(tenant string) {
}

func (as *AbstractStorable) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	return rdr.ReadString(c, cdc, "Id", &as.Id)
}

func (as *AbstractStorable) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
	_, ok := props["Id"]
	if ok {
		return rdr.ReadString(c, cdc, "Id", &as.Id)
	}
	return nil
}

func (as *AbstractStorable) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	return wtr.WriteString(c, cdc, "Id", &as.Id)
}

func (as *AbstractStorable) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
	_, ok := props["Id"]
	if ok {
		return wtr.WriteString(c, cdc, "Id", &as.Id)
	}
	return nil
}

type SoftDeleteStorable struct {
	*AbstractStorable `json:",inline"  bson:",inline" laatoo:"initialize=laatoo/server/data.AbstractStorable" protobuf:"group,62,opt,name=AbstractStorable,proto3"`
	Deleted           bool `json:"Deleted" bson:"Deleted" protobuf:"bytes,52,opt,name=deleted,proto3"`
}

func NewSoftDeleteStorable() *SoftDeleteStorable {
	return &SoftDeleteStorable{NewAbstractStorable(), false}
}
func (sds *SoftDeleteStorable) IsDeleted() bool {
	return sds.Deleted
}
func (sds *SoftDeleteStorable) SoftDeleteField() string {
	return "Deleted"
}

func (sds *SoftDeleteStorable) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	if err := rdr.ReadBool(c, cdc, "Deleted", &sds.Deleted); err != nil {
		return err
	}
	return sds.AbstractStorable.ReadAll(c, cdc, rdr)
}

func (sds *SoftDeleteStorable) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
	_, ok := props["Deleted"]
	if ok {
		err := rdr.ReadBool(c, cdc, "Deleted", &sds.Deleted)
		if err != nil {
			return err
		}
	}
	return sds.AbstractStorable.ReadProps(c, cdc, rdr, props)
}

func (sds *SoftDeleteStorable) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	if err := wtr.WriteBool(c, cdc, "Deleted", &sds.Deleted); err != nil {
		return err
	}
	return sds.AbstractStorable.WriteAll(c, cdc, wtr)
}

func (sds *SoftDeleteStorable) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
	_, ok := props["Id"]
	if ok {
		err := wtr.WriteBool(c, cdc, "Deleted", &sds.Deleted)
		if err != nil {
			return err
		}
	}
	return sds.AbstractStorable.WriteProps(c, cdc, wtr, props)
}

type HardDeleteAuditable struct {
	*AbstractStorable `json:",inline"  bson:",inline" laatoo:"initialize=laatoo/server/data.AbstractStorable" protobuf:"group,62,opt,name=AbstractStorable,proto3"`
	New               bool      `json:"IsNew" bson:"IsNew" protobuf:"bytes,53,opt,name=isnew,proto3"`
	CreatedBy         string    `json:"CreatedBy" bson:"CreatedBy" protobuf:"bytes,54,opt,name=createdby,proto3" gorm:"column:CreatedBy"`
	UpdatedBy         string    `json:"UpdatedBy" bson:"UpdatedBy" protobuf:"bytes,55,opt,name=updatedby,proto3" gorm:"column:UpdatedBy"`
	CreatedAt         time.Time `json:"CreatedAt" bson:"CreatedAt" protobuf:"bytes,56,opt,name=createdat,proto3" gorm:"column:CreatedAt"`
	UpdatedAt         time.Time `json:"UpdatedAt" bson:"UpdatedAt" protobuf:"bytes,57,opt,name=updatedat,proto3" gorm:"column:UpdatedAt"`
}

func NewHardDeleteAuditable() *HardDeleteAuditable {
	return &HardDeleteAuditable{AbstractStorable: NewAbstractStorable()}
}
func (hda *HardDeleteAuditable) IsNew() bool {
	return hda.New
}
func (hda *HardDeleteAuditable) PreSave(ctx core.RequestContext) error {
	hda.New = (hda.CreatedBy == "")
	return nil
}

func (hda *HardDeleteAuditable) SetCreatedAt(val time.Time) {
	hda.CreatedAt = val
}
func (hda *HardDeleteAuditable) GetCreatedAt() time.Time {
	return hda.CreatedAt
}

func (hda *HardDeleteAuditable) SetUpdatedAt(val time.Time) {
	hda.UpdatedAt = val
}
func (hda *HardDeleteAuditable) GetUpdatedAt() time.Time {
	return hda.UpdatedAt
}

func (hda *HardDeleteAuditable) SetUpdatedBy(val string) {
	hda.UpdatedBy = val
}
func (hda *HardDeleteAuditable) GetUpdatedBy() string {
	return hda.UpdatedBy
}

func (hda *HardDeleteAuditable) SetCreatedBy(val string) {
	hda.CreatedBy = val
}
func (hda *HardDeleteAuditable) GetCreatedBy() string {
	return hda.CreatedBy
}

func (hda *HardDeleteAuditable) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error
	if err = rdr.ReadBool(c, cdc, "IsNew", &hda.New); err != nil {
		return err
	}
	if err = rdr.ReadString(c, cdc, "CreatedBy", &hda.CreatedBy); err != nil {
		return err
	}
	if err = rdr.ReadString(c, cdc, "UpdatedBy", &hda.UpdatedBy); err != nil {
		return err
	}
	if err = rdr.ReadTime(c, cdc, "CreatedAt", &hda.CreatedAt); err != nil {
		return err
	}
	if err = rdr.ReadTime(c, cdc, "UpdatedAt", &hda.UpdatedAt); err != nil {
		return err
	}
	return hda.AbstractStorable.ReadAll(c, cdc, rdr)
}

func (hda *HardDeleteAuditable) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
	var err error
	_, ok := props["IsNew"]
	if ok {
		err = rdr.ReadBool(c, cdc, "IsNew", &hda.New)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedBy"]
	if ok {
		err = rdr.ReadString(c, cdc, "CreatedBy", &hda.CreatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedBy"]
	if ok {
		err = rdr.ReadString(c, cdc, "UpdatedBy", &hda.UpdatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedAt"]
	if ok {
		err = rdr.ReadTime(c, cdc, "CreatedAt", &hda.CreatedAt)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedAt"]
	if ok {
		err = rdr.ReadTime(c, cdc, "UpdatedAt", &hda.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return hda.AbstractStorable.ReadProps(c, cdc, rdr, props)
}

func (hda *HardDeleteAuditable) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error
	if err = wtr.WriteBool(c, cdc, "IsNew", &hda.New); err != nil {
		return err
	}
	if err = wtr.WriteString(c, cdc, "CreatedBy", &hda.CreatedBy); err != nil {
		return err
	}
	if err = wtr.WriteString(c, cdc, "UpdatedBy", &hda.UpdatedBy); err != nil {
		return err
	}
	if err = wtr.WriteTime(c, cdc, "CreatedAt", &hda.CreatedAt); err != nil {
		return err
	}
	if err = wtr.WriteTime(c, cdc, "UpdatedAt", &hda.UpdatedAt); err != nil {
		return err
	}
	return hda.AbstractStorable.WriteAll(c, cdc, wtr)
}

func (hda *HardDeleteAuditable) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
	_, ok := props["IsNew"]
	var err error
	if ok {
		err = wtr.WriteBool(c, cdc, "IsNew", &hda.New)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedBy"]
	if ok {
		err = wtr.WriteString(c, cdc, "CreatedBy", &hda.CreatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedBy"]
	if ok {
		err = wtr.WriteString(c, cdc, "UpdatedBy", &hda.UpdatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedAt"]
	if ok {
		err = wtr.WriteTime(c, cdc, "CreatedAt", &hda.CreatedAt)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedAt"]
	if ok {
		err = wtr.WriteTime(c, cdc, "UpdatedAt", &hda.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return hda.AbstractStorable.WriteProps(c, cdc, wtr, props)
}

type SoftDeleteAuditable struct {
	*SoftDeleteStorable `json:",inline"  bson:",inline" laatoo:"initialize=laatoo/server/data.SoftDeleteStorable" protobuf:"group,63,opt,name=SoftDeleteStorable,proto3"`
	New                 bool      `json:"IsNew" bson:"IsNew" protobuf:"bytes,53,opt,name=isnew,proto3"`
	CreatedBy           string    `json:"CreatedBy" bson:"CreatedBy" protobuf:"bytes,54,opt,name=createdby,proto3" gorm:"column:CreatedBy"`
	UpdatedBy           string    `json:"UpdatedBy" bson:"UpdatedBy" protobuf:"bytes,55,opt,name=updatedby,proto3" gorm:"column:UpdatedBy"`
	CreatedAt           time.Time `json:"CreatedAt" bson:"CreatedAt" protobuf:"bytes,56,opt,name=createdat,proto3" gorm:"column:CreatedAt"`
	UpdatedAt           time.Time `json:"UpdatedAt" bson:"UpdatedAt" protobuf:"bytes,57,opt,name=updatedat,proto3" gorm:"column:UpdatedAt"`
}

func NewSoftDeleteAuditable() *SoftDeleteAuditable {
	return &SoftDeleteAuditable{SoftDeleteStorable: NewSoftDeleteStorable()}
}
func (sda *SoftDeleteAuditable) IsNew() bool {
	return sda.New
}
func (sda *SoftDeleteAuditable) PreSave(ctx core.RequestContext) error {
	sda.New = (sda.CreatedBy == "")
	return nil
}

func (sda *SoftDeleteAuditable) SetCreatedAt(val time.Time) {
	sda.CreatedAt = val
}
func (sda *SoftDeleteAuditable) GetCreatedAt() time.Time {
	return sda.CreatedAt
}

func (sda *SoftDeleteAuditable) SetUpdatedAt(val time.Time) {
	sda.UpdatedAt = val
}
func (sda *SoftDeleteAuditable) GetUpdatedAt() time.Time {
	return sda.UpdatedAt
}

func (sda *SoftDeleteAuditable) SetUpdatedBy(val string) {
	sda.UpdatedBy = val
}
func (sda *SoftDeleteAuditable) GetUpdatedBy() string {
	return sda.UpdatedBy
}

func (sda *SoftDeleteAuditable) SetCreatedBy(val string) {
	sda.CreatedBy = val
}
func (sda *SoftDeleteAuditable) GetCreatedBy() string {
	return sda.CreatedBy
}

func (sda *SoftDeleteAuditable) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error
	if err = rdr.ReadBool(c, cdc, "IsNew", &sda.New); err != nil {
		return err
	}
	if err = rdr.ReadString(c, cdc, "CreatedBy", &sda.CreatedBy); err != nil {
		return err
	}
	if err = rdr.ReadString(c, cdc, "UpdatedBy", &sda.UpdatedBy); err != nil {
		return err
	}
	if err = rdr.ReadTime(c, cdc, "CreatedAt", &sda.CreatedAt); err != nil {
		return err
	}
	if err = rdr.ReadTime(c, cdc, "UpdatedAt", &sda.UpdatedAt); err != nil {
		return err
	}
	return sda.SoftDeleteStorable.ReadAll(c, cdc, rdr)
}

func (sda *SoftDeleteAuditable) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
	var err error
	_, ok := props["IsNew"]
	if ok {
		err = rdr.ReadBool(c, cdc, "IsNew", &sda.New)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedBy"]
	if ok {
		err = rdr.ReadString(c, cdc, "CreatedBy", &sda.CreatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedBy"]
	if ok {
		err = rdr.ReadString(c, cdc, "UpdatedBy", &sda.UpdatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedAt"]
	if ok {
		err = rdr.ReadTime(c, cdc, "CreatedAt", &sda.CreatedAt)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedAt"]
	if ok {
		err = rdr.ReadTime(c, cdc, "UpdatedAt", &sda.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return sda.SoftDeleteStorable.ReadProps(c, cdc, rdr, props)
}

func (sda *SoftDeleteAuditable) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error
	if err = wtr.WriteBool(c, cdc, "IsNew", &sda.New); err != nil {
		return err
	}
	if err = wtr.WriteString(c, cdc, "CreatedBy", &sda.CreatedBy); err != nil {
		return err
	}
	if err = wtr.WriteString(c, cdc, "UpdatedBy", &sda.UpdatedBy); err != nil {
		return err
	}
	if err = wtr.WriteTime(c, cdc, "CreatedAt", &sda.CreatedAt); err != nil {
		return err
	}
	if err = wtr.WriteTime(c, cdc, "UpdatedAt", &sda.UpdatedAt); err != nil {
		return err
	}
	return sda.SoftDeleteStorable.WriteAll(c, cdc, wtr)
}

func (sda *SoftDeleteAuditable) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
	_, ok := props["IsNew"]
	var err error
	if ok {
		err = wtr.WriteBool(c, cdc, "IsNew", &sda.New)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedBy"]
	if ok {
		err = wtr.WriteString(c, cdc, "CreatedBy", &sda.CreatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedBy"]
	if ok {
		err = wtr.WriteString(c, cdc, "UpdatedBy", &sda.UpdatedBy)
		if err != nil {
			return err
		}
	}
	_, ok = props["CreatedAt"]
	if ok {
		err = wtr.WriteTime(c, cdc, "CreatedAt", &sda.CreatedAt)
		if err != nil {
			return err
		}
	}
	_, ok = props["UpdatedAt"]
	if ok {
		err = wtr.WriteTime(c, cdc, "UpdatedAt", &sda.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return sda.SoftDeleteStorable.WriteProps(c, cdc, wtr, props)
}
