package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
)

type dataExporter struct {
	dataStor data.DataComponent
}

func dataExporterFactory(core.ServerContext) datapipeline.Exporter {
	return &dataExporter{}
}
func (exp *dataExporter) Initialize(ctx ctx.Context, conf config.Config) error {
	svcName, ok := conf.GetString(ctx, CONF_EXP_DATASVC)
	if !ok {
		return errors.MissingConf(ctx, CONF_EXP_DATASVC)
	}
	stor, err := ctx.(core.ServerContext).GetService(svcName)
	if err != nil {
		return err
	}
	exp.dataStor, ok = stor.(data.DataComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_EXP_DATASVC)
	}
	return nil
}

func (exp *dataExporter) WriteRecord(ctx core.RequestContext, dataChan datapipeline.DataChan, pipelineErrorChan datapipeline.ErrorChan, done chan bool) error {

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
				return nil
			}
		}
	}
	return nil
}

const (
	CONF_EXP_DATASVC = "exportdataservice"
)
