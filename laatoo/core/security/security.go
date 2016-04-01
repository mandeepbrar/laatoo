package security

import (
/*	"bytes"
		"encoding/json"
		"github.com/labstack/echo"
		"laatoosdk/auth"
		"laatoosdk/config"
		"laatoosdk/core"
		"laatoosdk/utils"
		"net/http"
		"io/ioutil"
	"fmt"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"reflect"*/
)

const (
	CONF_SERVICE_AUTHBYPASS = "bypassauth"
	CONF_SERVICE_USECORS    = "usecors"
	CONF_AUTH_MODE          = "settings.authorization.mode"
	CONF_AUTH_MODE_LOCAL    = "local"
	CONF_AUTH_MODE_REMOTE   = "remote"
	CONF_API_AUTH           = "settings.authorization.apiauth"
	CONF_ROLES_API          = "settings.authorization.rolesapi"
	CONF_SECURITY_SVC       = "settings.authorization.securitysvc"
	CONF_PERMISSIONS_API    = "settings.authorization.permissionsapi"
	CONF_API_PUBKEY         = "settings.authorization.pubkey"
	CONF_API_DOMAIN         = "settings.authorization.domain"
	CONF_SERVICE_CORSHOSTS  = "corshosts"
)

/*



//load role permissions if needed from another environment
func (env *Environment) loadRolePermissions(ctx *Context) error {
	//check the authenticatino mode
	mode := env.Config.GetString(CONF_AUTH_MODE)
	if mode == CONF_AUTH_MODE_REMOTE {
		//load permissions from remote system
		apiauth := env.Config.GetString(CONF_API_AUTH)
		if len(apiauth) == 0 {
			return errors.ThrowError(ctx, AUTH_MISSING_API)
		}
		//authenticate to the remote system using public key
		pubkey := env.Config.GetString(CONF_API_PUBKEY)
		domain := env.Config.GetString(CONF_API_DOMAIN)
		//encrypt system domain and send
		key, err := EncryptWithKey(pubkey, domain)
		if err != nil {
			return errors.WrapError(err)
		}
		client := ctx.HttpClient()
		form := &KeyAuth{Key: key}
		load, _ := json.Marshal(form)
		resp, err := client.Post(apiauth, "application/json", bytes.NewBuffer(load))
		if err != nil {
			return errors.WrapError(err)
		}
		log.Logger.Trace(ctx, "core.env.remoteroles", "Got Response for api key", "Response", resp.StatusCode)
		if resp.StatusCode != 200 {
			//if the remote system did not allow auth
			return errors.ThrowError(ctx, AUTH_APISEC_NOTALLOWED)
		} else {

			//get token from remote system
			token := resp.Header.Get(env.AuthHeader)
			log.Logger.Trace(ctx, "core.env.remoteroles", "Auth token for api key", "Token", token)

			//get the url for remote system
			rolesurl := env.Config.GetString(CONF_ROLES_API)
			if len(rolesurl) == 0 {
				return errors.ThrowError(ctx, CORE_ROLESAPI_NOT_FOUND)
			}
			//create remote system role
			roles, err := CreateCollection(ctx, env.SystemRole)
			if err != nil {
				return errors.WrapError(err)
			}
			req, err := http.NewRequest("GET", rolesurl, nil)
			req.Header.Add(env.AuthHeader, token)
			res, err := client.Do(req)
			if err != nil {
				return errors.WrapError(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			log.Logger.Trace(ctx, "core.env.remoteroles", "result for roles query", "body", body)
			err = json.Unmarshal(body, &roles)

			log.Logger.Trace(ctx, "core.env.remoteroles", "result for roles query", "err", err)
			if err != nil {
				return errors.WrapError(err)
			}
			log.Logger.Trace(ctx, "core.env.remoteroles", "result for roles query", "Status code", resp.StatusCode)
			//get the response
			if resp.StatusCode != 200 {
				return errors.ThrowError(ctx, CORE_ROLESAPI_NOT_FOUND)
			}
			//register the roles and permissions received from auth system
			env.RegisterRoles(ctx, roles)
		}
	} else {

	}
	return nil
}
*/
