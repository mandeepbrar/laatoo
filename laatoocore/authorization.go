package laatoocore

import (
	"laatoosdk/core"
	"laatoosdk/log"
	"net/http"
)

func (router *Router) authorize(ctx core.Context, conf map[string]interface{}) (bool, error) {
	if conf != nil {
		auth, authok := conf[CONF_AUTHORIZATION]
		if authok {
			authMap := auth.(map[string]interface{})
			log.Logger.Trace(ctx, "core.router", "Testing auth", "authMap", authMap)
			for k, v := range authMap {
				switch k {
				case "functional":
					log.Logger.Trace(ctx, "core.router", "Testing auth", "v", v)
					if !router.environment.HasPermission(ctx, v.(string)) {
						err := ctx.NoContent(http.StatusForbidden)
						return false, err
					}
				case "method":
					method, err := GetMethod(ctx, v.(string))
					if err == nil {
						retErr := method(ctx)
						if retErr != nil {
							return false, retErr
						}
					}
				}
			}
		}
	}
	return true, nil
}
