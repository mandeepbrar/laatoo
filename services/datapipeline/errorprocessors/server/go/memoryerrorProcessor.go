package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

type memoryErrorsProcessor struct {
	core.Service
	records []*datapipeline.PipelineRecord
}

func (proc *memoryErrorsProcessor) Initialize(ctx core.ServerContext, conf config.Config) error {
	proc.records = make([]*datapipeline.PipelineRecord, 0)
	return nil
}

func (proc *memoryErrorsProcessor) ProcessErrorRecord(ctx core.RequestContext, rec *datapipeline.PipelineRecord) {
	log.Error(ctx, "Received error record in memory", "rec", rec)
	if rec.Err != nil {
		proc.records = append(proc.records, rec)
	}
	return
}

func (proc *memoryErrorsProcessor) GetErrorRecords(ctx core.RequestContext) []*datapipeline.PipelineRecord {
	return proc.records
}

func (proc *memoryErrorsProcessor) GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan) error {
	recordsToServe := proc.records
	log.Error(ctx, "Count of records in memory", "recordsToServe", len(recordsToServe))
	proc.records = make([]*datapipeline.PipelineRecord, 0)

	for _, record := range recordsToServe {
		dataChan <- datapipeline.NewPipelineRecord(record.Input, nil)
	}
	return nil
}
