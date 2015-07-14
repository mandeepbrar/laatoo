package ginauth_rbac

import (
	"github.com/storageutils"
)

const (
	ROLESTORER = "RoleStorer"
)

func (rb *rbac) configStorer() error {
	roleStorerInt, ok := rb.app.Configuration[ROLESTORER]
	if ok {
		rb.RoleStorer = roleStorerInt.(storageutils.Storer)
	} else {
		rb.RoleStorer = storageutils.NewMemoryStorer("Role")
	}
	rb.app.Logger.Info("Using role storer %s", rb.RoleStorer.GetName())
	return nil
}
