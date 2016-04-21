package core

import (
	glctx "golang.org/x/net/context"
	"net/http"
)

/*application and engine types*/
const (
	CONF_SERVERTYPE_STANDALONE = "STANDALONE"
	CONF_SERVERTYPE_GOOGLEAPP  = "GOOGLE_APP"
	CONF_ENGINE_HTTP           = "http"
	CONF_ENGINE_TCP            = "tcp"
)

type ServerElementType int

type TopicListener func(ctx Context, topic string, message interface{})

const (
	ServerElementEngine ServerElementType = iota
	ServerElementEnvironment
	ServerElementLoader
	ServerElementServiceFactory
	ServerElementServiceManager
	ServerElementChannel
	ServerElementChannelManager
	ServerElementFactoryManager
	ServerElementApplication
	ServerElementApplet
	ServerElementService
	ServerElementServiceResponseHandler
	ServerElementServer
	ServerElementSecurityHandler
	ServerElementMessagingManager
	ServerElementOpen1
	ServerElementOpen2
	ServerElementOpen3
)

type ContextMap map[ServerElementType]ServerElement

type ServerElement interface {
	Context
}

type ServerContext interface {
	Context
	GetServerType() string
	GetElement() ServerElement
	GetServerElement(ServerElementType) ServerElement
	GetService(alias string) (Service, error)
	GetElementType() ServerElementType
	NewContext(name string) ServerContext
	NewContextWithElements(name string, elements ContextMap, primaryElement ServerElementType) ServerContext
	SubContext(name string) ServerContext
	SubContextWithElement(name string, primaryElement ServerElementType) ServerContext
	CreateNewRequest(name string, engineCtx interface{}) RequestContext
	GetAppengineContext(ctx RequestContext) glctx.Context
	GetCloudContext(ctx RequestContext, scope string) glctx.Context
	HttpClient(ctx RequestContext) *http.Client
	GetOAuthContext(ctx Context) glctx.Context
	CreateCollection(objectName string, args MethodArgs) (interface{}, error)
	CreateObject(objectName string, args MethodArgs) (interface{}, error)
	GetObjectCollectionCreator(objectName string) (ObjectCollectionCreator, error)
	GetObjectCreator(objectName string) (ObjectCreator, error)
	Publish(topic string, message interface{}) error
	Subscribe(topics []string, lstnr TopicListener) error
}
