package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type restExporter struct {
}

func restExporterFactory(core.ServerContext) datapipeline.Exporter {
	return &restExporter{}
}
func (exp *restExporter) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (exp *restExporter) WriteRecord(ctx core.RequestContext, output interface{}) error {
	return nil
}
