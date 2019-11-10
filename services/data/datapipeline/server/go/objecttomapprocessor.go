package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type objectToMapProcessor struct {
}

func objecToMapFactory(core.ServerContext) datapipeline.Processor {
	return &objectToMapProcessor{}
}
func (proc *objectToMapProcessor) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (proc *objectToMapProcessor) Transform(ctx core.RequestContext, input interface{}) (interface{}, error) {
	return nil, nil
}
