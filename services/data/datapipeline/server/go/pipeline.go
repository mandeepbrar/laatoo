package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_OBJECT_TO_IMPORT = "object"
	CONF_DATASOURCETYPE   = "datasourcetype"
	CONF_DATADESTTYPE     = "datadestinationtype"
	CONF_PROCTYPE         = "processortype"
	CONF_ERRORPROCTYPE    = "errorprocessortype"
)

func newPipelineService(ctx core.ServerContext, fac *dataPipelineFactory) (*pipelineService, error) {
	return &pipelineService{pipelineFac: fac}, nil
}

type pipelineService struct {
	core.Service
	conf           config.Config
	pipelineFac    *dataPipelineFactory
	objType        string
	objFac         core.ObjectFactory
	objDataSvc     data.DataComponent
	dataSource     string
	importer       datapipeline.Importer
	exporter       datapipeline.Exporter
	processor      datapipeline.Processor
	errorProcessor datapipeline.ErrorProcessor
}

func (svc *pipelineService) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.conf = conf
	obj, ok := svc.GetStringConfiguration(ctx, CONF_OBJECT_TO_IMPORT)
	if ok {
		svc.objFac, ok = ctx.GetObjectFactory(obj)
		if !ok {
			return errors.BadConf(ctx, errors.CORE_ERROR_MISSING_CONF)
		}
	} else {
		return errors.MissingConf(ctx, CONF_OBJECT_TO_IMPORT)
	}
	/*
		datasvcName, ok := svc.GetStringConfiguration(ctx, CONF_DATASERVICE)
		if !ok {
			return errors.MissingConf(ctx, CONF_DATASERVICE)
		}
		objDataSvc, err := ctx.GetService(datasvcName)
		if err != nil {
			return errors.WrapErrorWithCode(ctx, err, errors.CORE_ERROR_BAD_CONF)
		}
		svc.objDataSvc, ok = objDataSvc.(data.DataComponent)
		if !ok {
			return errors.BadConf(ctx, CONF_DATASERVICE)
		}
	*/
	return nil
}

func (svc *pipelineService) Start(ctx core.ServerContext) error {
	dataSourceType, ok := svc.GetStringConfiguration(ctx, CONF_DATASOURCETYPE)
	if !ok {
		return errors.MissingConf(ctx, CONF_DATASOURCETYPE)
	}

	svc.importer, ok = svc.pipelineFac.getImporter(ctx, dataSourceType)
	if !ok {
		return errors.BadConf(ctx, CONF_DATASOURCETYPE)
	}
	err := svc.importer.Initialize(ctx, svc.conf)
	if err != nil {
		return err
	}

	dataDestinationType, ok := svc.GetStringConfiguration(ctx, CONF_DATADESTTYPE)
	if !ok {
		return errors.MissingConf(ctx, CONF_DATADESTTYPE)
	}

	svc.exporter, ok = svc.pipelineFac.getExporter(ctx, dataDestinationType)
	if !ok {
		return errors.BadConf(ctx, CONF_DATADESTTYPE)
	}
	err = svc.exporter.Initialize(ctx, svc.conf)
	if err != nil {
		return err
	}

	procType, ok := svc.GetStringConfiguration(ctx, CONF_PROCTYPE)
	if ok {
		svc.processor, ok = svc.pipelineFac.getProcessor(ctx, procType)
		if !ok {
			return errors.BadConf(ctx, CONF_PROCTYPE)
		}
		err = svc.processor.Initialize(ctx, svc.conf)
		if err != nil {
			return err
		}
	}

	errorProcType, ok := svc.GetStringConfiguration(ctx, CONF_ERRORPROCTYPE)
	if ok {
		svc.errorProcessor, ok = svc.pipelineFac.getErrorProcessor(ctx, errorProcType)
		if !ok {
			return errors.BadConf(ctx, CONF_ERRORPROCTYPE)
		}
		err = svc.errorProcessor.Initialize(ctx, svc.conf)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *pipelineService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *pipelineService) runPipeline(ctx core.RequestContext) error {
	inp := make(datapipeline.DataChan)

	inpDoneChannel := make(chan bool)
	go svc.importer.GetRecords(ctx, inp, inpDoneChannel)
	for {
		select {
		case dataObj := <-inp:
			{
				fmt.Println("Received data", dataObj)
				go func(reqCtx core.RequestContext, dataObj interface{}) {
					var procOutput interface{}
					var err error
					if svc.processor != nil {
						procOutput, err = svc.processor.Transform(reqCtx, dataObj)
						if err != nil {
							log.Warn(reqCtx, "Record could not be processed", "err", err)
							svc.errorProcessor.ProcessErrorRecord(ctx, dataObj, nil, err)
							return
						}
					} else {
						procOutput = dataObj
					}
					err = svc.exporter.WriteRecord(reqCtx, procOutput)
					if err != nil {
						log.Warn(reqCtx, "Record could not be exported", "err", err)
						svc.errorProcessor.ProcessErrorRecord(ctx, dataObj, procOutput, err)
						return
					}
				}(ctx, dataObj)
			}
		case <-inpDoneChannel:
			{
				fmt.Println("Complete")
				return nil
			}
		}
	}

	return nil
}

/*

func (svc *pipelineService) importAccount(ctx core.RequestContext, accdata map[string]interface{}, acct *GLAccount) error {
	parentAcctId, ok := accdata["parent"]
	var parentAcct *GLAccount
	if ok {
		cond, err := svc.glAccountSvc.CreateCondition(ctx, data.FIELDVALUE, map[string]interface{}{"Code": parentAcctId})
		if err != nil {
			return err
		}
		stordata, _, _, recs, err := svc.glAccountSvc.Get(ctx, cond, -1, -1, "", nil)
		if err != nil {
			return err
		}
		if recs > 0 {
			parentAcct = stordata[0].(*GLAccount)
		}
	}
	accId, ok := accdata["Id"]
	if !ok {
		return errors.MissingArg(ctx, "Id")
	}
	accDesc, _ := accdata["Name"]

	accToCreate := &GLAccount{Code: accId.(string), Description: "accDesc", Title: fmt.Sprintf("%s %s", accId, accDesc)}
	if parentAcct != nil {
		accToCreate.Parent = GLAccount_Ref{Id: parentAcct.GetId(), Name: parentAcct.Title}
	}
	err := svc.glAccountSvc.Save(ctx, accToCreate)
	return err
}
*/
