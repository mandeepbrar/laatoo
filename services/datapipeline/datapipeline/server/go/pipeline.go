package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_DATASOURCETYPE = "datasource"
	CONF_DATADESTTYPE   = "datadestination"
	CONF_PROCTYPE       = "processor"
	CONF_ERRORPROCTYPE  = "errorprocessor"
	CONF_DRIVERTYPE     = "driver"
	CONF_RETRYSERVICE   = "retrysource"
)

type pipelineService struct {
	core.Service
	importer       datapipeline.Importer
	exporter       datapipeline.Exporter
	processor      datapipeline.Processor
	errorProcessor datapipeline.ErrorProcessor
	driver         datapipeline.Driver
	retrySource    datapipeline.Importer
}

func (svc *pipelineService) Initialize(ctx core.ServerContext, conf config.Config) error {
	dataSource, ok := svc.GetStringConfiguration(ctx, CONF_DATASOURCETYPE)
	if !ok {
		return errors.MissingConf(ctx, CONF_DATASOURCETYPE)
	}

	dsvc, err := ctx.GetService(dataSource)
	if err != nil {
		return errors.BadConf(ctx, CONF_DATASOURCETYPE, "Err", err)
	} else {
		svc.importer, ok = dsvc.(datapipeline.Importer)
		if !ok {
			return errors.BadConf(ctx, CONF_DATASOURCETYPE)
		}
	}

	dataDestinationType, ok := svc.GetStringConfiguration(ctx, CONF_DATADESTTYPE)
	if !ok {
		return errors.MissingConf(ctx, CONF_DATADESTTYPE)
	}
	dsvc, err = ctx.GetService(dataDestinationType)
	if err != nil {
		return errors.BadConf(ctx, CONF_DATADESTTYPE, "Err", err)
	} else {
		svc.exporter, ok = dsvc.(datapipeline.Exporter)
		if !ok {
			return errors.BadConf(ctx, CONF_DATADESTTYPE)
		}
	}

	procType, ok := svc.GetStringConfiguration(ctx, CONF_PROCTYPE)
	log.Error(ctx, "Processor type", "svc", procType, "conf", conf)
	if ok {
		psvc, err := ctx.GetService(procType)
		if err != nil {
			return errors.BadConf(ctx, CONF_PROCTYPE, "Err", err)
		} else {
			svc.processor, ok = psvc.(datapipeline.Processor)
			if !ok {
				return errors.BadConf(ctx, CONF_PROCTYPE)
			}
		}
	}

	errorproc, ok := svc.GetStringConfiguration(ctx, CONF_ERRORPROCTYPE)
	if ok {
		epsvc, err := ctx.GetService(errorproc)
		if err != nil {
			return errors.BadConf(ctx, CONF_ERRORPROCTYPE, "Err", err)
		} else {
			svc.errorProcessor, ok = epsvc.(datapipeline.ErrorProcessor)
			if !ok {
				return errors.BadConf(ctx, CONF_ERRORPROCTYPE)
			}
		}
	}

	driverName, ok := svc.GetStringConfiguration(ctx, CONF_DRIVERTYPE)
	if !ok {
		svc.driver = &memoryDriver{}
	} else {
		driver, err := ctx.GetService(driverName)
		if err != nil {
			return errors.BadConf(ctx, CONF_DRIVERTYPE, "Err", err)
		} else {
			svc.driver, ok = driver.(datapipeline.Driver)
			if !ok {
				return errors.BadConf(ctx, CONF_DRIVERTYPE)
			}
		}
	}

	retryService, ok := svc.GetStringConfiguration(ctx, CONF_RETRYSERVICE)
	if ok {
		retrysvc, err := ctx.GetService(retryService)
		if err != nil {
			return errors.BadConf(ctx, CONF_RETRYSERVICE, "Err", err)
		} else {
			svc.retrySource, ok = retrysvc.(datapipeline.Importer)
			if !ok {
				return errors.BadConf(ctx, CONF_RETRYSERVICE)
			}
		}
	}

	return nil
}

func (svc *pipelineService) Invoke(ctx core.RequestContext) error {
	data, _ := ctx.GetStringMapParam("Data")
	retries, _ := ctx.GetIntParam("retries")
	go func() {
		err := svc.driver.Run(ctx, svc.importer, svc.exporter, svc.processor, svc.errorProcessor, data)
		log.Error(ctx, "Retrying pipeline", "err", err, "retrysrouce", svc.retrySource, "retries", retries)
		if err == nil && svc.retrySource != nil {
			for i := 0; i < retries; i++ {
				log.Debug(ctx, "Retrying pipeline", "retry", i)
				err = svc.driver.Run(ctx, svc.retrySource, svc.exporter, svc.processor, svc.errorProcessor, data)
				if err != nil {
					log.Error(ctx, "Error in pipeline", "err", err)
					break
				}
			}

		} else {
			log.Error(ctx, "Error in pipeline", "err", err)
		}
	}()
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
