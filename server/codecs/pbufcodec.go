package codecs

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"

	//"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
	//"github.com/golang/protobuf/proto"
)

var (
	unmarshellingError = fmt.Errorf("Unmarshelling Error")
	marshellingError   = fmt.Errorf("Marshelling Error")
)

func NewProtobufCodec() *ProtobufCodec {
	return &ProtobufCodec{}
}

type ProtobufCodec struct {
}

func (cdc *ProtobufCodec) Unmarshal(c ctx.Context, arr []byte, val interface{}) error {
	msg, ok := val.(proto.Message)
	if !ok {
		return unmarshellingError
	}
	/*srl, ok := val.(core.Serializable)
	if ok {
		return cdc.UnmarshalSerializable(c, arr, srl)
	}*/
	return proto.Unmarshal(arr, msg)
}

func (codec *ProtobufCodec) Marshal(c ctx.Context, obj interface{}) ([]byte, error) {
	msg, ok := obj.(proto.Message)
	if !ok {
		return nil, marshellingError
	}
	/*srl, ok := obj.(core.Serializable)
	if ok {
		return codec.MarshalSerializable(c, srl)
	}*/
	return proto.Marshal(msg)
}

func (codec *ProtobufCodec) UnmarshalSerializable(c ctx.Context, arr []byte, obj core.Serializable) error {
	return UnmarshalSerializable(c, codec, nil, arr, obj)
}

func (codec *ProtobufCodec) MarshalSerializable(c ctx.Context, obj core.Serializable) ([]byte, error) {
	return MarshalSerializable(c, codec, nil, obj)
}

func (codec *ProtobufCodec) UnmarshalSerializableProps(c ctx.Context, arr []byte, obj core.Serializable, props map[string]interface{}) error {
	return UnmarshalSerializableProps(c, codec, nil, arr, obj, props)
}

func (codec *ProtobufCodec) MarshalSerializableProps(c ctx.Context, obj core.Serializable, props map[string]interface{}) ([]byte, error) {
	return MarshalSerializableProps(c, codec, nil, obj, props)
}
