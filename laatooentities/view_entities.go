package laatooentities

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	views "laatooview"
	"net/http"
	"strconv"
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
	var err error
	entity := ctx.Query(VIEW_ENTITY)
	if entity == "" {
		return errors.ThrowHttpError(views.VIEW_ERROR_MISSING_ARG, ctx, VIEW_ENTITY)
	}
	pagesize := -1
	pagesizeVal := ctx.Query(data.VIEW_PAGESIZE)
	if pagesizeVal != "" {
		pagesize, err = strconv.Atoi(pagesizeVal)
		if err != nil {
			return err
		}
	}
	pagenum := 1
	pagenumVal := ctx.Query(data.VIEW_PAGENUM)
	if pagenumVal != "" {
		pagenum, err = strconv.Atoi(pagenumVal)
		if err != nil {
			return err
		}
	}
	args := ctx.Query(VIEW_ARGS)
	perm := fmt.Sprintf("View %s", entity)
	log.Logger.Debugf("Executing entity view %s with args %s and permission %s", entity, args, perm)
	if !laatoocore.IsAllowed(ctx, perm) {
		return errors.ThrowHttpError(laatoocore.AUTH_ERROR_SECURITY, ctx)
	}

	var argsMap map[string]interface{}

	if len(args) > 0 {
		byt := []byte(args)
		if err := json.Unmarshal(byt, &argsMap); err != nil {
			return err
		}
	}

	entities, totalrecs, recsreturned, err := dataStore.Get(entity, argsMap, pagesize, pagenum, "")
	if err != nil {
		return err
	}
	log.Logger.Debugf("Totalrecs %d RecsReturned %d", totalrecs, recsreturned)
	ctx.Response().Header().Set(data.VIEW_TOTALRECS, fmt.Sprint(totalrecs))
	ctx.Response().Header().Set(data.VIEW_RECSRETURNED, fmt.Sprint(recsreturned))
	return ctx.JSON(http.StatusOK, entities)
}
