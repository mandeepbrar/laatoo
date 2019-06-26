package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	CONF_STATIC_FILEPARAM          = "file"
	CONF_STATIC_REPLACEMENTS       = "replacements"
	CONF_STATIC_SEARCHSTRING       = "search"
	CONF_STATIC_REPLACEMENT_METHOD = "method"
	CONF_STATIC_FILE_PATH          = "path"
	CONF_STATIC_FILE_INFO          = "info"
	CONF_STATIC_FILES              = "files"
)

type File struct {
	path                string
	cachedcontent       bool
	Encoding            string
	Content             *[]byte
	lastModified        time.Time
	info                map[string]interface{}
	runtimereplacements bool
	replacements        []*replacement
}

type replacement struct {
	search string
	method string
}

type FileService struct {
	core.Service
	filesMap map[string]*File
	storage  components.StorageComponent
	name     string
}

func (fs *FileService) Initialize(ctx core.ServerContext, conf config.Config) error {
	/*fs.SetDescription(ctx, "Static files service")
	fs.AddStringConfigurations(ctx, []string{CONF_FILE_STORAGE}, nil)
	fs.AddConfigurations(ctx, map[string]string{CONF_STATIC_FILES: config.OBJECTTYPE_CONFIG})
	fs.AddStringParams(ctx, []string{CONF_STATIC_FILEPARAM}, nil)*/
	fs.filesMap = make(map[string]*File, 10)
	return nil
}

func (fs *FileService) Invoke(ctx core.RequestContext) error {
	filename, ok := ctx.GetStringParam(CONF_STATIC_FILEPARAM)
	log.Trace(ctx, "Param name for file", "filename", filename, "ok", ok)
	if ok {
		//filename := fn.GetValue().(string)
		log.Trace(ctx, "Providing file", "filename", filename)
		file, ok := fs.filesMap[filename]
		if ok {
			lastModTimeStr, ok := ctx.GetString(core.LastModified)
			log.Trace(ctx, "got header", "lastModTimeStr", lastModTimeStr)
			if ok {
				lastModTime, err := time.Parse(http.TimeFormat, lastModTimeStr)
				if err == nil {
					if !file.lastModified.After(lastModTime) {
						ctx.SetResponse(core.StatusNotModifiedResponse)
						return nil
					}
				}
			}
			if !file.cachedcontent {
				bytesArr, err := ioutil.ReadFile(file.path)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				content := &bytesArr
				if file.runtimereplacements {
					content = fs.runtimeReplace(ctx, content, file.replacements)
				}
				ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusServeBytes, content, file.info))
				return nil
			} else {
				content := file.Content
				if file.runtimereplacements {
					content = fs.runtimeReplace(ctx, content, file.replacements)
				}
				ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusServeBytes, content, file.info))
				return nil
			}
		} else {
			ctx.SetResponse(core.StatusNotFoundResponse)
			return nil
		}
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
		return nil
	}
}

func (fs *FileService) Start(ctx core.ServerContext) error {

	stg, _ := fs.GetStringConfiguration(ctx, CONF_FILE_STORAGE)
	stgSvc, err := ctx.GetService(stg)
	if err != nil {
		return err
	}
	fs.storage = stgSvc.(components.StorageComponent)

	conf, ok := fs.GetMapConfiguration(ctx, CONF_STATIC_FILES)
	if ok {
		filesConf := conf.(config.Config)
		filenames := filesConf.AllConfigurations(ctx)
		for _, filename := range filenames {
			fileconfig, _ := filesConf.GetSubConfig(ctx, filename)
			cacheStr, ok := fileconfig.GetString(ctx, CONF_STATIC_CACHE)
			cache := (cacheStr == "true")
			path, ok := fileconfig.GetString(ctx, CONF_STATIC_FILE_PATH)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATIC_FILE_PATH)
			}

			req := ctx.CreateSystemRequest("GetFilePath")
			path = fs.storage.GetFullPath(req, path)

			fil, err := os.Stat(path)
			if err != nil {
				return errors.WrapError(ctx, err, "Filepath", path)
			}
			mimetype := ""
			extension := filepath.Ext(fil.Name())
			if extension != "" {
				mimetype = mime.TypeByExtension(extension)
			}
			file := &File{cachedcontent: cache, path: path, lastModified: fil.ModTime(), info: make(map[string]interface{}, 2)}
			if mimetype != "" {
				file.info[core.ContentType] = mimetype
			}
			file.info[core.LastModified] = file.lastModified.UTC().Format(http.TimeFormat)

			replacements, ok := fileconfig.GetConfigArray(ctx, CONF_STATIC_REPLACEMENTS)
			if ok {
				file.runtimereplacements = ok
				file.replacements = make([]*replacement, len(replacements))
				for i, replaceConfig := range replacements {
					searchString, ok := replaceConfig.GetString(ctx, CONF_STATIC_SEARCHSTRING)
					if ok {
						replaceMethod, ok := replaceConfig.GetString(ctx, CONF_STATIC_REPLACEMENT_METHOD)
						if ok {
							file.replacements[i] = &replacement{searchString, replaceMethod}
						} else {
							return errors.BadConf(ctx, CONF_STATIC_REPLACEMENT_METHOD, "filename", filename)
						}
					} else {
						return errors.BadConf(ctx, CONF_STATIC_SEARCHSTRING, "filename", filename)
					}
				}
			}

			if cache {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				file.Content = &content
			}
			encoding, ok := fileconfig.GetString(ctx, core.ContentEncoding)
			if ok {
				file.info[core.ContentEncoding] = encoding
			}
			log.Trace(ctx, "Reading file", "file", file)
			fs.filesMap[filename] = file
		}
	}

	return nil
}

func (fs *FileService) runtimeReplace(ctx core.RequestContext, content *[]byte, replacements []*replacement) *[]byte {
	retVal := *content
	for _, replace := range replacements {
		if replace.method == "routeinfo" {
			p, _ := json.Marshal(ctx.EngineRequestParams())
			replaceString := []byte(fmt.Sprintf("<script>var _routeParams=%s;</script>", p))
			retVal = bytes.ReplaceAll(retVal, []byte(replace.search), replaceString)
		}
	}
	return &retVal
}
