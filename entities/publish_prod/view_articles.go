package publish_prod

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"net/http"
	//"laatoosdk/errors"
)

const (
	VIEW_NAME   = "view_articles"
	VIEW_ENTITY = "article"
)

type ArticlesView struct {
	Options map[string]interface{}
}

func NewArticlesView(conf map[string]interface{}) (interface{}, error) {
	return &ArticlesView{conf}, nil
}

func init() {
	laatoocore.RegisterObjectProvider(VIEW_NAME, NewArticlesView)
}

func (view *ArticlesView) Execute(dataStore data.DataService, ctx *echo.Context) error {
	entities, err := dataStore.Get(VIEW_ENTITY, nil)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, entities)
}
