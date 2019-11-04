package codecs

import "laatoo/sdk/server/core"

func GetCodec(encoding string) (core.Codec, bool) {
	switch encoding {
	case "json":
		return NewJsonCodec(), true
	case "fastjson":
		return NewFastJsonCodec(), true
	case "bin":
		return NewBinaryCodec(), true
	case "protobuf":
		return NewProtobufCodec(), true
	}
	return nil, false
}
