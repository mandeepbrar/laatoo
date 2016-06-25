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
				return engineContext.JSON(http.StatusOK, resp.Data)
			} else {
				return engineContext.NoContent(http.StatusOK)
			}
		case core.StatusServeFile:
			return engineContext.File(resp.Data.(string))
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
			_, err := engineContext.Write(bytestoreturn)
			if err != nil {
				return err
			}
			return nil
		case core.StatusNotModified:
			return engineContext.NoContent(http.StatusNotModified)
		case core.StatusUnauthorized:
			return engineContext.NoContent(http.StatusUnauthorized)
		case core.StatusNotFound:
			return engineContext.NoContent(http.StatusNotFound)
		case core.StatusBadRequest:
			return engineContext.NoContent(http.StatusBadRequest)
		case core.StatusRedirect:
			return engineContext.Redirect(http.StatusTemporaryRedirect, resp.Data.(string))
		}
	}
	return engineContext.NoContent(http.StatusOK)
}
