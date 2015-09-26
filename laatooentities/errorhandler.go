package laatooentities

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	ENTITY_ERROR_MISSING_ROUTER     = "Entity_Error_Missing_Router"
	ENTITY_ERROR_MISSING_ENTITY     = "Entity_Error_Missing_Entity"
	ENTITY_ERROR_WRONG_ENTITYCONFIG = "Entity_Error_Wrong_Entityconfig"

	ENTITY_ERROR_MISSING_NAME          = "Entity_Error_Missing_Name"
	ENTITY_ERROR_MISSING_DATASVC       = "Entity_Error_Missing_Data"
	ENTITY_ERROR_MISSING_METHODS       = "Entity_Error_Missing_Methods"
	ENTITY_ERROR_MISSING_PATH          = "Entity_Error_Missing_Path"
	ENTITY_ERROR_MISSING_METHOD        = "Entity_Error_Missing_Method"
	ENTITY_ERROR_INCORRECT_METHOD_CONF = "Entity_Error_Incorrect_Method_Conf"
	ENTITY_ERROR_CONF_INCORRECT        = "Entity_Error_Conf_Incorrect"
	ENTITY_VIEW_MISSING_ARG            = "Entity_View_Missing_Arg"
)

func init() {
	errors.RegisterCode(ENTITY_ERROR_MISSING_ROUTER, errors.FATAL, fmt.Errorf("Router not found in entity service."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_MISSING_ENTITY, errors.FATAL, fmt.Errorf("Entity not found in entity service."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_WRONG_ENTITYCONFIG, errors.FATAL, fmt.Errorf("Wrong config provided for entity."), LOGGING_CONTEXT)

	errors.RegisterCode(ENTITY_ERROR_MISSING_NAME, errors.FATAL, fmt.Errorf("Name of entity not provided."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_MISSING_DATASVC, errors.FATAL, fmt.Errorf("Data service not found for entity."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_MISSING_METHODS, errors.FATAL, fmt.Errorf("Methods not found for entity."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_MISSING_PATH, errors.FATAL, fmt.Errorf("Path not found for method."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_MISSING_METHOD, errors.FATAL, fmt.Errorf("Method not provided."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_INCORRECT_METHOD_CONF, errors.FATAL, fmt.Errorf("Incorrect conf provided for method."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_ERROR_CONF_INCORRECT, errors.FATAL, fmt.Errorf("Incorrect conf provided for entity."), LOGGING_CONTEXT)
	errors.RegisterCode(ENTITY_VIEW_MISSING_ARG, errors.FATAL, fmt.Errorf("Name of the entity not provided as an argument to the view."), LOGGING_CONTEXT)

}
