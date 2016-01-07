package data

import (
	"github.com/labstack/echo"
)

//Object stored by data service
type Storable interface {
	GetId() string
	SetId(string)
	PreSave(ctx *echo.Context) error
	PostSave(ctx *echo.Context) error
	PostLoad(ctx *echo.Context) error
	GetIdField() string
}

//Factory function for creating storable
type StorableCreator func() interface{}
