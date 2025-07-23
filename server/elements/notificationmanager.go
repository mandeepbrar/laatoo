package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type NotificationManager interface {
	core.ServerElement
	components.NotificationManager
}
