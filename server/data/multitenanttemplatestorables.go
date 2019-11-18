package data

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/utils"
	"reflect"
	"time"

	"github.com/twinj/uuid"
)

type AbstractStorableMT struct {
	Id     string      `json:"Id" protobuf:"bytes,51,opt,name=id,proto3" bson:"Id" sql:"type:varchar(100); primary key; unique;index" gorm:"primary_key"`
	Tenant string      `json:"Tenant" protobuf:"bytes,61,opt,name=tenant,proto3" bson:"Tenant" sql:"type:varchar(100);"`
	P_ref  interface{} `json:"-" bson:"-" sql:"-"`
}

func NewAbstractStorableMT() *AbstractStorableMT {
	return &AbstractStorableMT{Id: uuid.NewV4().String()}
}
func (as *AbstractStorableMT) Constructor() {
	if as.Id != "" {
		as.Id = uuid.NewV4().String()
	}
}
func (as *AbstractStorableMT) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}
func (as *AbstractStorableMT) GetId() string {
	return as.Id
}
func (as *AbstractStorableMT) SetId(val string) {
	as.Id = val
}

func (as *AbstractStorableMT) GetLabel(ctx core.RequestContext, i interface{}) string {
	stor := i.(data.Storable)
	c := stor.Config()
	if c != nil && c.LabelField != "" {
		v := reflect.ValueOf(stor).Elem()
		f := v.FieldByName(c.LabelField)
		return f.String()
	}
	return ""
}

func (as *AbstractStorableMT) PreSave(ctx core.RequestContext) error {
	return nil
}
func (as *AbstractStorableMT) PostSave(ctx core.RequestContext) error {
	return nil
}
func (as *AbstractStorableMT) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (as *AbstractStorableMT) SetValues(ctx core.RequestContext, obj interface{}, val map[string]interface{}) error {
	return utils.SetObjectFields(ctx, obj, val, nil, nil)
}
func (as *AbstractStorableMT) IsDeleted() bool {
	return false
}
func (as *AbstractStorableMT) Delete() {
}

func (as *AbstractStorableMT) IsMultitenant() bool {
	return true
}
func (as *AbstractStorableMT) GetTenant() string {
	return as.Tenant
}
func (as *AbstractStorableMT) SetTenant(tenant string) {
	as.Tenant = tenant
}

func (as *AbstractStorableMT) Join(item data.Storable) {
}
func (as *AbstractStorableMT) Config() *data.StorableConfig {
	return nil
}

func (as *AbstractStorableMT) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	return rdr.ReadString(c, cdc, "Id", &as.Id)
}

func (as *AbstractStorableMT) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
	_, ok := props["Id"]
	if ok {
		return rdr.ReadString(c, cdc, "Id", &as.Id)
	}
	return nil
}

func (as *AbstractStorableMT) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	return wtr.WriteString(c, cdc, "Id", &as.Id)
}

func (as *AbstractStorableMT) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
	_, ok := props["Id"]
	if ok {
		return wtr.WriteString(c, cdc, "Id", &as.Id)
	}
	return nil
}

type SoftDeleteStorableMT struct {
	*AbstractStorableMT `json:",inline" bson:",inline" laatoo:"initialize=laatoo/server/data.AbstractStorableMT" protobuf:"group,65,opt,name=AbstractStorableMT,proto3"`
	Deleted             bool `json:"Deleted" bson:"Deleted"`
}

func NewSoftDeleteStorableMT() *SoftDeleteStorableMT {
	return &SoftDeleteStorableMT{NewAbstractStorableMT(), false}
}
func (sds *SoftDeleteStorableMT) IsDeleted() bool {
	return sds.Deleted
}
func (sds *SoftDeleteStorableMT) SoftDeleteField() string {
	return "Deleted"
}

func (sds *SoftDeleteStorableMT) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	if err := rdr.ReadBool(c, cdc, "Deleted", &sds.Deleted); err != nil {
		return err
	}
	return sds.AbstractStorableMT.ReadAll(c, cdc, rdr)
}

func (sds *SoftDeleteStorableMT) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
	_, ok := props["Deleted"]
	if ok {
		err := rdr.ReadBool(c, cdc, "Deleted", &sds.Deleted)
		if err != nil {
			return err
		}
	}
	return sds.AbstractStorableMT.ReadProps(c, cdc, rdr, props)
}

func (sds *SoftDeleteStorableMT) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	if err := wtr.WriteBool(c, cdc, "Deleted", &sds.Deleted); err != nil {
		return err
	}
	return sds.AbstractStorableMT.WriteAll(c, cdc, wtr)
}

func (sds *SoftDeleteStorableMT) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
	_, ok := props["Id"]
	if ok {
		err := wtr.WriteBool(c, cdc, "Deleted", &sds.Deleted)
		if err != nil {
			return err
		}
	}
	return sds.AbstractStorableMT.WriteProps(c, cdc, wtr, props)
}

type HardDeleteAuditableMT struct {
	*AbstractStorableMT `json:",inline"  bson:",inline" laatoo:"initialize=laatoo/server/data.AbstractStorableMT" protobuf:"group,65,opt,name=AbstractStorableMT,proto3"`
	New                 bool      `json:"IsNew" bson:"IsNew" protobuf:"bytes,53,opt,name=isnew,proto3"`
	CreatedBy           string    `json:"CreatedBy" bson:"CreatedBy" protobuf:"bytes,54,opt,name=createdby,proto3" gorm:"column:CreatedBy"`
	UpdatedBy           string    `json:"UpdatedBy" bson:"UpdatedBy" protobuf:"bytes,55,opt,name=updatedby,proto3" gorm:"column:UpdatedBy"`
	CreatedAt           time.Time `json:"CreatedAt" bson:"CreatedAt" protobuf:"bytes,56,opt,name=createdat,proto3" gorm:"column:CreatedAt"`
	UpdatedAt           time.Time `json:"UpdatedAt" bson:"UpdatedAt" protobuf:"bytes,57,opt,name=updatedat,proto3" gorm:"column:UpdatedAt"`
}

func NewHardDeleteAuditableMT() *HardDeleteAuditableMT {
	return &HardDeleteAuditableMT{AbstractStorableMT: NewAbstractStorableMT()}
}
func (hda *HardDeleteAuditableMT) IsNew() bool {
	return hda.New
}
func (hda *HardDeleteAuditableMT) PreSave(ctx core.RequestContext) error {
	hda.New = (hda.CreatedBy == "")
	return nil
}

func (hda *HardDeleteAuditableMT) SetCreatedAt(val time.Time) {
	hda.CreatedAt = val
}
func (hda *HardDeleteAuditableMT) GetCreatedAt() time.Time {
	return hda.CreatedAt
}

func (hda *HardDeleteAuditableMT) SetUpdatedAt(val time.Time) {
	hda.UpdatedAt = val
}
func (hda *HardDeleteAuditableMT) GetUpdatedAt() time.Time {
	return hda.UpdatedAt
}

func (hda *HardDeleteAuditableMT) SetUpdatedBy(val string) {
	hda.UpdatedBy = val
}
func (hda *HardDeleteAuditableMT) GetUpdatedBy() string {
	return hda.UpdatedBy
}

func (hda *HardDeleteAuditableMT) SetCreatedBy(val string) {
	hda.CreatedBy = val
}
func (hda *HardDeleteAuditableMT) GetCreatedBy() string {
	return hda.CreatedBy
}

func (hda *HardDeleteAuditableMT) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
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
	return hda.AbstractStorableMT.ReadAll(c, cdc, rdr)
}

func (hda *HardDeleteAuditableMT) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
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
	return hda.AbstractStorableMT.ReadProps(c, cdc, rdr, props)
}

func (hda *HardDeleteAuditableMT) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
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
	return hda.AbstractStorableMT.WriteAll(c, cdc, wtr)
}

func (hda *HardDeleteAuditableMT) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
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
	return hda.AbstractStorableMT.WriteProps(c, cdc, wtr, props)
}

type SoftDeleteAuditableMT struct {
	*SoftDeleteStorableMT `json:",inline" bson:",inline" laatoo:"initialize=laatoo/server/data.SoftDeleteStorableMT" protobuf:"group,66,opt,name=SoftDeleteStorableMT,proto3"`
	New                   bool      `json:"IsNew" bson:"IsNew" protobuf:"bytes,53,opt,name=isnew,proto3"`
	CreatedBy             string    `json:"CreatedBy" bson:"CreatedBy" protobuf:"bytes,54,opt,name=createdby,proto3" gorm:"column:CreatedBy"`
	UpdatedBy             string    `json:"UpdatedBy" bson:"UpdatedBy" protobuf:"bytes,55,opt,name=updatedby,proto3" gorm:"column:UpdatedBy"`
	CreatedAt             time.Time `json:"CreatedAt" bson:"CreatedAt" protobuf:"bytes,56,opt,name=createdat,proto3" gorm:"column:CreatedAt"`
	UpdatedAt             time.Time `json:"UpdatedAt" bson:"UpdatedAt" protobuf:"bytes,57,opt,name=updatedat,proto3" gorm:"column:UpdatedAt"`
}

func NewSoftDeleteAuditableMT() *SoftDeleteAuditableMT {
	return &SoftDeleteAuditableMT{SoftDeleteStorableMT: NewSoftDeleteStorableMT()}
}
func (hda *SoftDeleteAuditableMT) IsNew() bool {
	return hda.New
}
func (hda *SoftDeleteAuditableMT) PreSave(ctx core.RequestContext) error {
	hda.New = (hda.CreatedBy == "")
	return nil
}
func (sda *SoftDeleteAuditableMT) SetCreatedAt(val time.Time) {
	sda.CreatedAt = val
}
func (sda *SoftDeleteAuditableMT) GetCreatedAt() time.Time {
	return sda.CreatedAt
}

func (sda *SoftDeleteAuditableMT) SetUpdatedAt(val time.Time) {
	sda.UpdatedAt = val
}
func (sda *SoftDeleteAuditableMT) GetUpdatedAt() time.Time {
	return sda.UpdatedAt
}

func (sda *SoftDeleteAuditableMT) SetUpdatedBy(val string) {
	sda.UpdatedBy = val
}
func (sda *SoftDeleteAuditableMT) GetUpdatedBy() string {
	return sda.UpdatedBy
}

func (sda *SoftDeleteAuditableMT) SetCreatedBy(val string) {
	sda.CreatedBy = val
}
func (sda *SoftDeleteAuditableMT) GetCreatedBy() string {
	return sda.CreatedBy
}

func (sda *SoftDeleteAuditableMT) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
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
	return sda.SoftDeleteStorableMT.ReadAll(c, cdc, rdr)
}

func (sda *SoftDeleteAuditableMT) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
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
	return sda.SoftDeleteStorableMT.ReadProps(c, cdc, rdr, props)
}

func (sda *SoftDeleteAuditableMT) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
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
	return sda.SoftDeleteStorableMT.WriteAll(c, cdc, wtr)
}

func (sda *SoftDeleteAuditableMT) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
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
	return sda.SoftDeleteStorableMT.WriteProps(c, cdc, wtr, props)
}
