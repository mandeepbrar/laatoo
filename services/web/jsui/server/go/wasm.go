package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"os"
	"path"
	"strings"
)

/*
// #cgo LDFLAGS: -L ./lib/lib -lbinaryenmerge -lwasm -lasmjs -lemscripten-optimizer -lpasses -lir -lcfg -lsupport -lwasm -lstdc++ -lm -ldl -lbindgenlib
// #include <stdlib.h>
//char* mergeFiles(int argc, const char* argv[], const char* outputName);
//int bindgen(char* input, char* wasm_var, char* out_dir);
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

func (svc *UI) mergeWasm(fileNames []string, output string) {

	//fileNames := []string{file1, file2}
	cargs := C.makeCharArray(C.int(len(fileNames)))
	defer C.freeCharArray(cargs, C.int(len(fileNames)))
	for i, s := range fileNames {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}
	C.mergeFiles(2, cargs, C.CString(output))
}
*/

type wasmModule struct {
	Name string
	Data string
}

func (svc *UI) writeWasmFile(ctx core.ServerContext, baseDir string) error {
	ctx = ctx.SubContext("Write Wasm File")
	//wasmFileCont := new(bytes.Buffer)

	if err := svc.mergeWasmFilesToJson(ctx, baseDir); err != nil {
		return err
	}

	if err := svc.mergeJSFiles(ctx, baseDir); err != nil {
		return err
	}

	/*
		wasmfile := path.Join(baseDir, FILES_DIR, WASM_DIR, "tmp.wasm")
		filePaths := []string{}
		//first := true
		for _, path := range svc.wasmFiles {
			filePaths = append(filePaths, path)
		}
		svc.mergeWasm(filePaths, wasmfile)

		res := C.bindgen(C.CString(wasmfile), C.CString(svc.wasmModName), C.CString(path.Join(baseDir, FILES_DIR, WASM_DIR)))
		if res < 0 {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST, "Error", "Wasm file bindgen failed")
		}

		wasm_bg := path.Join(baseDir, FILES_DIR, WASM_DIR, "tmp_bg.wasm")

		ex, _, _ := utils.FileExists(wasm_bg)
		if ex {
			if err := os.Rename(wasm_bg, path.Join(baseDir, FILES_DIR, WASM_DIR, svc.mergedwasmfile)); err != nil {
				return err
			} else {
				f, err := os.Open(path.Join(baseDir, FILES_DIR, WASM_DIR, "tmp.js"))
				defer f.Close()
				if err != nil {
					return errors.WrapError(ctx, err)
				} else {
					js, err := ioutil.ReadAll(f)
					if err != nil {
						return err
					}
					svc.uiFiles[svc.wasmModName] = js
				}
			}
		}
	*/
	return nil
}

func (svc *UI) mergeWasmFilesToJson(ctx core.ServerContext, baseDir string) error {
	wasmArr := make([]*wasmModule, 0)

	written := make(map[string]bool)

	var appendMod func(string) error

	appendMod = func(name string) error {
		if written[name] {
			return nil
		}
		log.Error(ctx, "Writing mod", "name", name)
		deps := svc.modDeps[name]
		log.Error(ctx, "Writing mod", "deps", deps)
		for _, dep := range deps {
			if err := appendMod(dep); err != nil {
				return err
			}
		}
		wasmfile, ok := svc.wasmFiles[name]
		//do nothing if dependency not found in wasm as it may be dependent on js
		if ok {
			cont, err := ioutil.ReadFile(wasmfile)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			b64str := base64.StdEncoding.EncodeToString(cont)
			wasmArr = append(wasmArr, &wasmModule{name, b64str})
			written[name] = true
		}
		return nil
	}

	for name, _ := range svc.wasmFiles {
		if err := appendMod(name); err != nil {
			return err
		}
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

func (svc *UI) mergeJSFiles(ctx core.ServerContext, baseDir string) error {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprintln(buf, "function wasmBGImports() {var wasm;const __exports = {};__exports.__wasmInit=function(ins){wasm=ins;console.log('memory', wasm);};")
	skipCachedDec := false
	skipUint8Mem := false
	skipStringFunc := false
	for _, filePath := range svc.wasmImportFiles {
		f, err := os.Open(filePath)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		lc := 0
		skipLines := 4
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			if skipLines > 0 {
				skipLines = skipLines - 1
				continue
			}
			lc++
			text := sc.Text()
			if strings.Contains(text, "function init(wasm_path)") {
				break
			}
			if strings.Contains(text, "let cachedDecoder = new") {
				if skipCachedDec {
					continue
				} else {
					skipCachedDec = true
				}
			}

			if strings.Contains(text, "let cachegetUint8Memory = null;") {
				if skipUint8Mem {
					skipLines = 6
					continue
				} else {
					skipUint8Mem = true
				}
			}
			if strings.Contains(text, "function getStringFromWasm(ptr, len)") {
				if skipStringFunc {
					skipLines = 2
					continue
				} else {
					skipStringFunc = true
				}
			}

			fmt.Fprintln(buf, text)
		}
	}
	fmt.Fprintln(buf, "return __exports;}")
	fmt.Println(buf.String())
	/*wasmImportsFile := path.Join(baseDir, FILES_DIR, WASM_DIR, svc.mergedwasmimports)
	err := ioutil.WriteFile(wasmImportsFile, buf.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/
	svc.wasmImportScript = buf.Bytes()
	return nil
}
