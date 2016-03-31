package static

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net/http"
	"os"
	"time"
)

const (
	CONF_STATICSVC_FILEBUNDLE = "filebundle"
	CONF_STATIC_BUNDLEPARAM   = "bundle"
	CONF_STATIC_FILEBUNDLES   = "bundles"
)

type BundledFile struct {
	filepath string `json:"-"`
	Content  *[]byte
	Info     config.Config
}

type Bundle struct {
	files        map[string]*BundledFile
	fullcontent  bool       `json:"-"`
	lastModified *time.Time `json:"-"`
}

type BundledFileService struct {
	conf       config.Config
	bundlesMap map[string]*Bundle
}

func (bs *BundledFileService) Initialize(ctx core.ServerContext) error {
	bs.bundlesMap = make(map[string]*Bundle, 10)
	bundlesConf, ok := bs.conf.GetSubConfig(CONF_STATIC_FILEBUNDLES)
	if ok {
		bundlenames := bundlesConf.AllConfigurations()
		for _, bundlename := range bundlenames {
			bundleconfig, _ := bundlesConf.GetSubConfig(bundlename)
			bundle, err := buildBundle(ctx, bundleconfig)
			if err != nil {
				return err
			}
			bs.bundlesMap[bundlename] = bundle
		}
	}
	log.Logger.Info(ctx, "Bundle service initialized")
	return nil
}

func (bs *BundledFileService) Invoke(ctx core.RequestContext) error {
	bundlename, ok := ctx.GetString(CONF_STATIC_BUNDLEPARAM)
	log.Logger.Trace(ctx, "Get Bundle Method called", "Bundle", bundlename)
	if ok {
		bundle, ok := bs.bundlesMap[bundlename]
		if ok {
			lastModTimeStr, ok := ctx.GetString(core.LastModified)
			if ok {
				lastModTime, err := time.Parse(http.TimeFormat, lastModTimeStr)
				if err == nil {
					if !bundle.lastModified.After(lastModTime) {
						ctx.SetResponse(core.StatusNotModifiedResponse)
						return nil
					}
				}
			}
			if !bundle.fullcontent {
				newbundle, err := buildFullContentBundle(ctx, bundle)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, newbundle, nil))
			} else {
				ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, bundle, nil))
			}
		} else {
			ctx.SetResponse(core.StatusNotFoundResponse)
		}
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
	}
	return nil
}

func (bs *BundledFileService) GetConf() config.Config {
	return bs.conf
}

func (bs *BundledFileService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}

func buildBundle(ctx core.ServerContext, bundleconfig config.Config) (*Bundle, error) {
	filenames := bundleconfig.AllConfigurations()
	cacheStr, _ := bundleconfig.GetString(CONF_STATIC_CACHE)
	cache := (cacheStr == "true")
	var lastModified *time.Time
	bundleFiles := make(map[string]*BundledFile, len(filenames))
	for _, filename := range filenames {
		fileconfig, _ := bundleconfig.GetSubConfig(filename)
		path, ok := fileconfig.GetString(CONF_STATIC_FILE_PATH)
		if !ok {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATIC_FILE_PATH)
		}
		fil, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		fileTime := fil.ModTime()
		if lastModified == nil {
			lastModified = &fileTime
		} else {
			if lastModified.After(fileTime) {
				lastModified = &fileTime
			}
		}
		info, _ := fileconfig.GetSubConfig(CONF_STATIC_FILE_INFO)
		bundledFile, err := buildBundledFile(ctx, path, info, cache)
		if err != nil {
			return nil, err
		}
		bundleFiles[filename] = bundledFile
	}
	return &Bundle{files: bundleFiles, fullcontent: cache, lastModified: lastModified}, nil
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
