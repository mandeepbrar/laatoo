package search

/*
import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/server/components/search"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_SEARCHADAPTER_SERVICES   = "searchadapter"
	CONF_SEARCHADAPTER_SEARCH_SVC = "search_svc"
	CONF_SVC_INDEX                = "INDEX"
	CONF_SVC_DELETEINDEX          = "DELETE"
	CONF_SVC_SEARCH               = "SEARCH"
)

func init() {
	objects.RegisterObject(CONF_SEARCHADAPTER_SERVICES, createSearchAdapterFactory, nil)
}

type SearchAdapterFactory struct {
	SearchIndex       search.SearchComponent
	searchServiceName string
}

func createSearchAdapterFactory() interface{} {
	return &SearchAdapterFactory{}
}

func (ss *SearchAdapterFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	searchsvc, ok := conf.GetString(CONF_SEARCHADAPTER_SEARCH_SVC)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Configuration", CONF_SEARCHADAPTER_SEARCH_SVC)
	}
	ss.searchServiceName = searchsvc
	return nil
}

//Create the services configured for factory.
func (es *SearchAdapterFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newSearchAdapterService(ctx, name, method, es)
}

//The services start serving when this method is called
func (ss *SearchAdapterFactory) Start(ctx core.ServerContext) error {
	searchSvc, err := ctx.GetService(ss.searchServiceName)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, err, "Name", es.searchServiceName)
	}
	es.SearchIndex = searchSvc.(search.SearchComponent)
	return nil
}

type searchAdapterService struct {
	name        string
	method      string
	svcfunc     core.ServiceFunc
	conf        config.Config
	fac         *SearchAdapterFactory
	SearchIndex search.SearchComponent
}

func newSearchAdapterService(ctx core.ServerContext, name string, method string, fac *SearchAdapterFactory) (*searchAdapterService, error) {
	ss := &searchAdapterService{name: name, fac: fac, method: method}
	//exported methods
	switch method {
	/*case CONF_SVC_INDEX:
		ss.svcfunc = ss.INDEX
	case CONF_SVC_DELETEINDEX:
		ss.svcfunc = ss.DELETEINDEX
	case CONF_SVC_SEARCH:
		ss.svcfunc = ss.SEARCH
	default:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Wrong Service method", method)
	}
	//cache, _ := conf.GetBool(CONF_DATA_CACHEABLE)
	//ds.cache = cache
	log.Trace(ctx, "Created Search Adapter service", "Svc Name", name, "Method", method)
	return ds, nil
}

func (ss *searchAdapterService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Search adapter initialize")
	ss.conf = conf
	return nil
}

func (ss *searchAdapterService) Start(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Search adapter start")
	ss.SearchIndex = ss.fac.SearchIndex
	return nil
}

func (ss *searchAdapterService) Invoke(ctx core.RequestContext) error {
	return ss.svcfunc(ctx)
}

func (ss *searchAdapterService) INDEX(ctx core.RequestContext) error {
	ctx = ctx.SubContext("INDEX")
	return nil
}
func (ss *searchAdapterService) DELETEINDEX(ctx core.RequestContext) error {
	ctx = ctx.SubContext("DELETEINDEX")
	return nil
}
func (ss *searchAdapterService) SEARCH(ctx core.RequestContext) error {
	ctx = ctx.SubContext("SEARCH")
	return nil
}
*/
