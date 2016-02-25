// +build !appengine

package context

import (
	"crypto/tls"
	"github.com/labstack/echo"
	"net/http"
)

func HttpClient(ctx *echo.Context) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
