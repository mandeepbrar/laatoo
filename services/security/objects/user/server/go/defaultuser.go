package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	uc = &data.StorableConfig{
		PreSave:    true,
		PostSave:   true,
		PostLoad:   true,
		Auditable:  true,
		Collection: "User",
		Cacheable:  true,
	}
)

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
	data.Storable `laatoo:"softdelete, auditable, multitenant"`
	Username      string           `json:"Username" form:"Username" bson:"Username"`
	Password      string           `json:"Password" form:"Password" bson:"Password"`
	Roles         []string         `json:"Roles" bson:"Roles"`
	Permissions   []string         `json:"Permissions" bson:"Permissions"`
	Email         string           `json:"Email" bson:"Email"`
	Name          string           `json:"Name" bson:"Name"`
	Picture       string           `json:"Picture" bson:"Picture"`
	Account       data.StorableRef `json:"Account" bson:"Account"`
	Realm         string           `json:"Realm" bson:"Realm"`
	Status        int              `json:"Status" bson:"Status"`
}

//Creates object
func (usr *DefaultUser) Initialize(ctx ctx.Context, conf config.Config) error {
	usr.Storable.Initialize(ctx, conf)
	if conf != nil {
		id, ok := conf.GetString(ctx, "Id")
		if ok {
			usr.SetId(id)
		}
		username, ok := conf.GetString(ctx, "Username")
		if ok {
			usr.Username = username
		}
		roles, ok := conf.GetStringArray(ctx, "Roles")
		if ok {
			usr.SetRoles(roles)
		}
		realm, ok := conf.GetString(ctx, "Realm")
		if ok {
			usr.Realm = realm
		}
		picture, ok := conf.GetString(ctx, "Picture")
		if ok {
			usr.Picture = picture
		}
		password, ok := conf.GetString(ctx, "Password")
		if ok {
			usr.Password = password
		}
		email, ok := conf.GetString(ctx, "Email")
		if ok {
			usr.Email = email
		}
		name, ok := conf.GetString(ctx, "Name")
		if ok {
			usr.Name = name
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
	ent.Storable.PreSave(ctx)
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

func (usr *DefaultUser) setPassword(password string) {
	usr.Password = password
}
func (usr *DefaultUser) GetRoles() ([]string, error) {
	return usr.Roles, nil
}

func (usr *DefaultUser) SetRoles(roles []string) error {
	usr.Roles = roles
	return nil
}

func (usr *DefaultUser) GetRealm() string {
	return usr.Realm
}

func (usr *DefaultUser) SetAccount(data data.StorableRef) error {
	usr.Account = data
	return nil
}

func (usr *DefaultUser) GetAccount() data.StorableRef {
	return usr.Account
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

func (usr *DefaultUser) GetStatus() int {
	return usr.Status
}

func (usr *DefaultUser) SetStatus(val int) {
	usr.Status = val
}

func (usr *DefaultUser) GetTenant() string {
	return usr.Storable.GetTenant()
}

func (usr *DefaultUser) LoadClaims(claims map[string]interface{}) {
	usr.SetId(claims["UserId"].(string))
	usrName, ok := claims["UserName"]
	if ok {
		usr.Username = usrName.(string)
	}
	name, ok := claims["Name"]
	if ok {
		usr.Name = name.(string)
	}
	picture, ok := claims["Picture"]
	if ok {
		usr.Picture = picture.(string)
	}
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

func (ent *DefaultUser) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

	if err = rdr.ReadString(c, cdc, "Username", &ent.Username); err != nil {
		return err
	}

	if err = rdr.ReadString(c, cdc, "Password", &ent.Password); err != nil {
		return err
	}

	if err = rdr.ReadArray(c, cdc, "Roles", &ent.Roles); err != nil {
		return err
	}

	if err = rdr.ReadArray(c, cdc, "Permissions", &ent.Permissions); err != nil {
		return err
	}

	if err = rdr.ReadString(c, cdc, "Email", &ent.Email); err != nil {
		return err
	}

	if err = rdr.ReadString(c, cdc, "Name", &ent.Name); err != nil {
		return err
	}

	if err = rdr.ReadString(c, cdc, "Picture", &ent.Picture); err != nil {
		return err
	}

	if err = rdr.ReadString(c, cdc, "Picture", &ent.Picture); err != nil {
		return err
	}

	if err = rdr.ReadString(c, cdc, "Realm", &ent.Realm); err != nil {
		return err
	}

	if err = rdr.ReadInt(c, cdc, "Status", &ent.Status); err != nil {
		return err
	}

	ent.Account = data.StorableRef{}
	if err = rdr.ReadObject(c, cdc, "Account", &ent.Account); err != nil {
		return err
	}

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *DefaultUser) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

	if err = wtr.WriteString(c, cdc, "Username", &ent.Username); err != nil {
		return err
	}

	if err = wtr.WriteArray(c, cdc, "Roles", &ent.Roles); err != nil {
		return err
	}

	if err = wtr.WriteArray(c, cdc, "Permissions", &ent.Permissions); err != nil {
		return err
	}

	if err = wtr.WriteString(c, cdc, "Email", &ent.Email); err != nil {
		return err
	}

	if err = wtr.WriteString(c, cdc, "Name", &ent.Name); err != nil {
		return err
	}

	if err = wtr.WriteString(c, cdc, "Picture", &ent.Picture); err != nil {
		return err
	}

	if err = wtr.WriteString(c, cdc, "Picture", &ent.Picture); err != nil {
		return err
	}

	if err = wtr.WriteString(c, cdc, "Realm", &ent.Realm); err != nil {
		return err
	}

	if err = wtr.WriteInt(c, cdc, "Status", &ent.Status); err != nil {
		return err
	}

	if err = wtr.WriteObject(c, cdc, "Account", &ent.Account); err != nil {
		return err
	}

	return ent.Storable.WriteAll(c, cdc, wtr)
}
