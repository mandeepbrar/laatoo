package static

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type BundledFile struct {
	filepath string `json:"-"`
	Content  *[]byte
	Info     config.Config
}

type Bundle struct {
	files       map[string]*BundledFile
	fullcontent bool `json:"-"`
}

func CreateFileBundleSvc(ctx core.ServerContext, name string, conf config.Config) (core.ServiceFunc, error) {
	bundlesMap := make(map[string]*Bundle, 10)
	bundlesConf, ok := conf.GetSubConfig(CONF_STATIC_FILEBUNDLES)
	if ok {
		bundlenames := bundlesConf.AllConfigurations()
		for _, bundlename := range bundlenames {
			bundleconfig, _ := bundlesConf.GetSubConfig(bundlename)
			bundle, err := buildBundle(ctx, bundleconfig)
			if err != nil {
				return nil, err
			}
			bundlesMap[bundlename] = bundle
		}
	}
	log.Logger.Info(ctx, "Bundle service created")
	return func(bundlesMap map[string]*Bundle) core.ServiceFunc {
		return func(ctx core.RequestContext) error {
			log.Logger.Info(ctx, "Method called")
			bundlename, ok := ctx.GetString(CONF_STATIC_BUNDLEPARAM)
			if ok {
				bundle, ok := bundlesMap[bundlename]
				if ok {
					if !bundle.fullcontent {
						newbundle, err := buildFullContentBundle(ctx, bundle)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
						ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, newbundle))
					} else {
						ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, bundle))
					}
				} else {
					ctx.SetResponse(core.NewServiceResponse(core.StatusNotFound, nil))
				}
			} else {
				ctx.SetResponse(core.NewServiceResponse(core.StatusNotFound, nil))
			}
			return nil
		}
	}(bundlesMap), nil
	return nil, nil
}

func buildBundle(ctx core.ServerContext, bundleconfig config.Config) (*Bundle, error) {
	filenames := bundleconfig.AllConfigurations()
	cacheStr, _ := bundleconfig.GetString(CONF_STATIC_CACHE)
	cache := (cacheStr == "true")
	bundleFiles := make(map[string]*BundledFile, len(filenames))
	for _, filename := range filenames {
		fileconfig, _ := bundleconfig.GetSubConfig(filename)
		path, ok := fileconfig.GetString(CONF_STATIC_FILE_PATH)
		if !ok {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATIC_FILE_PATH)
		}
		info, _ := fileconfig.GetSubConfig(CONF_STATIC_FILE_INFO)
		bundledFile, err := buildBundledFile(ctx, path, info, cache)
		if err != nil {
			return nil, err
		}
		bundleFiles[filename] = bundledFile
	}
	return &Bundle{files: bundleFiles, fullcontent: cache}, nil
}

func buildBundledFile(ctx core.Context, path string, info config.Config, readContent bool) (*BundledFile, error) {
	bundledFile := &BundledFile{filepath: path, Info: info}
	if readContent {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		bundledFile.Content = &content
	}
	return bundledFile, nil
}

func buildFullContentBundle(ctx core.RequestContext, bundle *Bundle) (*Bundle, error) {
	bundleFiles := make(map[string]*BundledFile, len(bundle.files))
	for filename, bundledFile := range bundle.files {
		bundledFile, err := buildBundledFile(ctx, bundledFile.filepath, bundledFile.Info, true)
		if err != nil {
			return nil, err
		}
		bundleFiles[filename] = bundledFile
	}
	return &Bundle{files: bundleFiles, fullcontent: true}, nil

}
