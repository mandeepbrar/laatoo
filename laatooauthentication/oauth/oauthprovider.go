package ginauth_oauth

import (
	//"github.com/gin-gonic/gin"
	//"github.com/ginauth"
	"golang.org/x/oauth2"
)

// Provider needs to be implemented for each 3rd party authentication provider
// e.g. Facebook, Twitter, etc...
type OAuthProvider interface {
	Name() string
	GetConfig() *oauth2.Config
	FetchUser(interface{}) (OAuthUser, error)
}
