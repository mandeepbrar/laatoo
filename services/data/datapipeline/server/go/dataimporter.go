package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type dataImporter struct {
}

func dataImporterFactory(core.ServerContext) datapipeline.Importer {
	return &dataImporter{}
}
func (imp *dataImporter) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (imp *dataImporter) GetRecords(ctx core.RequestContext, dataChan datapipeline.DataChan, done chan bool) error {
	return nil
}
