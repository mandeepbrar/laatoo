package laatoopages

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	PAGE_ERROR_MISSING_ROUTER     = "Page_Error_Missing_Router"
	PAGE_ERROR_MISSING_PUBLICDIR  = "Page_Error_Missing_PublicDir"
	PAGE_ERROR_MISSING_PAGEPATH   = "Page_Error_Missing_Pagepath"
	PAGE_ERROR_MISSING_DEST       = "Page_Error_Missing_Dest"
	PAGE_ERROR_PAGES_NOT_PROVIDED = "Page_Error_Pages_Not_Provided"
	PAGE_ERROR_WRONG_PARAMS       = "Page_Error_Wrong_Params"
)

func init() {
	errors.RegisterCode(PAGE_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in page service."))
	errors.RegisterCode(PAGE_ERROR_MISSING_PUBLICDIR, errors.ERROR, fmt.Errorf("Public directory not provided to page service."))
	errors.RegisterCode(PAGE_ERROR_MISSING_PAGEPATH, errors.ERROR, fmt.Errorf("Path of a page has not been provided."))
	errors.RegisterCode(PAGE_ERROR_MISSING_DEST, errors.ERROR, fmt.Errorf("Dest of a page has not been provided."))
	errors.RegisterCode(PAGE_ERROR_PAGES_NOT_PROVIDED, errors.ERROR, fmt.Errorf("Pages have not been provided to the page service."))
	errors.RegisterCode(PAGE_ERROR_WRONG_PARAMS, errors.ERROR, fmt.Errorf("Wrong parameters have been provided for the page."))
}
