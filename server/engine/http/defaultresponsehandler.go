package http

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"laatoo/server/engine/http/net"
	"net/http"
	"strings"
)

type defaultResponseHandler struct {
}

func DefaultResponseHandler(ctx core.ServerContext) *defaultResponseHandler {
	return nil
}

func (rh *defaultResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response) error {
	log.Trace(ctx, "Returning request with default response handler", "resp", resp)
	if resp == nil {
		resp = core.StatusSuccessResponse
	}
	engineContext := ctx.EngineRequestContext().(net.WebContext)
	if resp != nil {
		var respData interface{}
		log.Trace(ctx, "Returning request with status", "Status", resp.Status)
		switch resp.Status {
		case core.StatusSuccess:
			if resp.Data != nil {
				respData = resp.Data["Data"]
				keyNames := make([]string, len(resp.Data))
				i := 0
				for key, val := range resp.Data {
					if key != "Data" {
						engineContext.SetHeader(key, fmt.Sprint(val))
						keyNames[i] = key
						i++
					}
				}
				engineContext.SetHeader("Access-Control-Expose-Headers", strings.Join(keyNames, ","))
				log.Trace(ctx, "Returned data", "data", respData)
				return engineContext.JSON(http.StatusOK, respData)
			} else {
				log.Debug(ctx, "Returning request without content")
				return engineContext.NoContent(http.StatusOK)
			}
		case core.StatusServeFile:
			if resp.Data != nil {
				respData = resp.Data["Data"]
				fil := respData.(string)
				log.Debug(ctx, "Returning serve file", "file", fil)
				return engineContext.File(fil)
			}
		case core.StatusServeBytes:
			if resp.Data != nil {
				respData = resp.Data["Data"]
				log.Trace(ctx, " service returning bytes")
				val, ok := resp.Data[core.ContentType]
				if ok {
					engineContext.SetHeader(core.ContentType, fmt.Sprint(val))
				}
				val, ok = resp.Data[core.ContentEncoding]
				if ok {
					engineContext.SetHeader(core.ContentEncoding, fmt.Sprint(val))
				}
				val, ok = resp.Data[core.LastModified]
				if ok {
					engineContext.SetHeader(core.LastModified, fmt.Sprint(val))
				}
				bytestoreturn := *respData.(*[]byte)
				log.Debug(ctx, "Returning bytes", "length", len(bytestoreturn))
				_, err := engineContext.Write(bytestoreturn)
				if err != nil {
					return err
				}
			}
			return nil
		case core.StatusNotModified:
			log.Debug(ctx, "Returning not modified")
			return engineContext.NoContent(http.StatusNotModified)
		case core.StatusUnauthorized:
			log.Debug(ctx, "Returning unauthorized")
			return engineContext.NoContent(http.StatusUnauthorized)
		case core.StatusNotFound:
			log.Debug(ctx, "Returning not found")
			return engineContext.NoContent(http.StatusNotFound)
		case core.StatusBadRequest:
			log.Debug(ctx, "Returning bad request")
			return engineContext.NoContent(http.StatusBadRequest)
		case core.StatusRedirect:
			log.Debug(ctx, "Returning redirect")
			return engineContext.Redirect(http.StatusTemporaryRedirect, respData.(string))
		}
	}
	log.Debug(ctx, "Returning request without content")
	return engineContext.NoContent(http.StatusOK)
}

func (proxy *defaultResponseHandler) Reference() core.ServerElement {
	anotherref := proxy
	return anotherref
}

func (proxy *defaultResponseHandler) GetProperty(name string) interface{} {
	return nil
}

func (proxy *defaultResponseHandler) GetName() string {
	return "DefaultResponseHandler"
}
func (proxy *defaultResponseHandler) GetType() core.ServerElementType {
	return core.ServerElementServiceResponseHandler
}
