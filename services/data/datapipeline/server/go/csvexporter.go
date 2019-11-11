package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
)

const (
	CONF_EXP_STORAGE   = "exportstorageservice"
	CONF_EXP_BUCKET    = "exportstoragebucket"
	CONF_EXP_CSVINPUT  = "csvoutputfile"
	CONF_EXP_CSVHEADER = "exportcsvhasheaders"
)

type csvExporter struct {
	stor       components.StorageComponent
	storBucket string
	csvFile    string
	headers    bool
}

func csvExporterFactory(core.ServerContext) datapipeline.Exporter {
	return &csvExporter{}
}
func (exp *csvExporter) Initialize(ctx ctx.Context, conf config.Config) error {
	svcName, ok := conf.GetString(ctx, CONF_EXP_STORAGE)
	if !ok {
		return errors.MissingConf(ctx, CONF_EXP_STORAGE)
	}
	stor, err := ctx.(core.ServerContext).GetService(svcName)
	if err != nil {
		return errors.WrapErrorWithCode(ctx, err, errors.CORE_ERROR_BAD_CONF)
	}
	exp.stor, ok = stor.(components.StorageComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_EXP_STORAGE)
	}
	exp.storBucket, _ = conf.GetString(ctx, CONF_EXP_BUCKET)

	exp.csvFile, ok = conf.GetString(ctx, CONF_EXP_CSVINPUT)
	if !ok {
		return errors.MissingConf(ctx, CONF_EXP_CSVINPUT)
	}
	exp.headers, ok = conf.GetBool(ctx, CONF_EXP_CSVHEADER)
	if !ok {
		exp.headers = true
	}
	return nil
}

func (exp *csvExporter) WriteRecord(ctx core.RequestContext, dataChan datapipeline.DataChan, pipelineErrorChan datapipeline.ErrorChan, done chan bool) error {
	/*
		inpRdr, err := imp.stor.Open(ctx, imp.storBucket, imp.csvFile)
		if err != nil {
			return err
		}


		for {
			select {
			case output := <-dataChan:
				{
					outputObj, ok := output.(data.Storable)
					if !ok {
						pipelineErrorChan <- datapipeline.NewPipelineError(errors.BadRequest(ctx, "Error", "Object is not a storable"), nil, output)
						continue
					}
					err := exp.dataStor.Save(ctx, outputObj)
					if err != nil {
						pipelineErrorChan <- datapipeline.NewPipelineError(err, nil, output)
						continue
					}
					pipelineErrorChan <- nil

				}
			case <-done:
				{
					return
				}
			}
		}*/

	//no support yet from storage layer
	return nil
}
