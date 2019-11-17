package codecs

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"reflect"
	"time"

	"github.com/valyala/fastjson"
)

func NewJsonReader(c ctx.Context, arr []byte) (core.SerializableReader, error) {
	var p fastjson.Parser
	v, err := p.ParseBytes(arr)
	if err != nil {
		return nil, err
	}
	return &JsonReader{root: v}, nil
}

func internalJsonReader(c ctx.Context, val core.SerializableReader) (core.SerializableReader, error) {
	jsonref := val.(*JsonReader)
	return &JsonReader{root: jsonref.root}, nil
}

type JsonReader struct {
	//	keys []string
	root *fastjson.Value
}

func (rdr *JsonReader) Bytes() []byte {
	return rdr.root.MarshalTo(nil)
}

func (rdr *JsonReader) ReadBytes(ctx ctx.Context, cdc core.Codec, prop string) ([]byte, error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil, errNoKey
	}
	return v.MarshalTo(nil), nil
}

func (rdr *JsonReader) Read(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableReader, error) {
	//keys := append(rdr.keys, prop)
	v := rdr.root.Get(prop)
	return &JsonReader{root: v}, nil
}

func (rdr *JsonReader) ReadInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) error {
	intV := rdr.root.GetInt(prop)
	*val = intV
	return nil
}

func (rdr *JsonReader) ReadInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) error {
	intV := rdr.root.GetInt64(prop)
	*val = intV
	return nil
}

func (rdr *JsonReader) ReadString(ctx ctx.Context, cdc core.Codec, prop string, val *string) error {
	byts := rdr.root.GetStringBytes(prop)
	if byts != nil {
		*val = string(byts)
	}
	return nil
}

func (rdr *JsonReader) ReadFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) error {
	fltV := rdr.root.GetFloat64(prop)
	*val = float32(fltV)
	return nil
}

func (rdr *JsonReader) ReadFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) error {
	fltV := rdr.root.GetFloat64(prop)
	*val = fltV
	return nil
}

func (rdr *JsonReader) ReadBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) error {
	boolV := rdr.root.GetBool(prop)
	*val = boolV
	return nil
}

func (rdr *JsonReader) ReadObject(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) (err error) {
	//ctx.(core.RequestContext).GetCodec("json", objtype)
	//keys := append(rdr.keys, prop)

	v := rdr.root.Get(prop)
	if v == nil {
		return errNoKey
	}

	srl, ok := val.(core.Serializable)
	if ok {
		err = cdc.UnmarshalReader(ctx, rdr, srl)
		if err != nil {
			return err
		}
	} else {
		byts := v.MarshalTo(nil)
		err = cdc.Unmarshal(ctx, byts, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rdr *JsonReader) ReadMap(ctx ctx.Context, cdc core.Codec, prop string, val *map[string]interface{}) error {
	//valMap := *val
	v := rdr.root.Get(prop)
	if v != nil {
		return errNoKey
	}
	bys := v.MarshalTo(nil)
	return cdc.Unmarshal(ctx, bys, val)
}

func (rdr *JsonReader) ReadArray(ctx ctx.Context, cdc core.Codec, prop string, objType string, val interface{}) error {
	reqCtx := ctx.(core.RequestContext)
	arrV := rdr.root.GetArray(prop)
	if arrV != nil {
		return errNoKey
	}
	collection, err := reqCtx.CreateCollection(objType, len(arrV))
	if err != nil {
		v := rdr.root.Get(prop)
		bys := v.MarshalTo(nil)
		return cdc.Unmarshal(ctx, bys, val)
	} else {
		collVal := reflect.ValueOf(collection)
		for i, v := range arrV {
			rdr := &JsonReader{root: v}
			obj, err := reqCtx.CreateObject(objType)
			if err != nil {
				return err
			}
			srl, ok := obj.(core.Serializable)
			if ok {
				err = cdc.UnmarshalReader(ctx, rdr, srl)
				if err != nil {
					return err
				}
				iVal := collVal.Index(i)
				iVal.Set(reflect.ValueOf(obj))
			}
		}
		reflect.ValueOf(val).Set(collVal)
	}
	return nil
}

func (rdr *JsonReader) ReadTime(ctx ctx.Context, cdc core.Codec, prop string, val *time.Time) error {
	v := rdr.root.Get(prop)
	if v != nil {
		return errNoKey
	}
	bys := v.MarshalTo(nil)
	tim, err := time.Parse(time.RFC3339, string(bys))
	if err != nil {
		return err
	}
	*val = tim
	return nil
}
