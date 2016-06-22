package security

import (
	"fmt"
	"laatoo/sdk/errors"
)

/*
import (
	"fmt"
	"laatoocore"
	"laatoo/sdk/errors"
	"net/http"
)

const (
	AUTH_ERROR_MISSING_ROUTER             = "Auth_Error_Missing_Router"
	AUTH_ERROR_INITIALIZING_TYPE          = "Auth_Error_Initializing_Type"
	AUTH_ERROR_OAUTH_MISSING_CLIENTID     = "Auth_Error_Oauth_Missing_Clientid"
	AUTH_ERROR_OAUTH_MISSING_CLIENTSECRET = "Auth_Error_Oauth_Missing_Clientsecret"
	AUTH_ERROR_OAUTH_MISSING_AUTHURL      = "Auth_Error_Oauth_Missing_Authurl"
	AUTH_ERROR_OAUTH_MISSING_CALLBACKURL  = "Auth_Error_Oauth_Missing_Callbackurl"
	AUTH_ERROR_OAUTH_MISSING_TYPE         = "Auth_Error_Oauth_Missing_Type"

	AUTH_ERROR_KEYAUTH_MISSING_PVTKEY     = "Auth_Error_Keyauth_Missing_PvtKey"
	AUTH_ERROR_USER_VALIDATION_FAILED     = "Auth_Error_User_Validation_Failed"
	AUTH_ERROR_AUTH_COMPLETION_FAILED     = "Auth_Error_Auth_Completion_Failed"
	AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH = "Auth_Error_Internal_Server_Error_Auth"
	AUTH_ERROR_JWT_CREATION               = "Auth_Error_JWT_Creation"
	AUTH_ERROR_WRONG_PASSWORD             = "Auth_Error_Wrong_Password"
	AUTH_ERROR_USER_NOT_FOUND             = "Auth_Error_User_Not_Found"
	AUTH_ERROR_INCORRECT_REQ_FORMAT       = "Auth_Error_Incorrect_Req_Format"
	AUTH_ERROR_OAUTH_MISSING_PROFILEURL   = "Auth_Error_Oauth_Missing_Profileurl"
	AUTH_ERROR_INVALID_PASSWORD           = "Auth_Error_Invalid_Password"
	AUTH_ERROR_NOT_ALLOWED                = "Auth_Error_Not_Allowed"

)

func init() {
	errors.RegisterCode(AUTH_ERROR_MISSING_ROUTER, errors.FATAL, fmt.Errorf("Router not found in authentication service."))
	errors.RegisterCode(AUTH_ERROR_INITIALIZING_TYPE, errors.FATAL, fmt.Errorf("Auth Type could not be initialized."))
	errors.RegisterCode(AUTH_ERROR_OAUTH_MISSING_CLIENTID, errors.FATAL, fmt.Errorf("Client id not provided for oauth site."))
	errors.RegisterCode(AUTH_ERROR_OAUTH_MISSING_CLIENTSECRET, errors.FATAL, fmt.Errorf("Client secret not provided for oauth site."))
	errors.RegisterCode(AUTH_ERROR_OAUTH_MISSING_AUTHURL, errors.FATAL, fmt.Errorf("Auth url not provided for oauth site."))
	errors.RegisterCode(AUTH_ERROR_OAUTH_MISSING_CALLBACKURL, errors.FATAL, fmt.Errorf("Callback url not provided for oauth site."))
	errors.RegisterCode(AUTH_ERROR_OAUTH_MISSING_PROFILEURL, errors.FATAL, fmt.Errorf("Profile url not provided for oauth site."))
	errors.RegisterCode(AUTH_ERROR_OAUTH_MISSING_TYPE, errors.FATAL, fmt.Errorf("Type not provided for oauth site."))
	errors.RegisterCode(AUTH_ERROR_KEYAUTH_MISSING_PVTKEY, errors.FATAL, fmt.Errorf("Private key could not be loaded for keyauth."))


	errors.RegisterCode(AUTH_ERROR_USER_VALIDATION_FAILED, errors.INFO, echo.NewHTTPError(http.StatusUnauthorized, "User Validation Failed."))
	errors.RegisterErrorHandler(AUTH_ERROR_USER_VALIDATION_FAILED, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_INVALID_PASSWORD, errors.INFO, echo.NewHTTPError(http.StatusBadRequest, "Invalid Password."))

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

	errors.RegisterCode(AUTH_ERROR_NOT_ALLOWED, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "Transaction is not allowed"))
	errors.RegisterErrorHandler(AUTH_ERROR_NOT_ALLOWED, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_INCORRECT_REQ_FORMAT, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "Request was not in a correct format"))
	errors.RegisterErrorHandler(AUTH_ERROR_INCORRECT_REQ_FORMAT, laatoocore.AuthError)

	errors.RegisterCode(AUTH_ERROR_DOMAIN_NOT_ALLOWED, errors.ERROR, echo.NewHTTPError(http.StatusUnauthorized, "Domain not allowed by system"))
	errors.RegisterErrorHandler(AUTH_ERROR_DOMAIN_NOT_ALLOWED, laatoocore.AuthError)
}
*/
const (
	AUTH_ERROR_DOMAIN_NOT_ALLOWED         = "Auth_Error_Domain_Not_Allowed"
	AUTH_ERROR_JWT_CREATION               = "Auth_Error_JWT_Creation"
	AUTH_ERROR_MISSING_USER_DATA_SERVICE  = "Auth_Error_Missing_User_Data_Service"
	AUTH_ERROR_INCORRECT_SECURITY_HANDLER = "Auth_Error_Incorrect_Security_Handler"
	AUTH_ERROR_USER_EXISTS                = "Auth_Error_User_Exists"
	AUTH_ERROR_ENC_ERROR                  = "Auth_Error_Enc_Error"
)

func init() {
	errors.RegisterCode(AUTH_ERROR_DOMAIN_NOT_ALLOWED, errors.ERROR, fmt.Errorf("Domain not allowed by system"))
	errors.RegisterCode(AUTH_ERROR_JWT_CREATION, errors.ERROR, fmt.Errorf("Could not create JWT Token."))
	errors.RegisterCode(AUTH_ERROR_MISSING_USER_DATA_SERVICE, errors.FATAL, fmt.Errorf("User data service not provided to authentication service."))
	errors.RegisterCode(AUTH_ERROR_USER_EXISTS, errors.ERROR, fmt.Errorf("User already exists."))
	errors.RegisterCode(AUTH_ERROR_INCORRECT_SECURITY_HANDLER, errors.ERROR, fmt.Errorf("Security handler has not been correctly configured."))
	errors.RegisterCode(AUTH_ERROR_ENC_ERROR, errors.ERROR, fmt.Errorf("Internal Server error in encrypting password."))
}
