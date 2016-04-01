package core

import (
	"laatoo/sdk/config"
)

type SecurityHandler interface {
	Initialize(ctx ServerContext, conf config.Config) error
	HasPermission(RequestContext, string) bool
	GetRolePermissions(role []string) ([]string, bool)
}
