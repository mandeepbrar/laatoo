package laatoopages

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net/http"
)

const (
	CONF_PAGE_PATH          = "path"
	CONF_PAGE_TYPE          = "type"
	CONF_PAGE_DEST          = "destination"
	CONF_PAGE_AUTH          = "authentication"
	CONF_PAGE_PERM          = "permission"
	CONF_PAGE_PARAMS        = "params"
	CONF_PAGE_TYPE_FILE     = "file"
	CONF_PAGE_TYPE_REDIRECT = "redirect"
)

func createPage(conf map[string]interface{}, router *echo.Group, pagesDir string) error {
	var pagePath, pageDest string
	pagePerm := ""
	pageType := CONF_PAGE_TYPE_FILE
	pageAuth := false

	pagepathInt, ok := conf[CONF_PAGE_PATH]
	if !ok {
		return errors.ThrowError(PAGE_ERROR_MISSING_PAGEPATH)
	}
	pagePath = pagepathInt.(string)

	pagetypeInt, ok := conf[CONF_PAGE_TYPE]
	if ok {
		pageType = pagetypeInt.(string)
	}

	pagedestInt, ok := conf[CONF_PAGE_DEST]
	if !ok {
		return errors.ThrowError(PAGE_ERROR_MISSING_DEST)
	}
	pageDest = pagedestInt.(string)

	pageauthInt, ok := conf[CONF_PAGE_AUTH]
	if ok {
		if pageauthInt.(string) == "true" {
			pageAuth = true
		}
	}

	pagepermInt, ok := conf[CONF_PAGE_PERM]
	if ok {
		pagePerm = pagepermInt.(string)
	}

	if pageType == CONF_PAGE_TYPE_FILE {
		dest := fmt.Sprint(pagesDir, "/", pageDest)
		router.ServeFile(pagePath, dest)

		params, ok := conf[CONF_PAGE_PARAMS]
		if ok {
			jsonObj, err := json.Marshal(params)
			if err != nil {
				return errors.RethrowError(PAGE_ERROR_WRONG_PARAMS, err)
			}
			confPage := fmt.Sprint(pagePath, "/conf")

			router.Get(confPage, func(ctx *echo.Context) error {
				ctx.String(http.StatusOK, string(jsonObj))
				return nil
			})
		}
	}

	log.Logger.Infof("pagePerm %s", pagePerm)
	log.Logger.Infof("pageAuth %s", pageAuth)
	log.Logger.Infof("pageType %s", pageType)
	return nil
}
