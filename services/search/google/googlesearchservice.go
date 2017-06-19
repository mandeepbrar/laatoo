package main

import (
	"fmt"
	"laatoo/sdk/components/search"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"strconv"

	googlesearch "google.golang.org/appengine/search"
)

const (
	CONF_GOOGLESEARCH_SVC = "googlesearch"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_GOOGLESEARCH_SVC, Object: GoogleSearchService{}}}
}

type GoogleSearchService struct {
	indexName    string
	numOfResults int
}

func (gs *GoogleSearchService) Initialize(ctx core.ServerContext, conf config.Config) error {
	index, ok := conf.GetString(search.CONF_INDEX)
	if !ok {
		return errors.MissingConf(ctx, search.CONF_INDEX)
	}
	gs.indexName = index
	num, ok := conf.GetString(search.CONF_NUMOFRESULTS)
	var err error
	if !ok {
		gs.numOfResults = 15
	} else {
		gs.numOfResults, err = strconv.Atoi(num)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
func (gs *GoogleSearchService) Invoke(ctx core.RequestContext) error {
	query := ctx.GetRequest().(string)
	res, err := gs.Search(ctx, query)
	if err != nil {
		return err
	}
	ctx.SetResponse(core.SuccessResponse(res))
	return nil
}
func (gs *GoogleSearchService) Start(ctx core.ServerContext) error {
	return nil
}

func (gs *GoogleSearchService) Index(ctx core.RequestContext, s search.Searchable) error {
	ctx = ctx.SubContext("GoogleSearch_Index")
	appengineCtx := ctx.GetAppengineContext()
	index, err := googlesearch.Open(gs.indexName)
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
	index, err := googlesearch.Open(gs.indexName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	id = fmt.Sprintf("%s_%s", stype, id)
	bs := new(search.BaseSearchDocument)
	err = index.Get(appengineCtx, id, bs)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	utils.SetObjectFields(bs, u)
	log.Logger.Info(ctx, "Creating index ***********", "bs", bs, "u", u)
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
	index, err := googlesearch.Open(gs.indexName)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	res := make([]search.Searchable, gs.numOfResults)
	i := 0
	for t := index.Search(appengineCtx, query, nil); i < gs.numOfResults; i++ {
		var bs search.BaseSearchDocument
		_, err := t.Next(&bs)
		if err == googlesearch.Done {
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
	index, err := googlesearch.Open(gs.indexName)
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
