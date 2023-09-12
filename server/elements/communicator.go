package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type Communicator interface {
	core.ServerElement
	components.Communicator
}
