package laatoologin

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	LOGIN_ERROR_MISSING_ROUTER    = "Login_Error_Missing_Router"
	LOGIN_ERROR_MISSING_PUBLICDIR = "Login_Error_Missing_PublicDir"
)

func init() {
	errors.RegisterCode(LOGIN_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in login service."))
	errors.RegisterCode(LOGIN_ERROR_MISSING_PUBLICDIR, errors.ERROR, fmt.Errorf("Public directory not provided to login service."))
}
