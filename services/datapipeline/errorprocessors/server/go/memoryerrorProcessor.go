package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

type MemoryErrorsProcessor struct {
	core.Service
	records []*datapipeline.PipelineRecord
}

func (proc *MemoryErrorsProcessor) Initialize(ctx core.ServerContext, conf config.Config) error {
	proc.records = make([]*datapipeline.PipelineRecord, 0)
	return nil
}

func (proc *MemoryErrorsProcessor) ProcessErrorRecord(ctx core.RequestContext, rec *datapipeline.PipelineRecord) {
	log.Error(ctx, "Received error record in memory", "rec", rec)
	if rec.Err != nil {
		proc.records = append(proc.records, rec)
	}
	return
}

func (proc *MemoryErrorsProcessor) GetErrorRecords(ctx core.RequestContext) []*datapipeline.PipelineRecord {
	return proc.records
}

func (proc *MemoryErrorsProcessor) GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan) error {
	recordsToServe := proc.records
	log.Error(ctx, "Count of records in memory", "recordsToServe", len(recordsToServe))
	proc.records = make([]*datapipeline.PipelineRecord, 0)

	for _, record := range recordsToServe {
		dataChan <- datapipeline.NewPipelineRecord(record.Input, nil)
	}
	return nil
}
