package ginauth_rbac

import (
	"github.com/storageutils"
)

type RbacUser interface {
	GetId() string
	GetRoles() (storageutils.StringSet, error)
	SetRoles(roles storageutils.StringSet) error
	GetPermissions() (permissions storageutils.StringSet, err error)
	SetPermissions(permissions storageutils.StringSet) error
	LoadPermissions(storageutils.Storer) error
	AddRole(role string) error
	RemoveRole(role string) error
}
