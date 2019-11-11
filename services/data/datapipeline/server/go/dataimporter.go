package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
)

type dataImporter struct {
	dataStor data.DataComponent
}

func dataImporterFactory(core.ServerContext) datapipeline.Importer {
	return &dataImporter{}
}
func (imp *dataImporter) Initialize(ctx ctx.Context, conf config.Config) error {
	svcName, ok := conf.GetString(ctx, CONF_INP_DATASVC)
	if !ok {
		return errors.MissingConf(ctx, CONF_INP_DATASVC)
	}
	stor, err := ctx.(core.ServerContext).GetService(svcName)
	if err != nil {
		return err
	}
	imp.dataStor, ok = stor.(data.DataComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_INP_DATASVC)
	}
	return nil
}

func (imp *dataImporter) GetRecords(ctx core.RequestContext, dataChan datapipeline.DataChan, done chan bool) error {
	recs, _, _, reccount, err := imp.dataStor.GetList(ctx, -1, -1, "", nil)
	if err != nil {
		return err
	}

	for i := 0; i < reccount; i++ {
		dataChan <- recs[i]
	}
	done <- true
	return nil
}

const (
	CONF_INP_DATASVC = "importdataservice"
)
