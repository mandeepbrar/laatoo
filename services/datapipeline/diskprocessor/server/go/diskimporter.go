package main

import (
	"bufio"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_OBJECT_TO_IMPORT = "importobject"
	CONF_IMP_STORAGE      = "importstorageservice"
	CONF_IMP_BUCKET       = "importstoragebucket"
	CONF_IMP_FILE         = "importfile"
	CONF_IMP_CODEC        = "importcodec"
)

type DiskImporter struct {
	core.Service
	stor       components.StorageComponent
	storBucket string
	file       string
	objFac     core.ObjectFactory
	codec      core.Codec
	delim      byte
}

func (imp *DiskImporter) Initialize(ctx core.ServerContext, conf config.Config) error {
	svcName, ok := conf.GetString(ctx, CONF_IMP_STORAGE)
	if !ok {
		return errors.MissingConf(ctx, CONF_IMP_STORAGE)
	}
	stor, err := ctx.(core.ServerContext).GetService(svcName)
	if err != nil {
		return errors.WrapErrorWithCode(ctx, err, errors.CORE_ERROR_BAD_CONF)
	}
	imp.stor, ok = stor.(components.StorageComponent)
	if !ok {
		return errors.BadConf(ctx, CONF_IMP_STORAGE)
	}
	imp.storBucket, _ = conf.GetString(ctx, CONF_IMP_BUCKET)

	codecName, ok := conf.GetString(ctx, CONF_IMP_CODEC)
	if ok {
		codec, ok := ctx.GetCodec(codecName)
		if ok {
			imp.codec = codec
		}
	}

	if imp.codec == nil {
		imp.codec, _ = ctx.GetCodec("json")
	}

	imp.file, _ = conf.GetString(ctx, CONF_IMP_FILE)

	obj, ok := conf.GetString(ctx, CONF_OBJECT_TO_IMPORT)
	if ok {
		imp.objFac, ok = ctx.(core.ServerContext).GetObjectFactory(obj)
		if !ok {
			return errors.BadConf(ctx, CONF_OBJECT_TO_IMPORT)
		}
	} else {
		return errors.MissingConf(ctx, CONF_OBJECT_TO_IMPORT)
	}

	return nil
}

func (imp *DiskImporter) GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan) error {
	var file string
	fileInt, ok := initData["inputfile"]

	if !ok {
		file = imp.file
	} else {
		file = fileInt.(string)
	}
	if file == "" {
		return errors.BadRequest(ctx, "Error", "input file not provided")
	}

	inpRdr, err := imp.stor.Open(ctx, imp.storBucket, file)
	if err != nil {
		return err
	}
	defer inpRdr.Close()

	rdr := bufio.NewReader(inpRdr)

	for {
		record, err := rdr.ReadBytes(imp.delim)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		obj := imp.objFac.CreateObject(ctx)

		err = imp.codec.Unmarshal(ctx, record, obj)
		if err != nil {
			return err
		}

		dataChan <- datapipeline.NewPipelineRecord(obj, nil)
	}
	log.Error(ctx, "Closing get records")
	return nil
}
