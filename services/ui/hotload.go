package main

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"

	"github.com/fsnotify/fsnotify"
)

func (svc *UI) addWatch(ctx core.ServerContext, mod, file, dir string, actionF func(ctx core.ServerContext, mod, file, dir string) error) error {
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)

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

	<-done
	return nil
}

func (svc *UI) reloadAppFile(ctx core.ServerContext, mod, file, dir string) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)
	err := svc.writeAppFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) reloadVendorFile(ctx core.ServerContext, mod, file, dir string) error {
	baseDir, _ := ctx.GetString(config.MODULEDIR)
	err := svc.writeVendorFile(ctx, baseDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
