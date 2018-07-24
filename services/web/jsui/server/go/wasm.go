package main

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/utils"
	"path"
)

// #cgo LDFLAGS: -L ./lib/lib -lbinaryenmerge -lwasm -lasmjs -lemscripten-optimizer -lpasses -lir -lcfg -lsupport -lwasm -lstdc++ -lm
// #include <stdlib.h>
//char* mergeFiles(int argc, const char* argv[], const char* outputName);
//static char** makeCharArray(int size) {
//	return calloc(sizeof(char*), size);
//}
//static void setArrayString(char **a, char *s, int n) {
//	a[n] = s;
//}
//static void freeCharArray(char **a, int size) {
//	int i;
//	for (i = 0; i < size; i++)
//			free(a[i]);
//	free(a);
//}
import "C"

func (svc *UI) mergeWasm(file1, file2 string) {

	fileNames := []string{file1, file2}
	cargs := C.makeCharArray(C.int(len(fileNames)))
	defer C.freeCharArray(cargs, C.int(len(fileNames)))
	for i, s := range fileNames {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}
	/*
		cArray := C.malloc(C.size_t(len(fileNames)) * C.size_t(unsafe.Sizeof(uintptr(0))))
		a := (*[1<<30 - 1]*C.char)(cArray)

		for idx, str := range fileNames {
			a[idx] = C.CString(str)
		}*/

	C.mergeFiles(2, cargs, C.CString(file1))
}

/*
type wasmModule struct {
	Name string
	Data string
}*/

func (svc *UI) writeWasmFile(ctx core.ServerContext, baseDir string) error {
	//wasmFileCont := new(bytes.Buffer)

	/*wasmArr := make([]*wasmModule, 0)

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
	}*/

	wasmfile := path.Join(baseDir, FILES_DIR, WASM_DIR, svc.mergedwasmfile)
	first := true
	for _, path := range svc.wasmFiles {
		if first {
			err := utils.CopyFile(path, wasmfile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			first = false
		} else {
			svc.mergeWasm(wasmfile, path)
		}
	}

	return nil
}
