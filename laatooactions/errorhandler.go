package laatooactions

import (
	"fmt"
	"laatoosdk/errors"
)

const (
	ACTION_ERROR_MISSING_ROUTER        = "Action_Error_Missing_Router"
	ACTION_ERROR_MISSING_ACTIONS       = "Action_Error_Missing_Actions"
	ACTION_ERROR_INCORRECT_ACTION_CONF = "Action_Error_Incorrect_Action_Conf"
	ACTION_ERROR_MISSING_ACTION_NAME   = "Action_Error_Missing_Action_Name"
	ACTION_ERROR_MISSING_ACTION_PATH   = "Action_Error_Missing_Action_Path"
	ACTION_ERROR_MISSING_ACTION_TYPE   = "Action_Error_Missing_Action_Type"
)

func init() {
	errors.RegisterCode(ACTION_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in action service."))
	errors.RegisterCode(ACTION_ERROR_MISSING_ACTIONS, errors.PANIC, fmt.Errorf("Actions not provided for the action service."))
	errors.RegisterCode(ACTION_ERROR_INCORRECT_ACTION_CONF, errors.PANIC, fmt.Errorf("Incorrect config for actions."))
	errors.RegisterCode(ACTION_ERROR_MISSING_ACTION_NAME, errors.PANIC, fmt.Errorf("Name not provided to the action service."))
	errors.RegisterCode(ACTION_ERROR_MISSING_ACTION_PATH, errors.PANIC, fmt.Errorf("Path not provided for action."))
	errors.RegisterCode(ACTION_ERROR_MISSING_ACTION_TYPE, errors.PANIC, fmt.Errorf("Action type not provided for action."))
}
