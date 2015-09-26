package laatooauthentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"laatoocore"
	"laatoosdk/utils"
	"strings"
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
func (ent *DefaultUser) PreSave() error {
	err := ent.encryptPassword()
	if err != nil {
		return err
	}
	return nil
}
func (ent *DefaultUser) PostLoad() error {
	return nil
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
func (usr *DefaultUser) SetPermissions(permissions []string) {
	usr.Permissions = permissions
}
func (usr *DefaultUser) AddRole(role string) error {
	usr.Roles = append(usr.Roles, role)
	return nil
}
func (usr *DefaultUser) RemoveRole(role string) error {
	usr.Roles = utils.Remove(usr.Roles, role)
	return nil
}

func (usr *DefaultUser) SetJWTClaims(token *jwt.Token) {
	token.Claims["Roles"] = strings.Join(usr.Roles, ",")
}

func (usr *DefaultUser) LoadJWTClaims(token *jwt.Token) {
	rolesInt := token.Claims["Roles"]
	if rolesInt != nil {
		usr.Roles = strings.Split(rolesInt.(string), ",")
	}
}

func init() {
	laatoocore.RegisterObjectProvider(laatoocore.DEFAULT_USER, NewUser)
}

func NewUser(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	return &DefaultUser{}, nil
}

func (usr *DefaultUser) encryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.SetPassword(string(hash))
	return nil
}
