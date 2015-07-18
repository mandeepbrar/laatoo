package ginauth_user

import (
	"laatoo/commonobjects"
	rbac "laatoo/usermanagement/rbac"
)

type AppUser struct {
	Id          string                 `json:"Id" form:"Id" bson:"Id"`
	Password    string                 `json:"Password" form:"Password" bson:"Password"`
	Roles       storageutils.StringSet `json:"Roles" bson:"Roles"`
	Permissions storageutils.StringSet `json:"Permissions" bson:"Permissions"`
}

func (usr *AppUser) GetId() string {
	return usr.Id
}
func (usr *AppUser) SetId(id string) {
	usr.Id = id
}
func (usr *AppUser) GetIdField() string {
	return "Id"
}
func (usr *AppUser) GetPassword() string {
	return usr.Password
}
func (usr *AppUser) SetPassword(password string) {
	usr.Password = password
}
func (usr *AppUser) GetRoles() (commonobjects.StringSet, error) {
	return usr.Roles, nil
}
func (usr *AppUser) SetRoles(roles commonobjects.StringSet) error {
	usr.Roles = roles
	return nil
}
func (usr *AppUser) GetPermissions() (permissions commonobjects.StringSet, err error) {
	return usr.Permissions, nil
}
func (usr *AppUser) SetPermissions(permissions commonobjects.StringSet) error {
	usr.Permissions = permissions
	return nil
}
func (usr *AppUser) LoadPermissions(storer commonobjects.Storer) error {
	roles := usr.Roles
	usr.Permissions = commonobjects.NewStringSet([]string{})
	for k, _ := range roles {
		roleInt, err := storer.GetById(k)
		if err == nil {
			role := roleInt.(*rbac.Role)
			usr.Permissions.Join(role.Permissions)
		}
	}
	return nil
}
func (usr *AppUser) AddRole(role string) error {
	usr.Roles.Add(role)
	return nil
}
func (usr *AppUser) RemoveRole(role string) error {
	usr.Roles.Remove(role)
	return nil
}

func NewAppUser() interface{} {
	return &AppUser{}
}
