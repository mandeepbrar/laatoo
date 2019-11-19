package codecs

import (
	"errors"
	"io"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
)

var (
	errNoCodec     = errors.New("No codec found")
	errWrongObject = errors.New("Wrong Object Type")
)

func GetCodec(encoding string) (core.Codec, bool) {
	switch encoding {
	case "json":
		return NewJsonCodec(), true
	case "bin":
		//return NewBinaryCodec(), true
	case "protobuf":
		//return NewProtobufCodec(), true
	}
	return nil, false
}

type objReaderFromBytes func(c ctx.Context, arr []byte) (core.SerializableReader, error)
type objReaderFromStream func(c ctx.Context, rdr io.Reader) (core.SerializableReader, error)
type objReaderInternal func(c ctx.Context, val core.SerializableReader) (core.SerializableReader, error)
type writerCreator func(c ctx.Context) (core.SerializableWriter, error)
type writerFromStreamCreator func(c ctx.Context, wtr io.Writer) (core.SerializableWriter, error)

func UnmarshalSerializable(c ctx.Context, cdc core.Codec, rdrCreator objReaderFromBytes, arr []byte, obj core.Serializable) error {
	rdr, err := rdrCreator(c, arr)
	if err != nil {
		return err
	}
	return obj.ReadAll(c, cdc, rdr)
}

func UnmarshalStreamProps(c ctx.Context, cdc core.Codec, rdrCreator objReaderFromStream, reader io.Reader, obj core.Serializable, props map[string]interface{}) error {
	rdr, err := rdrCreator(c, reader)
	if err != nil {
		return err
	}
	return obj.ReadAll(c, cdc, rdr)
}

func UnmarshalSerializableFromStream(c ctx.Context, cdc core.Codec, rdrCreator objReaderFromStream, reader io.Reader, obj core.Serializable) error {
	rdr, err := rdrCreator(c, reader)
	if err != nil {
		return err
	}
	return obj.ReadAll(c, cdc, rdr)
}

func UnmarshalSerializableProps(c ctx.Context, cdc core.Codec, rdrCreator objReaderFromBytes, arr []byte, obj core.Serializable, props map[string]interface{}) error {
	rdr, err := rdrCreator(c, arr)
	if err != nil {
		return err
	}
	return obj.ReadAll(c, cdc, rdr)
}

func MarshalSerializable(c ctx.Context, cdc core.Codec, wtrCreator writerCreator, obj core.Serializable) ([]byte, error) {
	wtr, err := wtrCreator(c)
	if err != nil {
		return nil, err
	}
	err = wtr.Start()
	if err != nil {
		return nil, err
	}
	err = obj.WriteAll(c, cdc, wtr)
	if err != nil {
		return nil, err
	}
	return wtr.Bytes(), nil
}

func MarshalSerializableProps(c ctx.Context, cdc core.Codec, creator writerCreator, obj core.Serializable, props map[string]interface{}) ([]byte, error) {
	wtr, err := creator(c)
	if err != nil {
		return nil, err
	}
	err = obj.WriteAll(c, cdc, wtr)
	if err != nil {
		return nil, err
	}
	return wtr.Bytes(), nil
}

func MarshalSerializableToStream(c ctx.Context, cdc core.Codec, wtrCreator writerFromStreamCreator, obj core.Serializable, writer io.Writer) error {
	wtr, err := wtrCreator(c, writer)
	if err != nil {
		return err
	}
	defer wtr.Close()
	log.Error(c, "Marshalling start")
	err = wtr.Start()
	if err != nil {
		return err
	}
	err = obj.WriteAll(c, cdc, wtr)
	if err != nil {
		return err
	}
	log.Error(c, "Marshalling to stream ends")
	return nil
}

func MarshalStreamProps(c ctx.Context, cdc core.Codec, creator writerFromStreamCreator, obj core.Serializable, writer io.Writer, props map[string]interface{}) error {
	wtr, err := creator(c, writer)
	if err != nil {
		return err
	}
	defer wtr.Close()
	err = wtr.Start()
	if err != nil {
		return err
	}
	err = obj.WriteAll(c, cdc, wtr)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalReader(c ctx.Context, cdc core.Codec, intl objReaderInternal, rdr core.SerializableReader, obj core.Serializable) error {
	newrdr, err := intl(c, rdr)
	if err != nil {
		return err
	}
	return obj.ReadAll(c, cdc, newrdr)
}
