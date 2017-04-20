package static

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	CONF_STATICSVC_FILEBUNDLE = "filebundle"
	CONF_STATIC_BUNDLEPARAM   = "bundle"
	CONF_STATIC_BUNDLEFILES   = "files"
	CONF_STATIC_FILEBUNDLES   = "bundles"
	CONF_STATIC_MINIFY        = "minify"
)

type BundledFile struct {
	filepath string `json:"-"`
	Content  string
	Info     config.Config
}

type Bundle struct {
	Files        map[string]*BundledFile
	fullcontent  bool       `json:"-"`
	lastModified *time.Time `json:"-"`
}

type BundledFileService struct {
	bundlesMap map[string]*Bundle
	name       string
}

func (bs *BundledFileService) Initialize(ctx core.ServerContext, conf config.Config) error {
	bs.bundlesMap = make(map[string]*Bundle, 10)
	bundlesConf, ok := conf.GetSubConfig(CONF_STATIC_FILEBUNDLES)
	if ok {
		bundlenames := bundlesConf.AllConfigurations()
		for _, bundlename := range bundlenames {
			bundleconfig, _ := bundlesConf.GetSubConfig(bundlename)
			bundle, err := buildBundle(ctx, bundleconfig)
			if err != nil {
				return err
			}
			bs.bundlesMap[bundlename] = bundle
			log.Logger.Debug(ctx, "Created Bundle", "Name", bundlename)
		}
	}
	log.Logger.Info(ctx, "Bundle service initialized")
	return nil
}

func (bs *BundledFileService) Invoke(ctx core.RequestContext) error {
	bundlename, ok := ctx.GetString(CONF_STATIC_BUNDLEPARAM)
	if ok {
		bundle, ok := bs.bundlesMap[bundlename]
		log.Logger.Trace(ctx, "Get Bundle Method called", "Bundle", bundlename, "BundleFound", ok)
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
			log.Logger.Trace(ctx, "Get Bundle Method called", "Bundle", bundlename, "BundleFound", ok)
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

func (svc *BundledFileService) Start(ctx core.ServerContext) error {
	return nil
}

func buildBundle(ctx core.ServerContext, bundleconfig config.Config) (*Bundle, error) {
	bundlefiles, ok := bundleconfig.GetSubConfig(CONF_STATIC_BUNDLEFILES)
	if !ok {
		return nil, nil
	}
	minifyStr, _ := bundleconfig.GetString(CONF_STATIC_MINIFY)
	minify := (minifyStr == "true")
	cacheStr, _ := bundleconfig.GetString(CONF_STATIC_CACHE)
	cache := (cacheStr == "true")
	var lastModified *time.Time
	filenames := bundlefiles.AllConfigurations()
	bundleFiles := make(map[string]*BundledFile, len(filenames))
	for _, filename := range filenames {
		log.Logger.Trace(ctx, "Reading file for bundle", "Name", filename)
		fileconfig, _ := bundlefiles.GetSubConfig(filename)
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
		bundledFile, err := buildBundledFile(ctx, path, info, cache, minify)
		if err != nil {
			return nil, err
		}
		bundleFiles[filename] = bundledFile
	}
	return &Bundle{Files: bundleFiles, fullcontent: cache, lastModified: lastModified}, nil
}

func buildBundledFile(ctx core.Context, path string, info config.Config, readContent bool, minifyfiles bool) (*BundledFile, error) {
	bundledFile := &BundledFile{filepath: path, Info: info}
	if readContent {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		if minifyfiles {
			extension := filepath.Ext(path)
			log.Logger.Trace(ctx, "Minifying file", "extension", extension, "path", path)
			if extension == ".html" {
				m := minify.New()
				m.AddFunc("text/html", html.Minify)
				content, err = m.Bytes("text/html", content)
				if err != nil {
					return nil, errors.WrapError(ctx, err)
				}

			}
		}
		bundledFile.Content = string(content)
	}
	return bundledFile, nil
}

func buildFullContentBundle(ctx core.RequestContext, bundle *Bundle) (*Bundle, error) {
	bundleFiles := make(map[string]*BundledFile, len(bundle.Files))
	for filename, bundledFile := range bundle.Files {
		bundledFile, err := buildBundledFile(ctx, bundledFile.filepath, bundledFile.Info, true, false)
		if err != nil {
			return nil, err
		}
		bundleFiles[filename] = bundledFile
	}
	return &Bundle{Files: bundleFiles, fullcontent: true}, nil

}
