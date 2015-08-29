package laatooauthentication

import (
	"encoding/json"
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/log"
	"net/http"
)

const (
	VIEW_ROLE        = "view_roles"
	VIEW_ENTITY_ROLE = "default_role"
	VIEW_ARGS        = "args"
)

type RolesView struct {
	Options map[string]interface{}
}

func NewRolesView(conf map[string]interface{}) (interface{}, error) {
	return &RolesView{conf}, nil
}

func init() {
	laatoocore.RegisterObjectProvider(VIEW_ROLE, NewRolesView)
}

func (view *RolesView) Execute(dataStore data.DataService, ctx *echo.Context) error {
	args := ctx.Query(VIEW_ARGS)
	log.Logger.Debugf("Executing roles view %s with args %s", args)

	var argsMap map[string]interface{}

	if len(args) > 0 {
		byt := []byte(args)
		if err := json.Unmarshal(byt, &argsMap); err != nil {
			return err
		}
	}

	entities, err := dataStore.Get(VIEW_ENTITY_ROLE, argsMap)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, entities)
}
