package laatoocore

import (
	"laatoosdk/core"
	"laatoosdk/log"
	"net/http"
)

func (router *Router) authorize(ctx core.Context, conf map[string]interface{}) (bool, error) {
	if ctx.IsAdmin() {
		return true, nil
	}
	if conf != nil {
		auth, authok := conf[CONF_AUTHORIZATION]
		if authok {
			authMap := auth.(map[string]interface{})
			for k, v := range authMap {
				switch k {
				case "functional":
					if !router.environment.HasPermission(ctx, v.(string)) {
						log.Logger.Trace(ctx, "core.router", "Denying permission to user", "permission", v)
						err := ctx.NoContent(http.StatusForbidden)
						if err != nil {
							return false, nil
						}
						return false, nil
					}
				case "method":
					methodConf := v.(map[string]interface{})
					methodName := methodConf["methodname"]
					if methodName != nil {
						method, err := GetMethod(ctx, methodName.(string))
						if err == nil {
							retErr := method(ctx, methodConf)
							if retErr != nil {
								err := ctx.NoContent(http.StatusForbidden)
								if err != nil {
									return false, nil
								}
								return false, nil
							}
							return true, nil
						}
					}
					return false, nil
				}
			}
		}
	}
	return true, nil
}
