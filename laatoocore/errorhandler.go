package laatoocore

import (
	"fmt"
	"github.com/labstack/echo"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/utils"
	"net/http"
)

const (
	CORE_ERROR_MISSING_SERVICE_NAME   = "Core_Missing_Service_Name"
	CORE_ERROR_SERVICE_CREATION       = "Core_Service_Creation"
	CORE_ERROR_SERVICE_INITIALIZATION = "Core_Service_Initialization"
	CORE_ERROR_PUBSUB_INITIALIZATION  = "Core_Error_Pubsub_Initialization"
	CORE_ERROR_SERVICE_NOT_FOUND      = "Core_Service_Not_Found"
	CORE_ERROR_SERVICE_NOT_STARTED    = "Core_Service_Not_Started"
	CORE_ERROR_PROVIDER_NOT_FOUND     = "Core_Provider_Not_Found"
	CORE_ERROR_OBJECT_NOT_CREATED     = "Core_Object_Not_Created"
	CORE_ERROR_NO_CACHE_SVC           = "Core_No_Cache_Svc"
	CORE_ENVIRONMENT_NOT_CREATED      = "Core_Environment_Not_Created"
	CORE_ENVIRONMENT_NOT_INITIALIZED  = "Core_Environment_Not_Initialized"
	CORE_ERROR_NOCOMMSVC              = "Core_Error_Nocommsvc"
	CORE_ERROR_INVALID_HTTPMETHOD     = "Core_Error_Invalid_Httpmethod"
	CORE_SERVERADD_NOT_FOUND          = "Core_ServerAdd_Not_Found"
	CORE_ROLESAPI_NOT_FOUND           = "Core_Rolesapi_Not_Found"
	CORE_PERMAPI_NOT_FOUND            = "Core_Permapi_Not_Found"
	CORE_ROLES_INIT_ERROR             = "Core_Roles_Init_Error"

	AUTH_ERROR_WRONG_SIGNING_METHOD   = "Auth_Error_Wrong_Signing_Method"
	AUTH_ERROR_USEROBJECT_NOT_CREATED = "Auth_Error_User_Object_Not_Created"
	AUTH_ERROR_HEADER_NOT_FOUND       = "Auth_Error_Header_Not_Found"
	AUTH_ERROR_INVALID_TOKEN          = "Auth_Error_Invalid_Token"
	AUTH_ERROR_SECURITY               = "Auth_Error_Security"
	AUTH_MISSING_API                  = "Auth_Missing_Api"
	AUTH_APISEC_NOTALLOWED            = "Auth_Apisec_notallowed"
)

func init() {
	errors.RegisterCode(CORE_ERROR_MISSING_SERVICE_NAME, errors.FATAL, fmt.Errorf("Service name is missing or wrong."), "core")
	errors.RegisterCode(CORE_ERROR_SERVICE_CREATION, errors.FATAL, fmt.Errorf("Service could not be created."), "core")
	errors.RegisterCode(CORE_ERROR_SERVICE_INITIALIZATION, errors.FATAL, fmt.Errorf("Service could not be initialized."), "core")
	errors.RegisterCode(CORE_ERROR_SERVICE_NOT_FOUND, errors.FATAL, fmt.Errorf("Service not found."), "core")
	errors.RegisterCode(CORE_ERROR_PUBSUB_INITIALIZATION, errors.FATAL, fmt.Errorf("Pubsub could not be initialized."), "core")
	errors.RegisterCode(CORE_ERROR_SERVICE_NOT_STARTED, errors.FATAL, fmt.Errorf("Service not started."), "core")
	errors.RegisterCode(CORE_ERROR_PROVIDER_NOT_FOUND, errors.FATAL, fmt.Errorf("Provider for object not found."), "core")
	errors.RegisterCode(CORE_ERROR_OBJECT_NOT_CREATED, errors.FATAL, fmt.Errorf("Object could not be created from provider."), "core")
	errors.RegisterCode(CORE_ERROR_NO_CACHE_SVC, errors.FATAL, fmt.Errorf("No cache service has been configured."), "core")
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_CREATED, errors.FATAL, fmt.Errorf("Environment could not be created."), "core")
	errors.RegisterCode(CORE_ENVIRONMENT_NOT_INITIALIZED, errors.FATAL, fmt.Errorf("Environment could not be initialized."), "core")
	errors.RegisterCode(CORE_SERVERADD_NOT_FOUND, errors.FATAL, fmt.Errorf("Server address not provided."), "core")
	errors.RegisterCode(CORE_ROLESAPI_NOT_FOUND, errors.FATAL, fmt.Errorf("Roles api not provided."), "core")
	errors.RegisterCode(CORE_PERMAPI_NOT_FOUND, errors.FATAL, fmt.Errorf("Permission api for registering permissions to remote auth server not provided."), "core")
	errors.RegisterCode(CORE_ROLES_INIT_ERROR, errors.FATAL, fmt.Errorf("Roles could not be initialized."), "core")
	errors.RegisterCode(AUTH_MISSING_API, errors.FATAL, fmt.Errorf("Missing Info for authentication."), "core")
	errors.RegisterCode(AUTH_APISEC_NOTALLOWED, errors.FATAL, fmt.Errorf("System could not authenticate for Apis."), "core")
	errors.RegisterCode(CORE_ERROR_INVALID_HTTPMETHOD, errors.FATAL, fmt.Errorf("Invalid http method."), "core")

	errors.RegisterCode(AUTH_ERROR_WRONG_SIGNING_METHOD, errors.WARNING, echo.NewHTTPError(http.StatusUnauthorized, "Wrong signing method used in the token."), "core")
	errors.RegisterErrorHandler(AUTH_ERROR_WRONG_SIGNING_METHOD, AuthError)

	errors.RegisterCode(AUTH_ERROR_USEROBJECT_NOT_CREATED, errors.ERROR, echo.NewHTTPError(http.StatusInternalServerError, "User Object Could not be created."), "core")
	errors.RegisterErrorHandler(AUTH_ERROR_USEROBJECT_NOT_CREATED, AuthError)

	errors.RegisterCode(AUTH_ERROR_HEADER_NOT_FOUND, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "Auth header not found."), "core")
	errors.RegisterErrorHandler(AUTH_ERROR_HEADER_NOT_FOUND, AuthError)

	errors.RegisterCode(AUTH_ERROR_INVALID_TOKEN, errors.WARNING, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token."), "core")
	errors.RegisterErrorHandler(AUTH_ERROR_INVALID_TOKEN, AuthError)

	errors.RegisterCode(CORE_ERROR_NOCOMMSVC, errors.ERROR, fmt.Errorf("Communication Service not enabled."), "core")

	errors.RegisterCode(AUTH_ERROR_SECURITY, errors.WARNING, echo.NewHTTPError(http.StatusUnauthorized, "Not allowed."), "core")
}

func AuthError(ctx core.Context, err *errors.Error, info ...interface{}) bool {
	ctx.Set("User", nil)
	//ctx.Response().Header().Set(svc.AuthHeader, "")
	utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_AUTH_FAILED, ctx})
	return false
}
