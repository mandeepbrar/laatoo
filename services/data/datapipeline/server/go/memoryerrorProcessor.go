package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type memoryErrorsProcessor struct {
	records []interface{}
}

func memoryErrorsProcessorFactory(core.ServerContext) datapipeline.ErrorProcessor {
	return &memoryErrorsProcessor{make([]interface{}, 0)}
}

func (proc *memoryErrorsProcessor) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (proc *memoryErrorsProcessor) ProcessErrorRecord(ctx core.RequestContext, err *datapipeline.PipelineErrorRecord) {
	proc.records = append(proc.records, err)
	return
}
