package laatooauthentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/utils"
	"strings"
)

type DefaultUser struct {
	Id          string   `json:"Id" form:"Id" bson:"Id"`
	Password    string   `json:"Password" form:"Password" bson:"Password"`
	Roles       []string `json:"Roles" bson:"Roles"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
	Email       string   `json:"Email"`
	Name        string   `json:"Name"`
	Picture     string   `json:"Picture"`
	Gender      string   `json:"Gender"`
	CreatedBy   string   `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string   `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn   string   `json:"UpdatedOn" bson:"UpdatedOn"`
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
func (ent *DefaultUser) PreSave(ctx core.Context) error {
	passlen := len(ent.Password)
	//pass length > 15 will indicate previously encrypted value as passwords > 15 chars are not suppported
	//hack to prevent password from updating if a new one hasnt been provided
	if passlen > 0 {
		if passlen > 15 {
			return errors.ThrowError(ctx, AUTH_ERROR_INVALID_PASSWORD)
		}
		err := ent.encryptPassword()
		if err != nil {
			return err
		}
	}
	return nil
}
func (ent *DefaultUser) PostSave(ctx core.Context) error {
	return nil
}
func (ent *DefaultUser) PostLoad(ctx core.Context) error {
	//ent.Password = ""
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

/*func (usr *DefaultUser) GetPermissions() (permissions []string, err error) {
	return usr.Permissions, nil
}
func (usr *DefaultUser) SetPermissions(permissions []string) {
	usr.Permissions = permissions
}*/
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

func (usr *DefaultUser) GetEmail() string {
	return usr.Email
}
func (usr *DefaultUser) GetName() string {
	return usr.Name
}
func (usr *DefaultUser) GetPicture() string {
	return usr.Picture
}
func (usr *DefaultUser) GetGender() string {
	return usr.Gender
}

func (usr *DefaultUser) LoadJWTClaims(token *jwt.Token) {
	usr.SetId(token.Claims["UserId"].(string))
	rolesInt := token.Claims["Roles"]
	if rolesInt != nil {
		usr.Roles = strings.Split(rolesInt.(string), ",")
	}
}

func init() {
	laatoocore.RegisterObjectProvider(laatoocore.DEFAULT_USER, NewUser)
}

func NewUser(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
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

func (ent *DefaultUser) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *DefaultUser) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *DefaultUser) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *DefaultUser) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *DefaultUser) GetCreatedBy() string {
	return ent.CreatedBy
}
