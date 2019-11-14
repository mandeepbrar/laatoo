package main

import (
	"encoding/csv"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_EXP_STORAGE   = "exportstorageservice"
	CONF_EXP_BUCKET    = "exportstoragebucket"
	CONF_EXP_CSV       = "csvoutputfile"
	CONF_EXP_CSVHEADER = "exportcsvheaders"
)

type csvExporter struct {
	core.Service
	stor       components.StorageComponent
	storBucket string
	csvFile    string
	headers    []string
}

func (exp *csvExporter) Initialize(ctx core.ServerContext, conf config.Config) error {
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

	exp.csvFile, _ = conf.GetString(ctx, CONF_EXP_CSV)

	exp.headers, ok = conf.GetStringArray(ctx, CONF_EXP_CSVHEADER)
	if !ok {
		return errors.MissingConf(ctx, CONF_EXP_CSVHEADER)
	}

	return nil
}

func (exp *csvExporter) WriteRecord(ctx core.RequestContext, initData map[string]interface{}, inputDataChan datapipeline.DataChan, outputDataChan datapipeline.DataChan) error {

	//func (exp *csvExporter) WriteRecord(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan, pipelineErrorChan datapipeline.ErrorChan, done chan bool) error {

	file, ok := initData["outputfile"]
	if !ok {
		file = exp.csvFile
	}
	if file == "" {
		return errors.BadRequest(ctx, "Error", "output csv not provided")
	}

	outpurWrtr, err := exp.stor.OpenForWrite(ctx, exp.storBucket, file.(string))
	if err != nil {
		return err
	}
	defer outpurWrtr.Close()

	writer := csv.NewWriter(outpurWrtr)

	err = writer.Write(exp.headers)
	if err != nil {
		return err
	}

	for {
		select {
		case pipeRec := <-inputDataChan:
			{
				outputObj, ok := pipeRec.TransformedData.(map[string]interface{})
				log.Debug(ctx, "Recived object for saving", "save", outputObj)
				if !ok {
					pipeRec.Err = errors.BadRequest(ctx, "Error", "Object is not a storable")
				} else {

					val := make([]string, len(exp.headers))
					for i, name := range exp.headers {
						mapval, ok := outputObj[name]
						if ok {
							val[i] = fmt.Sprint(mapval)
						}
					}

					err := writer.Write(val)
					if err != nil {
						pipeRec.Err = err
					}
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
