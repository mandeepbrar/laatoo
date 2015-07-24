package user

import (
	jwt "github.com/dgrijalva/jwt-go"
	"laatoosdk/data"
	"laatoosdk/utils"
)

type RbacUser interface {
	GetId() string
	SetId(string)
	GetIdField() string
	GetRoles() (utils.StringSet, error)
	SetRoles(roles utils.StringSet) error
	GetPermissions() (permissions utils.StringSet, err error)
	LoadPermissions(data.DataService) error
	AddRole(role string) error
	RemoveRole(role string) error
	SetJWTClaims(*jwt.Token)
	LoadJWTClaims(*jwt.Token)
}
