package errors

import (
	"fmt"
)

const (
	CORE_ERROR_MISSING_SERVICE = "Core_Missing_Service"
	CORE_ERROR_MISSING_ARG     = "Core_Missing_Arg"
	CORE_ERROR_RES_NOT_FOUND   = "Core_Resource_Not_Found"
	CORE_ERROR_TYPE_MISMATCH   = "Core_Type_Mismatch"
	CORE_ERROR_NOT_IMPLEMENTED = "Core_Not_Implemented"
)

func init() {
	RegisterCode(CORE_ERROR_MISSING_SERVICE, FATAL, fmt.Errorf("Expected service is missing."), "core.errors")
	RegisterCode(CORE_ERROR_MISSING_ARG, FATAL, fmt.Errorf("All arguments have not been provided for the call."), "core.errors")
	RegisterCode(CORE_ERROR_RES_NOT_FOUND, FATAL, fmt.Errorf("Requested resource was not found."), "core.errors")
	RegisterCode(CORE_ERROR_TYPE_MISMATCH, FATAL, fmt.Errorf("Type Mismatch."), "core.errors")
	RegisterCode(CORE_ERROR_NOT_IMPLEMENTED, FATAL, fmt.Errorf("Method has not been implemented by this service."), "core.errors")
}
