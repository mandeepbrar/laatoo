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
	VIEW_USER        = "view_users"
	VIEW_ENTITY_USER = "default_user"
)

type UsersView struct {
	Options map[string]interface{}
}

func NewUsersView(conf map[string]interface{}) (interface{}, error) {
	return &UsersView{conf}, nil
}

func init() {
	laatoocore.RegisterObjectProvider(VIEW_USER, NewUsersView)
}

func (view *UsersView) Execute(dataStore data.DataService, ctx *echo.Context) error {
	args := ctx.Query(VIEW_ARGS)
	log.Logger.Debugf("Executing user view %s with args %s", args)

	var argsMap map[string]interface{}

	if len(args) > 0 {
		byt := []byte(args)
		if err := json.Unmarshal(byt, &argsMap); err != nil {
			return err
		}
	}

	entities, err := dataStore.Get(VIEW_ENTITY_USER, argsMap)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, entities)
}
