package laatoofiles

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	FILE_ERROR_MISSING_ROUTER   = "File_Error_Missing_Router"
	FILE_ERROR_MISSING_FILEDIR  = "File_Error_Missing_FilesDir"
	FILE_ERROR_MISSING_FILESURL = "File_Error_Missing_FilesUrl"
)

func init() {
	errors.RegisterCode(FILE_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in media service."))
	errors.RegisterCode(FILE_ERROR_MISSING_FILEDIR, errors.ERROR, fmt.Errorf("File directory not provided to file service."))
	errors.RegisterCode(FILE_ERROR_MISSING_FILESURL, errors.ERROR, fmt.Errorf("URL for files not provided to file service."))
}
