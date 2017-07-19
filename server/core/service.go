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
)

type objectType int

const (
	stringmap objectType = iota
	bytes
	files
	stringtype
	custom
)

type service struct {
	name                        string
	service                     core.Service
	conf                        config.Config
	factory                     server.Factory
	owner                       *serviceManager
	middleware                  []*service
	paramValues                 map[string]interface{}
	info                        *core.ServiceInfo
	svrContext                  *serverContext
	dataObjectCreator           core.ObjectCreator
	dataObjectCollectionCreator core.ObjectCollectionCreator
	dataObjectType              objectType
	codecs                      map[string]core.Codec
}

func (svc *service) initialize(ctx core.ServerContext, conf config.Config) error {
	svc.codecs = map[string]core.Codec{"json": codecs.NewJsonCodec(), "bin": codecs.NewBinaryCodec(), "proto": codecs.NewProtobufCodec()}
	svc.info = svc.service.Info()
	if err := svc.processInfo(ctx); err != nil {
		return err
	}

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

	err := svc.service.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	//svc.info = info
	return nil
}

func (svc *service) start(ctx core.ServerContext) error {
	middleware := make([]*service, 0)
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
	return svc.service.Start(ctx)
}

func (svc *service) processInfo(ctx core.ServerContext) error {
	switch svc.info.Request.DataType {
	case constants.CONF_OBJECT_STRINGMAP:
		svc.dataObjectType = stringmap
	case constants.CONF_OBJECT_BYTES:
		svc.dataObjectType = bytes
	case constants.CONF_OBJECT_STRING:
		svc.dataObjectType = stringtype
	case constants.CONF_OBJECT_FILES:
		svc.dataObjectType = files
	default:
		svc.dataObjectType = custom
	}

	if !svc.info.Request.Streaming && (svc.dataObjectType == custom) {
		if svc.info.Request.IsCollection {
			dataObjectCollectionCreator, err := ctx.GetObjectCollectionCreator(svc.info.Request.DataType)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", svc.info.Request.DataType)
			}
			svc.dataObjectCollectionCreator = dataObjectCollectionCreator
		} else {
			dataObjectCreator, err := ctx.GetObjectCreator(svc.info.Request.DataType)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", svc.info.Request.DataType)
			}
			svc.dataObjectCreator = dataObjectCreator
		}
	}
	log.Trace(ctx, "Processed info for service", "name", svc.name, "conf", svc.info)
	return nil
}

func (svc *service) handleRequest(ctx core.RequestContext, vals map[string]interface{}, body []byte) (*core.ServiceResponse, error) {
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

	if !svc.info.Request.Streaming {
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
			if svc.info.Request.IsCollection {
				reqData = svc.dataObjectCollectionCreator(5)
			} else {
				reqData = svc.dataObjectCreator()
			}
		}
	}
	if err := codec.Unmarshal(body, reqData); err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	req := ctx.CreateRequest()
	req.SetBody(reqData)

	if err := svc.populateParams(ctx, vals, req); err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	return svc.invoke(ctx, req)
}

func (svc *service) handleStreamedRequest(ctx core.RequestContext, vals map[string]interface{}, body io.ReadCloser) (*core.ServiceResponse, error) {
	req := ctx.CreateRequest()
	req.SetBody(body)
	if err := svc.populateParams(ctx, vals, req); err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	return svc.invoke(ctx, req)
}

func (svc *service) populateParams(ctx core.RequestContext, vals map[string]interface{}, request core.ServiceRequest) error {
	reqParams := make(core.ServiceParamsMap)
	for _, param := range svc.info.Request.Params {
		reqParam := &core.ServiceParam{}
		*reqParam = *param
		val, ok := vals[param.Name]
		if ok {
			reqParam.Value = val
		} else {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "Argument", param.Name)
		}
		reqParams[param.Name] = reqParam
	}
	request.SetParams(reqParams)
	return nil
}

func (svc *service) invoke(ctx core.RequestContext, request core.ServiceRequest) (*core.ServiceResponse, error) {
	for _, svcStruct := range svc.middleware {
		log.Trace(ctx, "Invoking middleware service", "name", svcStruct.name)
		/*req, err := svc.createRequest(ctx, svcStruct, request)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}*/
		res, err := svc.service.Invoke(ctx, request)
		if err != nil {
			log.Trace(ctx, "got error ", "service name", svc.name)
			return nil, errors.WrapError(ctx, err)
		}
		if res != nil {
			log.Trace(ctx, "Got response ", "service name", svc.name)
			return res, nil
		}
	}
	return nil, nil
}
