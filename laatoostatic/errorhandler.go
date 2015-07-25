package laatoostatic

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	STATIC_ERROR_MISSING_ROUTER    = "Static_Error_Missing_Router"
	STATIC_ERROR_MISSING_PUBLICDIR = "Static_Error_Missing_PublicDir"
)

func init() {
	errors.RegisterCode(STATIC_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in static file service."))
	errors.RegisterCode(STATIC_ERROR_MISSING_PUBLICDIR, errors.ERROR, fmt.Errorf("Public directory not provided to static file service."))
}
