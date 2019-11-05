package main

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

type HelloWorld struct {
	core.Service
}

func (hw *HelloWorld) Invoke(ctx core.RequestContext) error {
	inter, ok := ctx.GetParamValue("Data")
	log.Error(ctx, "Hello world invoked", "data", inter, "found", ok)
	return nil
}
