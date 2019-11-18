package jsui

import (
	"bytes"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"path"
)

func (svc *UI) writeCssFile(ctx core.ServerContext, baseDir string) error {
	cssFileCont := new(bytes.Buffer)
	for _, cont := range svc.cssFiles {
		_, err := cssFileCont.Write(cont)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	cssfile := path.Join(baseDir, FILES_DIR, CSS_DIR, svc.mergedcssfile)
	err := ioutil.WriteFile(cssfile, cssFileCont.Bytes(), 0755)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
