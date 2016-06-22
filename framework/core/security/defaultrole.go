package security

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"

	"github.com/twinj/uuid"
)

func init() {
	objects.RegisterObjectFactory(config.DEFAULT_ROLE, &RoleFactory{})
}

//interface that needs to be implemented by any object provider in a system
type RoleFactory struct {
}

//Initialize the object factory
func (rf *RoleFactory) Initialize(ctx core.ServerContext, config config.Config) error {
	return nil
}
func (rf *RoleFactory) Start(ctx core.ServerContext) error {
	return nil
}

//Creates object
func (rf *RoleFactory) CreateObject(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	role := &Role{}
	role.Id = uuid.NewV4().String()
	return role, nil
}

//Creates collection
func (rf *RoleFactory) CreateObjectCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	rolecollection := make([]Role, length)
	return &rolecollection, nil
}

type Role struct {
	Id          string   `json:"Id" form:"Id" bson:"Id"`
	Role        string   `json:"Role" form:"Role" bson:"Role"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
	Deleted     bool     `json:"Deleted" bson:"Deleted"`
	CreatedBy   string   `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string   `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn   string   `json:"UpdatedOn" bson:"UpdatedOn"`
	Realm       string   `json:"Realm" bson:"Realm"`
}

func (r *Role) GetId() string {
	return r.Id
}
func (r *Role) SetId(id string) {
	r.Id = id
}
func (r *Role) GetIdField() string {
	return "Id"
}

func (r *Role) GetObjectType() string {
	return config.DEFAULT_ROLE
}

func (r *Role) GetPermissions() []string {
	return r.Permissions
}

func (r *Role) SetPermissions(permissions []string) {
	r.Permissions = permissions
}
func (ent *Role) PreSave(ctx core.RequestContext) error {
	return nil
}
func (ent *Role) PostSave(ctx core.RequestContext) error {
	return nil
}
func (ent *Role) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (ent *Role) GetName() string {
	return ent.Role
}
func (ent *Role) SetName(val string) {
	ent.Role = val
}
func (ent *Role) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *Role) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *Role) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *Role) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *Role) GetCreatedBy() string {
	return ent.CreatedBy
}
