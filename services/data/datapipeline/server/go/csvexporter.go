package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type csvExporter struct {
}

func csvExporterFactory(core.ServerContext) datapipeline.Exporter {
	return &csvExporter{}
}
func (exp *csvExporter) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (exp *csvExporter) WriteRecord(ctx core.RequestContext, output interface{}) error {
	return nil
}
