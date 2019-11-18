package main

import (
	"encoding/csv"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_IMP_STORAGE   = "importstorageservice"
	CONF_IMP_BUCKET    = "importstoragebucket"
	CONF_IMP_CSVINPUT  = "csvinputfile"
	CONF_IMP_CSVHEADER = "importcsvhasheaders"
)

type CsvImporter struct {
	core.Service
	stor       components.StorageComponent
	storBucket string
	csvFile    string
	headers    bool
}

func (imp *CsvImporter) Initialize(ctx core.ServerContext, conf config.Config) error {
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

	imp.csvFile, _ = conf.GetString(ctx, CONF_IMP_CSVINPUT)

	imp.headers, ok = conf.GetBool(ctx, CONF_IMP_CSVHEADER)
	if !ok {
		imp.headers = true
	}
	return nil
}

func (imp *CsvImporter) GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan) error {
	var file string
	fileInt, ok := initData["inputfile"]
	if !ok {
		file = imp.csvFile
	} else {
		file = fileInt.(string)
	}
	if file == "" {
		return errors.BadRequest(ctx, "Error", "input csv not provided")
	}

	inpRdr, err := imp.stor.Open(ctx, imp.storBucket, file)
	if err != nil {
		return err
	}
	defer inpRdr.Close()

	var header []string
	rdr := csv.NewReader(inpRdr)
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Error(ctx, "Could not read input data", "err", err)
			break
		}
		if imp.headers {
			if header == nil {
				empty := false
				for _, k := range record {
					if k == "" {
						empty = true
						break
					}
				}
				if !empty {
					header = record
					checkExcelBytes(header)
				}
			} else {
				if len(record) < len(header) {
					continue
				}
				data := map[string]string{}
				for i := range header {

					data[header[i]] = record[i]
				}
				log.Debug(ctx, "Putting record in pipeline with data", "data", data)
				dataChan <- datapipeline.NewPipelineRecord(data, nil)
			}
		} else {
			dataChan <- datapipeline.NewPipelineRecord(record, nil)
		}
	}
	log.Error(ctx, "Closing get records")
	return nil
}

func checkExcelBytes(header []string) {
	firstColumn := header[0]
	if len(firstColumn) > 3 {
		if firstColumn[0] == 239 {
			header[0] = firstColumn[3:]
		}
	}
}
