package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
)

type logErrorsProcessor struct {
}

func logErrorProcFactory(core.ServerContext) datapipeline.ErrorProcessor {
	return &logErrorsProcessor{}
}

func (proc *logErrorsProcessor) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (proc *logErrorsProcessor) ProcessErrorRecord(ctx core.RequestContext, err *datapipeline.PipelineErrorRecord) {
	log.Error(ctx, "Error encountered in processing record", "error", err)
	return
}
