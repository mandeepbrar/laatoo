package static

import (
	"fmt"
	"io/ioutil"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/services"
)

type File struct {
	path        string
	fullcontent bool
	Content     *[]byte
}

const (
	CONF_STATIC_SERVICEFACTORY = "static_files"
	CONF_STATICSVC_DIRECTORY   = "directory"
	CONF_STATIC_FILEPARAM      = "file"
	CONF_STATIC_FILE_PATH      = "path"
	CONF_STATICSVC_FILEBUNDLE  = "filebundle"
	CONF_STATIC_FILE_INFO      = "info"
	CONF_STATIC_BUNDLEPARAM    = "bundle"
	CONF_STATIC_FILEBUNDLES    = "bundles"
	CONF_STATIC_FILES          = "files"
	CONF_STATIC_CACHE          = "cache"
	CONF_STATICSVC_FILE        = "file"
	CONF_STATIC_DIR            = "directory"
)

//Environment hosting an application
type StaticServiceFactory struct {
	*services.DefaultFactory
}

//Initialize service, register provider with laatoo
func init() {
	registry.RegisterServiceFactoryProvider(CONF_STATIC_SERVICEFACTORY, NewStaticServiceFactory)
}

//factory method returns the service object to the environment
func NewStaticServiceFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Info(ctx, "Creating static service")
	svc := &StaticServiceFactory{}
	svc.DefaultFactory = services.NewDefaultFactory(SvcCreate)
	return svc, nil
}

func SvcCreate(ctx core.ServerContext, name string, conf config.Config) (core.ServiceFunc, error) {
	switch name {
	/*** Provides service for serving any files in a directory*****/
	case CONF_STATICSVC_DIRECTORY:
		{
			return CreateDirectorySvc(ctx, name, conf)
		}
	/*** Provides service for serving files whose path has been specified*****/
	case CONF_STATICSVC_FILE:
		{
			return CreateFilesSvc(ctx, name, conf)
		}
	case CONF_STATICSVC_FILEBUNDLE:
		{
			return CreateFileBundleSvc(ctx, name, conf)
		}
	}
	return nil, nil
}

func CreateDirectorySvc(ctx core.ServerContext, name string, conf config.Config) (core.ServiceFunc, error) {
	dir, ok := conf.GetString(CONF_STATICSVC_DIRECTORY)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATICSVC_DIRECTORY)
	}
	log.Logger.Info(ctx, "Static service created", " Path", dir)
	return func(ctx core.RequestContext) error {
		filename, ok := ctx.GetString(CONF_STATIC_FILEPARAM)
		if ok {
			ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, fmt.Sprintf("%s/%s", dir, filename)))
		} else {
			ctx.SetResponse(core.NewServiceResponse(core.StatusNotFound, nil))
		}
		return nil
	}, nil
}

func CreateFilesSvc(ctx core.ServerContext, name string, conf config.Config) (core.ServiceFunc, error) {
	filesMap := make(map[string]*File, 10)
	filesConf, ok := conf.GetSubConfig(CONF_STATIC_FILES)
	if ok {
		filenames := filesConf.AllConfigurations()
		for _, filename := range filenames {
			fileconfig, _ := filesConf.GetSubConfig(filename)
			cacheStr, ok := fileconfig.GetString(CONF_STATIC_CACHE)
			cache := (cacheStr == "true")
			path, ok := fileconfig.GetString(CONF_STATIC_FILE_PATH)
			if !ok {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATIC_FILE_PATH)
			}
			file := &File{fullcontent: cache, path: path}
			if cache {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					return nil, errors.WrapError(ctx, err)
				}
				file.Content = &content
			}
			filesMap[filename] = file
		}
	}
	return func(filesMap map[string]*File) core.ServiceFunc {
		return func(ctx core.RequestContext) error {
			filename, ok := ctx.GetString(CONF_STATIC_FILEPARAM)
			if ok {
				file, ok := filesMap[filename]
				if ok {
					if !file.fullcontent {
						content, err := ioutil.ReadFile(file.path)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
						ctx.SetResponse(core.NewServiceResponse(core.StatusServeBytes, &content))
					} else {
						ctx.SetResponse(core.NewServiceResponse(core.StatusServeBytes, file.Content))
					}
				} else {
					ctx.SetResponse(core.NewServiceResponse(core.StatusNotFound, nil))
				}
			} else {
				ctx.SetResponse(core.NewServiceResponse(core.StatusNotFound, nil))
			}
			return nil
		}
	}(filesMap), nil
	return nil, nil
}

//The services start serving when this method is called
func (ds *StaticServiceFactory) StartServices(ctx core.ServerContext) error {
	return nil
}
