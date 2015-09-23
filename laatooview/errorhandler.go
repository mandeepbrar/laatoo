package laatooview

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	VIEW_ERROR_MISSING_ROUTER      = "View_Error_Missing_Router"
	VIEW_ERROR_MISSING_DATASVC     = "View_Error_Missing_DataSvc"
	VIEW_ERROR_MISSING_VIEWS       = "View_Error_Missing_Views"
	VIEW_ERROR_INCORRECT_VIEW_CONF = "View_Error_Incorrect_View_Conf"
	VIEW_ERROR_MISSING_VIEWNAME    = "View_Error_Incorrect_Viewname"
	VIEW_ERROR_MISSING_VIEWPATH    = "View_Error_Incorrect_Viewpath"
	VIEW_ERROR_MISSING_VIEW        = "View_Error_Incorrect_View"
	VIEW_ERROR_MISSING_ARG         = "View_Error_Missing_Arg"
)

func init() {
	errors.RegisterCode(VIEW_ERROR_MISSING_ROUTER, errors.ERROR, fmt.Errorf("Router not found in view service."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_MISSING_DATASVC, errors.FATAL, fmt.Errorf("Data service not provided for view."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEWS, errors.FATAL, fmt.Errorf("Views not provided for the view service."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_INCORRECT_VIEW_CONF, errors.FATAL, fmt.Errorf("Incorrect config for views."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEWNAME, errors.FATAL, fmt.Errorf("Name not provided for view."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEWPATH, errors.FATAL, fmt.Errorf("Path not provided for view."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEW, errors.ERROR, fmt.Errorf("No such view."), LOGGING_CONTEXT)
	errors.RegisterCode(VIEW_ERROR_MISSING_ARG, errors.ERROR, fmt.Errorf("Missing argument."), LOGGING_CONTEXT)
}
