package server

import (
	"fmt"
	"laatoo/sdk/errors"
)

const (
	CORE_ENVIRONMENT_NOT_CREATED       = "Core_Environment_Not_Created"
	CORE_ENVIRONMENT_NOT_INITIALIZED   = "Core_Environment_Not_Initialized"
	CORE_FACTORY_NOT_CREATED           = "Core_Factory_Not_Created"
	CORE_ERROR_SERVICE_CREATION        = "Core_Service_Creation"
	CORE_ERROR_NOCOMMSVC               = "Core_Error_Nocommsvc"
	CORE_ERROR_SERVICES_NOT_STARTED    = "Core_Services_Not_Started"
	CORE_ERROR_NO_CACHE_SVC            = "Core_No_Cache_Svc"
	CORE_ERROR_INCORRECT_DELIVERY_CONF = "Core_Error_Incorrect_Delivery_Conf"
)

func init() {
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_CREATED, errors.FATAL, fmt.Errorf("Environment could not be created."))
	errors.RegisterCode(CORE_ERROR_NOCOMMSVC, errors.ERROR, fmt.Errorf("Communication Service not enabled."))
	errors.RegisterCode(CORE_FACTORY_NOT_CREATED, errors.ERROR, fmt.Errorf("Factory could not be created."))
	errors.RegisterCode(CORE_ERROR_SERVICE_CREATION, errors.FATAL, fmt.Errorf("Service could not be created."))
	errors.RegisterCode(CORE_ERROR_SERVICES_NOT_STARTED, errors.FATAL, fmt.Errorf("Services could not be started."))
	errors.RegisterCode(CORE_ERROR_NO_CACHE_SVC, errors.FATAL, fmt.Errorf("No cache service has been configured."))
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_INITIALIZED, errors.FATAL, fmt.Errorf("Environment could not be initialized."))
	errors.RegisterCode(CORE_ERROR_INCORRECT_DELIVERY_CONF, errors.FATAL, fmt.Errorf("Delivery methods could not be configured."))
}
