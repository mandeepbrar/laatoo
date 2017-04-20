package service

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type transformerService struct {
	serviceNames        map[string]string
	serviceMap          map[string]core.Service
	transformerFunc     core.ServiceFunc
	transformerFuncName string
	splitParams         bool
}

const (
	CONF_TRANSFORMERSERVICE_NAME  = "__transformerservice__"
	CONF_TRANSFORMER_FUNC         = "transformer"
	CONF_TRANSFORMER_SPLIT_PARAMS = "splitparams"
)

func init() {
	objects.RegisterObject(CONF_TRANSFORMERSERVICE_NAME, createTransformerService, nil)
}

func createTransformerService() interface{} {
	return &transformerService{}
}

func (ds *transformerService) Initialize(ctx core.ServerContext, conf config.Config) error {
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
	ds.transformerFuncName, _ = conf.GetString(CONF_TRANSFORMER_FUNC)
	ds.splitParams, _ = conf.GetBool(CONF_TRANSFORMER_SPLIT_PARAMS)
	return nil
}

//The services start serving when this method is called
func (ds *transformerService) Start(ctx core.ServerContext) error {
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

func (ds *transformerService) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Transformer Service")
	var argsMap map[string]interface{}
	if ds.splitParams {
		body := ctx.GetRequest().(*map[string]interface{})
		argsMap = *body
		log.Logger.Trace(ctx, "Transformer Service. Split Params", "argsMap", argsMap)
	}
	retval := make(map[string]*core.ServiceResponse, len(ds.serviceMap))
	for svcName, svc := range ds.serviceMap {
		reqctx := ctx.SubContext(svcName)
		if ds.splitParams {
			v := argsMap[svcName]
			sArg := v.(map[string]interface{})
			log.Logger.Trace(reqctx, "Service ", "Name", svcName, "Args", sArg)
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
		}
		err := svc.Invoke(reqctx)
		if err != nil {
			return errors.WrapError(reqctx, err)
		}
		resp := reqctx.GetResponse()
		retval[svcName] = resp
	}
	if ds.transformerFunc != nil {
		transformCtx := ctx.SubContext("transformer")
		transformCtx.Set("Source", retval)
		err := ds.transformerFunc(transformCtx)
		if err != nil {
			return errors.WrapError(transformCtx, err)
		}
		ctx.SetResponse(transformCtx.GetResponse())
	} else {
		ctx.SetResponse(core.SuccessResponse(retval))
	}
	return nil
}
