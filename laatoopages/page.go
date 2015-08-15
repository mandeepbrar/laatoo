package laatoopages

import (
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net/http"
)

const (
	CONF_PAGE_PATH            = "path"
	CONF_PAGE_TYPE            = "type"
	CONF_PAGE_DEST            = "destination"
	CONF_PAGE_AUTH            = "authentication"
	CONF_PAGE_PERM            = "permission"
	CONF_PAGE_PARTIALS        = "partials"
	CONF_PAGE_PARTIALPATH     = "path"
	CONF_PAGE_PARTIALTEMPLATE = "template"
	CONF_PAGE_TYPE_FILE       = "file"
	CONF_PAGE_TYPE_REDIRECT   = "redirect"
)

func (svc *PageService) createPage(conf map[string]interface{}, router *echo.Group, pagesDir string) error {
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

		partialsInt, ok := conf[CONF_PAGE_PARTIALS]
		if ok {
			log.Logger.Infof("partialsInt %s", partialsInt)
			partialsConf, ok := partialsInt.(map[string]interface{})
			if !ok {
				return errors.ThrowError(PAGE_ERROR_WRONG_PARTIALS)
			}
			partialPages := make([]*entities.Partial, len(partialsConf))
			i := 0
			for partialPageName, val := range partialsConf {
				//get the config for the page with given alias
				partialConf := val.(map[string]interface{})
				obj := new(entities.Partial)
				obj.Name = partialPageName
				obj.Path, ok = partialConf[CONF_PAGE_PARTIALPATH].(string)
				if !ok {
					return errors.ThrowError(PAGE_ERROR_WRONG_PARTIALPATH)
				}
				templateFile, ok := partialConf[CONF_PAGE_PARTIALTEMPLATE].(string)
				if !ok {
					return errors.ThrowError(PAGE_ERROR_WRONG_PARTIALFILE)
				}
				content, err := ioutil.ReadFile(templateFile)
				if err != nil {
					return errors.RethrowError(PAGE_ERROR_WRONG_PARTIALS, err, CONF_PAGE_SERVICENAME, partialPageName)
				} else {
					obj.Template = string(content)
				}
				partialPages[i] = obj
				i++
			}

			confURL := fmt.Sprint(pagePath, "/conf")
			log.Logger.Infof("partialsURL %s", confURL)

			router.Get(confURL, func(ctx *echo.Context) error {
				config, err := svc.actionSvc.Execute(CONF_PAGE_GETALLACTIONS_METHOD, nil)
				if err != nil {
					return err
				}
				config["partials"] = partialPages
				ctx.JSON(http.StatusOK, config)
				return nil
			})
		}
	}

	log.Logger.Infof("pagePerm %s", pagePerm)
	log.Logger.Infof("pageAuth %s", pageAuth)
	log.Logger.Infof("pageType %s", pageType)
	return nil
}
