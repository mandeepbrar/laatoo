package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type mapToObjectProcessor struct {
}

func mapToObjectFactory(core.ServerContext) datapipeline.Processor {
	return &mapToObjectProcessor{}
}
func (proc *mapToObjectProcessor) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (proc *mapToObjectProcessor) Transform(ctx core.RequestContext, input interface{}) (interface{}, error) {
	return nil, nil
}
