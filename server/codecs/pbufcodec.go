package codecs

import (
	"fmt"

	"github.com/golang/protobuf/proto"
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

func (cdc *ProtobufCodec) Unmarshal(arr []byte, val interface{}) error {
	msg, ok := val.(proto.Message)
	if !ok {
		return unmarshellingError
	}
	return proto.Unmarshal(arr, msg)
}

func (codec *ProtobufCodec) Marshal(obj interface{}) ([]byte, error) {
	msg, ok := obj.(proto.Message)
	if !ok {
		return nil, marshellingError
	}
	return proto.Marshal(msg)
}
