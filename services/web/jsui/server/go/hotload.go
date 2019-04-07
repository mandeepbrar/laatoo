package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

/*
func (svc *UI) addWatch(ctx core.ServerContext, mod, file, dir string, actionF func(ctx core.ServerContext, mod, file, dir string, cont []byte) error) error {
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	//defer watcher.Close()
	svc.watchers = append(svc.watchers, watcher)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)

				cont, err := ioutil.ReadFile(file)
				if err != nil {
					log.Error(ctx, "Error encountered during hot reload", "err", err)
				}

				err = actionF(svc.svrCtx, mod, file, dir, cont)
				if err != nil {
					log.Error(ctx, "Error encountered during hot reload", "err", err)
				}
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(file); err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}
*/

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
