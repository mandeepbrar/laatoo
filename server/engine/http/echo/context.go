package echo

import (
	"io"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"net/http"

	"github.com/labstack/echo"
)

type EchoContext struct {
	baseCtx echo.Context
}

func (echctx *EchoContext) JSON(status int, data interface{}) error {
	return echctx.baseCtx.JSON(status, data)
}
func (echctx *EchoContext) NoContent(status int) error {
	return echctx.baseCtx.NoContent(status)
}
func (echctx *EchoContext) File(file string) error {
	return echctx.baseCtx.File(file)
}
func (echctx *EchoContext) SetHeader(headerName string, headerVal string) {
	echctx.baseCtx.Response().Header().Set(headerName, headerVal)
}
func (echctx *EchoContext) SetCookie(cookie *http.Cookie) {
	echctx.SetCookie(cookie)
}

func (echctx *EchoContext) WriteHeader(status int) {
	echctx.baseCtx.Response().WriteHeader(status)
}
func (echctx *EchoContext) Write(bytes []byte) (int, error) {
	return echctx.baseCtx.Response().Write(bytes)
}
func (echctx *EchoContext) Redirect(status int, path string) error {
	return echctx.baseCtx.Redirect(status, path)
}
func (echctx *EchoContext) GetHeader(header string) string {
	return echctx.baseCtx.Request().Header.Get(header)
}
func (echctx *EchoContext) GetRouteParamNames() []string {
	return echctx.baseCtx.ParamNames()
}
func (echctx *EchoContext) GetRouteParam(paramname string) string {
	if paramname == "__0" {
		return echctx.baseCtx.ParamValues()[0]
	}
	return echctx.baseCtx.Param(paramname)
}
func (echctx *EchoContext) GetRouteParamByIndex(index int) string {
	return echctx.baseCtx.ParamValues()[index]
}
func (echctx *EchoContext) GetQueryParams() map[string][]string {
	return echctx.baseCtx.QueryParams()
}
func (echctx *EchoContext) GetQueryParam(paramname string) string {
	return echctx.baseCtx.QueryParam(paramname)
}
func (echctx *EchoContext) Bind(data interface{}) error {
	return echctx.baseCtx.Bind(data)
}
func (echctx *EchoContext) GetRequestStream() (io.Reader, error) {
	return echctx.baseCtx.Request().Body, nil
}
func (echctx *EchoContext) GetRequest() *http.Request {
	return echctx.baseCtx.Request()
}

func (echctx *EchoContext) GetBody() ([]byte, error) {
	return ioutil.ReadAll(echctx.baseCtx.Request().Body)
}
func (echctx *EchoContext) GetFiles() (map[string]*core.MultipartFile, error) {
	form, err := echctx.baseCtx.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := make(map[string]*core.MultipartFile, len(form.File))
	for _, headers := range form.File {
		fil := &core.MultipartFile{}
		mpfile, err := headers[0].Open()
		if err != nil {
			return nil, err
		}
		fil.File = mpfile
		fil.FileName = headers[0].Filename
		fil.MimeType = headers[0].Header.Get("Content-Type")
		files[fil.FileName] = fil
	}
	return files, nil
}
