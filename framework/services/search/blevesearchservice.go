package search

import (
	"fmt"
	searchsdk "laatoo/sdk/components/search"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"reflect"
	"strconv"

	"github.com/blevesearch/bleve"
	bsearch "github.com/blevesearch/bleve/search"
)

const (
	CONF_BLEVESEARCH_INDEX        = "index"
	CONF_BLEVESEARCH_NUMOFRESULTS = "results"
)

type BleveSearchService struct {
	indexName    string
	numOfResults int
	index        bleve.Index
}

func (bs *BleveSearchService) Initialize(ctx core.ServerContext, conf config.Config) error {
	index, ok := conf.GetString(CONF_INDEX)
	if !ok {
		return errors.MissingConf(ctx, CONF_INDEX)
	}
	var err error
	bs.indexName = index
	num, ok := conf.GetString(CONF_NUMOFRESULTS)
	if !ok {
		bs.numOfResults = 15
	} else {
		bs.numOfResults, err = strconv.Atoi(num)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
func (bs *BleveSearchService) Invoke(ctx core.RequestContext) error {
	query := ctx.GetRequest().(string)
	res, err := bs.Search(ctx, query)
	if err != nil {
		return err
	}
	ctx.SetResponse(core.SuccessResponse(res))
	return nil
}
func (bs *BleveSearchService) Start(ctx core.ServerContext) error {
	ind, err := bleve.Open(bs.indexName)
	if err != nil {
		ind, err = bleve.New(bs.indexName, bleve.NewIndexMapping())
		log.Logger.Info(ctx, "Creating index ***********", "index", bs.indexName, "Err", err)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	bs.index = ind
	return nil
}

func (bs *BleveSearchService) Index(ctx core.RequestContext, s searchsdk.Searchable) error {
	ctx = ctx.SubContext("BleveSearch_Index")
	log.Logger.Trace(ctx, "Writing doc ")

	id := fmt.Sprintf("%s_%s", s.GetType(), s.GetId())
	err := bs.index.Index(id, s)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Written to index ", "index", bs.index.Name(), "id", id)
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
		log.Logger.Trace(ctx, "Updated to index ", "index", bs.index.Name())
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
	log.Logger.Trace(ctx, "search results ")
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
