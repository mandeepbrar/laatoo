package laatoocore

import (
	"fmt"
	"github.com/labstack/echo"
	"laatoosdk/errors"
	"laatoosdk/utils"
	"net/http"
)

const (
	CORE_ERROR_MISSING_SERVICE_NAME   = "Core_Missing_Service_Name"
	CORE_ERROR_SERVICE_CREATION       = "Core_Service_Creation"
	CORE_ERROR_SERVICE_INITIALIZATION = "Core_Service_Initialization"
	CORE_ERROR_SERVICE_NOT_FOUND      = "Core_Service_Not_Found"
	CORE_ERROR_SERVICE_NOT_STARTED    = "Core_Service_Not_Started"
	CORE_ERROR_PROVIDER_NOT_FOUND     = "Core_Provider_Not_Found"
	CORE_ERROR_OBJECT_NOT_CREATED     = "Core_Object_Not_Created"
	CORE_ENVIRONMENT_NOT_CREATED      = "Core_Environment_Not_Created"
	CORE_ENVIRONMENT_NOT_INITIALIZED  = "Core_Environment_Not_Initialized"
	CORE_SERVERADD_NOT_FOUND          = "Core_ServerAdd_Not_Found"

	AUTH_ERROR_WRONG_SIGNING_METHOD   = "Auth_Error_Wrong_Signing_Method"
	AUTH_ERROR_USEROBJECT_NOT_CREATED = "Auth_Error_User_Object_Not_Created"
	AUTH_ERROR_HEADER_NOT_FOUND       = "Auth_Error_Header_Not_Found"
	AUTH_ERROR_INVALID_TOKEN          = "Auth_Error_Invalid_Token"
)

func init() {
	errors.RegisterCode(CORE_ERROR_MISSING_SERVICE_NAME, errors.FATAL, fmt.Errorf("Service name is missing or wrong."))
	errors.RegisterCode(CORE_ERROR_SERVICE_CREATION, errors.FATAL, fmt.Errorf("Service could not be created."))
	errors.RegisterCode(CORE_ERROR_SERVICE_INITIALIZATION, errors.FATAL, fmt.Errorf("Service could not be initialized."))
	errors.RegisterCode(CORE_ERROR_SERVICE_NOT_FOUND, errors.FATAL, fmt.Errorf("Service not found."))
	errors.RegisterCode(CORE_ERROR_SERVICE_NOT_STARTED, errors.FATAL, fmt.Errorf("Service not started."))
	errors.RegisterCode(CORE_ERROR_PROVIDER_NOT_FOUND, errors.FATAL, fmt.Errorf("Provider for object not found."))
	errors.RegisterCode(CORE_ERROR_OBJECT_NOT_CREATED, errors.FATAL, fmt.Errorf("Object could not be created from provider."))
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_CREATED, errors.FATAL, fmt.Errorf("Environment could not be created."))
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_INITIALIZED, errors.FATAL, fmt.Errorf("Environment could not be initialized."))
	errors.RegisterCode(CORE_SERVERADD_NOT_FOUND, errors.FATAL, fmt.Errorf("Server address not provided."))

	errors.RegisterCode(AUTH_ERROR_WRONG_SIGNING_METHOD, errors.WARNING, echo.NewHTTPError(http.StatusUnauthorized, "Wrong signing method used in the token."))
	errors.RegisterErrorHandler(AUTH_ERROR_WRONG_SIGNING_METHOD, AuthError)

	errors.RegisterCode(AUTH_ERROR_USEROBJECT_NOT_CREATED, errors.ERROR, echo.NewHTTPError(http.StatusInternalServerError, "User Object Could not be created."))
	errors.RegisterErrorHandler(AUTH_ERROR_USEROBJECT_NOT_CREATED, AuthError)

	errors.RegisterCode(AUTH_ERROR_HEADER_NOT_FOUND, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "Auth header not found."))
	errors.RegisterErrorHandler(AUTH_ERROR_HEADER_NOT_FOUND, AuthError)

	errors.RegisterCode(AUTH_ERROR_INVALID_TOKEN, errors.WARNING, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token."))
	errors.RegisterErrorHandler(AUTH_ERROR_INVALID_TOKEN, AuthError)

}

func AuthError(err *errors.Error, ctxMap map[string]interface{}, info ...string) bool {
	ctxInt, ok := ctxMap["Context"]
	if !ok {
		return false
	}
	ctx := ctxInt.(*echo.Context)
	ctx.Set("User", nil)
	//ctx.Response().Header().Set(svc.AuthHeader, "")
	utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_AUTH_FAILED, ctx})
	return false
}
