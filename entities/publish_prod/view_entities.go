package publish_prod

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	views "laatooview"
	"net/http"
)

const (
	VIEW_NAME   = "view_entities"
	VIEW_ENTITY = "entity"
)

type EntitiesView struct {
	Options map[string]interface{}
}

func NewEntitiesView(conf map[string]interface{}) (interface{}, error) {
	return &EntitiesView{conf}, nil
}

func init() {
	laatoocore.RegisterObjectProvider(VIEW_NAME, NewEntitiesView)
}

func (view *EntitiesView) Execute(dataStore data.DataService, ctx *echo.Context) error {
	entity := ctx.Query(VIEW_ENTITY)
	if entity == "" {
		return errors.ThrowHttpError(views.VIEW_ERROR_MISSING_ARG, ctx, VIEW_ENTITY)
	}
	entities, err := dataStore.Get(entity, nil)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, entities)
}
