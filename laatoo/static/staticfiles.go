package static

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type staticFiles struct {
	name      string
	directory string
}

func (svc *staticFiles) Initialize(ctx core.ServerContext, conf config.Config) error {
	dir, ok := conf.GetString(CONF_STATICSVC_DIRECTORY)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_STATICSVC_DIRECTORY)
	}
	svc.directory = dir
	return nil
}

func (svc *staticFiles) Invoke(ctx core.RequestContext) error {
	filename, ok := ctx.GetString(CONF_STATIC_FILEPARAM)
	if ok {
		ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, fmt.Sprintf("%s/%s", svc.directory, filename), nil))
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
	}
	return nil
}

func (svc *staticFiles) Start(ctx core.ServerContext) error {
	return nil
}
