package security

import (
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
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
	return &Role{}, nil
}

//Creates collection
func (rf *RoleFactory) CreateObjectCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	rolecollection := make([]Role, length)
	return &rolecollection, nil
}

type Role struct {
	Role        string   `json:"Role" form:"Role" bson:"Role"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
	Deleted     bool     `json:"Deleted" bson:"Deleted"`
	CreatedBy   string   `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string   `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn   string   `json:"UpdatedOn" bson:"UpdatedOn"`
}

func (r *Role) GetId() string {
	return r.Role
}
func (r *Role) SetId(id string) {
	r.Role = id
}
func (r *Role) GetIdField() string {
	return "Role"
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
