package errors

import (
	"fmt"
	sdkerrors "laatoo/sdk/errors"
)

const (
	CORE_ERROR_SESSION_NOT_FOUND = "Core_Session_Not_Found"
)

func init() {
	sdkerrors.RegisterCode(CORE_ERROR_SESSION_NOT_FOUND, sdkerrors.INFO, fmt.Errorf("Session nout found."))
}
