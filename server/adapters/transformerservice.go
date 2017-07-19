package adapters

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/constants"
)

const (
	CONF_TRANSFORMER_FUNC         = "transformer"
	CONF_TRANSFORMER_SPLIT_PARAMS = "splitparams"
)

type TransformerService struct {
	serviceNames        map[string]string
	serviceMap          map[string]core.Service
	transformerFunc     core.ServiceFunc
	transformerFuncName string
	splitParams         bool
}

func (ds *TransformerService) Initialize(ctx core.ServerContext, conf config.Config) error {
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
	ds.transformerFuncName, _ = conf.GetString(CONF_TRANSFORMER_FUNC)
	ds.splitParams, _ = conf.GetBool(CONF_TRANSFORMER_SPLIT_PARAMS)
	return nil
}

//The services start serving when this method is called
func (ds *TransformerService) Start(ctx core.ServerContext) error {
	for k, v := range ds.serviceNames {
		svc, err := ctx.GetService(v)
		if err != nil {
			return err
		}
		ds.serviceMap[k] = svc
	}
	if ds.transformerFuncName != "" {
		f, err := ctx.GetMethod(ds.transformerFuncName)
		if err != nil {
			return errors.BadConf(ctx, CONF_TRANSFORMER_FUNC, "Error", err)
		}
		ds.transformerFunc = f
	}
	return nil
}
func (ds *TransformerService) Info() *core.ServiceInfo {
	return &core.ServiceInfo{
		Request: core.RequestInfo{DataType: constants.CONF_OBJECT_STRINGMAP}}
}

func (ds *TransformerService) Invoke(ctx core.RequestContext, req core.ServiceRequest) (*core.ServiceResponse, error) {
	ctx = ctx.SubContext("Transformer Service")
	var argsMap map[string]interface{}
	if ds.splitParams {
		body := req.GetBody().(*map[string]interface{})
		argsMap = *body
		log.Trace(ctx, "Transformer Service. Split Params", "argsMap", argsMap)
	}
	retval := make(map[string]*core.ServiceResponse, len(ds.serviceMap))
	for svcName, svc := range ds.serviceMap {
		reqctx := ctx.SubContext(svcName)
		svcreq := reqctx.CreateRequest()
		if ds.splitParams {
			v := argsMap[svcName]
			sArg := v.(map[string]interface{})
			log.Trace(reqctx, "Service ", "Name", svcName, "Args", sArg)
			reqbody, ok := sArg["Body"]
			if ok {
				reqmap, ok := reqbody.(map[string]interface{})
				if !ok {
					reqmap = make(map[string]interface{})
				}
				//reqctx.SetRequest(&reqmap)
				svcreq.SetBody(reqmap)
			}
			params, ok := sArg["Params"]
			if ok {
				paramsMap, ok := params.(map[string]interface{})
				if ok {
					reqparams := make(core.ServiceParamsMap)
					for k, v := range paramsMap {
						reqparams.AddParam(k, v, "", false)
						//						reqctx.Set(k, v)
					}
					svcreq.SetParams(reqparams)
				}
			}
		}
		resp, err := svc.Invoke(reqctx, svcreq)
		if err != nil {
			return nil, errors.WrapError(reqctx, err)
		}
		retval[svcName] = resp
	}
	if ds.transformerFunc != nil {
		transformCtx := ctx.SubContext("transformer")
		transformreq := transformCtx.CreateRequest()
		transformreq.SetBody(retval)
		//transformCtx.Set("Source", retval)
		return ds.transformerFunc(transformCtx, transformreq)
	} else {
		return core.SuccessResponse(retval), nil
	}
}
