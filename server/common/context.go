package common

import (
	"fmt"
	"laatoo/sdk/core"
	"log"
	"net/http"
	"strconv"
	"time"

	glctx "golang.org/x/net/context"

	"github.com/twinj/uuid"
)

type Context struct {
	gaeReq       *http.Request
	appengineCtx glctx.Context
	Id           string
	Name         string
	Path         string
	ParamsStore  map[string]interface{}
	Parent       *Context
	creationTime time.Time
}

func NewContext(name string) *Context {
	return &Context{Name: name, Path: name, Id: uuid.NewV1().String(), ParamsStore: make(map[string]interface{}), creationTime: time.Now()}
}

func (ctx *Context) GetId() string {
	return ctx.Id
}

func (ctx *Context) GetParent() core.Context {
	return ctx.Parent
}

func (ctx *Context) GetPath() string {
	return ctx.Path
}

func (ctx *Context) GetName() string {
	return ctx.Name
}

func (ctx *Context) SetName(name string) {
	ctx.Name = name
}
func (ctx *Context) SetGaeReq(req *http.Request) {
	ctx.gaeReq = req
	ctx.appengineCtx = GetAppengineContext(ctx)
}

func (ctx *Context) GetCreationTime() time.Time {
	return ctx.creationTime
}

//completes a request
func (ctx *Context) GetElapsedTime() time.Duration {
	return time.Now().Sub(ctx.creationTime)
}

func (ctx *Context) SubCtx(name string) core.Context {
	return &Context{Name: name, Path: fmt.Sprintf("%s  -> @%s", ctx.Path, name), Parent: ctx, ParamsStore: ctx.ParamsStore, Id: ctx.Id,
		creationTime: ctx.creationTime, gaeReq: ctx.gaeReq, appengineCtx: ctx.appengineCtx}
}

func (ctx *Context) NewCtx(name string) core.Context {
	duplicateMap := make(map[string]interface{}, len(ctx.ParamsStore))
	for k, v := range ctx.ParamsStore {
		duplicateMap[k] = v
	}
	return &Context{Name: name, Path: name, Parent: ctx, ParamsStore: duplicateMap, Id: uuid.NewV1().String(),
		creationTime: time.Now(), gaeReq: ctx.gaeReq, appengineCtx: ctx.appengineCtx}
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	val, ok := ctx.ParamsStore[key]
	return val, ok
}

func (ctx *Context) GetString(key string) (string, bool) {
	valInt, ok := ctx.Get(key)
	if ok {
		val, ok := valInt.(string)
		if ok {
			return val, true
		}
	}
	return "", false
}
func (ctx *Context) GetBool(key string) (bool, bool) {
	val, found := ctx.Get(key)
	if found {
		boolval, ok := val.(bool)
		if ok {
			return boolval, true
		}
		strval, ok := val.(string)
		if ok {
			boolval, err := strconv.ParseBool(strval)
			if err == nil {
				return boolval, true
			}
		}
	}
	return false, false
}

func (ctx *Context) GetInt(key string) (int, bool) {
	valInt, ok := ctx.Get(key)
	if ok {
		val, ok := valInt.(string)
		if ok {
			intVal, err := strconv.Atoi(val)
			if err == nil {
				return intVal, true
			}
		}
	}
	return -1, false
}
func (ctx *Context) GetStringArray(key string) ([]string, bool) {
	valInt, ok := ctx.Get(key)
	if ok {
		val, ok := valInt.([]string)
		if ok {
			return val, true
		}
	}
	return nil, false
}

func (ctx *Context) Set(key string, val interface{}) {
	ctx.ParamsStore[key] = val
}
func (ctx *Context) GetAppengineContext() glctx.Context {
	return ctx.appengineCtx
}
func (ctx *Context) HttpClient() *http.Client {
	return HttpClient(ctx)
}
func (ctx *Context) GetOAuthContext() glctx.Context {
	return GetOAuthContext(ctx)
}
func (ctx *Context) LogTrace(msg string, args ...interface{}) {
	log.Println(msg)
}
func (ctx *Context) LogDebug(msg string, args ...interface{}) {
	log.Println(msg)
}
func (ctx *Context) LogInfo(msg string, args ...interface{}) {
	log.Println(msg)
}
func (ctx *Context) LogWarn(msg string, args ...interface{}) {
	log.Println(msg)
}
func (ctx *Context) LogError(msg string, args ...interface{}) {
	log.Println(msg)
}
func (ctx *Context) LogFatal(msg string, args ...interface{}) {
	log.Println(msg)
}
