package http

import (
	"laatoo/core/engine/http/net"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"net/http"
)

func (router *Router) HandleResponse(ctx core.RequestContext) error {
	log.Logger.Trace(ctx, "Returning request with default response handler")
	resp := ctx.GetResponse()
	engineContext := ctx.EngineContext().(net.WebContext)
	if resp != nil {
		log.Logger.Trace(ctx, "Returning request with status", "Status", resp.Status)
		switch resp.Status {
		case core.StatusSuccess:
			if resp.Data != nil {
				/****TODO***********/
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
						engineContext.SetHeader(core.ContentType, val.(string))
					case core.LastModified:
						engineContext.SetHeader(core.LastModified, val.(string))
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
			return engineContext.NoContent(http.StatusForbidden)
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
