package laatooentities

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	ENTITY_ERROR_MISSING_ROUTER     = "Entity_Error_Missing_Router"
	ENTITY_ERROR_MISSING_ENTITIES   = "Entity_Error_Missing_Entities"
	ENTITY_ERROR_WRONG_ENTITYCONFIG = "Entity_Error_Wrong_Entityconfig"
)

func init() {
	errors.RegisterCode(ENTITY_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in entity service."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_ENTITIES, errors.PANIC, fmt.Errorf("Entities not found in entity service."))
	errors.RegisterCode(ENTITY_ERROR_WRONG_ENTITYCONFIG, errors.PANIC, fmt.Errorf("Wrong config provided for entity."))
}
