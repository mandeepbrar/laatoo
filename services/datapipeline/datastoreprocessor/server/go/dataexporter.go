package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type dataExporter struct {
	core.Service
	dataStor data.DataComponent
}

func (exp *dataExporter) Initialize(ctx core.ServerContext, conf config.Config) error {
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
		return errors.BadConf(ctx, CONF_EXP_DATASVC, "Message", "Service is not data component")
	}
	return nil
}

func (exp *dataExporter) WriteRecord(ctx core.RequestContext, initData map[string]interface{}, inputDataChan datapipeline.DataChan, outputDataChan datapipeline.DataChan) error {

	for {
		select {
		case pipeRec := <-inputDataChan:
			{
				outputObj, ok := pipeRec.TransformedData.(data.Storable)
				log.Debug(ctx, "Recived object for saving", "save", outputObj)
				if !ok {
					pipeRec.Err = errors.BadRequest(ctx, "Error", "Object is not a storable")
				} else {
					err := exp.dataStor.Save(ctx, outputObj)
					if err != nil {
						pipeRec.Err = err
					}
					log.Error(ctx, "Successfully saved saving", "save", outputObj)
				}
				outputDataChan <- pipeRec
			}
		case <-ctx.Done():
			{
				log.Debug(ctx, "Done with data export")
				return nil
			}
		}
	}
	return nil
}

const (
	CONF_EXP_DATASVC = "exportdataservice"
)
