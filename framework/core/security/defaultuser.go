package security

import (
	"github.com/twinj/uuid"
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	objects.RegisterObjectFactory(config.DEFAULT_USER, &UserFactory{})
}

//interface that needs to be implemented by any object provider in a system
type UserFactory struct {
}

//Initialize the object factory
func (rf *UserFactory) Initialize(ctx core.ServerContext, config config.Config) error {
	return nil
}

func (rf *UserFactory) Start(ctx core.ServerContext) error {
	return nil
}

//Creates object
func (rf *UserFactory) CreateObject(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	usr := &DefaultUser{}
	usr.Id = uuid.NewV4().String()
	if args != nil {
		id, ok := args["Id"]
		if ok {
			usr.setId(id.(string))
		}
		roles, ok := args["Roles"]
		if ok {
			usr.setRoles(roles.([]string))
		}
		realm, ok := args["Realm"]
		if ok {
			usr.Realm = realm.(string)
		}

	}
	return usr, nil
}

//Creates collection
func (rf *UserFactory) CreateObjectCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	usercollection := make([]DefaultUser, length)
	return &usercollection, nil
}

type DefaultUser struct {
	Id          string   `json:"Id" form:"Id" bson:"Id"`
	Username    string   `json:"Username" form:"Username" bson:"Username"`
	Password    string   `json:"Password" form:"Password" bson:"Password"`
	Roles       []string `json:"Roles" bson:"Roles"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
	Email       string   `json:"Email" bson:"Email"`
	Name        string   `json:"Name" bson:"Name"`
	Deleted     bool     `json:"Deleted" bson:"Deleted"`
	Picture     string   `json:"Picture" bson:"Picture"`
	Gender      string   `json:"Gender" bson:"Gender"`
	CreatedBy   string   `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string   `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn   string   `json:"UpdatedOn" bson:"UpdatedOn"`
	Realm       string   `json:"Realm" bson:"Realm"`
}

func (usr *DefaultUser) GetId() string {
	return usr.Id
}

func (usr *DefaultUser) setId(id string) {
	usr.Id = id
}
func (usr *DefaultUser) GetIdField() string {
	return "Id"
}
func (usr *DefaultUser) GetUsernameField() string {
	return "Username"
}
func (ent *DefaultUser) PreSave(ctx core.RequestContext) error {
	err := ent.encryptPassword()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
func (ent *DefaultUser) GetObjectType() string {
	return config.DEFAULT_USER
}
func (ent *DefaultUser) PostSave(ctx core.RequestContext) error {
	return nil
}
func (ent *DefaultUser) PostLoad(ctx core.RequestContext) error {
	//ent.Password = ""
	return nil
}
func (usr *DefaultUser) GetPassword() string {
	return usr.Password
}

func (usr *DefaultUser) ClearPassword() {
	usr.Password = ""
}

func (usr *DefaultUser) setPassword(password string) {
	usr.Password = password
}
func (usr *DefaultUser) GetRoles() ([]string, error) {
	return usr.Roles, nil
}

func (usr *DefaultUser) setRoles(roles []string) error {
	usr.Roles = roles
	return nil
}

func (usr *DefaultUser) GetRealm() string {
	return usr.Realm
}

func (usr *DefaultUser) GetPermissions() (permissions []string, err error) {
	return usr.Permissions, nil
}

func (usr *DefaultUser) SetPermissions(permissions []string) {
	usr.Permissions = permissions
}
func (usr *DefaultUser) addRole(role string) error {
	usr.Roles = append(usr.Roles, role)
	return nil
}
func (usr *DefaultUser) removeRole(role string) error {
	usr.Roles = utils.Remove(usr.Roles, role)
	return nil
}

func (usr *DefaultUser) PopulateJWTToken(token *jwt.Token) {
	token.Claims["Roles"] = strings.Join(usr.Roles, ",")
	token.Claims["Name"] = usr.Name
	token.Claims["Picture"] = usr.Picture
}

func (usr *DefaultUser) GetEmail() string {
	return usr.Email
}
func (usr *DefaultUser) GetName() string {
	return usr.Name
}
func (usr *DefaultUser) GetUserName() string {
	return usr.Username
}

func (usr *DefaultUser) GetPicture() string {
	return usr.Picture
}
func (usr *DefaultUser) GetGender() string {
	return usr.Gender
}

func (usr *DefaultUser) LoadJWTClaims(token *jwt.Token) {
	usr.setId(token.Claims["UserId"].(string))
	usr.Name = token.Claims["Name"].(string)
	usr.Picture = token.Claims["Picture"].(string)
	rolesInt := token.Claims["Roles"]
	if rolesInt != nil {
		usr.Roles = strings.Split(rolesInt.(string), ",")
	}
}

func (usr *DefaultUser) encryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.setPassword(string(hash))
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
