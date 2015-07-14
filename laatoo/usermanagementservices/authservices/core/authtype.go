package ginauth

import (
	"github.com/gin-gonic/gin"
)

type AuthType interface {
	InitializeType(authStart gin.HandlerFunc, authFailed gin.HandlerFunc, authSuccessful gin.HandlerFunc) error
	ValidateUser(*gin.Context) error
	CompleteAuthentication(*gin.Context) error
}
