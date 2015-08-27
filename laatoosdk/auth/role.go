package auth

import (
	"laatoosdk/utils"
)

type Role interface {
	GetId() string
	SetId(string)
	GetIdField() string
	GetPermissions() utils.StringSet
}
