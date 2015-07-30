package user

import (
	jwt "github.com/dgrijalva/jwt-go"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/utils"
)

const (
	CONF_DEFAULT_USER = "default_user"
)

type DefaultUser struct {
	Id          string          `json:"Id" form:"Id" bson:"Id"`
	Password    string          `json:"Password" form:"Password" bson:"Password"`
	Roles       utils.StringSet `json:"Roles" bson:"Roles"`
	Permissions utils.StringSet `json:"Permissions" bson:"Permissions"`
}

func (usr *DefaultUser) GetId() string {
	return usr.Id
}
func (usr *DefaultUser) SetId(id string) {
	usr.Id = id
}
func (usr *DefaultUser) GetIdField() string {
	return "Id"
}
func (usr *DefaultUser) GetPassword() string {
	return usr.Password
}
func (usr *DefaultUser) SetPassword(password string) {
	usr.Password = password
}
func (usr *DefaultUser) GetRoles() (utils.StringSet, error) {
	return usr.Roles, nil
}
func (usr *DefaultUser) SetRoles(roles utils.StringSet) error {
	usr.Roles = roles
	return nil
}
func (usr *DefaultUser) GetPermissions() (permissions utils.StringSet, err error) {
	return usr.Permissions, nil
}
func (usr *DefaultUser) LoadPermissions(roleStorer data.DataService) error {
	roles := usr.Roles
	usr.Permissions = utils.NewStringSet([]string{})
	for k, _ := range roles {
		roleInt, err := roleStorer.GetById(CONF_DEFAULT_USER, k)
		if err == nil {
			role := roleInt.(RbacRole)
			usr.Permissions.Join(role.GetPermissions())
		}
	}
	return nil
}
func (usr *DefaultUser) AddRole(role string) error {
	usr.Roles.Add(role)
	return nil
}
func (usr *DefaultUser) RemoveRole(role string) error {
	usr.Roles.Remove(role)
	return nil
}

func (usr *DefaultUser) SetJWTClaims(*jwt.Token) {

}

func (usr *DefaultUser) LoadJWTClaims(*jwt.Token) {

}

func init() {
	laatoocore.RegisterObjectProvider(CONF_DEFAULT_USER, NewUser)
}

func NewUser(conf map[string]interface{}) (interface{}, error) {
	return &DefaultUser{}, nil
}
