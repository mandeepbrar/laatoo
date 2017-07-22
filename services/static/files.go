package main

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	CONF_STATIC_FILEPARAM = "file"
	CONF_STATIC_FILE_PATH = "path"
	CONF_STATIC_FILE_INFO = "info"
	CONF_STATICSVC_FILE   = "file"
	CONF_STATIC_FILES     = "files"
)

type File struct {
	path         string
	fullcontent  bool
	Encoding     string
	Content      *[]byte
	lastModified time.Time
	info         map[string]interface{}
}

type FileService struct {
	filesMap map[string]*File
	name     string
}

func (fs *FileService) Initialize(ctx core.ServerContext, conf config.Config) error {
	fs.filesMap = make(map[string]*File, 10)
	filesConf, ok := conf.GetSubConfig(CONF_STATIC_FILES)
	if ok {
		filenames := filesConf.AllConfigurations()
		for _, filename := range filenames {
			fileconfig, _ := filesConf.GetSubConfig(filename)
			cacheStr, ok := fileconfig.GetString(CONF_STATIC_CACHE)
			cache := (cacheStr == "true")
			path, ok := fileconfig.GetString(CONF_STATIC_FILE_PATH)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATIC_FILE_PATH)
			}
			fil, err := os.Stat(path)
			if err != nil {
				return err
			}
			mimetype := ""
			extension := filepath.Ext(fil.Name())
			if extension != "" {
				mimetype = mime.TypeByExtension(extension)
			}
			file := &File{fullcontent: cache, path: path, lastModified: fil.ModTime(), info: make(map[string]interface{}, 2)}
			if mimetype != "" {
				file.info[core.ContentType] = mimetype
			}
			file.info[core.LastModified] = file.lastModified.UTC().Format(http.TimeFormat)
			if cache {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				file.Content = &content
			}
			encoding, ok := fileconfig.GetString(core.ContentEncoding)
			if ok {
				file.info[core.ContentEncoding] = encoding
			}
			log.Trace(ctx, "Reading file", "file", file)
			fs.filesMap[filename] = file
		}
	}
	return nil
}

func (fs *FileService) Info() *core.ServiceInfo {
	return &core.ServiceInfo{Description: "Static files service",
		Request: core.BuildRequestInfo("", []string{CONF_STATIC_FILEPARAM})}
}

func (fs *FileService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	fn, ok := req.GetParam(CONF_STATIC_FILEPARAM)
	if ok {
		filename := fn.Value.(string)
		log.Trace(ctx, "Providing file", "filename", filename)
		file, ok := fs.filesMap[filename]
		if ok {
			lastModTimeStr, ok := ctx.GetString(core.LastModified)
			log.Trace(ctx, "got header", "lastModTimeStr", lastModTimeStr)
			if ok {
				lastModTime, err := time.Parse(http.TimeFormat, lastModTimeStr)
				if err == nil {
					if !file.lastModified.After(lastModTime) {
						return core.StatusNotModifiedResponse, nil
					}
				}
			}
			if !file.fullcontent {
				content, err := ioutil.ReadFile(file.path)
				if err != nil {
					return nil, errors.WrapError(ctx, err)
				}
				return core.NewServiceResponse(core.StatusServeBytes, &content, file.info), nil
			} else {
				return core.NewServiceResponse(core.StatusServeBytes, file.Content, file.info), nil
			}
		} else {
			return core.StatusNotFoundResponse, nil
		}
	} else {
		return core.StatusNotFoundResponse, nil
	}
}

func (fs *FileService) Start(ctx core.ServerContext) error {
	return nil
}
