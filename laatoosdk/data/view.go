package data

import (
	"github.com/labstack/echo"
)

const (
	VIEW_PAGENUM      = "pagenum"
	VIEW_PAGESIZE     = "pagesize"
	VIEW_RECSRETURNED = "records"
	VIEW_TOTALRECS    = "totalrecords"
)

type View interface {
	Execute(*echo.Context, DataService) error
}
