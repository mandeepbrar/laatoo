package datapipeline

import "laatoo/sdk/server/core"

type DataChan chan interface{}

type PipelineErrorRecord struct {
	Err    error
	Input  interface{}
	Output interface{}
}

type ErrorChan chan *PipelineErrorRecord

func NewPipelineError(err error, input interface{}, output interface{}) *PipelineErrorRecord {
	return &PipelineErrorRecord{err, input, output}
}

type Importer interface {
	core.Initializable
	GetRecords(ctx core.RequestContext, dataChan DataChan, done chan bool) error
}

type Exporter interface {
	core.Initializable
	WriteRecord(ctx core.RequestContext, dataChan DataChan, errorChan ErrorChan, done chan bool) error
}

type ErrorProcessor interface {
	core.Initializable
	ProcessErrorRecord(ctx core.RequestContext, rec *PipelineErrorRecord)
}

type Processor interface {
	core.Initializable
	Transform(ctx core.RequestContext, input interface{}) (interface{}, error)
}

type Driver interface {
	core.Initializable
	Run(ctx core.RequestContext, importer Importer, exporter Exporter, processor Processor, errorProcessor ErrorProcessor) error
}

type DataPipelineRegisterar interface {
	RegisterDriver(ctx core.ServerContext, name string, driverfac func(ctx core.ServerContext) Driver)
	RegisterImporter(ctx core.ServerContext, name string, importerfac func(ctx core.ServerContext) Importer)
	RegisterExporter(ctx core.ServerContext, name string, exporterfac func(ctx core.ServerContext) Exporter)
	RegisterProcessor(ctx core.ServerContext, name string, procfac func(ctx core.ServerContext) Processor)
	RegisterErrorProcessor(ctx core.ServerContext, name string, procfac func(ctx core.ServerContext) ErrorProcessor)
}
