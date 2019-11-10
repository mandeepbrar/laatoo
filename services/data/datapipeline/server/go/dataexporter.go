package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type dataExporter struct {
}

func dataExporterFactory(core.ServerContext) datapipeline.Exporter {
	return &dataExporter{}
}
func (exp *dataExporter) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (exp *dataExporter) WriteRecord(ctx core.RequestContext, output interface{}) error {
	return nil
}
