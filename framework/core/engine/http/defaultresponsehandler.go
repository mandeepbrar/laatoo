package http

import (
	"fmt"
	"laatoo/framework/core/common"
	"laatoo/framework/core/engine/http/net"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"net/http"
	"strings"
)

type defaultResponseHandler struct {
	*common.Context
}

func DefaultResponseHandler(ctx core.ServerContext) *defaultResponseHandler {
	return nil
}

func (rh *defaultResponseHandler) HandleResponse(ctx core.RequestContext) error {
	log.Logger.Trace(ctx, "Returning request with default response handler")
	resp := ctx.GetResponse()
	if resp == nil {
		resp = core.StatusSuccessResponse
	}
	engineContext := ctx.EngineRequestContext().(net.WebContext)
	if resp != nil {
		log.Logger.Trace(ctx, "Returning request with status", "Status", resp.Status)
		switch resp.Status {
		case core.StatusSuccess:
			if resp.Data != nil {
				if resp.Info != nil {
					keyNames := make([]string, len(resp.Info))
					i := 0
					for key, val := range resp.Info {
						engineContext.SetHeader(key, fmt.Sprint(val))
						keyNames[i] = key
						i++
					}
					engineContext.SetHeader("Access-Control-Expose-Headers", strings.Join(keyNames, ","))
				}
				log.Logger.Debug(ctx, "Returning request with data", "time", ctx.GetElapsedTime())
				log.Logger.Trace(ctx, "Returned data", "data", resp.Data)
				return engineContext.JSON(http.StatusOK, resp.Data)
			} else {
				log.Logger.Debug(ctx, "Returning request without content")
				return engineContext.NoContent(http.StatusOK)
			}
		case core.StatusServeFile:
			fil := resp.Data.(string)
			log.Logger.Debug(ctx, "Returning serve file", "file", fil)
			return engineContext.File(fil)
		case core.StatusServeBytes:
			log.Logger.Trace(ctx, " service returning bytes")
			if resp.Info != nil {
				for key, val := range resp.Info {
					switch key {
					case core.ContentType:
						engineContext.SetHeader(core.ContentType, fmt.Sprint(val))
					case core.ContentEncoding:
						log.Logger.Trace(ctx, " sending encoding", "inf", val)
						engineContext.SetHeader(core.ContentEncoding, fmt.Sprint(val))
					case core.LastModified:
						engineContext.SetHeader(core.LastModified, fmt.Sprint(val))
					}
				}
			}
			bytestoreturn := *resp.Data.(*[]byte)
			log.Logger.Trace(ctx, "Returning bytes", "length", len(bytestoreturn))
			log.Logger.Debug(ctx, "Returning bytes")
			_, err := engineContext.Write(bytestoreturn)
			if err != nil {
				return err
			}
			return nil
		case core.StatusNotModified:
			log.Logger.Debug(ctx, "Returning not modified")
			return engineContext.NoContent(http.StatusNotModified)
		case core.StatusUnauthorized:
			log.Logger.Debug(ctx, "Returning unauthorized")
			return engineContext.NoContent(http.StatusUnauthorized)
		case core.StatusNotFound:
			log.Logger.Debug(ctx, "Returning not found")
			return engineContext.NoContent(http.StatusNotFound)
		case core.StatusBadRequest:
			log.Logger.Debug(ctx, "Returning bad request")
			return engineContext.NoContent(http.StatusBadRequest)
		case core.StatusRedirect:
			log.Logger.Debug(ctx, "Returning redirect")
			return engineContext.Redirect(http.StatusTemporaryRedirect, resp.Data.(string))
		}
	}
	log.Logger.Debug(ctx, "Returning request without content")
	return engineContext.NoContent(http.StatusOK)
}
