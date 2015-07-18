package authentication

import (
	"github.com/gin-gonic/gin"
)

//Interface that has to be implemented by every auth type interface
type AuthType interface {
	//Initializes the authentication type module
	InitializeType(authStart gin.HandlerFunc, authFailed gin.HandlerFunc, authSuccessful gin.HandlerFunc) error
	//Called to validate the user by providing context
	ValidateUser(*gin.Context) error
	//Completes authentication
	CompleteAuthentication(*gin.Context) error
}
