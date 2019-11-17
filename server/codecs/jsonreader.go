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
	v, err := p.Parse(arr)
	if err != nil {
		return err
	}
	return &JsonReader{root: v, keys: []string{}}, nil
}

func internalJsonReader(c ctx.Context, val core.SerializableReader) (core.SerializableReader, error) {
	jsonref := val.(*fastjson.Value)
	return &JsonReader{root: *jsonref, keys: []string{}}, nil
}

type JsonReader struct {
	keys []string
	root fastjson.Value
}

func (rdr *JsonReader) Read(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableReader, error) {
	keys := append(rdr.keys, prop)
	v := rdr.root.Get(rdr.keys, prop)
	return &JsonReader{root: v, keys: keys}, nil
}

func (rdr *JsonReader) ReadByte(ctx ctx.Context, cdc core.Codec, prop string, val *byte) error {
	byts := rdr.root.GetByte(rdr.keys, prop)
	if byts != nil {
		*val = byts[0]
	}
	return nil
}

func (rdr *JsonReader) ReadBytes(ctx ctx.Context, cdc core.Codec) ([]byte, error) {
	return rdr.root.MarshalTo(nil), nil
}

func (rdr *JsonReader) ReadInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) error {
	intV := rdr.root.GetInt(rdr.keys, prop)
	*val = intV
	return nil
}

func (rdr *JsonReader) ReadInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) error {
	intV := rdr.root.GetInt64(rdr.keys, prop)
	*val = intV
	return nil
}

func (rdr *JsonReader) ReadString(ctx ctx.Context, cdc core.Codec, prop string, val *string) error {
	byts := rdr.root.GetStringBytes(rdr.keys, prop)
	if byts != nil {
		*val = string(byts)
	}
	return nil
}

func (rdr *JsonReader) ReadFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) error {
	fltV := rdr.root.GetFloat64(rdr.keys, prop)
	*val = float32(fltV)
	return nil
}

func (rdr *JsonReader) ReadFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) error {
	fltV := rdr.root.GetFloat64(rdr.keys, prop)
	*val = fltV
	return nil
}

func (rdr *JsonReader) ReadBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) error {
	boolV := rdr.root.GetBool(rdr.keys, prop)
	*val = boolV
	return nil
}

func (rdr *JsonReader) ReadObject(ctx ctx.Context, cdc core.Codec, prop, objtype string) (core.SerializableReader, error) {
	//ctx.(core.RequestContext).GetCodec("json", objtype)
	//keys := append(rdr.keys, prop)
	v := rdr.root.Get(rdr.keys, prop)
	if v == nil {
		return nil, errNoKey
	}
	return &JsonReader{root: v, keys: []string{}}, nil
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
	arrV := rdr.root.GetArray(rdr.keys, prop)
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
			rdr := &JsonReader{root: *v, keys: []string{}}
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
