package security

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	//"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/auth"
	//"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"net/http"
	"reflect"
)

const (
	CONF_SECURITY_REMOTEAUTHSERVER = "remoteauthserver"
	CONF_SECURITY_REMOTEROLESURL   = "rolesservice"
	CONF_SECURITY_DOMAINIDENTIFIER = "identifier"
)

type remoteSecurityHandler struct {
	adminRole        string
	authHeader       string
	domainidentifier []byte
	pvtKey           []byte
	rolesMap         map[string]auth.Role
	roleObject       string
	remoteAuthServer string
	rolesService     string
	rolePermissions  map[string]bool
}

func NewRemoteSecurityHandler(ctx core.ServerContext, conf config.Config, adminrole string, authHeader string, roleObject string) (SecurityPlugin, error) {
	rsh := &remoteSecurityHandler{adminRole: adminrole, authHeader: authHeader, roleObject: roleObject}
	rsh.rolesMap = make(map[string]auth.Role, 10)
	//map containing roles and permissions
	rsh.rolePermissions = make(map[string]bool, 50)

	remoteAuthServer, ok := conf.GetString(CONF_SECURITY_REMOTEAUTHSERVER)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_SECURITY_REMOTEAUTHSERVER)
	}
	rsh.remoteAuthServer = remoteAuthServer

	rolesService, ok := conf.GetString(CONF_SECURITY_REMOTEROLESURL)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_SECURITY_REMOTEROLESURL)
	}
	rsh.rolesService = rolesService

	pvtKeyPath, ok := conf.GetString(config.CONF_PVTKEYPATH)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_PVTKEYPATH)
	}
	pvtKey, err := utils.LoadPrivateKey(pvtKeyPath)
	if err != nil {
		return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "conf", config.CONF_PVTKEYPATH)
	}
	/*rsh.pvtKey, err = x509.MarshalPKIXPublicKey(pvtKey)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}*/

	domainidentifier, ok := conf.GetString(CONF_SECURITY_DOMAINIDENTIFIER)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_DOMAINIDENTIFIER)
	}
	h := sha256.New()
	h.Write([]byte(domainidentifier))
	d := h.Sum(nil)
	signedIdentifier, err := rsa.SignPKCS1v15(rand.Reader, pvtKey, crypto.SHA256, d)
	if err != nil {
		fmt.Errorf("could not sign request: %v", err)
	}
	rsh.domainidentifier = signedIdentifier
	return rsh, nil
}

func (rsh *remoteSecurityHandler) Start(ctx core.ServerContext) error {
	err := rsh.loadRoles(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (rsh *remoteSecurityHandler) HasPermission(ctx core.RequestContext, perm string) bool {
	return hasPermission(ctx, perm, rsh.rolePermissions)
}

func (rsh *remoteSecurityHandler) loadRoles(ctx core.ServerContext) error {
	client := ctx.HttpClient()
	log.Logger.Error(ctx, "******************client", "client", client, "remoteAuthServer", rsh.remoteAuthServer)
	resp, err := client.Post(rsh.remoteAuthServer, "application/json", bytes.NewBuffer(rsh.domainidentifier))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if resp.StatusCode != 200 {
		//if the remote system did not allow auth
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Could not reach remote server", rsh.remoteAuthServer, "Status", resp.StatusCode)
	} else {
		//get token from remote system
		token := resp.Header.Get(rsh.authHeader)
		data := make(map[string]interface{})
		dat, _ := json.Marshal(data)
		log.Logger.Error(ctx, "******************client", "data", string(dat))

		req, err := http.NewRequest("POST", rsh.rolesService, bytes.NewBuffer(dat))
		req.Header.Add(rsh.authHeader, token)
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		body, err := ioutil.ReadAll(res.Body)
		roles, _ := ctx.CreateCollection(rsh.roleObject, 0, nil)
		err = json.Unmarshal(body, roles)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		//get the response
		if resp.StatusCode != 200 {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Could not get roles from remote server", rsh.remoteAuthServer, "Status", resp.StatusCode)
		}

		rolesVal := reflect.ValueOf(roles).Elem()
		for i := 0; i < rolesVal.Len(); i++ {
			role := rolesVal.Index(i).Addr().Interface().(auth.Role)
			id := role.GetId()
			log.Logger.Error(ctx, "****************** loaded role", "id", id)
			permissions := role.GetPermissions()
			for _, perm := range permissions {
				key := fmt.Sprintf("%s#%s", id, perm)
				rsh.rolePermissions[key] = true
			}
		}

	}
	return nil
}
