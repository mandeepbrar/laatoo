// +build !appengine

package context

import (
	"github.com/labstack/echo"
	"net/http"
)

func HttpClient(ctx *echo.Context) *http.Client {
	return &http.Client{}
}
