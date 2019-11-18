package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	searchsdk "laatoo/sdk/server/components/search"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
	"strconv"

	"github.com/blevesearch/bleve"
	bsearch "github.com/blevesearch/bleve/search"
)

const (
	CONF_BLEVESEARCH_SVC = "blevesearch"
)

type BleveSearchService struct {
	core.Service
	numOfResults int
	index        bleve.Index
}

func (bs *BleveSearchService) Initialize(ctx core.ServerContext, conf config.Config) error {

	num, _ := bs.GetConfiguration(ctx, searchsdk.CONF_INDEX)
	var err error
	bs.numOfResults, err = strconv.Atoi(num.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	bs.AddStringParam(ctx, "query")

	/*bs.SetDescription(ctx, "Bleve search service")
	bs.SetRequestType(ctx, config.CONF_OBJECT_STRING, false, false)
	bs.AddStringConfigurations(ctx, []string{searchsdk.CONF_INDEX, searchsdk.CONF_NUMOFRESULTS}, []string{"", "15"})
	*/
	return nil
}

func (bs *BleveSearchService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	query, _ := req.GetStringParam(ctx, "query")
	res, err := bs.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return core.SuccessResponse(res), nil
}

func (bs *BleveSearchService) Start(ctx core.ServerContext) error {

	indexName, _ := bs.GetStringConfiguration(ctx, searchsdk.CONF_INDEX)

	ind, err := bleve.Open(indexName)
	if err != nil {
		ind, err = bleve.New(indexName, bleve.NewIndexMapping())
		log.Info(ctx, "Creating index ***********", "index", indexName, "Err", err)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	bs.index = ind

	return nil
}

func (bs *BleveSearchService) Index(ctx core.RequestContext, s searchsdk.Searchable) error {
	ctx = ctx.SubContext("BleveSearch_Index")
	log.Trace(ctx, "Writing doc ")

	id := fmt.Sprintf("%s_%s", s.GetType(), s.GetId())
	err := bs.index.Index(id, s)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "Written to index ", "index", bs.index.Name(), "id", id)
	return nil
}

func (bs *BleveSearchService) UpdateIndex(ctx core.RequestContext, id string, stype string, u map[string]interface{}) error {
	ctx = ctx.SubContext("BleveSearch_UpdateIndex")
	id = fmt.Sprintf("%s_%s", stype, id)
	bquery := bleve.NewDocIDQuery([]string{id})
	search := bleve.NewSearchRequest(bquery)
	searchResults, err := bs.index.Search(search)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if len(searchResults.Hits) > 0 {
		bd := bs.createDocument(ctx, searchResults.Hits[0])
		err := bs.index.Index(id, bd)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		log.Trace(ctx, "Updated to index ", "index", bs.index.Name())
	}
	return nil
}

//search a searchable document
func (bs *BleveSearchService) Search(ctx core.RequestContext, query string) ([]searchsdk.Searchable, error) {
	ctx = ctx.SubContext("BleveSearch_Search")

	// search for some text
	bquery := bleve.NewMatchQuery(query)
	search := bleve.NewSearchRequest(bquery)
	searchResults, err := bs.index.Search(search)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "search results ")
	res := make([]searchsdk.Searchable, bs.numOfResults)
	i := 0
	for _, val := range searchResults.Hits {
		res[i] = bs.createDocument(ctx, val)
		i++
	}
	return res[0:i], nil
}

func (bs *BleveSearchService) createDocument(ctx core.RequestContext, val *bsearch.DocumentMatch) searchsdk.Searchable {
	doc, _ := bs.index.Document(val.ID)
	bd := &searchsdk.BaseSearchDocument{}
	bdval := reflect.ValueOf(bd).Elem()
	for _, field := range doc.Fields {
		fname := field.Name()
		valField := bdval.FieldByName(fname)
		if valField.Kind() == reflect.String {
			fval := string(field.Value())
			valField.SetString(fval)
		}
	}
	return bd
}

//Delete a searchable document
func (bs *BleveSearchService) Delete(ctx core.RequestContext, s searchsdk.Searchable) error {
	ctx = ctx.SubContext("BleveSearch_Delete")
	id := fmt.Sprintf("%s_%s", s.GetType(), s.GetId())
	err := bs.index.Delete(id)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
