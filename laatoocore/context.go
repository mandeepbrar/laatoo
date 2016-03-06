package laatoocore

import (
	"github.com/labstack/echo"
	glctx "golang.org/x/net/context"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/errors"
	"net/http"
)

type Context struct {
	Context     *echo.Context
	Conf        map[string]interface{}
	User        auth.User
	environment *Environment
}

func (ctx *Context) Request() *http.Request {
	return ctx.Context.Request()
}

func (ctx *Context) ResponseWriter() http.ResponseWriter {
	return ctx.Context.Response().Writer()
}

func (ctx *Context) SetHeader(key string, val string) {
	ctx.Context.Response().Header().Set(key, val)
}

func (ctx *Context) JSON(code int, val interface{}) error {
	return ctx.Context.JSON(code, val)
}

func (ctx *Context) HTML(code int, format string, val ...interface{}) error {
	return ctx.Context.HTML(code, format, val)
}

func (ctx *Context) Redirect(code int, url string) error {
	return ctx.Context.Redirect(code, url)
}

func (ctx *Context) Get(key string) interface{} {
	return ctx.Context.Get(key)
}

func (ctx *Context) GetConf() map[string]interface{} {
	return ctx.Conf
}

func (ctx *Context) Bind(i interface{}) error {
	return ctx.Context.Bind(i)
}

func (ctx *Context) Set(key string, val interface{}) {
	ctx.Context.Set(key, val)
}

func (ctx *Context) Param(key string) string {
	return ctx.Context.Param(key)
}
func (ctx *Context) ParamByIndex(index int) string {
	return ctx.Context.P(index)
}

func (ctx *Context) Query(key string) string {
	return ctx.Context.Query(key)
}

func (ctx *Context) GetVariable(variable string) interface{} {
	return ctx.environment.GetVariable(variable)
}

func (ctx *Context) GetService(alias string) (core.Service, error) {
	return ctx.environment.GetService(ctx, alias)
}

func (ctx *Context) HasPermission(perm string) bool {
	return ctx.environment.HasPermission(ctx, perm)
}

func (ctx *Context) NoContent(errorcode int) error {
	return ctx.Context.NoContent(errorcode)
}

func (ctx *Context) SubscribeTopic(topic string, handler core.TopicListener) error {
	return ctx.environment.SubscribeTopic(ctx, topic, handler)
}

func (ctx *Context) PublishMessage(topic string, message interface{}) error {
	return ctx.environment.PublishMessage(ctx, topic, message)
}

func (ctx *Context) PutInCache(key string, item interface{}) error {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.PutObject(ctx, key, item)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}

func (ctx *Context) GetFromCache(key string, val interface{}) error {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.GetObject(ctx, key, val)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}

func (ctx *Context) GetMultiFromCache(keys []string, val map[string]interface{}) error {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.GetMulti(ctx, keys, val)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}

func (ctx *Context) DeleteFromCache(key string) error {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.Delete(ctx, key)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}

func (ctx *Context) HttpClient() *http.Client {
	return HttpClient(ctx)
}
func (ctx *Context) GetAppengineContext() glctx.Context {
	return GetAppengineContext(ctx)
}
func (ctx *Context) GetUser() auth.User {
	return ctx.User
}
func (ctx *Context) SetUser(usr auth.User) {
	ctx.User = usr
}

func (ctx *Context) GetCloudContext(scope string) glctx.Context {
	return GetCloudContext(ctx, scope)
}
