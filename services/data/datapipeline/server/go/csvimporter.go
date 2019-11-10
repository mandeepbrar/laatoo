package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type csvImporter struct {
}

func csvImporterFactory(core.ServerContext) datapipeline.Importer {
	return &csvImporter{}
}
func (imp *csvImporter) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (imp *csvImporter) GetRecords(ctx core.RequestContext, dataChan datapipeline.DataChan, done chan bool) error {
	return nil
}
