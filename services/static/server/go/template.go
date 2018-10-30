package main

import (
	"bytes"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	CONF_STATIC_TEMPFILEDIR       = "templatefilesdirectory"
	CONF_STATIC_PROCESSEDFILESDIR = "processedfilesdirectory"
	CONF_FILES                    = "inputfiles"
	CONF_TEMPLATE_MAP             = "templatesmap"
)

type TemplateFileService struct {
	core.Service
	tempFilesDir  string
	procFilesDir  string
	inputFilesMap config.Config
	templatesMap  config.Config
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
	inputFilesMap, ok := svc.GetMapConfiguration(ctx, CONF_FILES)
	if ok {
		svc.inputFilesMap = inputFilesMap
	}
	templatesMap, ok := svc.GetMapConfiguration(ctx, CONF_TEMPLATE_MAP)
	if ok {
		svc.templatesMap = templatesMap
	}

	return svc.processTemplates(ctx)
}

func (svc *TemplateFileService) processTemplates(ctx core.ServerContext) error {
	contextVar := func(variable string) string {
		val, _ := ctx.GetString(variable)
		return val
	}
	fileContent := func(fileIdentifier string) string {
		if svc.inputFilesMap != nil {
			depPath, _ := svc.inputFilesMap.GetString(ctx, fileIdentifier)
			depContent, err := ioutil.ReadFile(depPath)
			if err != nil {
				return "File Not found" + depPath
			}
			return string(depContent)
		} else {
			return "Files map not provided"
		}
	}
	funcMap := template.FuncMap{"var": contextVar, "file": fileContent}
	if svc.templatesMap == nil {
		return svc.processAllFiles(ctx, funcMap)
	} else {
		return svc.createTargetFiles(ctx, funcMap)
	}
}

func (svc *TemplateFileService) processAllFiles(ctx core.ServerContext, funcMap template.FuncMap) error {
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

func (svc *TemplateFileService) createTargetFiles(ctx core.ServerContext, funcMap template.FuncMap) error {
	targetFiles := svc.templatesMap.AllConfigurations(ctx)
	for _, targetFile := range targetFiles {
		templateFile, _ := svc.templatesMap.GetString(ctx, targetFile)
		tempFilePath := path.Join(svc.tempFilesDir, templateFile)
		temp, err := template.New(templateFile).Funcs(funcMap).ParseFiles(tempFilePath)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		result := new(bytes.Buffer)
		anon := struct{}{}
		err = temp.Execute(result, anon)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		dest := path.Join(svc.procFilesDir, targetFile)
		return ioutil.WriteFile(dest, result.Bytes(), 0700)
	}
	return nil
}
