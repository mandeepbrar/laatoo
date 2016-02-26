package laatoostatic

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	STATIC_ERROR_MISSING_ROUTER         = "Static_Error_Missing_Router"
	STATIC_ERROR_MISSING_PUBLICDIR      = "Static_Error_Missing_PublicDir"
	IMAGE_ERROR_MISSING_FILESVC         = "Image_Error_Missing_Filesvc"
	IMAGE_ERROR_DISP_MODES_NOT_PROVIDED = "Image_Error_Disp_Modes_Not_Provided"
)

func init() {
	errors.RegisterCode(STATIC_ERROR_MISSING_ROUTER, errors.ERROR, fmt.Errorf("Router not found in static file service."), LOGGING_CONTEXT)
	errors.RegisterCode(STATIC_ERROR_MISSING_PUBLICDIR, errors.ERROR, fmt.Errorf("Public directory not provided to static file service."), LOGGING_CONTEXT)
	errors.RegisterCode(IMAGE_ERROR_MISSING_FILESVC, errors.ERROR, fmt.Errorf("File Service not provided to image service."), IMAGE_LOGGING_CONTEXT)
	errors.RegisterCode(IMAGE_ERROR_DISP_MODES_NOT_PROVIDED, errors.ERROR, fmt.Errorf("Display Modes not provided to image service."), IMAGE_LOGGING_CONTEXT)
}
