package security

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

var (
	rc = &data.StorableConfig{
		IdField:         "Id",
		Type:            config.DEFAULT_ROLE,
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Role",
		Cacheable:       false,
		NotifyNew:       false,
		NotifyUpdates:   false,
	}
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
	role.Init()
	return role, nil
}

//Creates collection
func (rf *RoleFactory) CreateObjectCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	rolecollection := make([]Role, length)
	return &rolecollection, nil
}

type Role struct {
	data.SoftDeleteAuditable
	Role        string   `json:"Role" form:"Role" bson:"Role"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
	Realm       string   `json:"Realm" bson:"Realm"`
}

func (r *Role) Config() *data.StorableConfig {
	return rc
}

func (r *Role) GetPermissions() []string {
	return r.Permissions
}

func (r *Role) SetPermissions(permissions []string) {
	r.Permissions = permissions
}

func (ent *Role) GetName() string {
	return ent.Role
}
func (ent *Role) SetName(val string) {
	ent.Role = val
}
