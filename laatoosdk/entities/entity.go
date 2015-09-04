package entities

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	CONF_ENTITY_NAME     = "name"
	CONF_ENTITY_DATA_SVC = "data_svc"
	CONF_ENTITY_TYPE     = "type"
	CONF_ENTITY_ID       = "id"
	CONF_ENTITY_METHODS  = "methods"
	CONF_ENTITY_FIELDS   = "fields"
	CONF_ENTITY_METHOD   = "method"
	CONF_ENTITY_PATH     = "path"
	CONF_ENTITY_PERM     = "permission"
)

const (
	ENTITY_ERROR_MISSING_ROUTER        = "Entity_Error_Missing_Router"
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
	errors.RegisterCode(ENTITY_ERROR_MISSING_NAME, errors.PANIC, fmt.Errorf("Name of entity not provided."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_DATASVC, errors.PANIC, fmt.Errorf("Data service not found for entity."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_METHODS, errors.PANIC, fmt.Errorf("Methods not found for entity."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_PATH, errors.PANIC, fmt.Errorf("Path not found for method."))
	errors.RegisterCode(ENTITY_ERROR_MISSING_METHOD, errors.PANIC, fmt.Errorf("Method not provided."))
	errors.RegisterCode(ENTITY_ERROR_INCORRECT_METHOD_CONF, errors.PANIC, fmt.Errorf("Incorrect conf provided for method."))
	errors.RegisterCode(ENTITY_ERROR_CONF_INCORRECT, errors.PANIC, fmt.Errorf("Incorrect conf provided for entity."))
}

//Object stored by data service
type Entity interface {
	PreSave() error
	PostSave() error
	PreLoad() error
	PostLoad() error
}
