package laatoofiles

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	LOGGING_CONTEXT             = "fileservice"
	FILE_ERROR_MISSING_ROUTER   = "File_Error_Missing_Router"
	FILE_ERROR_MISSING_FILEDIR  = "File_Error_Missing_FilesDir"
	FILE_ERROR_MISSING_FILESURL = "File_Error_Missing_FilesUrl"
	FILE_ERROR_MISSING_BUCKET   = "File_Error_Missing_Bucket"
)

func init() {
	errors.RegisterCode(FILE_ERROR_MISSING_ROUTER, errors.ERROR, fmt.Errorf("Router not found in media service."), LOGGING_CONTEXT)
	errors.RegisterCode(FILE_ERROR_MISSING_FILEDIR, errors.ERROR, fmt.Errorf("File directory not provided to file service."), LOGGING_CONTEXT)
	errors.RegisterCode(FILE_ERROR_MISSING_FILESURL, errors.ERROR, fmt.Errorf("URL for files not provided to file service."), LOGGING_CONTEXT)
	errors.RegisterCode(FILE_ERROR_MISSING_BUCKET, errors.ERROR, fmt.Errorf("Bucket not provided to google cloud storage file service."), LOGGING_CONTEXT)
}
