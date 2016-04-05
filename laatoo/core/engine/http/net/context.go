package net

import (
	"laatoo/sdk/core"
)

type WebContext interface {
	core.EngineRequestContext
	GetHeader(header string) string
	GetRouteParam(paramname string) string
	GetQueryParams() map[string][]string
	GetQueryParam(paramname string) string
	Bind(data interface{}) error
	GetBody() ([]byte, error)
	JSON(status int, data interface{}) error
	NoContent(status int) error
	File(file string) error
	SetHeader(headerName string, headerVal string)
	Write(bytes []byte) (int, error)
	Redirect(status int, path string) error
}
