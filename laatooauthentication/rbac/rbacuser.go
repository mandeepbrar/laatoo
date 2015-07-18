package ginauth_rbac

import (
	"laatoo/commonobjects"
)

type RbacUser interface {
	GetId() string
	GetRoles() (commonobjects.StringSet, error)
	SetRoles(roles commonobjects.StringSet) error
	GetPermissions() (permissions commonobjects.StringSet, err error)
	SetPermissions(permissions commonobjects.StringSet) error
	LoadPermissions(commonobjects.Storer) error
	AddRole(role string) error
	RemoveRole(role string) error
}
