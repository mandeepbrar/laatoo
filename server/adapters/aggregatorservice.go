package adapters

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
)

type ServiceAggregator struct {
	core.Service
	serviceMap map[string]server.Service
}

func (svc *ServiceAggregator) Initialize(ctx core.ServerContext) error {
	svc.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	svc.AddConfigurations(ctx, map[string]string{constants.CONF_SERVICES: config.CONF_OBJECT_CONFIG})
	return nil
}

//The services start serving when this method is called
func (svc *ServiceAggregator) Start(ctx core.ServerContext) error {
	c, _ := svc.GetConfiguration(ctx, constants.CONF_SERVICES)
	svcConfig := c.(config.Config)
	svcMgr := ctx.GetServerElement(core.ServerElementServiceManager).(server.ServiceManager)
	svcs := svcConfig.AllConfigurations()
	svc.serviceMap = make(map[string]server.Service)
	for _, svcAlias := range svcs {
		svcConf, _ := svcConfig.GetSubConfig(svcAlias)
		name, _ := svcConf.GetString("Name")
		namedsvc, err := svcMgr.GetService(ctx, name)
		if err != nil {
			return err
		}
		svc.serviceMap[svcAlias] = namedsvc
	}
	return nil
}

func (svc *ServiceAggregator) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	ctx = ctx.SubContext("Aggregator Service")
	body := req.GetBody().(*map[string]interface{})
	argsMap := *body
	retval := make(map[string]*core.Response, len(svc.serviceMap))
	log.Trace(ctx, "Aggregator args", "argsMap", argsMap)
	for k, v := range argsMap {
		reqctx := ctx.SubContext(k)
		sArg := v.(map[string]interface{})
		log.Trace(reqctx, "Service ", "Name", k, "Args", sArg)
		reqbody, ok := sArg["Body"]
		var reqmap, reqparams map[string]interface{}
		if ok {
			reqmap, ok = reqbody.(map[string]interface{})
			if !ok {
				reqmap = make(map[string]interface{})
			}
			//reqctx.SetRequest(&reqmap)
		}
		params, ok := sArg["Params"]
		if ok {
			paramsMap, ok := params.(map[string]interface{})
			if ok {
				reqparams := make(map[string]interface{})
				for k, v := range paramsMap {
					reqparams[k] = v
					//reqctx.Set(k, v)
				}
			}
		}
		svcToInvoke, ok := svc.serviceMap[k]
		if !ok {
			return nil, errors.BadRequest(ctx)
		}
		resp, err := svcToInvoke.HandleRequest(reqctx, reqparams, reqmap)
		if err != nil {
			return nil, errors.WrapError(reqctx, err)
		}
		retval[k] = resp
	}
	return core.SuccessResponse(retval), nil
}
