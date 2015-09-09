package laatoodata

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	DATA_ERROR_MISSING_CONNECTION_STRING = "Data_Error_Missing_Connection_String"
	DATA_ERROR_MISSING_DATABASE          = "Data_Error_Missing_Database"
	DATA_ERROR_CONNECTION                = "Data_Error_Connection"
	DATA_ERROR_MISSING_OBJECTS           = "Data_Error_Missing_Objects"
	DATA_ERROR_NOT_IMPLEMENTED           = "Data_Error_Not_Implemented"
	DATA_ERROR_MISSING_COLLECTION        = "Data_Error_Missing_Collection"
	DATA_ERROR_NOT_FOUND                 = "Data_Error_Not_Found"
)

func init() {
	errors.RegisterCode(DATA_ERROR_MISSING_CONNECTION_STRING, errors.ERROR, fmt.Errorf("Connection string not provided for the database."))
	errors.RegisterCode(DATA_ERROR_MISSING_DATABASE, errors.ERROR, fmt.Errorf("Database name not provided"))
	errors.RegisterCode(DATA_ERROR_MISSING_OBJECTS, errors.ERROR, fmt.Errorf("Name of the objects stored in the database not provided."))
	errors.RegisterCode(DATA_ERROR_CONNECTION, errors.ERROR, fmt.Errorf("Could not connect to the database."))
	errors.RegisterCode(DATA_ERROR_NOT_IMPLEMENTED, errors.ERROR, fmt.Errorf("Method not implemented for the service."))
	errors.RegisterCode(DATA_ERROR_MISSING_COLLECTION, errors.ERROR, fmt.Errorf("Collection name not present for the object."))
	errors.RegisterCode(DATA_ERROR_NOT_FOUND, errors.ERROR, fmt.Errorf("Data not found for criteria."))
}
