package publish_prod

import (
	"encoding/json"
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	views "laatooview"
	"net/http"
)

const (
	VIEW_NAME   = "view_entities"
	VIEW_ENTITY = "entity"
	VIEW_ARGS   = "args"
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
	args := ctx.Query(VIEW_ARGS)
	log.Logger.Debugf("Executing entity view %s with args %s", entity, args)

	var argsMap map[string]interface{}

	if len(args) > 0 {
		byt := []byte(args)
		if err := json.Unmarshal(byt, &argsMap); err != nil {
			return err
		}
	}

	entities, err := dataStore.Get(entity, argsMap)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, entities)
}
