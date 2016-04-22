package server

import (
	"laatoo/sdk/core"
)

type SecurityHandler interface {
	core.ServerElement
	HasPermission(core.RequestContext, string) bool
	GetRolePermissions(ctx core.RequestContext, role []string) ([]string, bool)
	AuthenticateRequest(ctx core.RequestContext) error
}
