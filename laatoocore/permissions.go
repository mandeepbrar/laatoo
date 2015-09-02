package laatoocore

import (
	"laatoosdk/utils"
)

var (
	Permissions = utils.NewStringSet([]string{})
)

//register the object factory in the global register
func RegisterPermissions(perm []string) {
	Permissions.Append(perm)
}

func ListAllPermissions() []string {
	return Permissions.Values()
}
