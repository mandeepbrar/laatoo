package search

/*
import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type SearchFactory struct {
}

const (
	CONF_SEARCH_NAME      = "search"
	CONF_GOOGLESEARCH_SVC = "googlesearch"
	CONF_BLEVESEARCH_SVC  = "blevesearch"
	CONF_INDEX            = "index"
	CONF_NUMOFRESULTS     = "results"
)

func init() {
	objects.Register(CONF_SEARCH_NAME, SearchFactory{})
}

//Create the services configured for factory.
func (sf *SearchFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case CONF_GOOGLESEARCH_SVC:
		{
			return &GoogleSearchService{}, nil
		}
	case CONF_BLEVESEARCH_SVC:
		{
			return &BleveSearchService{}, nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *SearchFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *SearchFactory) Start(ctx core.ServerContext) error {
	return nil
}
*/
