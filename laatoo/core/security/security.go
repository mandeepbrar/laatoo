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
