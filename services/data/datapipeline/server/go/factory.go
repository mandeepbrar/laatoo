package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
)

type dataPipelineFactory struct {
	core.ServiceFactory
	importers  map[string]func(ctx core.ServerContext) datapipeline.Importer
	exporters  map[string]func(ctx core.ServerContext) datapipeline.Exporter
	processors map[string]func(ctx core.ServerContext) datapipeline.Processor
	errorProcs map[string]func(ctx core.ServerContext) datapipeline.ErrorProcessor
	drivers    map[string]func(ctx core.ServerContext) datapipeline.Driver
}

func (fac *dataPipelineFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	fac.importers = make(map[string]func(ctx core.ServerContext) datapipeline.Importer)
	fac.exporters = make(map[string]func(ctx core.ServerContext) datapipeline.Exporter)
	fac.processors = make(map[string]func(ctx core.ServerContext) datapipeline.Processor)
	fac.errorProcs = make(map[string]func(ctx core.ServerContext) datapipeline.ErrorProcessor)
	fac.drivers = make(map[string]func(ctx core.ServerContext) datapipeline.Driver)

	fac.RegisterErrorProcessor(ctx, "logErrors", logErrorProcFactory)
	fac.RegisterErrorProcessor(ctx, "memErrors", memoryErrorsProcessorFactory)
	fac.RegisterExporter(ctx, "csvExporter", csvExporterFactory)
	fac.RegisterExporter(ctx, "restExporter", restExporterFactory)
	fac.RegisterExporter(ctx, "dataExporter", dataExporterFactory)
	fac.RegisterImporter(ctx, "csvImporter", csvImporterFactory)
	fac.RegisterImporter(ctx, "dataImporter", dataImporterFactory)
	fac.RegisterImporter(ctx, "restImporter", restImporterFactory)
	fac.RegisterProcessor(ctx, "objectToMap", objecToMapFactory)
	fac.RegisterProcessor(ctx, "mapToObject", mapToObjectFactory)
	fac.RegisterDriver(ctx, "memoryDriver", memoryDriverFactory)

	return nil
}

//Create the services configured for factory.
func (fac *dataPipelineFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newPipelineService(ctx, fac)
}

func (fac *dataPipelineFactory) getImporter(ctx core.ServerContext, name string) (datapipeline.Importer, bool) {
	impfac, ok := fac.importers[name]
	if !ok {
		return nil, false
	}
	return impfac(ctx), ok
}

func (fac *dataPipelineFactory) getExporter(ctx core.ServerContext, name string) (datapipeline.Exporter, bool) {
	expfac, ok := fac.exporters[name]
	if !ok {
		return nil, false
	}
	return expfac(ctx), ok
}

func (fac *dataPipelineFactory) getProcessor(ctx core.ServerContext, name string) (datapipeline.Processor, bool) {
	procfac, ok := fac.processors[name]
	if !ok {
		return nil, false
	}
	return procfac(ctx), ok
}
func (fac *dataPipelineFactory) getDriver(ctx core.ServerContext, name string) (datapipeline.Driver, bool) {
	driverfac, ok := fac.drivers[name]
	if !ok {
		return nil, false
	}
	return driverfac(ctx), ok
}
func (fac *dataPipelineFactory) getErrorProcessor(ctx core.ServerContext, name string) (datapipeline.ErrorProcessor, bool) {
	errprocfac, ok := fac.errorProcs[name]
	if !ok {
		return nil, false
	}
	return errprocfac(ctx), ok
}

func (fac *dataPipelineFactory) RegisterImporter(ctx core.ServerContext, name string, importerfac func(ctx core.ServerContext) datapipeline.Importer) {
	fac.importers[name] = importerfac
}

func (fac *dataPipelineFactory) RegisterExporter(ctx core.ServerContext, name string, exporterfac func(ctx core.ServerContext) datapipeline.Exporter) {
	fac.exporters[name] = exporterfac
}

func (fac *dataPipelineFactory) RegisterProcessor(ctx core.ServerContext, name string, procfac func(ctx core.ServerContext) datapipeline.Processor) {
	fac.processors[name] = procfac
}
func (fac *dataPipelineFactory) RegisterErrorProcessor(ctx core.ServerContext, name string, errprocfac func(ctx core.ServerContext) datapipeline.ErrorProcessor) {
	fac.errorProcs[name] = errprocfac
}
func (fac *dataPipelineFactory) RegisterDriver(ctx core.ServerContext, name string, driverfac func(ctx core.ServerContext) datapipeline.Driver) {
	fac.drivers[name] = driverfac
}
