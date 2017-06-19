package service

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/objects"
)

type serviceAggregator struct {
	serviceNames map[string]string
	serviceMap   map[string]core.Service
}

const (
	CONF_SERVICEAGGREGATOR_NAME = "__serviceaggregator__"
)

func init() {
	objects.RegisterObject(CONF_SERVICEAGGREGATOR_NAME, createServiceAggregator, nil)
}

func createServiceAggregator() interface{} {
	return &serviceAggregator{}
}

func (ds *serviceAggregator) Initialize(ctx core.ServerContext, conf config.Config) error {
	svcConfig, ok := conf.GetSubConfig(CONF_SERVICES)
	if ok {
		svcs := svcConfig.AllConfigurations()
		ds.serviceNames = make(map[string]string)
		ds.serviceMap = make(map[string]core.Service)
		for _, svcAlias := range svcs {
			svcConf, _ := svcConfig.GetSubConfig(svcAlias)
			name, ok := svcConf.GetString("Name")
			if ok {
				ds.serviceNames[svcAlias] = name
			}
		}
	}
	return nil
}

//The services start serving when this method is called
func (ds *serviceAggregator) Start(ctx core.ServerContext) error {
	for k, v := range ds.serviceNames {
		svc, err := ctx.GetService(v)
		if err != nil {
			return err
		}
		ds.serviceMap[k] = svc
	}
	return nil
}

func (ds *serviceAggregator) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Aggregator Service")
	body := ctx.GetRequest().(*map[string]interface{})
	argsMap := *body
	retval := make(map[string]*core.ServiceResponse, len(ds.serviceMap))
	log.Logger.Trace(ctx, "Aggregator", "argsMap", argsMap)
	for k, v := range argsMap {
		reqctx := ctx.SubContext(k)
		sArg := v.(map[string]interface{})
		log.Logger.Trace(reqctx, "Service ", "Name", k, "Args", sArg)
		reqbody, ok := sArg["Body"]
		if ok {
			reqmap, ok := reqbody.(map[string]interface{})
			if !ok {
				reqmap = make(map[string]interface{})
			}
			reqctx.SetRequest(&reqmap)
		}
		params, ok := sArg["Params"]
		if ok {
			paramsMap, ok := params.(map[string]interface{})
			if ok {
				for k, v := range paramsMap {
					reqctx.Set(k, v)
				}
			}
		}
		svc, ok := ds.serviceMap[k]
		if !ok {
			return errors.BadRequest(ctx)
		}
		err := svc.Invoke(reqctx)
		if err != nil {
			return errors.WrapError(reqctx, err)
		}
		resp := reqctx.GetResponse()
		retval[k] = resp
	}
	ctx.SetResponse(core.SuccessResponse(retval))
	return nil
}
