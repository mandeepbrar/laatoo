package main

import (
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"sync"
)

type MemoryDriver struct {
}

func (proc *MemoryDriver) processErrorRecord(ctx core.RequestContext, errorProcessor datapipeline.ErrorProcessor, rec *datapipeline.PipelineRecord) {
	wg := rec.PipelineData.(*sync.WaitGroup)
	defer wg.Done()
	if errorProcessor != nil {
		errorProcessor.ProcessErrorRecord(ctx, rec)
	} else {
		if rec.Err != nil {
			log.Error(ctx, "Error encountered in processing record", "error", rec.Err, " input ", rec.Input)
		} else {
			log.Error(ctx, "Record inserted cleanly", "input", rec.Input)
		}
	}
}

func (proc *MemoryDriver) Run(ctx core.RequestContext, importer datapipeline.Importer, exporter datapipeline.Exporter,
	processor datapipeline.Processor, errorProcessor datapipeline.ErrorProcessor, initData map[string]interface{}) error {

	myctx, cancelFunc := ctx.WithCancel()

	ctx = myctx.(core.RequestContext)

	inp := make(datapipeline.DataChan, 10)
	exportInputChan := make(datapipeline.DataChan, 10)
	exportOutputChan := make(datapipeline.DataChan, 10)

	pipelineErrChan := make(chan error)

	defer func() {
		cancelFunc()
		close(inp)
		close(pipelineErrChan)
		close(exportInputChan)
		close(exportOutputChan)
	}()

	var wg sync.WaitGroup

	go func(ctx core.RequestContext, initData map[string]interface{}, inp datapipeline.DataChan, pipelineErrChan chan error) {
		err := importer.GetRecords(ctx, initData, inp)
		if err != nil {
			log.Error(ctx, "Error encountered while opening read stream of data pipeline", "Error", err)
			pipelineErrChan <- err
		}
		pipelineErrChan <- nil
	}(ctx, initData, inp, pipelineErrChan)

	go func(ctx core.RequestContext, initData map[string]interface{}, exportInputChan datapipeline.DataChan, exportOutputChan datapipeline.DataChan,
		pipelineErrChan chan error) {
		err := exporter.WriteRecord(ctx, initData, exportInputChan, exportOutputChan)
		log.Debug(ctx, "Write record cleanup", "Error", err)
		if err != nil {
			log.Error(ctx, "Error encountered while opening write stream of data pipeline", "Error", err)
			pipelineErrChan <- err
		}
		pipelineErrChan <- nil
	}(ctx, initData, exportInputChan, exportOutputChan, pipelineErrChan)

	go func(ctx core.RequestContext, inp datapipeline.DataChan, exportInputChan datapipeline.DataChan,
		exportOutputChan datapipeline.DataChan, pipelineErrChan chan error, processor datapipeline.Processor,
		errorProcessor datapipeline.ErrorProcessor, wg *sync.WaitGroup) {

		err := proc.processLoop(ctx, inp, exportInputChan, exportOutputChan, pipelineErrChan, processor, errorProcessor, wg)
		log.Debug(ctx, "Write record cleanup", "Error", err)
		if err != nil {
			log.Debug(ctx, "Error encountered while opening write stream of data pipeline", "Error", err)
			pipelineErrChan <- err
		}
		pipelineErrChan <- nil
	}(ctx, inp, exportInputChan, exportOutputChan, pipelineErrChan, processor, errorProcessor, &wg)

	err := <-pipelineErrChan
	wg.Wait()
	cancelFunc()
	if err != nil {
		return err
	} else {
		//wait for processor loop and writer to come back from cancellation
		<-pipelineErrChan
		<-pipelineErrChan
		return nil
	}

	return nil
}

func (proc *MemoryDriver) processLoop(ctx core.RequestContext, inp datapipeline.DataChan, exportInputChan datapipeline.DataChan,
	exportOutputChan datapipeline.DataChan, pipelineErrChan chan error, processor datapipeline.Processor,
	errorProcessor datapipeline.ErrorProcessor, wg *sync.WaitGroup) error {

	for {
		select {
		case precord := <-inp:
			{
				wg.Add(1)
				precord.PipelineData = wg
				log.Debug(ctx, "Received input record", "input record", precord)
				if precord.Err == nil {
					transform := func(reqCtx core.RequestContext, rec *datapipeline.PipelineRecord, processor datapipeline.Processor,
						exportInputChan datapipeline.DataChan) {

						var procOutput interface{}
						var err error

						if processor != nil {
							procOutput, err = processor.Transform(reqCtx, rec.Input)
						} else {
							procOutput = rec.Input
						}

						log.Debug(ctx, "Output from processor", "input", rec, "output", procOutput)

						if err != nil {
							rec.Err = err
							log.Warn(reqCtx, "Record could not be processed", "err", err)
							proc.processErrorRecord(ctx, errorProcessor, rec)
						} else {
							if procOutput == nil {
								rec.Err = errors.InternalError(ctx, "Processor generated a nil output")
								proc.processErrorRecord(ctx, errorProcessor, rec)
							} else {
								rec.TransformedData = procOutput
								log.Debug(ctx, "Sending to output channel", "obj", procOutput)
								exportInputChan <- rec
							}
						}
					}

					//not parallel processed at this time
					go transform(ctx, precord, processor, exportInputChan)
				} else {
					proc.processErrorRecord(ctx, errorProcessor, precord)
				}
			}
		case precord := <-exportOutputChan:
			{
				proc.processErrorRecord(ctx, errorProcessor, precord)
			}
		case <-ctx.Done():
			{
				log.Debug(ctx, "Closed write record")
				return nil
			}
		}
	}
	return nil
}
