package entities

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net/http"
)

func ParseConfig(conf map[string]interface{}, svc EntityService, entityName string, entityCreator laatoocore.ObjectFactory) (string, *echo.Group, error) {
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return "", nil, errors.ThrowError(ENTITY_ERROR_MISSING_ROUTER)
	}
	router := routerInt.(*echo.Group)

	entitydatasvcInt, ok := conf[CONF_ENTITY_DATA_SVC]
	if !ok {
		return "", nil, errors.ThrowError(ENTITY_ERROR_MISSING_DATASVC, svc.GetName())
	}

	entitymethodsInt, ok := conf[CONF_ENTITY_METHODS]
	if !ok {
		return "", nil, errors.ThrowError(ENTITY_ERROR_MISSING_METHODS, svc.GetName())
	}

	entityMethods, ok := entitymethodsInt.(map[string]interface{})
	if !ok {
		return "", nil, errors.ThrowError(ENTITY_ERROR_MISSING_METHODS, svc.GetName())
	}

	for name, val := range entityMethods {

		methodConfig, ok := val.(map[string]interface{})
		if !ok {
			return "", nil, errors.ThrowError(ENTITY_ERROR_INCORRECT_METHOD_CONF, svc.GetName(), name)
		}

		pathInt, ok := methodConfig[CONF_ENTITY_PATH]
		if !ok {
			return "", nil, errors.ThrowError(ENTITY_ERROR_MISSING_PATH, svc.GetName(), name)
		} else {
			methodInt, ok := methodConfig[CONF_ENTITY_METHOD]
			if !ok {
				return "", nil, errors.ThrowError(ENTITY_ERROR_MISSING_METHOD, svc.GetName(), name)
			}

			path := pathInt.(string)
			method := methodInt.(string)

			switch method {
			/*			case "list":
						router.Get(path, svc.ListArticle)*/
			case "get":
				router.Get(path, func(ctx *echo.Context) error {
					id := ctx.P(0)
					log.Logger.Debugf("Getting entity %s", id)
					ent, err := svc.GetDataStore().GetById(entityName, id)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ent)
				})
			case "post":
				router.Post(path, func(ctx *echo.Context) error {
					ent, err := entityCreator(nil)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					err = svc.GetDataStore().Save(entityName, ent)
					if err != nil {
						return err
					}
					return nil
				})
			case "put":
				router.Put(path, func(ctx *echo.Context) error {
					id := ctx.P(0)
					log.Logger.Debugf("Updating entity %s", id)
					ent, err := entityCreator(nil)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					err = svc.GetDataStore().Put(entityName, id, ent)
					if err != nil {
						return err
					}
					return nil
				})
			case "delete":
				router.Delete(path, func(ctx *echo.Context) error {
					id := ctx.P(0)
					log.Logger.Debugf("Deleting entity %s", id)
					err := svc.GetDataStore().Delete(entityName, id)
					if err != nil {
						return err
					}
					return nil
				})
			}
		}
	}
	return entitydatasvcInt.(string), router, nil
}
