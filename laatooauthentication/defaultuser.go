package laatooauthentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/data"
	"laatoosdk/utils"
)

type DefaultUser struct {
	Id          string   `json:"Id" form:"Id" bson:"Id"`
	Password    string   `json:"Password" form:"Password" bson:"Password"`
	Roles       []string `json:"Roles" bson:"Roles"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
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
func (usr *DefaultUser) GetRoles() ([]string, error) {
	return usr.Roles, nil
}
func (usr *DefaultUser) SetRoles(roles []string) error {
	usr.Roles = roles
	return nil
}
func (usr *DefaultUser) GetPermissions() (permissions []string, err error) {
	return usr.Permissions, nil
}
func (usr *DefaultUser) LoadPermissions(roleStorer data.DataService) error {
	roles := usr.Roles
	permissions := utils.NewStringSet([]string{})
	for _, k := range roles {
		roleInt, err := roleStorer.GetById(laatoocore.DEFAULT_USER, k)
		if err == nil {
			role := roleInt.(auth.Role)
			permissions.Join(role.GetPermissions())
		}
	}
	usr.Permissions = permissions.Values()
	return nil
}
func (usr *DefaultUser) AddRole(role string) error {
	usr.Roles = append(usr.Roles, role)
	return nil
}
func (usr *DefaultUser) RemoveRole(role string) error {
	usr.Roles = utils.Remove(usr.Roles, role)
	return nil
}

func (usr *DefaultUser) SetJWTClaims(*jwt.Token) {

}

func (usr *DefaultUser) LoadJWTClaims(*jwt.Token) {

}

func init() {
	laatoocore.RegisterObjectProvider(laatoocore.DEFAULT_USER, NewUser)
}

func NewUser(conf map[string]interface{}) (interface{}, error) {
	return &DefaultUser{}, nil
}
