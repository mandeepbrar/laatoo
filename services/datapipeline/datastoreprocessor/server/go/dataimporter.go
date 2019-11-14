package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type dataImporter struct {
	core.Service
	dataStor data.DataComponent
}

func (imp *dataImporter) Initialize(ctx core.ServerContext, conf config.Config) error {
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

func (imp *dataImporter) GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan) error {
	var queryCond interface{}
	var err error
	query, ok := initData["query"]
	if ok {
		queryMap, ok := query.(map[string]interface{})
		if ok {
			for k, v := range queryMap {
				queryCond, err = imp.dataStor.CreateCondition(ctx, data.FIELDVALUE, k, v)
				if err != nil {
					return err
				}
			}
		}
	}
	recs, _, _, reccount, err := imp.dataStor.Get(ctx, queryCond, -1, -1, "", nil)
	if err != nil {
		return err
	}

	for i := 0; i < reccount; i++ {
		dataChan <- datapipeline.NewPipelineRecord(recs[i], nil)
	}
	return nil
}

const (
	CONF_INP_DATASVC = "importdataservice"
)
