package common

import (
	googleContext "context"
	"fmt"
	laatooContext "laatoo/sdk/server/ctx"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/twinj/uuid"
)

type Context struct {
	googleContext.Context
	gaeReq       *http.Request
	appengineCtx googleContext.Context
	Name         string
	Path         string
	ParamsStore  map[string]interface{}
	Parent       *Context
	creationTime time.Time
}

func NewContext(name string) *Context {
	return &Context{Context: googleContext.WithValue(googleContext.Background(), "tId", uuid.NewV1().String()), Name: name, Path: name, ParamsStore: make(map[string]interface{}), creationTime: time.Now()}
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.Context.Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.Context.Done()
}

func (ctx *Context) Err() error {
	return ctx.Context.Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.Context.Value(key)
}

func (ctx *Context) WithCancel() (laatooContext.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithCancel(ctx)
	return ctx.withNewContext(newgooglectx), cancelFunc
}

func (ctx *Context) WithDeadline(timeout time.Time) (laatooContext.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithDeadline(ctx, timeout)
	return ctx.withNewContext(newgooglectx), cancelFunc
}

func (ctx *Context) WithTimeout(timeout time.Duration) (laatooContext.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithTimeout(ctx, timeout)
	return ctx.withNewContext(newgooglectx), cancelFunc
}

func (ctx *Context) WithValue(key, val interface{}) laatooContext.Context {
	newgooglectx := googleContext.WithValue(ctx, key, val)
	return ctx.withNewContext(newgooglectx)
}

func (ctx *Context) WithContext(parent googleContext.Context) laatooContext.Context {
	return ctx.withNewContext(parent)
}

func (ctx *Context) GetId() string {
	return ctx.Value("tId").(string)
}

func (ctx *Context) GetParent() laatooContext.Context {
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

func (ctx *Context) SubCtx(name string) laatooContext.Context {
	return ctx.subCtx(name, ctx.Context)
}

func (ctx *Context) withNewContext(googleCtx googleContext.Context) *Context {
	return ctx.subCtx(ctx.Name, googleCtx)
}

func (ctx *Context) subCtx(name string, googleCtx googleContext.Context) *Context {
	return &Context{Context: googleCtx, Name: name, Path: fmt.Sprintf("%s  -> %s", ctx.Path, name), Parent: ctx, ParamsStore: ctx.ParamsStore,
		creationTime: ctx.creationTime, gaeReq: ctx.gaeReq, appengineCtx: ctx.appengineCtx}
}

func (ctx *Context) NewCtx(name string) laatooContext.Context {
	duplicateMap := make(map[string]interface{}, len(ctx.ParamsStore))
	for k, v := range ctx.ParamsStore {
		duplicateMap[k] = v
	}
	return &Context{Context: googleContext.WithValue(googleContext.Background(), "tId", uuid.NewV1().String()), Name: name, Path: fmt.Sprintf("%s  :: @@%s", ctx.Path, name),
		Parent: ctx, ParamsStore: duplicateMap, creationTime: time.Now(), gaeReq: ctx.gaeReq, appengineCtx: ctx.appengineCtx}
}

func (ctx *Context) CompleteContext() {
	log.Println(fmt.Sprintf("Context complete %s\n Time elapsed %s\n", ctx.Path, ctx.GetElapsedTime()))
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

func (ctx *Context) SetVals(vals map[string]interface{}) {
	if vals != nil {
		for k, v := range vals {
			ctx.Set(k, v)
		}
	}
}

func (ctx *Context) GetAppengineContext() googleContext.Context {
	return ctx.appengineCtx
}
func (ctx *Context) HttpClient() *http.Client {
	return HttpClient(ctx)
}
func (ctx *Context) GetOAuthContext() googleContext.Context {
	return GetOAuthContext(ctx)
}
func (ctx *Context) Dump() {
	log.Println(ctx.Path, "---->")
	for k, v := range ctx.ParamsStore {
		log.Println(k, "  ", v)
	}
}
func (ctx *Context) LogTrace(msg string, args ...interface{}) {
	log.Println(msg, fmt.Sprint(args))
}
func (ctx *Context) LogDebug(msg string, args ...interface{}) {
	log.Println(msg, fmt.Sprint(args))
}
func (ctx *Context) LogInfo(msg string, args ...interface{}) {
	log.Println(msg, fmt.Sprint(args))
}
func (ctx *Context) LogWarn(msg string, args ...interface{}) {
	log.Println(msg, fmt.Sprint(args))
}
func (ctx *Context) LogError(msg string, args ...interface{}) {
	log.Println(msg, fmt.Sprint(args))
}
func (ctx *Context) LogFatal(msg string, args ...interface{}) {
	log.Println(msg, fmt.Sprint(args))
}
