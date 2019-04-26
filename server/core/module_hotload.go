package core

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"github.com/radovskyb/watcher"
	//"github.com/fsnotify/fsnotify"
)

func (modMgr *moduleManager) addWatch(ctx core.ServerContext, modName string, modDir string) error {
	ctx = ctx.SubContext("Watch " + modName)

	/*
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return errors.WrapError(ctx, err)
		}*/
	w := watcher.New()
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case event := <-w.Event:
				log.Error(ctx, "Module change event ", "event", event)
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Error(ctx, "Error while watching", err)
			case <-w.Closed:
				return
			}
		}
	}()

	//defer watcher.Close()
	modMgr.watchers = append(modMgr.watchers, w)

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(modDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Error(ctx, "Watching module directory for change ", "dir", modDir, "watchers", modMgr.watchers)

	return nil
}

/*cont, err := ioutil.ReadFile(file)
if err != nil {
	log.Error(ctx, "Error encountered during hot reload", "err", err)
}

err = actionF(svc.svrCtx, mod, file, dir, cont)
if err != nil {
	log.Error(ctx, "Error encountered during hot reload", "err", err)
}*/
// watch for errors

/*
//
go func() {
	for {
		select {
		// watch for events
		case event := <-watcher.Events:
			//fmt.Printf("EVENT! %#v\n", event)
			log.Error(ctx, "Module change event ", "event", event.Name, "op", event.Op)
			// watch for errors
		case err := <-watcher.Errors:
			fmt.Println("ERROR", err)
			log.Error(ctx, "Module change event error ", "err", err)
		}
	}
}()

// out of the box fsnotify can watch a single file, or a single directory

visit := func(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		log.Trace(ctx, "Watching directory for change ", "dir", path)
		if err := watcher.Add(path); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

err = filepath.Walk(modDir, visit)
if err != nil {
	log.Error(ctx, "Could not walk through hot directory")
}
*/
