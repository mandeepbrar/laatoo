package user

import (
	"laatoosdk/utils"
)

type RbacRole interface {
	GetId() string
	SetId(string)
	GetIdField() string
	GetPermissions() utils.StringSet
}
