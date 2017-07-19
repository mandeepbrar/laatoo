package adapters

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/constants"
)

type ServiceAggregator struct {
	serviceNames map[string]string
	serviceMap   map[string]core.Service
}

func (ds *ServiceAggregator) Initialize(ctx core.ServerContext, conf config.Config) error {
	svcConfig, ok := conf.GetSubConfig(constants.CONF_SERVICES)
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
func (ds *ServiceAggregator) Start(ctx core.ServerContext) error {
	for k, v := range ds.serviceNames {
		svc, err := ctx.GetService(v)
		if err != nil {
			return err
		}
		ds.serviceMap[k] = svc
	}
	return nil
}

func (ds *ServiceAggregator) Info() *core.ServiceInfo {
	return &core.ServiceInfo{
		Request: core.RequestInfo{DataType: constants.CONF_OBJECT_STRINGMAP}}
}

func (ds *ServiceAggregator) Invoke(ctx core.RequestContext, req core.ServiceRequest) (*core.ServiceResponse, error) {
	ctx = ctx.SubContext("Aggregator Service")
	body := req.GetBody().(*map[string]interface{})
	argsMap := *body
	retval := make(map[string]*core.ServiceResponse, len(ds.serviceMap))
	log.Trace(ctx, "Aggregator args", "argsMap", argsMap)
	for k, v := range argsMap {
		reqctx := ctx.SubContext(k)
		sArg := v.(map[string]interface{})
		log.Trace(reqctx, "Service ", "Name", k, "Args", sArg)
		svcreq := reqctx.CreateRequest()
		reqbody, ok := sArg["Body"]
		if ok {
			reqmap, ok := reqbody.(map[string]interface{})
			if !ok {
				reqmap = make(map[string]interface{})
			}
			svcreq.SetBody(reqmap)
			//reqctx.SetRequest(&reqmap)
		}
		params, ok := sArg["Params"]
		if ok {
			paramsMap, ok := params.(map[string]interface{})
			if ok {
				reqparams := make(core.ServiceParamsMap)
				for k, v := range paramsMap {
					reqparams.AddParam(k, v, "", false)
					//reqctx.Set(k, v)
				}
				svcreq.SetParams(reqparams)
			}
		}
		svc, ok := ds.serviceMap[k]
		if !ok {
			return nil, errors.BadRequest(ctx)
		}
		resp, err := svc.Invoke(reqctx, svcreq)
		if err != nil {
			return nil, errors.WrapError(reqctx, err)
		}
		retval[k] = resp
	}
	return core.SuccessResponse(retval), nil
}
