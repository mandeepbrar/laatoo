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

func (proc *logErrorsProcessor) ProcessErrorRecord(ctx core.RequestContext, input, output interface{}, err error) {
	args := []interface{}{"error", err, "Input", input}
	if output != nil {
		args = append(args, "Output", output)
	}
	log.Error(ctx, "Error encountered in processing record", args...)
	return
}
