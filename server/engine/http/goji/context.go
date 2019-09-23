package goji

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"net/http"

	"goji.io/pat"
)

type GojiContext struct {
	baseCtx context.Context
	req     *http.Request
	res     http.ResponseWriter
}

func (gojictx *GojiContext) JSON(status int, data interface{}) error {
	gojictx.res.WriteHeader(status)
	if data != nil {
		encoder := json.NewEncoder(gojictx.res)
		return encoder.Encode(data)
	}
	return nil
}
func (gojictx *GojiContext) NoContent(status int) error {
	gojictx.res.WriteHeader(status)
	return nil
}
func (gojictx *GojiContext) File(file string) error {
	http.ServeFile(gojictx.res, gojictx.req, file)
	return nil
}
func (gojictx *GojiContext) SetHeader(headerName string, headerVal string) {
	gojictx.res.Header().Add(headerName, headerVal)
}

func (gojictx *GojiContext) SetCookie(cookie *http.Cookie) {
	gojictx.res.SetCookie(cookie)
}


func (gojictx *GojiContext) Write(bytes []byte) (int, error) {
	return gojictx.res.Write(bytes)
}

func (gojictx *GojiContext) Redirect(status int, path string) error {
	http.Redirect(gojictx.res, gojictx.req, path, status)
	return nil
}
func (gojictx *GojiContext) GetHeader(header string) string {
	return gojictx.req.Header.Get(header)
}

func (gojictx *GojiContext) GetRouteParamNames() []string {
	/*all := gojictx.baseCtx.Value(pattern.AllVariables).(map[pattern.Variable]interface{})
	length := len(all)
	paramnames := make([]string, length)
	i := 0
	for k, v := range all {
		paramnames[i] = fmt.Sprint(k)
		i++
	}
	length := len(ginctx.baseCtx.Params)
	paramnames := make([]string, length)
	for i := 0; i < length; i++ {
		paramnames[i] = ginctx.baseCtx.Params[i].Key
	}
	return paramnames*/
	return []string{}
}
func (gojictx *GojiContext) GetRouteParam(paramname string) string {
	return pat.Param(gojictx.baseCtx, paramname)
}

/*
func (gojictx *GojiContext) GetRouteParamByIndex(index int) string {
	return "" //ginctx.baseCtx.Params[index].Value
}*/
func (gojictx *GojiContext) GetQueryParams() map[string][]string {
	return gojictx.req.URL.Query()
}
func (gojictx *GojiContext) GetQueryParam(paramname string) string {
	return gojictx.req.URL.Query().Get(paramname)
}
func (gojictx *GojiContext) Bind(data interface{}) error {
	dec := json.NewDecoder(gojictx.req.Body)
	return dec.Decode(data)
}
func (gojictx *GojiContext) GetBody() ([]byte, error) {
	return ioutil.ReadAll(gojictx.req.Body)
}
func (gojictx *GojiContext) GetRequestStream() (io.Reader, error) {
	return gojictx.req.Body, nil
}
func (gojictx *GojiContext) GetRequest() *http.Request {
	return gojictx.req
}

func (gojictx *GojiContext) GetFiles() (map[string]*core.MultipartFile, error) {
	err := gojictx.req.ParseMultipartForm(2000000000)
	if err != nil {
		return nil, err
	}
	form := gojictx.req.MultipartForm
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
