package gin

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type GinContext struct {
	baseCtx *gin.Context
}

func (ginctx *GinContext) JSON(status int, data interface{}) error {
	ginctx.baseCtx.JSON(status, data)
	return nil
}
func (ginctx *GinContext) NoContent(status int) error {
	ginctx.baseCtx.Status(status)
	return nil
}
func (ginctx *GinContext) File(file string) error {
	ginctx.baseCtx.File(file)
	return nil
}
func (ginctx *GinContext) SetHeader(headerName string, headerVal string) {
	ginctx.baseCtx.Header(headerName, headerVal)
}

func (ginctx *GinContext) Write(bytes []byte) (int, error) {
	return ginctx.baseCtx.Writer.Write(bytes)
}
func (ginctx *GinContext) Redirect(status int, path string) error {
	ginctx.baseCtx.Redirect(status, path)
	return nil
}
func (ginctx *GinContext) GetHeader(header string) string {
	return ginctx.baseCtx.Request.Header.Get(header)
}
func (ginctx *GinContext) GetRouteParamNames() []string {
	length := len(ginctx.baseCtx.Params)
	paramnames := make([]string, length)
	for i := 0; i < length; i++ {
		paramnames[i] = ginctx.baseCtx.Params[i].Key
	}
	return paramnames
}
func (ginctx *GinContext) GetRouteParam(paramname string) string {
	return ginctx.baseCtx.Param(paramname)
}
func (ginctx *GinContext) GetRouteParamByIndex(index int) string {
	return ginctx.baseCtx.Params[index].Value
}
func (ginctx *GinContext) GetQueryParams() map[string][]string {
	return ginctx.baseCtx.Request.URL.Query()
}
func (ginctx *GinContext) GetQueryParam(paramname string) string {
	return ginctx.baseCtx.Query(paramname)
}
func (ginctx *GinContext) Bind(data interface{}) error {
	return ginctx.baseCtx.Bind(data)
}
func (ginctx *GinContext) GetBody() ([]byte, error) {
	return ioutil.ReadAll(ginctx.baseCtx.Request.Body)
}
