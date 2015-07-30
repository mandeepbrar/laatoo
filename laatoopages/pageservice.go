package laatoopages

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	//"net/http"
)

const (
	CONF_PAGE_SERVICENAME = "page_service"
	CONF_PAGE_PAGESDIR    = "pagesdir"
	CONF_PAGE_PAGES       = "pages"
)

type PageService struct {
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_PAGE_SERVICENAME, PageServiceFactory)
}

//factory method returns the service object to the environment
func PageServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating page service")
	svc := &PageService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(PAGE_ERROR_MISSING_ROUTER)
	}
	pagesdir, ok := conf[CONF_PAGE_PAGESDIR]
	if !ok {
		return nil, errors.ThrowError(PAGE_ERROR_MISSING_PAGESDIR)
	}
	router := routerInt.(*echo.Group)
	log.Logger.Infof("Router %s", router)

	//get a map of all the pages
	pagesInt, ok := conf[CONF_PAGE_PAGES]
	if !ok {
		return nil, errors.ThrowError(PAGE_ERROR_PAGES_NOT_PROVIDED)
	}

	pages, ok := pagesInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(PAGE_ERROR_PAGES_NOT_PROVIDED)
	}
	for name, val := range pages {
		//get the config for the page with given alias
		pageConf := val.(map[string]interface{})
		//get the service name to be created for the alias
		log.Logger.Info("Creating page %s", name)
		//create page with provided conf
		createPage(pageConf, router, pagesdir.(string))
	}

	return svc, nil
}

//Provides the name of the service
func (svc *PageService) GetName() string {
	return CONF_PAGE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *PageService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *PageService) Serve() error {
	return nil
}

//Type of service
func (svc *PageService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}
