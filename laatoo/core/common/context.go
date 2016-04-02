package common

import (
	"fmt"
	"github.com/twinj/uuid"
)

type Context struct {
	Id   string
	Name string
}

func NewContext(name string) *Context {
	return &Context{Name: name, Id: uuid.NewV4().String()}
}

func (ctx *Context) GetId() string {
	return ctx.Id
}

func (ctx *Context) GetName() string {
	return ctx.Name
}

func (ctx *Context) DupCtx(name string) *Context {
	return &Context{Name: fmt.Sprintf("%s>%s", ctx.Name, name)}
}
