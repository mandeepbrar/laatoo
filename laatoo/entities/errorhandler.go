package entities

import (
	"fmt"
	"laatoo/sdk/errors"
)

const (
	ENTITY_ERROR_MISSING_ENTITY     = "Entity_Error_Missing_Entity"
	ENTITY_ERROR_WRONG_ENTITYCONFIG = "Entity_Error_Wrong_Entityconfig"
	ENTITY_ERROR_NOT_FOUND          = "Entity_Error_Not_Found"
	ENTITY_ERROR_NOT_ALLOWED        = "Entity_Error_Not_Allowed"
)

func init() {
	errors.RegisterCode(ENTITY_ERROR_MISSING_ENTITY, errors.FATAL, fmt.Errorf("Entity not found in entity service."))
	errors.RegisterCode(ENTITY_ERROR_WRONG_ENTITYCONFIG, errors.FATAL, fmt.Errorf("Wrong config provided for entity."))
	errors.RegisterCode(ENTITY_ERROR_NOT_ALLOWED, errors.DEBUG, fmt.Errorf("Access to the entity is not allowed."))
	errors.RegisterCode(ENTITY_ERROR_NOT_FOUND, errors.DEBUG, fmt.Errorf("Entity being accessed was not found."))

}
