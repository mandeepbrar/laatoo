package jsui

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type UIPlugin interface {
	UILoad(ctx core.ServerContext) map[string]config.Config
	LoadingComplete(ctx core.ServerContext) map[string]config.Config
}
