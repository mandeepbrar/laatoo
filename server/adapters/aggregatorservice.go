package adapters

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
)

type ServiceAggregator struct {
	core.Service
	serviceMap map[string]elements.Service
}

func (svc *ServiceAggregator) Describe(ctx core.ServerContext) error {
	svc.AddConfigurations(ctx, map[string]string{constants.CONF_SERVICES: config.OBJECTTYPE_CONFIG})
	return svc.AddParamWithType(ctx, "argsMap", config.OBJECTTYPE_STRINGMAP)
}

//The services start serving when this method is called
func (svc *ServiceAggregator) Start(ctx core.ServerContext) error {
	c, _ := svc.GetConfiguration(ctx, constants.CONF_SERVICES)
	svcConfig := c.(config.Config)
	svcMgr := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
	svcs := svcConfig.AllConfigurations(ctx)
	svc.serviceMap = make(map[string]elements.Service)
	for _, svcAlias := range svcs {
		svcConf, _ := svcConfig.GetSubConfig(ctx, svcAlias)
		name, _ := svcConf.GetString(ctx, "Name")
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
	argsMap, ok := req.GetStringMapParam(ctx, "argsMap")
	if !ok {
		return core.BadRequestResponse("Argsmap not provided"), nil
	}
	//argsMap := *body
	retval := make(map[string]*core.Response, len(svc.serviceMap))
	log.Trace(ctx, "Aggregator SetRequestTypeargs", "argsMap", argsMap)
	for k, v := range argsMap {
		reqctx := ctx.SubContext(k)
		sArg := v.(map[string]interface{})
		log.Trace(reqctx, "Service ", "Name", k, "Args", sArg)
		/*		reqbody, ok := sArg["Body"]
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
				}*/
		svcToInvoke, ok := svc.serviceMap[k]
		if !ok {
			return nil, errors.BadRequest(ctx)
		}
		resp, err := svcToInvoke.HandleRequest(reqctx, sArg)
		if err != nil {
			return nil, errors.WrapError(reqctx, err)
		}
		retval[k] = resp
	}
	return core.SuccessResponse(retval), nil
}
