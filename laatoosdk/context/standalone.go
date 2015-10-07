// +build !appengine

package context

import (
	"net/http"
)

func HttpClient(ctx interface{}) *http.Client {
	return &http.Client{}
}
