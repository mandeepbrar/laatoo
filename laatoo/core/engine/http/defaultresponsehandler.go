package http

import (
	"github.com/labstack/echo"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"net/http"
)

func (router *Router) HandleResponse(ctx core.RequestContext, resp *core.ServiceResponse) error {
	engineContext := ctx.EngineContext().(echo.Context)
	engineResp := engineContext.Response()
	if resp != nil {
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
						engineResp.Header().Set(core.ContentType, val.(string))
					case core.LastModified:
						engineResp.Header().Set(core.LastModified, val.(string))
					}
				}
			}
			engineContext.Response().WriteHeader(http.StatusOK)
			bytestoreturn := *resp.Data.(*[]byte)
			_, err := engineContext.Response().Write(bytestoreturn)
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
		case core.StatusRedirect:
			return engineContext.Redirect(http.StatusTemporaryRedirect, resp.Data.(string))
		}
	}
	return engineContext.NoContent(http.StatusOK)
}
