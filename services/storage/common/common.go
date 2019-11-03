package storagecommon

import (
	"io"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"github.com/twinj/uuid"
)

func SaveFiles(ctx core.RequestContext, svc components.StorageComponent, bucket string) error {
	val, _ := ctx.GetParamValue("Data")
	files := *val.(*map[string]*core.MultipartFile)
	urls := map[string]string{}
	i := 0
	for filNam, fil := range files {
		defer fil.File.Close()
		fileName := uuid.NewV4().String()
		log.Debug(ctx, "writing file", "name", fileName, "MimeType", fil.MimeType)
		url, err := svc.SaveFile(ctx, bucket, fil.File, fileName, fil.MimeType)
		if err != nil {
			log.Debug(ctx, "Error while invoking upload", "err", err)
			return errors.WrapError(ctx, err)
		}
		urls[filNam] = url
		i++
	}
	ctx.SetResponse(core.SuccessResponse(urls))
	return nil
}

func CopyFile(ctx core.RequestContext, svc components.StorageComponent, bucket, fileName string, dest io.WriteCloser) error {
	// Source file
	src, err := svc.Open(ctx, bucket, fileName)
	if err != nil {
		return err
	}
	defer src.Close()

	if _, err = io.Copy(dest, src); err != nil {
		return err
	}
	return nil
}
