package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"laatoosdk/data"
)

type RbacUser interface {
	GetId() string
	SetId(string)
	GetIdField() string
	GetRoles() ([]string, error)
	SetRoles(roles []string) error
	GetPermissions() (permissions []string, err error)
	LoadPermissions(data.DataService) error
	AddRole(role string) error
	RemoveRole(role string) error
	SetJWTClaims(*jwt.Token)
	LoadJWTClaims(*jwt.Token)
}
