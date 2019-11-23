package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/search"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"strconv"

	gaesearch "google.golang.org/appengine/search"
)

const (
	CONF_GOOGLESEARCH_SVC = "googlesearch"
)

type GoogleSearchService struct {
	core.Service
	indexName    string
	numOfResults int
}

func (gs *GoogleSearchService) Initialize(ctx core.ServerContext, conf config.Config) error {

	/*	gs.SetDescription(ctx, "Google search service")
		gs.SetRequestType(ctx, config.CONF_OBJECT_STRING, false, false)
		gs.AddStringConfigurations(ctx, []string{search.CONF_INDEX, search.CONF_NUMOFRESULTS}, []string{"", "15"})
	*/

	index, ok := gs.GetConfiguration(ctx, search.CONF_INDEX)
	if !ok {
		return errors.MissingConf(ctx, search.CONF_INDEX)
	}
	gs.indexName = index.(string)

	num, _ := gs.GetConfiguration(ctx, search.CONF_INDEX)
	var err error
	gs.numOfResults, err = strconv.Atoi(num.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	gs.AddStringParam(ctx, "query")

	return nil
}

func (gs *GoogleSearchService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	query, _ := req.GetStringParam(ctx, "query")
	res, err := gs.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return core.SuccessResponse(res), nil
}

func (gs *GoogleSearchService) Index(ctx core.RequestContext, s search.Searchable) error {
	ctx = ctx.SubContext("GoogleSearch_Index")
	appengineCtx := ctx.GetAppengineContext()
	index, err := gaesearch.Open(gs.indexName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	id := fmt.Sprintf("%s_%s", s.GetType(), s.GetId())
	_, err = index.Put(appengineCtx, id, s)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (gs *GoogleSearchService) UpdateIndex(ctx core.RequestContext, id string, stype string, u map[string]interface{}) error {
	ctx = ctx.SubContext("GoogleSearch_UpdateIndex")
	appengineCtx := ctx.GetAppengineContext()
	index, err := gaesearch.Open(gs.indexName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	id = fmt.Sprintf("%s_%s", stype, id)
	bs := new(search.BaseSearchDocument)
	err = index.Get(appengineCtx, id, bs)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = utils.SetObjectFields(ctx, bs, u, nil, nil)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Info(ctx, "Creating index ***********", "bs", bs, "u", u)
	_, err = index.Put(appengineCtx, id, bs)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

//Index a searchable document
func (gs *GoogleSearchService) Search(ctx core.RequestContext, query string) ([]search.Searchable, error) {
	ctx = ctx.SubContext("GoogleSearch_Search")
	appengineCtx := ctx.GetAppengineContext()
	index, err := gaesearch.Open(gs.indexName)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	res := make([]search.Searchable, gs.numOfResults)
	i := 0
	for t := index.Search(appengineCtx, query, nil); i < gs.numOfResults; i++ {
		var bs search.BaseSearchDocument
		_, err := t.Next(&bs)
		if err == gaesearch.Done {
			break
		}
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		res[i] = &bs
	}
	return res[0:i], nil
}

//Delete a searchable document
func (gs *GoogleSearchService) Delete(ctx core.RequestContext, s search.Searchable) error {
	ctx = ctx.SubContext("GoogleSearch_Delete")
	appengineCtx := ctx.GetAppengineContext()
	index, err := gaesearch.Open(gs.indexName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	id := fmt.Sprintf("%s_%s", s.GetType(), s.GetId())
	err = index.Delete(appengineCtx, id)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}