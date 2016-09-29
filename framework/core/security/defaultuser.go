package security

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	uc = &data.StorableConfig{
		IdField:         "Id",
		Type:            config.DEFAULT_USER,
		SoftDeleteField: "Deleted",
		PreSave:         true,
		PostSave:        false,
		PostLoad:        true,
		Auditable:       true,
		Collection:      "User",
		Cacheable:       true,
		NotifyNew:       false,
		NotifyUpdates:   false,
	}
)

func init() {
	objects.Register(config.DEFAULT_USER, DefaultUser{})
}

/*
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



//Creates collection
func (rf *UserFactory) CreateObjectCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	usercollection := make([]DefaultUser, length)
	return &usercollection, nil
}*/

type DefaultUser struct {
	data.SoftDeleteAuditable `bson:",inline"`
	Username                 string   `json:"Username" form:"Username" bson:"Username"`
	Password                 string   `json:"Password" form:"Password" bson:"Password"`
	Roles                    []string `json:"Roles" bson:"Roles"`
	Permissions              []string `json:"Permissions" bson:"Permissions"`
	Email                    string   `json:"Email" bson:"Email"`
	Name                     string   `json:"Name" bson:"Name"`
	Picture                  string   `json:"Picture" bson:"Picture"`
	Realm                    string   `json:"Realm" bson:"Realm"`
}

//Creates object
func (usr *DefaultUser) Init(ctx core.Context, args core.MethodArgs) error {
	usr.SoftDeleteAuditable.Init(ctx, args)
	if args != nil {
		id, ok := args["Id"]
		if ok {
			usr.SetId(id.(string))
		}
		username, ok := args["Username"]
		if ok {
			usr.Username = username.(string)
		}
		roles, ok := args["Roles"]
		if ok {
			usr.setRoles(roles.([]string))
		}
		realm, ok := args["Realm"]
		if ok {
			usr.Realm = realm.(string)
		}
		picture, ok := args["Picture"]
		if ok {
			usr.Picture = picture.(string)
		}
		password, ok := args["Password"]
		if ok {
			usr.Password = password.(string)
		}
		email, ok := args["Email"]
		if ok {
			usr.Email = email.(string)
		}
		name, ok := args["Name"]
		if ok {
			usr.Name = name.(string)
		}

	}
	return nil
}

func (r *DefaultUser) Config() *data.StorableConfig {
	return uc
}

func (usr *DefaultUser) GetUsernameField() string {
	return "Username"
}
func (ent *DefaultUser) PreSave(ctx core.RequestContext) error {
	ent.SoftDeleteAuditable.PreSave(ctx)
	err := ent.encryptPassword()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (ent *DefaultUser) PostLoad(ctx core.RequestContext) error {
	ent.Password = ""
	return nil
}
func (usr *DefaultUser) GetPassword() string {
	return usr.Password
}

func (usr *DefaultUser) ClearPassword() {
	usr.Password = ""
}

func (ent *DefaultUser) IsDeleted() bool {
	return ent.Deleted
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

func (usr *DefaultUser) PopulateClaims(claims map[string]interface{}) {
	claims["Roles"] = strings.Join(usr.Roles, ",")
	claims["UserName"] = usr.Username
	claims["Name"] = usr.Name
	claims["Picture"] = usr.Picture
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

func (usr *DefaultUser) LoadClaims(claims map[string]interface{}) {
	usr.SetId(claims["UserId"].(string))
	usr.Username = claims["UserName"].(string)
	usr.Name = claims["Name"].(string)
	usr.Picture = claims["Picture"].(string)
	rolesInt := claims["Roles"]
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
