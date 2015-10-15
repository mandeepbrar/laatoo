package laatooentities

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	//views "laatooview"
	"laatoosdk/service"
	"net/http"
	"strconv"
)

const (
	VIEW_NAME   = "view_entities"
	VIEW_ENTITY = "entity"
	VIEW_ARGS   = "args"
)

type EntitiesView struct {
	Options        map[string]interface{}
	serviceContext service.ServiceContext
	entity         string
}

func NewEntitiesView(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	serviceContext := ctx.(service.ServiceContext)
	entityInt, ok := conf[VIEW_ENTITY]
	if !ok {
		return nil, errors.ThrowError(ctx, ENTITY_VIEW_MISSING_ARG, "Entity", VIEW_ENTITY)
	}
	return &EntitiesView{Options: conf, serviceContext: serviceContext, entity: entityInt.(string)}, nil
}

func init() {
	laatoocore.RegisterObjectProvider(VIEW_NAME, NewEntitiesView)
}

func (view *EntitiesView) Execute(ctx *echo.Context, dataStore data.DataService) error {
	var err error
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
	perm := fmt.Sprintf("View %s", view.entity)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Executing entity view", "Entity", view.entity, "Args", args, "Permission", perm)
	if !view.serviceContext.IsAllowed(ctx, perm) {
		return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
	}

	var argsMap map[string]interface{}

	if len(args) > 0 {
		byt := []byte(args)
		if err := json.Unmarshal(byt, &argsMap); err != nil {
			return err
		}
	}

	entities, totalrecs, recsreturned, err := dataStore.Get(ctx, view.entity, argsMap, pagesize, pagenum, "")
	if err != nil {
		return err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Executed View", "Entity", view.entity, "Totalrecs", totalrecs, "RecsReturned", recsreturned)
	ctx.Response().Header().Set(data.VIEW_TOTALRECS, fmt.Sprint(totalrecs))
	ctx.Response().Header().Set(data.VIEW_RECSRETURNED, fmt.Sprint(recsreturned))
	return ctx.JSON(http.StatusOK, entities)
}
