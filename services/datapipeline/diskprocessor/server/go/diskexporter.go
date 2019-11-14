package main

import (
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
	CONF_EXP_CSVHEADER = "exportcsvhasheaders"
)

type diskExporter struct {
	core.Service
	stor       components.StorageComponent
	storBucket string
	file       string
	codec      core.Codec
	delim      byte
}

func (exp *diskExporter) Initialize(ctx core.ServerContext, conf config.Config) error {
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

	codecName, ok := conf.GetString(ctx, CONF_IMP_CODEC)
	if ok {
		codec, ok := ctx.GetCodec(codecName)
		if ok {
			exp.codec = codec
		}
	}

	if exp.codec == nil {
		exp.codec, _ = ctx.GetCodec("json")
	}

	exp.file, _ = conf.GetString(ctx, CONF_IMP_FILE)

	return nil
}

func (exp *diskExporter) WriteRecord(ctx core.RequestContext, initData map[string]interface{}, inputDataChan datapipeline.DataChan, outputDataChan datapipeline.DataChan) error {

	file, ok := initData["outputfile"]
	if !ok {
		file = exp.file
	}
	if file == "" {
		return errors.BadRequest(ctx, "Error", "output file not provided")
	}

	outpurWrtr, err := exp.stor.OpenForWrite(ctx, exp.storBucket, file.(string))
	if err != nil {
		return err
	}
	defer outpurWrtr.Close()

	for {
		select {
		case pipeRec := <-inputDataChan:
			{
				byts, err := exp.codec.Marshal(ctx, pipeRec.TransformedData)
				if err != nil {
					pipeRec.Err = err
				} else {
					_, err := outpurWrtr.Write(byts)
					if err != nil {
						pipeRec.Err = err
					} else {
						_, err := outpurWrtr.Write([]byte{exp.delim})
						if err != nil {
							pipeRec.Err = err
						}
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
