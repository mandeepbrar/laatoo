package laatoopages

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	PAGE_ERROR_MISSING_ROUTER         = "Page_Error_Missing_Router"
	PAGE_ERROR_MISSING_PAGESDIR       = "Page_Error_Missing_PagesDir"
	PAGE_ERROR_MISSING_PAGEPATH       = "Page_Error_Missing_Pagepath"
	PAGE_ERROR_MISSING_DEST           = "Page_Error_Missing_Dest"
	PAGE_ERROR_PAGES_NOT_PROVIDED     = "Page_Error_Pages_Not_Provided"
	PAGE_ERROR_WRONG_PARTIALS         = "Page_Error_Wrong_Partials"
	PAGE_ERROR_WRONG_PARTIALPATH      = "Page_Error_Wrong_Partialpath"
	PAGE_ERROR_WRONG_PARTIALFILE      = "Page_Error_Wrong_Partialfile"
	PAGE_ERROR_ACTIONSVC_NOT_PROVIDED = "Page_Error_Actionsvc_Not_Provided"
)

func init() {
	errors.RegisterCode(PAGE_ERROR_MISSING_ROUTER, errors.FATAL, fmt.Errorf("Router not found in page service."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_MISSING_PAGESDIR, errors.FATAL, fmt.Errorf("Pages directory not provided to page service."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_MISSING_PAGEPATH, errors.FATAL, fmt.Errorf("Path of a page has not been provided."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_MISSING_DEST, errors.FATAL, fmt.Errorf("Dest of a page has not been provided."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_PAGES_NOT_PROVIDED, errors.FATAL, fmt.Errorf("Pages directive has not been provided to the page service."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_WRONG_PARTIALS, errors.FATAL, fmt.Errorf("Wrong data provided for partials of the page."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_WRONG_PARTIALPATH, errors.FATAL, fmt.Errorf("Wrong path provided for partial."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_WRONG_PARTIALFILE, errors.FATAL, fmt.Errorf("Wrong file provided for partial."), LOGGING_CONTEXT)
	errors.RegisterCode(PAGE_ERROR_ACTIONSVC_NOT_PROVIDED, errors.ERROR, fmt.Errorf("Actions Service not provided to Page Service."), LOGGING_CONTEXT)
}
