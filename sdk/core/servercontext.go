package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/ctx"
)

/*application and engine types*/
const (
	CONF_SERVERTYPE_STANDALONE = "STANDALONE"
	CONF_SERVERTYPE_GOOGLEAPP  = "GOOGLE_APP"
	CONF_ENGINE_HTTP           = "http"
	CONF_ENGINE_TCP            = "tcp"
)

type ServerElementType int

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
	ServerElementRulesManager
	ServerElementService
	ServerElementServiceResponseHandler
	ServerElementServer
	ServerElementSecurityHandler
	ServerElementMessagingManager
	ServerElementModuleManager
	ServerElementTaskManager
	ServerElementCacheManager
	ServerElementModule
	ServerElementLogger
	ServerElementOpen1
	ServerElementOpen2
	ServerElementOpen3
)

type ContextMap map[ServerElementType]ServerElement

type ServerElement interface {
	Reference() ServerElement
	GetProperty(string) interface{}
	GetName() string
	GetType() ServerElementType
}

type ServerContext interface {
	ctx.Context
	GetServerElement(ServerElementType) ServerElement
	GetService(alias string) (Service, error)
	//NewContext(name string) ServerContext
	SubContext(name string) ServerContext
	GetServerProperties() map[string]interface{}
	CreateNewRequest(name string, params interface{}) RequestContext
	CreateCollection(objectName string, length int) (interface{}, error)
	CreateObject(objectName string) (interface{}, error)
	GetObjectCollectionCreator(objectName string) (ObjectCollectionCreator, error)
	GetObjectCreator(objectName string) (ObjectCreator, error)
	GetObjectMetadata(objectName string) (Info, error)
	CreateSystemRequest(name string) RequestContext
	SubscribeTopic(topics []string, lstnr MessageListener) error
	CreateConfig() config.Config
	ReadConfigData(data []byte, funcs map[string]interface{}) (config.Config, error)
	ReadConfig(file string, funcs map[string]interface{}) (config.Config, error)
}
