package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"path"
)

type wasmModule struct {
	Name string
	Data string
}

func (svc *UI) writeWasmFile(ctx core.ServerContext, baseDir string) error {
	//wasmFileCont := new(bytes.Buffer)

	wasmArr := make([]*wasmModule, 0)

	written := make(map[string]bool)

	var appendMod func(string)

	appendMod = func(name string) {
		if written[name] {
			return
		}
		log.Error(ctx, "Writing mod", "name", name)
		deps := svc.modDeps[name]
		log.Error(ctx, "Writing mod", "deps", deps)
		for _, dep := range deps {
			appendMod(dep)
		}
		cont, ok := svc.wasmFiles[name]
		//do nothing if dependency not found in wasm as it may be dependent on js
		if ok {
			b64str := base64.StdEncoding.EncodeToString(cont)
			wasmArr = append(wasmArr, &wasmModule{name, b64str})
			written[name] = true
		}
	}

	/*for name, cont := range svc.wasmFiles {

		b64str := base64.StdEncoding.EncodeToString(cont)
		wasmMap[name] = b64str
	}*/
	for name, _ := range svc.wasmFiles {
		appendMod(name)
	}
	//deps := svc.modDeps[name]
	log.Error(ctx, "Wasm array last", "wasmArr", wasmArr)

	filesCont, err := json.Marshal(wasmArr)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Writing wasm file ", "file name", svc.mergedwasmfile, "len", len(wasmArr))
	wasmfile := path.Join(baseDir, FILES_DIR, WASM_DIR, svc.mergedwasmfile)
	err = ioutil.WriteFile(wasmfile, filesCont, 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
