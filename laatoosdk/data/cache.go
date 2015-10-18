package data

import (
	"github.com/labstack/echo"
)

type Cache interface {
	PutObject(ctx *echo.Context, key string, item interface{}) error
	GetObject(ctx *echo.Context, key string, val interface{}) error
	GetMulti(ctx *echo.Context, keys []string, val map[string]interface{}) error
	Delete(ctx *echo.Context, key string) error
}
