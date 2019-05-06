package common

import (
	"fmt"
	"laatoo/server/engine/http/echo"
	"laatoo/server/engine/http/gin"
	"laatoo/server/engine/http/net"
	"net/http"

	"github.com/rs/cors"
)

type WebFWAdapter struct {
	framework net.Webframework
	fwName    string
	handlers  map[string]*handlerController
}

func NewAdapter(engName, fwName string) (*WebFWAdapter, error) {
	adapter := &WebFWAdapter{fwName: fwName, handlers: make(map[string]*handlerController)}
	switch fwName {
	case "Echo":
		adapter.framework = &echo.EchoWebFramework{}
	default:
		adapter.framework = &gin.GinWebFramework{Name: engName}
		/*	case "Goji":
			eng.framework = &goji.GojiWebFramework{}*/
	}
	return adapter, nil
}

func (adapter *WebFWAdapter) Initialize() error {
	return adapter.framework.Initialize()
}

func (adapter *WebFWAdapter) GetRootHandler() http.Handler {
	return adapter.framework.GetRootHandler()
}

func (adapter *WebFWAdapter) GetParentRouter(path string) net.Router {
	return adapter.framework.GetParentRouter(path)
}

func (adapter *WebFWAdapter) StartServer(address string) error {
	return adapter.framework.StartServer(address)
}

func (adapter *WebFWAdapter) StartSSLServer(address string, certpath string, keypath string) error {
	return adapter.framework.StartSSLServer(address, certpath, keypath)
}

func (adapter *WebFWAdapter) SetupCors(router net.Router, allowedOrigins []string, corsOptionsPath string) error {
	corsMw := cors.New(cors.Options{
		AllowedOrigins:     allowedOrigins,
		AllowedHeaders:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders:     []string{"*"},
		OptionsPassthrough: true,
		AllowCredentials:   true,
	})
	switch adapter.fwName {
	case "Echo":
		if corsOptionsPath == "" {
			corsOptionsPath = "/*"
		}
		adapter.UseMW(router, corsMw.Handler)
	case "Gin":
		if corsOptionsPath == "" {
			corsOptionsPath = "/*f"
		}
		adapter.UseMiddleware(router, corsMw.HandlerFunc)
	case "Goji":
		if corsOptionsPath == "" {
			corsOptionsPath = "/*"
		}
		adapter.UseMW(router, corsMw.Handler)
	}
	router.Options(corsOptionsPath, func(webctx net.WebContext) error {
		webctx.NoContent(200)
		return nil
	})
	return nil
}

func (adapter *WebFWAdapter) UseMW(router net.Router, handler func(http.Handler) http.Handler) {
	router.UseMW(handler)
}
func (adapter *WebFWAdapter) UseMiddleware(router net.Router, handler http.HandlerFunc) {
	router.UseMiddleware(handler)
}

func (adapter *WebFWAdapter) Group(router net.Router, path string) net.Router {
	return router.Group(path)
}
func (adapter *WebFWAdapter) Use(router net.Router, handler net.HandlerFunc) {
	router.Use(handler)
}
func (adapter *WebFWAdapter) Get(router net.Router, path string, handler net.HandlerFunc) error {
	adaptedHandler, err := adapter.createController("Get", path, handler)
	if err != nil {
		return err
	}
	router.Get(path, adaptedHandler)
	return nil
}
func (adapter *WebFWAdapter) Options(router net.Router, path string, handler net.HandlerFunc) error {
	adaptedHandler, err := adapter.createController("Options", path, handler)
	if err != nil {
		return err
	}
	router.Options(path, adaptedHandler)
	return nil
}
func (adapter *WebFWAdapter) Post(router net.Router, path string, handler net.HandlerFunc) error {
	adaptedHandler, err := adapter.createController("Post", path, handler)
	if err != nil {
		return err
	}
	router.Post(path, adaptedHandler)
	return nil
}
func (adapter *WebFWAdapter) Put(router net.Router, path string, handler net.HandlerFunc) error {
	adaptedHandler, err := adapter.createController("Put", path, handler)
	if err != nil {
		return err
	}
	router.Put(path, adaptedHandler)
	return nil
}
func (adapter *WebFWAdapter) Delete(router net.Router, path string, handler net.HandlerFunc) error {
	adaptedHandler, err := adapter.createController("Delete", path, handler)
	if err != nil {
		return err
	}
	router.Delete(path, adaptedHandler)
	return nil
}
func (adapter *WebFWAdapter) RemovePath(router net.Router, path string, method string) error {
	key := fmt.Sprintf("%s_%s", method, path)
	controllerStruct, ok := adapter.handlers[key]
	if ok {
		controllerStruct.handler = nil
	}
	return nil //router.RemovePath(path, method)
}

func (adapter *WebFWAdapter) createController(method, path string, handler net.HandlerFunc) (net.HandlerFunc, error) {
	key := fmt.Sprintf("%s_%s", method, path)
	controllerStruct, ok := adapter.handlers[key]
	if !ok {
		controllerStruct = &handlerController{handler: handler}
		adapter.handlers[key] = controllerStruct
	} else {
		if controllerStruct.handler == nil {
			controllerStruct.handler = handler
		} else {
			return nil, fmt.Errorf("Handler already exists for path %s method %s", path, method)
		}

	}
	return adapter.controllerFunc(controllerStruct), nil
}

func (adapter *WebFWAdapter) controllerFunc(ctrl *handlerController) net.HandlerFunc {
	return func(wcx net.WebContext) error {
		if ctrl.handler != nil {
			return ctrl.handler(wcx)
		}
		return fmt.Errorf("No such path")
	}
}

type handlerController struct {
	handler net.HandlerFunc
}
