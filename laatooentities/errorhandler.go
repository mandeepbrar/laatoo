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
)

func init() {
	errors.RegisterCode(ENTITY_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in entity service."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_ENTITY, errors.PANIC, fmt.Errorf("Entity not found in entity service."))
	errors.RegisterCode(ENTITY_ERROR_WRONG_ENTITYCONFIG, errors.PANIC, fmt.Errorf("Wrong config provided for entity."))

	errors.RegisterCode(ENTITY_ERROR_MISSING_NAME, errors.PANIC, fmt.Errorf("Name of entity not provided."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_DATASVC, errors.PANIC, fmt.Errorf("Data service not found for entity."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_METHODS, errors.PANIC, fmt.Errorf("Methods not found for entity."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_PATH, errors.PANIC, fmt.Errorf("Path not found for method."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_METHOD, errors.PANIC, fmt.Errorf("Method not provided."))
	errors.RegisterCode(ENTITY_ERROR_INCORRECT_METHOD_CONF, errors.PANIC, fmt.Errorf("Incorrect conf provided for method."))
	errors.RegisterCode(ENTITY_ERROR_CONF_INCORRECT, errors.PANIC, fmt.Errorf("Incorrect conf provided for entity."))

}
