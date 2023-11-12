package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type ActivityManager interface {
	core.ServerElement
	components.ActivityManager
}
