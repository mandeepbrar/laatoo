package net

import (
	"net/http"
)

type Webframework interface {
	Initialize() error
	GetRootHandler() http.Handler
	GetParentRouter() Router
	StartServer(address string) error
	StartSSLServer(address string, certpath string, keypath string) error
}
