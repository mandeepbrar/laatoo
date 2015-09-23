package laatoocodegen

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	CODEGEN_ERROR_MISSING_ROUTER    = "Codegen_Error_Missing_Router"
	CODEGEN_ERROR_MISSING_PUBLICDIR = "Codegen_Error_Missing_PublicDir"
)

func init() {
	errors.RegisterCode(CODEGEN_ERROR_MISSING_ROUTER, errors.FATAL, fmt.Errorf("Router not found in codegen service."), LOGGING_CONTEXT)
}
