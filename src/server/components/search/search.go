package search

import "laatoo/sdk/server/core"

const (
	CONF_INDEX        = "searchindex"
	CONF_NUMOFRESULTS = "searchresults"
)

type SearchComponent interface {
	//Index a searchable document
	Index(ctx core.RequestContext, bucket string, s Searchable) error
	//Update index
	UpdateIndex(ctx core.RequestContext, bucket string, id string, stype string, u map[string]interface{}) error
	//Index a searchable document
	Search(ctx core.RequestContext, bucket string, query string) ([]Searchable, error)
	//Delete a searchable document
	Delete(ctx core.RequestContext, bucket string, s Searchable) error
}
