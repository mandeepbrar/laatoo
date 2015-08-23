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
	errors.RegisterCode(VIEW_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in view service."))
	errors.RegisterCode(VIEW_ERROR_MISSING_DATASVC, errors.PANIC, fmt.Errorf("Data service not provided for view."))
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEWS, errors.PANIC, fmt.Errorf("Views not provided for the view service."))
	errors.RegisterCode(VIEW_ERROR_INCORRECT_VIEW_CONF, errors.PANIC, fmt.Errorf("Incorrect config for views."))
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEWNAME, errors.PANIC, fmt.Errorf("Name not provided for view."))
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEWPATH, errors.PANIC, fmt.Errorf("Path not provided for view."))
	errors.RegisterCode(VIEW_ERROR_MISSING_VIEW, errors.PANIC, fmt.Errorf("No such view."))
	errors.RegisterCode(VIEW_ERROR_MISSING_ARG, errors.PANIC, fmt.Errorf("Missing argument."))
}
