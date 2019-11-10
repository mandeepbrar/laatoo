package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type restImporter struct {
}

func restImporterFactory(core.ServerContext) datapipeline.Importer {
	return &restImporter{}
}
func (imp *restImporter) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (imp *restImporter) GetRecords(ctx core.RequestContext, dataChan datapipeline.DataChan, done chan bool) error {
	return nil
}
