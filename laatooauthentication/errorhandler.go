package laatooauthentication

import (
	"fmt"
	"github.com/labstack/echo"
	"laatoosdk/errors"
	"laatoosdk/utils"
	"net/http"
)

const (
	AUTH_ERROR_MISSING_ROUTER            = "Auth_Error_Missing_Router"
	AUTH_ERROR_MISSING_USER_DATA_SERVICE = "Auth_Error_Missing_User_Data_Service"
	AUTH_ERROR_INITIALIZING_TYPE         = "Auth_Error_Initializing_Type"

	AUTH_ERROR_WRONG_SIGNING_METHOD       = "Auth_Error_Wrong_Signing_Method"
	AUTH_ERROR_USEROBJECT_NOT_CREATED     = "Auth_Error_User_Object_Not_Created"
	AUTH_ERROR_HEADER_NOT_FOUND           = "Auth_Error_Header_Not_Found"
	AUTH_ERROR_USER_VALIDATION_FAILED     = "Auth_Error_User_Validation_Failed"
	AUTH_ERROR_AUTH_COMPLETION_FAILED     = "Auth_Error_Auth_Completion_Failed"
	AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH = "Auth_Error_Internal_Server_Error_Auth"
	AUTH_ERROR_JWT_CREATION               = "Auth_Error_JWT_Creation"
	AUTH_ERROR_WRONG_PASSWORD             = "Auth_Error_Wrong_Password"
	AUTH_ERROR_USER_NOT_FOUND             = "Auth_Error_User_Not_Found"
	AUTH_ERROR_INCORRECT_REQ_FORMAT       = "Auth_Error_Incorrect_Req_Format"
)

func init() {
	errors.RegisterCode(AUTH_ERROR_MISSING_ROUTER, errors.PANIC, fmt.Errorf("Router not found in authentication service."))
	errors.RegisterCode(AUTH_ERROR_MISSING_USER_DATA_SERVICE, errors.PANIC, fmt.Errorf("User data service not provided to authentication service."))
	errors.RegisterCode(AUTH_ERROR_INITIALIZING_TYPE, errors.PANIC, fmt.Errorf("Auth Type could not be initialized."))

	errors.RegisterCode(AUTH_ERROR_WRONG_SIGNING_METHOD, errors.WARNING, echo.NewHTTPError(http.StatusUnauthorized, "Wrong signing method used in the token."))
	errors.RegisterErrorHandler(AUTH_ERROR_WRONG_SIGNING_METHOD, AuthError)

	errors.RegisterCode(AUTH_ERROR_USEROBJECT_NOT_CREATED, errors.ERROR, echo.NewHTTPError(http.StatusInternalServerError, "User Object Could not be created."))
	errors.RegisterErrorHandler(AUTH_ERROR_USEROBJECT_NOT_CREATED, AuthError)

	errors.RegisterCode(AUTH_ERROR_HEADER_NOT_FOUND, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "Auth header not found."))
	errors.RegisterErrorHandler(AUTH_ERROR_HEADER_NOT_FOUND, AuthError)

	errors.RegisterCode(AUTH_ERROR_USER_VALIDATION_FAILED, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "User Validation Failed."))
	errors.RegisterErrorHandler(AUTH_ERROR_USER_VALIDATION_FAILED, AuthError)

	errors.RegisterCode(AUTH_ERROR_AUTH_COMPLETION_FAILED, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "Could not complete authentication of user."))
	errors.RegisterErrorHandler(AUTH_ERROR_AUTH_COMPLETION_FAILED, AuthError)

	errors.RegisterCode(AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, errors.INFO, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error completing authentication."))
	errors.RegisterErrorHandler(AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, AuthError)

	errors.RegisterCode(AUTH_ERROR_JWT_CREATION, errors.ERROR, echo.NewHTTPError(http.StatusInternalServerError, "Could not create JWT Token."))
	errors.RegisterErrorHandler(AUTH_ERROR_JWT_CREATION, AuthError)

	errors.RegisterCode(AUTH_ERROR_WRONG_PASSWORD, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "User name or password entered by you is wrong"))
	errors.RegisterErrorHandler(AUTH_ERROR_WRONG_PASSWORD, AuthError)

	errors.RegisterCode(AUTH_ERROR_USER_NOT_FOUND, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "User name or password entered by you is wrong"))
	errors.RegisterErrorHandler(AUTH_ERROR_USER_NOT_FOUND, AuthError)

	errors.RegisterCode(AUTH_ERROR_INCORRECT_REQ_FORMAT, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "Request for login was not in a correct format"))
	errors.RegisterErrorHandler(AUTH_ERROR_INCORRECT_REQ_FORMAT, AuthError)
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
