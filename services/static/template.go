package main

import (
	"bytes"
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	CONF_STATIC_TEMPFILEDIR       = "templatefilesdirectory"
	CONF_STATIC_PROCESSEDFILESDIR = "processedfilesdirectory"
)

type TemplateFileService struct {
	core.Service
	tempFilesDir string
	procFilesDir string
}

func (svc *TemplateFileService) Initialize(ctx core.ServerContext, conf config.Config) error {
	filesdir, ok := svc.GetStringConfiguration(ctx, CONF_STATIC_TEMPFILEDIR)
	if ok {
		svc.tempFilesDir = filesdir
	} else {
		baseDir, ok := ctx.GetString(config.MODULEDIR)
		if !ok {
			baseDir, _ = ctx.GetString(config.BASEDIR)
		}
		svc.tempFilesDir = path.Join(baseDir, "files")
	}
	procfilesdir, ok := svc.GetStringConfiguration(ctx, CONF_STATIC_PROCESSEDFILESDIR)
	if ok {
		svc.procFilesDir = procfilesdir
	} else {
		svc.procFilesDir = svc.tempFilesDir
	}
	return nil
}

func (svc *TemplateFileService) Start(ctx core.ServerContext) error {
	mymethod := func(variable string) string {
		val, _ := svc.GetStringConfiguration(ctx, variable)
		return val
	}
	funcMap := template.FuncMap{"var": mymethod}
	return filepath.Walk(svc.tempFilesDir, func(filepath string, f os.FileInfo, err error) error {
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		if !f.IsDir() {
			if strings.HasSuffix(f.Name(), ".tpl") {
				temp, err := template.New(f.Name()).Funcs(funcMap).ParseFiles(filepath)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				result := new(bytes.Buffer)
				anon := struct{}{}
				err = temp.Execute(result, anon)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				dest := path.Join(svc.procFilesDir, strings.TrimSuffix(f.Name(), ".tpl"))
				return ioutil.WriteFile(dest, result.Bytes(), 0700)
			}
		}
		return nil
	})
}
