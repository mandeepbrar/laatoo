package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
)

type memoryDriver struct {
}

func memoryDriverFactory(core.ServerContext) datapipeline.Driver {
	return &memoryDriver{}
}
func (proc *memoryDriver) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}
func (proc *memoryDriver) Run(ctx core.RequestContext, importer datapipeline.Importer, exporter datapipeline.Exporter, processor datapipeline.Processor, errorProcessor datapipeline.ErrorProcessor) error {
	inp := make(datapipeline.DataChan)
	outputChan := make(datapipeline.DataChan)
	errChan := make(datapipeline.ErrorChan)
	outputDoneChannel := make(chan bool)
	inpDoneChannel := make(chan bool)
	defer func() {
		close(inp)
		close(outputChan)
		close(inpDoneChannel)
		close(outputDoneChannel)
		close(errChan)
	}()

	go func(ctx core.RequestContext, inp datapipeline.DataChan, inpDoneChannel chan bool) {
		err := importer.GetRecords(ctx, inp, inpDoneChannel)
		if err != nil {
			log.Error(ctx, "Error encountered while opening read stream of data pipeline", "Error", err)
			inpDoneChannel <- true
			return
		}
	}(ctx, inp, inpDoneChannel)

	go func(ctx core.RequestContext, outputChan datapipeline.DataChan, pipelineErrorChan datapipeline.ErrorChan, outputdone, inpDoneChannel chan bool) {
		err := exporter.WriteRecord(ctx, outputChan, pipelineErrorChan, outputdone)
		if err != nil {
			log.Error(ctx, "Error encountered while opening write stream of data pipeline", "Error", err)
			outputdone <- true
			inpDoneChannel <- true
		}
	}(ctx, outputChan, errChan, outputDoneChannel, inpDoneChannel)

	for {
		select {
		case dataObj := <-inp:
			{
				fmt.Println("Received data", dataObj)
				proc := func(reqCtx core.RequestContext, dataObj interface{}, outputChan datapipeline.DataChan, pipelineErrorChan datapipeline.ErrorChan) {
					var procOutput interface{}
					var err error
					if processor != nil {
						procOutput, err = processor.Transform(reqCtx, dataObj)
						if err != nil {
							log.Warn(reqCtx, "Record could not be processed", "err", err)
							errorProcessor.ProcessErrorRecord(ctx, datapipeline.NewPipelineError(err, dataObj, nil))
							return
						}
					} else {
						procOutput = dataObj
					}
					outputChan <- procOutput

					errRec := <-errChan
					if errRec != nil {
						errRec.Input = dataObj
						log.Warn(reqCtx, "Record could not be exported", "err", errRec.Err)
						errorProcessor.ProcessErrorRecord(ctx, errRec)
					}
				}
				//not parallel processed at this time
				proc(ctx, dataObj, outputChan, errChan)
			}
		case <-inpDoneChannel:
			{
				outputDoneChannel <- true
				fmt.Println("Complete")
				return nil
			}
		}
	}

	return nil
}
