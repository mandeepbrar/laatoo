package net

import (
	"io"
	"laatoo/sdk/server/core"
	"net/http"
)

type WebContext interface {
	GetHeader(header string) string
	GetRouteParam(paramname string) string
	GetRouteParamNames() []string
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
	GetRequestStream() (io.Reader, error)
	GetRequest() *http.Request
	GetFiles() (map[string]*core.MultipartFile, error)
}
