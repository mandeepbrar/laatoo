package elements

import (
	"laatoo/sdk/server/core"
)

type Isolation interface {
	core.ServerElement
	GetIsolationId() string
	//GetApplet(name string) (Applet, bool)
}
