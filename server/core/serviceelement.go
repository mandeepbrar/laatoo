package core

import (
	"io"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/codecs"
	"laatoo/server/constants"
	"reflect"
)

type objectType int

const (
	stringmap objectType = iota
	bytes
	files
	inttype
	stringtype
	stringarr
	booltype
	custom
	none
)

type serverService struct {
	name                        string
	service                     core.Service
	conf                        config.Config
	factory                     server.Factory
	owner                       *serviceManager
	middleware                  []*serverService
	paramValues                 map[string]interface{}
	impl                        *serviceImpl
	info                        core.ServiceInfo
	svrContext                  *serverContext
	dataObjectCreator           core.ObjectCreator
	dataObjectCollectionCreator core.ObjectCollectionCreator
	dataObjectType              objectType
	codecs                      map[string]core.Codec
}

func (svc *serverService) initialize(ctx core.ServerContext, conf config.Config) error {
	//inject service implementation into
	//every service
	impl := newServiceImpl()
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

	err := svc.service.Initialize(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	svc.codecs = map[string]core.Codec{"json": codecs.NewJsonCodec(), "bin": codecs.NewBinaryCodec(), "proto": codecs.NewProtobufCodec()}
	svc.info = svc.service.Info()
	if err := svc.processInfo(ctx, conf, svc.info); err != nil {
		return err
	}

	impl.state = Initialized

	/*svc.metaParams = make(core.ServiceParamsMap)
	svc.paramValues = make(map[string]interface{})
	svcParamsConf, ok := conf.GetSubConfig(constants.CONF_SVCPARAMS)
	if ok {
		staticValuesConf, ok := conf.GetSubConfig(constants.CONF_SVCPARAMS_STATIC)
		if ok {
			values := staticValuesConf.AllConfigurations()
			for _, paramName := range values {
				val, _ = staticValuesConf.Get(paramname)
				svc.paramValues[paramName] = val
			}
		}
	}*/

	//svc.info = info
	return nil
}

func (svc *serverService) start(ctx core.ServerContext) error {
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
	err := svc.service.Start(ctx)
	if err == nil {
		svc.impl.state = Started
	}
	return err
}

func (svc *serverService) processInfo(ctx core.ServerContext, svcconf config.Config, info core.ServiceInfo) error {
	reqInfo := info.GetRequestInfo()

	svcs := info.GetRequiredServices()
	if err := svc.injectServices(ctx, svcconf, svcs); err != nil {
		return err
	}

	if err := svc.impl.processInfo(ctx, svcconf); err != nil {
		return err
	}

	datatype := reqInfo.GetDataType()
	switch datatype {
	case "":
		svc.dataObjectType = none
	case config.CONF_OBJECT_STRINGMAP:
		svc.dataObjectType = stringmap
	case config.CONF_OBJECT_BYTES:
		svc.dataObjectType = bytes
	case config.CONF_OBJECT_STRING:
		svc.dataObjectType = stringtype
	case config.CONF_OBJECT_STRINGARR:
		svc.dataObjectType = stringarr
	case config.CONF_OBJECT_BOOL:
		svc.dataObjectType = booltype
	case config.CONF_OBJECT_FILES:
		svc.dataObjectType = files
	default:
		svc.dataObjectType = custom
	}

	if !reqInfo.IsStream() && (svc.dataObjectType == custom) {
		if reqInfo.IsCollection() {
			dataObjectCollectionCreator, err := ctx.GetObjectCollectionCreator(datatype)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", datatype)
			}
			svc.dataObjectCollectionCreator = dataObjectCollectionCreator
		} else {
			dataObjectCreator, err := ctx.GetObjectCreator(datatype)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", datatype)
			}
			svc.dataObjectCreator = dataObjectCreator
		}
	}
	log.Trace(ctx, "Processed info for service", "name", svc.name, "conf", svc.info)
	return nil
}

func (svc *serverService) injectServices(ctx core.ServerContext, svcconf config.Config, svcsToInject map[string]string) error {
	if svcsToInject == nil {
		return nil
	}
	for confName, fieldName := range svcsToInject {
		svcName, ok := svcconf.GetString(confName)
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

func (svc *serverService) handleEncodedRequest(ctx *requestContext, vals map[string]interface{}, body []byte) (*core.Response, error) {
	if svc.dataObjectType == none {
		return svc.handleRequest(ctx, vals, nil)
	}

	codecname := "json"
	co, ok := vals["encoding"]
	if ok {
		codecname = co.(string)
	}
	codec, ok := svc.codecs[codecname]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_CODEC_NOT_FOUND)
	}

	var reqData interface{}

	reqInfo := svc.service.Info().GetRequestInfo()

	if !reqInfo.IsStream() {
		switch svc.dataObjectType {
		case stringmap:
			mapobj := make(map[string]interface{}, 10)
			reqData = &mapobj
		case bytes:
			reqData = body
		case stringtype:
			reqData = ""
		case files:
			reqData = ""
			//////not required
		default:
			if reqInfo.IsCollection() {
				reqData = svc.dataObjectCollectionCreator(5)
			} else {
				reqData = svc.dataObjectCreator()
			}
		}
	}
	if err := codec.Unmarshal(body, reqData); err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return svc.handleRequest(ctx, vals, reqData)
}

func (svc *serverService) handleRequest(ctx *requestContext, vals map[string]interface{}, body interface{}) (*core.Response, error) {
	req := ctx.createRequest()
	req.setBody(body)
	if err := svc.populateParams(ctx, vals, req); err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	ctx.req = req
	err := svc.invoke(ctx)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return ctx.GetResponse(), nil
}

func (svc *serverService) handleStreamedRequest(ctx *requestContext, vals map[string]interface{}, body io.ReadCloser) (*core.Response, error) {
	req := ctx.createRequest()
	req.setBody(body)
	if err := svc.populateParams(ctx, vals, req); err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	ctx.req = req
	err := svc.invoke(ctx)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return ctx.GetResponse(), nil
}

func (svc *serverService) populateParams(ctx *requestContext, vals map[string]interface{}, req *request) error {
	reqParams := make(map[string]core.Param)
	reqInfo := svc.service.Info().GetRequestInfo()
	params := reqInfo.GetParams()
	for name, svcParam := range params {
		reqParam := &param{}
		*reqParam = *svcParam.(*param)
		val, ok := vals[name]
		if ok {
			switch svcParam.GetDataType() {
			case config.CONF_OBJECT_STRINGMAP:
				reqParam.value, ok = val.(map[string]interface{})
			case config.CONF_OBJECT_INT:
				reqParam.value, ok = val.(int)
			case config.CONF_OBJECT_STRING:
				strval, ok := val.(string)
				if ok {
					reqParam.value = strval
				}
			case config.CONF_OBJECT_STRINGARR:
				reqParam.value, ok = val.([]string)
			case config.CONF_OBJECT_BOOL:
				reqParam.value, ok = val.(bool)
			default:
				reqParam.value = val
			}
			if !ok {
				return errors.BadArg(ctx, name)
			}
		} else {
			return errors.MissingArg(ctx, name)
		}
		reqParams[name] = reqParam
	}
	req.setParams(reqParams)
	return nil
}

func (svc *serverService) invoke(ctx core.RequestContext) error {
	for _, svcStruct := range svc.middleware {
		log.Trace(ctx, "Invoking middleware service", "name", svcStruct.name)
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
			log.Trace(ctx, "Got response ", "service name", svc.name)
			return nil
		}
	}
	return nil
}
