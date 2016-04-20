package server

import (
	"laatoo/sdk/core"
)

type SecurityHandler interface {
	core.ServerElement
	HasPermission(core.RequestContext, string) bool
	GetRolePermissions(ctx core.RequestContext, role []string) ([]string, bool)
	CreateSystemRequest(ctx core.ServerContext, name string) core.RequestContext
	AuthenticateRequest(ctx core.RequestContext) error
}
