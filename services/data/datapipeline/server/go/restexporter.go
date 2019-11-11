package main

import (
	"encoding/json"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"net/http"
	"time"
)

type restExporter struct {
	restEndpoint string
	objFac       core.ObjectFactory
}

func restExporterFactory(core.ServerContext) datapipeline.Exporter {
	return &restExporter{}
}
func (exp *restExporter) Initialize(ctx ctx.Context, conf config.Config) error {
	resEndpoint, ok := conf.GetString(ctx, CONF_OBJECT_TO_EXPORT)
	if !ok {
		return errors.MissingConf(ctx, CONF_OBJECT_TO_EXPORT)
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

func (exp *restExporter) WriteRecord(ctx core.RequestContext, dataChan datapipeline.DataChan, pipelineErrorChan datapipeline.ErrorChan, done chan bool) error {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	for {
		select {
		case output := <-dataChan:
			{
				pr, pw := io.Pipe()
				go func(pw io.Writer, pipelineErrorChan datapipeline.ErrorChan) {
					err := json.NewEncoder(pw).Encode(output)
					if err != nil {
						log.Error(ctx, "Error encoding", "error", err)
						pipelineErrorChan <- datapipeline.NewPipelineError(err, nil, output)
					}
				}(pw, pipelineErrorChan)

				_, err := netClient.Post(exp.restEndpoint, "application/json", pr)
				if err != nil {
					pipelineErrorChan <- datapipeline.NewPipelineError(err, nil, output)
				}

				pipelineErrorChan <- nil
			}
		case <-done:
			{
				return nil
			}
		}
	}

}

const (
	CONF_OBJECT_TO_EXPORT  = "exportobject"
	CONF_EXP_REST_ENDPOINT = "exportrestendpoint"
	CONF_EXP_LIST          = "exportlist"
	CONF_EXP_REST_METHOD   = "exportmethod"
)
