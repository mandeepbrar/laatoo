package laatoormi

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	RMI_ERROR_MISSING_ROUTER        = "Rmi_Error_Missing_Router"
	RMI_ERROR_MISSING_METHODS       = "Rmi_Error_Missing_Methods"
	RMI_ERROR_INCORRECT_METHOD_CONF = "Rmi_Error_Incorrect_Method_Conf"
	RMI_ERROR_MISSING_DATASVC       = "Rmi_Error_Missing_Data"
	RMI_ERROR_MISSING_PATH          = "Rmi_Error_Missing_Path"
	RMI_ERROR_MISSING_METHODNAME    = "Rmi_Error_Missing_Methodname"
)

func init() {
	errors.RegisterCode(RMI_ERROR_MISSING_ROUTER, errors.FATAL, fmt.Errorf("Router not found in rmi service."), LOGGING_CONTEXT)
	errors.RegisterCode(RMI_ERROR_MISSING_METHODS, errors.FATAL, fmt.Errorf("Methods not found in rmi service."), LOGGING_CONTEXT)
	errors.RegisterCode(RMI_ERROR_INCORRECT_METHOD_CONF, errors.FATAL, fmt.Errorf("Incorrect method conf in rmi service."), LOGGING_CONTEXT)
	errors.RegisterCode(RMI_ERROR_MISSING_DATASVC, errors.FATAL, fmt.Errorf("Data service not found for Rmi service."), LOGGING_CONTEXT)
	errors.RegisterCode(RMI_ERROR_MISSING_PATH, errors.FATAL, fmt.Errorf("Path not provided for method in rmi service."), LOGGING_CONTEXT)
	errors.RegisterCode(RMI_ERROR_MISSING_METHODNAME, errors.FATAL, fmt.Errorf("Method name not provided for method in rmi service."), LOGGING_CONTEXT)
}
