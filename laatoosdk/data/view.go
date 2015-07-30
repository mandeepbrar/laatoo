package data

import (
	"github.com/labstack/echo"
)

type View interface {
	Execute(DataService, *echo.Context) error
}
