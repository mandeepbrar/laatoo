package laatoocore

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	CORE_ERROR_MISSING_SERVICE_NAME   = "Core_Missing_Service_Name"
	CORE_ERROR_SERVICE_CREATION       = "Core_Service_Creation"
	CORE_ERROR_SERVICE_INITIALIZATION = "Core_Service_Initialization"
	CORE_ERROR_SERVICE_NOT_FOUND      = "Core_Service_Not_Found"
	CORE_ERROR_SERVICE_NOT_STARTED    = "Core_Service_Not_Started"
	CORE_ERROR_PROVIDER_NOT_FOUND     = "Core_Provider_Not_Found"
	CORE_ERROR_OBJECT_NOT_CREATED     = "Core_Object_Not_Created"
	CORE_ENVIRONMENT_NOT_CREATED      = "Core_Environment_Not_Created"
	CORE_ENVIRONMENT_NOT_INITIALIZED  = "Core_Environment_Not_Initialized"
	CORE_SERVERADD_NOT_FOUND          = "Core_ServerAdd_Not_Found"
)

func init() {
	errors.RegisterCode(CORE_ERROR_MISSING_SERVICE_NAME, errors.FATAL, fmt.Errorf("Service name is missing or wrong."))
	errors.RegisterCode(CORE_ERROR_SERVICE_CREATION, errors.FATAL, fmt.Errorf("Service could not be created."))
	errors.RegisterCode(CORE_ERROR_SERVICE_INITIALIZATION, errors.FATAL, fmt.Errorf("Service could not be initialized."))
	errors.RegisterCode(CORE_ERROR_SERVICE_NOT_FOUND, errors.FATAL, fmt.Errorf("Service not found."))
	errors.RegisterCode(CORE_ERROR_SERVICE_NOT_STARTED, errors.FATAL, fmt.Errorf("Service not started."))
	errors.RegisterCode(CORE_ERROR_PROVIDER_NOT_FOUND, errors.FATAL, fmt.Errorf("Provider for object not found."))
	errors.RegisterCode(CORE_ERROR_OBJECT_NOT_CREATED, errors.FATAL, fmt.Errorf("Object could not be created from provider."))
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_CREATED, errors.FATAL, fmt.Errorf("Environment could not be created."))
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_INITIALIZED, errors.FATAL, fmt.Errorf("Environment could not be initialized."))
	errors.RegisterCode(CORE_SERVERADD_NOT_FOUND, errors.FATAL, fmt.Errorf("Server address not provided."))
}
