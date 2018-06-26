package main

import (
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"path"
)

func (svc *UI) processDependencies(ctx core.ServerContext, itemType string, itemName string, conf config.Config, modDir string) error {
	log.Trace(ctx, "Processing dependency", "depName", itemName)
	depFileName, ok := conf.GetString(ctx, "filename")
	if ok {
		depFile := path.Join(modDir, FILES_DIR, depFileName)
		ok, _, _ = utils.FileExists(depFile)
		if ok {
			cont, err := ioutil.ReadFile(depFile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			log.Info(ctx, "Added dependency", "dep", itemName)
			svc.vendorFiles[itemName] = cont
		} else {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Dependency File", depFileName)
		}
	} else {
		return errors.MissingConf(ctx, "filename")
	}
	return nil
}
