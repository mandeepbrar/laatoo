package elements

import (
	"laatoo/sdk/server/core"
)

type Solution interface {
	core.ServerElement
	GetPeers(filter string) ([]string, error)
}
