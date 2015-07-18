package laatoodesigner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"laatoocore"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	CONF_DESIGNER_SERVICENAME = "designer_service"
	CONF_DESIGNER_PUBLICDIR   = "publicdir"
)

//Environment hosting an application
type DesignerService struct {
	//alias for the service
	alias string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterServiceProvider(CONF_DESIGNER_SERVICENAME, DesignerServiceFactory)
}

//factory method returns the service object to the environment
func DesignerServiceFactory(alias string, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating designer service with alias %s", alias)
	svc := &DesignerService{alias: alias}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, fmt.Errorf("Designer Service: Router not found ")
	}
	publicdir, ok := conf[CONF_DESIGNER_PUBLICDIR]
	if !ok {
		return nil, fmt.Errorf("Designer Service: Public Dir not found ")
	}
	router := routerInt.(*gin.RouterGroup)
	log.Logger.Infof("Designer service starting with page path %s", publicdir)
	router.Static("/", publicdir.(string))
	return svc, nil
}

//Provides the name of the service
func (svc *DesignerService) GetName() string {
	return CONF_DESIGNER_SERVICENAME
}

//Provides the alias of the service
func (svc *DesignerService) GetAlias() string {
	return svc.alias
}

//Initialize the service. Consumer of a service passes the data
func (svc *DesignerService) Initialize(ctx interface{}) error {
	return nil
}

//The service starts serving when this method is called
func (svc *DesignerService) Serve() error {
	return nil
}

//Type of service
func (svc *DesignerService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}
