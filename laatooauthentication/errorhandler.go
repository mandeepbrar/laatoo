package laatooauthentication

import (
	"fmt"
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"net/http"
)

const (
	AUTH_ERROR_MISSING_ROUTER            = "Auth_Error_Missing_Router"
	AUTH_ERROR_MISSING_USER_DATA_SERVICE = "Auth_Error_Missing_User_Data_Service"
	AUTH_ERROR_INITIALIZING_TYPE         = "Auth_Error_Initializing_Type"

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

	errors.RegisterCode(AUTH_ERROR_USER_VALIDATION_FAILED, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "User Validation Failed."))
	errors.RegisterErrorHandler(AUTH_ERROR_USER_VALIDATION_FAILED, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_AUTH_COMPLETION_FAILED, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "Could not complete authentication of user."))
	errors.RegisterErrorHandler(AUTH_ERROR_AUTH_COMPLETION_FAILED, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, errors.INFO, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error completing authentication."))
	errors.RegisterErrorHandler(AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_JWT_CREATION, errors.ERROR, echo.NewHTTPError(http.StatusInternalServerError, "Could not create JWT Token."))
	errors.RegisterErrorHandler(AUTH_ERROR_JWT_CREATION, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_WRONG_PASSWORD, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "User name or password entered by you is wrong"))
	errors.RegisterErrorHandler(AUTH_ERROR_WRONG_PASSWORD, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_USER_NOT_FOUND, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "User name or password entered by you is wrong"))
	errors.RegisterErrorHandler(AUTH_ERROR_USER_NOT_FOUND, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_INCORRECT_REQ_FORMAT, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "Request for login was not in a correct format"))
	errors.RegisterErrorHandler(AUTH_ERROR_INCORRECT_REQ_FORMAT, laatoocore.AuthError)
}
