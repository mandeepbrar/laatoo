package jsui

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

func (svc *UI) reloadAppFile(ctx core.ServerContext, mod, file, dir string, cont []byte) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)
	svc.uiFiles[mod] = cont
	err := svc.writeAppFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) reloadVendorFile(ctx core.ServerContext, mod, file, dir string, cont []byte) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)
	svc.cssFiles[mod] = cont
	err := svc.writeVendorFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) reloadCSSFile(ctx core.ServerContext, mod, file, dir string, cont []byte) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)
	svc.vendorFiles[mod] = cont
	err := svc.writeCssFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) reloadRegistry(ctx core.ServerContext, itemType string) func(ctx core.ServerContext, mod, file, dir string, cont []byte) error {
	return func(ctx core.ServerContext, mod, file, dir string, cont []byte) error {
		err := svc.processRegItem(ctx, file, itemType, dir)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		return nil
	}
}
