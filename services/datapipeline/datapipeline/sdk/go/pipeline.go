package datapipeline

import (
	"laatoo/sdk/server/core"
)

type PipelineRecord struct {
	Err             error
	Input           interface{}
	TransformedData interface{}
	PipelineData    interface{}
}

func NewPipelineRecord(input interface{}, err error) *PipelineRecord {
	rec := &PipelineRecord{Input: input, Err: err}
	return rec
}

type DataChan chan *PipelineRecord

//type DataChan chan interface{}

type Importer interface {
	GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan DataChan) error
}

type Exporter interface {
	WriteRecord(ctx core.RequestContext, initData map[string]interface{}, inputDataChan DataChan, outputDataChan DataChan) error
}

type ErrorProcessor interface {
	ProcessErrorRecord(ctx core.RequestContext, rec *PipelineRecord)
}

type Processor interface {
	Transform(ctx core.RequestContext, input interface{}) (interface{}, error)
}

type Driver interface {
	Run(ctx core.RequestContext, importer Importer, exporter Exporter, processor Processor, errorProcessor ErrorProcessor, initData map[string]interface{}) error
}
