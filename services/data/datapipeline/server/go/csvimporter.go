package main

import (
	"encoding/csv"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
)

const (
	CONF_IMP_STORAGE   = "importstorageservice"
	CONF_IMP_BUCKET    = "importstoragebucket"
	CONF_IMP_CSVINPUT  = "csvinputfile"
	CONF_IMP_CSVHEADER = "importcsvhasheaders"
)

type csvImporter struct {
	stor       components.StorageComponent
	storBucket string
	csvFile    string
	headers    bool
}

func csvImporterFactory(core.ServerContext) datapipeline.Importer {
	return &csvImporter{}
}
func (imp *csvImporter) Initialize(ctx ctx.Context, conf config.Config) error {
	svcName, ok := conf.GetString(ctx, CONF_IMP_STORAGE)
	if !ok {
		return errors.MissingConf(ctx, CONF_IMP_STORAGE)
	}
	stor, err := ctx.(core.ServerContext).GetService(svcName)
	if err != nil {
		return errors.WrapErrorWithCode(ctx, err, errors.CORE_ERROR_BAD_CONF)
	}
	imp.stor, ok = stor.(components.StorageComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_IMP_STORAGE)
	}
	imp.storBucket, _ = conf.GetString(ctx, CONF_IMP_BUCKET)

	imp.csvFile, ok = conf.GetString(ctx, CONF_IMP_CSVINPUT)
	if !ok {
		return errors.MissingConf(ctx, CONF_IMP_CSVINPUT)
	}
	imp.headers, ok = conf.GetBool(ctx, CONF_IMP_CSVHEADER)
	if !ok {
		imp.headers = true
	}
	return nil
}

func (imp *csvImporter) GetRecords(ctx core.RequestContext, dataChan datapipeline.DataChan, done chan bool) error {
	inpRdr, err := imp.stor.Open(ctx, imp.storBucket, imp.csvFile)
	if err != nil {
		return err
	}

	var header []string
	rdr := csv.NewReader(inpRdr)
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			done <- true
			return nil
		}
		if err != nil {
			return err
		}
		if imp.headers {
			if header == nil {
				header = record
			} else {
				if len(record) < len(header) {
					continue
				}
				data := map[string]string{}
				for i := range header {
					data[header[i]] = record[i]
				}
				dataChan <- data
			}
		} else {
			dataChan <- record
		}
	}
	return nil
}
