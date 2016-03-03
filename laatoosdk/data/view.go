package data

import (
	"laatoosdk/core"
)

const (
	VIEW_PAGENUM      = "pagenum"
	VIEW_PAGESIZE     = "pagesize"
	VIEW_RECSRETURNED = "records"
	VIEW_TOTALRECS    = "totalrecords"
)

type View interface {
	Execute(core.Context, DataService) error
}
