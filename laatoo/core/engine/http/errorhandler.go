package http

import (
	"fmt"
	"laatoo/sdk/errors"
)

const (
	CORE_ERROR_INCORRECT_DELIVERY_CONF = "Core_Error_Incorrect_Delivery_Conf"
)

func init() {
	errors.RegisterCode(CORE_ERROR_INCORRECT_DELIVERY_CONF, errors.FATAL, fmt.Errorf("Delivery methods could not be configured."))
}
