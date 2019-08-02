package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/codecs"
	"laatoo/server/constants"
	"reflect"
)

type serverService struct {
	name             string
	objectName       string
	service          core.Service
	conf             config.Config
	factory          elements.Factory
	owner            *serviceManager
	middleware       []*serverService
	paramValues      map[string]interface{}
	impl             *serviceImpl
	svrContext       *serverContext
	codecs           map[string]core.Codec
	serviceToForward string
	forward          bool
}

func (svc *serverService) loadMetaData(ctx core.ServerContext) error {
	//inject service implementation into
	//every service
	impl := newServiceImpl(svc.name)
	svc.impl = impl
	var svcval core.Service
	svcval = impl
	val := reflect.ValueOf(svc.service)
	elem := val.Elem()
	fld := elem.FieldByName("Service")
	if fld.CanSet() {
		fld.Set(reflect.ValueOf(svcval))
	} else {
		return errors.TypeMismatch(ctx, "Service does not inherit from core.Service", svc.name)
	}

	svc.codecs = map[string]core.Codec{"json": codecs.NewJsonCodec(), "fastjson": codecs.NewFastJsonCodec(), "bin": codecs.NewBinaryCodec(), "proto": codecs.NewProtobufCodec()}

	ldr := ctx.GetServerElement(core.ServerElementLoader).(elements.ObjectLoader)
	md, _ := ldr.GetMetaData(ctx, svc.objectName)
	if md != nil {
		inf, ok := md.(*serviceInfo)
		if ok {
			impl.serviceInfo = inf.clone()
		}
	}
	err := svc.service.Describe(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	//svc.info = svc.service.Info()
	log.Trace(ctx, "Service info ", "Name", svc.name, "Info", svc.impl.serviceInfo)

	return nil
}

func (svc *serverService) initialize(ctx core.ServerContext, conf config.Config) error {

	log.Error(ctx, "initializing service with conf", "conf", conf)
	if err := svc.processInfo(ctx, conf, svc.impl); err != nil {
		return err
	}

	svc.serviceToForward, svc.forward = conf.GetString(ctx, constants.CONF_SERVICE_FORWARD)

	err := svc.service.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	svc.impl.state = Initialized

	return nil
}

func (svc *serverService) start(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Service Start")

	err := svc.service.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	middleware := make([]*serverService, 0)
	middlewareNames, ok := ctx.GetStringArray(constants.CONF_MIDDLEWARE)
	log.Trace(ctx, "Getting middleware", "middleware", middlewareNames, "service", svc.name)
	if ok {
		for _, mwName := range middlewareNames {
			log.Trace(ctx, "Adding middleware", "name", mwName, "service", svc.name)
			//			mwSvc, err := ctx.GetService(mwName)
			mwSvcStruct, err := svc.owner.proxy.GetService(ctx, mwName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			mwSvc := mwSvcStruct.(*serviceProxy).svc
			middleware = append(middleware, mwSvc)
		}
	}
	middleware = append(middleware, svc)
	svc.middleware = middleware
	svc.impl.state = Started
	return err
}

func (svc *serverService) processInfo(ctx core.ServerContext, svcconf config.Config, info core.ServiceInfo) error {

	svcs := info.GetRequiredServices()
	if err := svc.injectServices(ctx, svcconf, svcs); err != nil {
		return err
	}

	if err := svc.impl.processInfo(ctx, svcconf); err != nil {
		return err
	}

	log.Info(ctx, "Processed info for service", "name", svc.name)
	return nil
}

func (svc *serverService) injectServices(ctx core.ServerContext, svcconf config.Config, svcsToInject map[string]string) error {
	if svcsToInject == nil {
		return nil
	}
	for confName, fieldName := range svcsToInject {
		svcName, ok := svcconf.GetString(ctx, confName)
		if !ok {
			errors.MissingConf(ctx, confName)
		}
		svrsvc, err := svc.owner.getService(ctx, svcName)
		if err != nil {
			errors.MissingService(ctx, svcName)
		}
		svctoinject := svrsvc.Service()
		svcval := reflect.ValueOf(svc.service)
		fld := svcval.FieldByName(fieldName)
		if fld.IsNil() {
			errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Field", fieldName)
		}
		if fld.CanSet() {
			fld.Set(reflect.ValueOf(svctoinject))
		}
	}
	return nil
}

func (svc *serverService) handleRequest(ctx *requestContext, vals map[string]interface{}) (*core.Response, error) {
	ctx = ctx.SubContext(svc.name).(*requestContext)
	codecname := "json"
	co, ok := vals["encoding"]
	if ok {
		codecname = co.(string)
	}
	var codec core.Codec
	if codecname != "" {
		codec, ok = svc.codecs[codecname]
		if !ok {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_CODEC_NOT_FOUND)
		}
	}

	req := ctx.createRequest()
	if err := svc.populateParams(ctx, vals, req, codec); err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	ctx.req = req
	log.Trace(ctx, "Invoking service", "Request", req)
	err := svc.invoke(ctx)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resp := ctx.GetResponse()
	var data map[string]interface{}
	if resp != nil {
		data = resp.Data
	}
	log.Error(ctx, "handle request", "vals", vals, " data", data)
	if svc.forward {
		if resp.Status == core.StatusSuccess {
			params := utils.ShallowMergeMaps(vals, data)
			//dat := resp.Data
			params["encoding"] = ""
			log.Error(ctx, "handle request", "params", params)

			err := ctx.Forward(svc.serviceToForward, params)
			if err != nil {
				return nil, errors.WrapError(ctx, err)
			}
			resp = ctx.GetResponse()
		}
	}
	return resp, nil
}

func (svc *serverService) populateParams(ctx *requestContext, vals map[string]interface{}, req *request, codec core.Codec) error {
	encoded := (codec != nil)
	reqParams := make(map[string]core.Param)
	reqInfo := svc.impl.GetRequestInfo()
	params := reqInfo.ParamInfo()

	for name, svcParam := range params {
		reqParam := svcParam.(*param).clone()
		val, ok := vals[name]
		log.Error(ctx, "populate params", "param", svcParam, "val", val, "encoded", encoded)
		if (val != nil) && ok {
			err := reqParam.setValue(ctx, val, codec, encoded)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			if reqParam.IsRequired() {
				return errors.MissingArg(ctx, name)
			}
		}
		reqParams[name] = reqParam
	}

	log.Trace(ctx, "Populated params", "reqParams", reqParams, "params", params, "reqInfo", reqInfo)
	req.setParams(reqParams)
	return nil
}

func (svc *serverService) invoke(ctx core.RequestContext) error {
	for _, svcStruct := range svc.middleware {
		log.Trace(ctx, "Invoking middleware service", "name", svcStruct.name, "params configured", svcStruct.impl.request.params)
		/*req, err := svc.createRequest(ctx, svcStruct, request)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}*/
		err := svc.service.Invoke(ctx)
		if err != nil {
			log.Trace(ctx, "got error ", "service name", svc.name)
			return errors.WrapError(ctx, err)
		}

		res := ctx.GetResponse()
		if res != nil {
			log.Trace(ctx, "Got response ", "service name", svc.name, "res", res)
			return nil
		}
	}
	return nil
}

func (svc *serverService) stop(ctx core.ServerContext) error {
	return svc.service.Stop(ctx)
}
func (svc *serverService) unload(ctx core.ServerContext) error {
	return svc.service.Unload(ctx)
}
