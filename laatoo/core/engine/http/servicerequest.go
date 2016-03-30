package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"io/ioutil"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net/http"
)

func processServiceRequest(ctx core.ServerContext, method string, routename string, svc core.Service, service string, dataObjectName string, isdataObject bool, isdataCollection bool, dataObjectCreator core.ObjectCreator, dataObjectCollectionCreator core.ObjectCollectionCreator, routeParams map[string]int) core.HandlerFunc {
	return func(webctx core.RequestContext) error {
		var reqData interface{}
		var err error
		engineContext := webctx.EngineContext().(echo.Context)
		if method == CONF_ROUTE_METHOD_INVOKE {
			service = engineContext.QueryParam(CONF_ROUTE_SERVICE)
			if len(service) == 0 {
				return errors.ThrowError(webctx, errors.CORE_ERROR_MISSING_ARG, "Missing argument", CONF_ROUTE_SERVICE)
			}
			svc, err = webctx.GetService(service)
			if err != nil || svc == nil {
				return errors.RethrowError(webctx, errors.CORE_ERROR_BAD_ARG, err, "No such service has been created", service)
			}
		}
		log.Logger.Trace(webctx, "Received request ", "route", routename, "service", service)
		if isdataObject {
			if dataObjectName == CONF_STRINGMAP_DATA_OBJECT {
				mapobj := make(map[string]interface{}, 10)
				reqData = &mapobj
			} else {
				if isdataCollection {
					reqData, err = dataObjectCollectionCreator(webctx, nil)
					if err != nil {
						return errors.WrapError(webctx, err)
					}
				} else {
					reqData, err = dataObjectCreator(webctx, nil)
					if err != nil {
						return errors.WrapError(webctx, err)
					}
				}
			}
			err = engineContext.Bind(reqData)
			if err != nil {
				return errors.WrapError(webctx, err)
			}
			log.Logger.Trace(webctx, "Request data bound ", "data", reqData)
		} else {
			reqData, err = ioutil.ReadAll(engineContext.Request().(engine.Request).Body())
			if err != nil {
				return errors.WrapError(webctx, err)
			}
		}
		log.Logger.Trace(webctx, "Invoking service ", "router", routename, "service", service, "svc", svc)
		reqctx := webctx.SubContext(service, svc.GetConf())
		reqctx.SetRequestBody(reqData)
		if routeParams != nil {
			for param, index := range routeParams {
				paramVal := engineContext.P(index)
				reqctx.Set(param, paramVal)
			}
		} else {
			paramNames := engineContext.ParamNames()
			for _, param := range paramNames {
				paramVal := engineContext.Param(param)
				reqctx.Set(param, paramVal)
			}
		}
		queryParams := engineContext.QueryParams()
		for param, val := range queryParams {
			reqctx.Set(param, val)
		}
		err = svc.Invoke(reqctx)
		if err != nil {
			return errors.WrapError(webctx, err)
		}
		resp := reqctx.GetResponse()
		return processResponse(webctx, resp, engineContext)
	}
}

func processStreamServiceRequest(ctx core.ServerContext, method string, routename string, svc core.Service, service string, routeParams map[string]int) core.HandlerFunc {
	return func(webctx core.RequestContext) error {
		log.Logger.Trace(webctx, "Received request ", "route", routename, "service", service)
		engineContext := webctx.EngineContext().(echo.Context)
		reqStream := engineContext.Request().Body()
		var err error
		log.Logger.Trace(webctx, "Invoking service ", "router", routename, "service", service)
		reqctx := webctx.SubContext(service, svc.GetConf())
		reqctx.SetRequestBody(reqStream)
		if routeParams != nil {
			for param, index := range routeParams {
				paramVal := engineContext.P(index)
				reqctx.Set(param, paramVal)
			}
		} else {
			paramNames := engineContext.ParamNames()
			for _, param := range paramNames {
				paramVal := engineContext.Param(param)
				reqctx.Set(param, paramVal)
			}
		}
		queryParams := engineContext.QueryParams()
		for param, val := range queryParams {
			reqctx.Set(param, val)
		}
		err = svc.Invoke(reqctx)
		if err != nil {
			return errors.WrapError(webctx, err)
		}
		resp := reqctx.GetResponse()
		return processResponse(webctx, resp, engineContext)
	}
}

func processResponse(ctx core.RequestContext, resp *core.ServiceResponse, engineContext echo.Context) error {
	if resp != nil {
		switch resp.Status {
		case core.StatusSuccess:
			if resp.Data != nil {
				/****TODO***********/
				return engineContext.JSON(http.StatusOK, resp.Data)
			}
		case core.StatusServeFile:
			return engineContext.File(resp.Data.(string))
		case core.StatusServeBytes:
			log.Logger.Trace(ctx, " service returning bytes")
			bytestoreturn := *resp.Data.(*[]byte)
			_, err := engineContext.Response().Write(bytestoreturn)
			if err != nil {
				return err
			}
			engineContext.Response().WriteHeader(http.StatusOK)
			return nil
		case core.StatusUnauthorized:
			return engineContext.NoContent(http.StatusForbidden)
		case core.StatusNotFound:
			return engineContext.NoContent(http.StatusNotFound)
		case core.StatusRedirect:
			return engineContext.Redirect(http.StatusTemporaryRedirect, resp.Data.(string))
		}
	}
	return engineContext.NoContent(http.StatusOK)

}
