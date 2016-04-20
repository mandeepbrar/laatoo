package common

import (
	"fmt"
	"github.com/twinj/uuid"
	"laatoo/sdk/core"
	"strconv"
)

type Context struct {
	Id          string
	Name        string
	ParamsStore map[string]interface{}
	Parent      *Context
}

func NewContext(name string) *Context {
	return &Context{Name: name, Id: uuid.NewV4().String(), ParamsStore: make(map[string]interface{})}
}

func (ctx *Context) GetId() string {
	return ctx.Id
}

func (ctx *Context) GetParent() core.Context {
	return ctx.Parent
}

func (ctx *Context) GetName() string {
	return ctx.Name
}

func (ctx *Context) SetName(name string) {
	ctx.Name = name
}

func (ctx *Context) SubCtx(name string) core.Context {
	return &Context{Name: fmt.Sprintf("%s>%s", ctx.Name, name), Parent: ctx, ParamsStore: ctx.ParamsStore, Id: ctx.Id}
}

func (ctx *Context) NewCtx(name string) core.Context {
	duplicateMap := make(map[string]interface{}, len(ctx.ParamsStore))
	for k, v := range ctx.ParamsStore {
		duplicateMap[k] = v
	}
	return &Context{Name: fmt.Sprintf("%s:%s", ctx.Name, name), Parent: ctx, ParamsStore: duplicateMap, Id: uuid.NewV4().String()}
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
