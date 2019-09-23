package http

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/engine/http/net"
	"net/http"
)

func handleResponse(ctx core.RequestContext, resp *core.Response, handleMetaInfo func(net.WebContext, map[string]interface{}) error) error {
	if resp == nil {
		resp = core.StatusSuccessResponse
	}
	engineContext := ctx.EngineRequestContext().(net.WebContext)
	if resp != nil {
		log.Trace(ctx, "Returning request with status", "Status", resp.Status)
		switch resp.Status {
		case core.StatusSuccess:
			if resp.Data != nil {
				err := handleMetaInfo(engineContext, resp.MetaInfo)
				if err!=nil {
					return errors.WrapError(ctx, err)
				}
				return engineContext.JSON(http.StatusOK, resp.Data)
			} else {
				log.Debug(ctx, "Returning request without content")
				return engineContext.NoContent(http.StatusOK)
			}
		case core.StatusServeFile:
			if resp.Data != nil {
				fil := resp.Data.(string)
				log.Debug(ctx, "Returning serve file", "file", fil)
				return engineContext.File(fil)
			}
		case core.StatusServeBytes:
			if resp.Data != nil {
				log.Trace(ctx, " service returning bytes")
				val, ok := resp.MetaInfo[core.ContentType]
				if ok {
					engineContext.SetHeader(core.ContentType, fmt.Sprint(val))
				}
				val, ok = resp.MetaInfo[core.ContentEncoding]
				if ok {
					engineContext.SetHeader(core.ContentEncoding, fmt.Sprint(val))
				}
				val, ok = resp.MetaInfo[core.LastModified]
				if ok {
					engineContext.SetHeader(core.LastModified, fmt.Sprint(val))
				}
				bytestoreturn := *resp.Data.(*[]byte)
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
			return engineContext.Redirect(http.StatusTemporaryRedirect, resp.Data.(string))
		}
	}
	log.Debug(ctx, "Returning request without content")
	return engineContext.NoContent(http.StatusOK)
}
