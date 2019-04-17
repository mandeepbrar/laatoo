package core

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"github.com/fsnotify/fsnotify"
)

func (modMgr *moduleManager) addWatch(ctx core.ServerContext, modName string, modDir string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	//defer watcher.Close()
	modMgr.watchers = append(modMgr.watchers, watcher)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				log.Error(ctx, "Module change event ", "event", event)
				/*cont, err := ioutil.ReadFile(file)
				if err != nil {
					log.Error(ctx, "Error encountered during hot reload", "err", err)
				}

				err = actionF(svc.svrCtx, mod, file, dir, cont)
				if err != nil {
					log.Error(ctx, "Error encountered during hot reload", "err", err)
				}*/
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
				log.Error(ctx, "Module change event error ", "err", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(modDir); err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Watching module directory for change ", "dir", modDir)

	return nil
}
