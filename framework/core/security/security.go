package security

import (
	//	"laatoo/sdk/auth"
	"laatoo/sdk/core"
	/*	"bytes"
			"encoding/json"
			"github.com/labstack/echo"
			"laatoosdk/auth"
			"laatoosdk/config"
			"laatoosdk/utils"
			"net/http"
			"io/ioutil"
		"fmt"
		"laatoosdk/data"
		"laatoosdk/errors"
		"laatoosdk/log"
		"reflect"*/)

/*const (
	CONF_SERVICE_AUTHBYPASS = "bypassauth"
	CONF_AUTH_MODE          = "settings.authorization.mode"
	CONF_AUTH_MODE_LOCAL    = "local"
	CONF_AUTH_MODE_REMOTE   = "remote"
	CONF_API_AUTH           = "settings.authorization.apiauth"
	CONF_ROLES_API          = "settings.authorization.rolesapi"
	CONF_SECURITY_SVC       = "settings.authorization.securitysvc"
	CONF_PERMISSIONS_API    = "settings.authorization.permissionsapi"
	CONF_API_PUBKEY         = "settings.authorization.pubkey"
	CONF_API_DOMAIN         = "settings.authorization.domain"
)*/

type SecurityPlugin interface {
	Start(core.ServerContext) error
	HasPermission(core.RequestContext, string) bool
	AllPermissions(core.RequestContext) []string
}
