package main

import (
	"encoding/json"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"net/http"
	"time"
)

type RestExporter struct {
	core.Service
	restEndpoint string
	objFac       core.ObjectFactory
}

func (exp *RestExporter) Initialize(ctx core.ServerContext, conf config.Config) error {
	resEndpoint, ok := conf.GetString(ctx, CONF_EXP_REST_ENDPOINT)
	if !ok {
		return errors.MissingConf(ctx, CONF_EXP_REST_ENDPOINT)
	}
	exp.restEndpoint = resEndpoint

	/*method, ok := conf.GetString(ctx, CONF_EXP_REST_METHOD)
	if !ok {
		exp.method = "GET"
	} else {
		exp.method = method
	}
	*/
	/*	obj, ok := conf.GetString(ctx, CONF_OBJECT_TO_EXPORT)
		if ok {
			exp.objFac, ok = ctx.(core.ServerContext).GetObjectFactory(obj)
			if !ok {
				return errors.BadConf(ctx, CONF_OBJECT_TO_EXPORT)
			}
		} else {
			return errors.MissingConf(ctx, CONF_OBJECT_TO_EXPORT)
		}

		exp.list, _ = conf.GetBool(ctx, CONF_EXP_LIST)*/

	return nil
}

func (exp *RestExporter) WriteRecord(ctx core.RequestContext, initData map[string]interface{}, inputDataChan datapipeline.DataChan, outputDataChan datapipeline.DataChan) error {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	for {
		select {
		case pipeRec := <-inputDataChan:
			{
				pr, pw := io.Pipe()
				go func(pw io.Writer, pipeRec *datapipeline.PipelineRecord, outputDataChan datapipeline.DataChan) {
					err := json.NewEncoder(pw).Encode(pipeRec.TransformedData)
					if err != nil {
						log.Error(ctx, "Error encoding", "error", err)
						pipeRec.Err = err
						outputDataChan <- pipeRec
					}
				}(pw, pipeRec, outputDataChan)

				_, err := netClient.Post(exp.restEndpoint, "application/json", pr)
				pipeRec.Err = err
				outputDataChan <- pipeRec
			}
		case <-ctx.Done():
			{
				log.Error(ctx, "Done with write record")
				return nil
			}

		}
	}

}

const (
	CONF_EXP_REST_ENDPOINT = "exportrestendpoint"
)
