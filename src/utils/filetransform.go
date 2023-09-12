package utils

import (
	"io"
	"laatoo/sdk/constants"
	"laatoo/sdk/ctx"
	"path"
	"path/filepath"
)

type FileTransform func(io.Reader, io.Writer) error

func GetAbsFilePath(ctx ctx.Context, fpath string) string {
	if filepath.IsAbs(fpath) {
		return fpath
	} else {
		basedir, _ := ctx.GetString(constants.BASEDIR)
		return path.Join(basedir, fpath)
	}
}
