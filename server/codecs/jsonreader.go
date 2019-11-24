package codecs

import (
	"io"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
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

//Not full support for reader
func NewJsonStreamReader(c ctx.Context, rdr io.Reader) (core.SerializableReader, error) {
	arr, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, err
	}
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
	io.Reader
	//	keys []string
	root *fastjson.Value
}

func (rdr *JsonReader) Start() error {
	return nil
}

func (rdr *JsonReader) Bytes() []byte {
	return rdr.root.MarshalTo(nil)
}

func (rdr *JsonReader) ReadBytes(ctx ctx.Context, cdc core.Codec, prop string) ([]byte, error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil, nil
	}
	return v.MarshalTo(nil), nil
}

func (rdr *JsonReader) ReadProp(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableReader, error) {
	//keys := append(rdr.keys, prop)
	v := rdr.root.Get(prop)
	if v == nil {
		return nil, nil
	}
	return &JsonReader{root: v}, nil
}

func (rdr *JsonReader) ReadInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) (err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	*val, err = v.Int()
	return err
}

func (rdr *JsonReader) ReadInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) (err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	*val, err = v.Int64()
	return nil
}

func (rdr *JsonReader) ReadString(ctx ctx.Context, cdc core.Codec, prop string, val *string) error {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	byts, err := v.StringBytes()

	if byts != nil {
		*val = string(byts)
	}
	return err
}

func (rdr *JsonReader) ReadFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) (err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	flt, err := v.Float64()
	*val = float32(flt)
	return err
}

func (rdr *JsonReader) ReadFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) (err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	*val, err = v.Float64()
	return err
}

func (rdr *JsonReader) ReadBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) (err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	*val, err = v.Bool()
	return err
}

func (rdr *JsonReader) ReadObject(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) (err error) {
	log.Error(ctx, "Reading object", "prop", prop)
	v := rdr.root.Get(prop)
	if v == nil {
		return err
	}
	/*obj, _, err := rdr.createObject(ctx, cdc, val)
	if err != nil {
		return err
	}*/
	err = rdr.readObject(ctx, cdc, v, val)
	return err
}

func (rdr *JsonReader) ReadMap(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) error {
	//valMap := *val
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	bys := v.MarshalTo(nil)
	return cdc.Unmarshal(ctx, bys, val)
}

func (rdr *JsonReader) ReadArray(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) error {
	log.Error(ctx, "Reading array", "prop", prop)
	arrV := rdr.root.Get(prop)
	if arrV == nil {
		return nil
	}
	return rdr.readObject(ctx, cdc, arrV, val)
}

func (rdr *JsonReader) ReadTime(ctx ctx.Context, cdc core.Codec, prop string, val *time.Time) error {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	bys := v.GetStringBytes()
	tim, err := time.Parse(time.RFC3339, string(bys))
	if err != nil {
		return err
	}
	*val = tim
	return nil
}

func (rdr *JsonReader) readObject(ctx ctx.Context, cdc core.Codec, val *fastjson.Value, obj interface{}) (err error) {

	objVal := reflect.ValueOf(obj)
	if objVal.Kind() == reflect.Ptr {
		log.Error(ctx, "reading object kind", "objVal.Elem().Kind()", objVal.Elem().Kind())
		switch objVal.Elem().Kind() {
		case reflect.Struct:
			{
				if objVal.IsZero() {
					obj, _, err = rdr.createObject(ctx, cdc, obj)
					if err != nil {
						return err
					}
				}
			}
		case reflect.Array:
			{
				return rdr.readArray(ctx, cdc, val, obj)
			}
		case reflect.Slice:
			{
				return rdr.readArray(ctx, cdc, val, obj)
			}
		}
	}
	srl, ok := obj.(core.Serializable)
	log.Error(ctx, "Reading object", "type", objVal.Kind(), "srl", srl, "ok", ok)
	if ok {
		newrdr := &JsonReader{root: val}
		err = srl.ReadAll(ctx, cdc, newrdr)
	} else {
		switch objVal.Kind() {
		case reflect.Array:
			{
				log.Error(ctx, "reading array", "val", val)
				return rdr.readArray(ctx, cdc, val, obj)
			}
		case reflect.Slice:
			{
				log.Error(ctx, "reading slice", "val", val)
				return rdr.readArray(ctx, cdc, val, obj)
			}
		}
		byts := val.MarshalTo(nil)
		err = json.Unmarshal(byts, obj)
	}
	return
}

func (rdr *JsonReader) createObject(ctx ctx.Context, cdc core.Codec, val interface{}) (obj interface{}, objType string, err error) {
	objVal := reflect.ValueOf(val)
	if objVal.Kind() != reflect.Ptr {
		return nil, "", errWrongObject
	}

	if !objVal.Elem().IsNil() {
		return
	}

	reqCtx := ctx.(core.RequestContext)
	objType, reg, _ := reqCtx.GetRegName(val)
	if reg {
		obj, err = reqCtx.CreateObject(objType)
		if err != nil {
			return
		}
		objVal.Elem().Set(reflect.ValueOf(obj))
	} else {
		obj := reflect.New(objVal.Type())
		objVal.Elem().Set(reflect.ValueOf(obj))
	}

	return
}

func (rdr *JsonReader) createCollection(ctx ctx.Context, cdc core.Codec, val interface{}, len int) (objType string, reg bool, isPtr bool, err error) {
	collVal := reflect.ValueOf(val)
	if collVal.Kind() != reflect.Ptr && collVal.Elem().Kind() != reflect.Array && collVal.Elem().Kind() != reflect.Slice {
		return "", reg, isPtr, errWrongObject
	}

	reqCtx := ctx.(core.RequestContext)
	objType, reg, isPtr = reqCtx.GetRegName(val)
	log.Error(ctx, "reading array", "objType", objType, "val", val, "isPtr", isPtr)
	var collection interface{}
	if reg {
		if isPtr {
			collection, err = reqCtx.CreateObjectPointersCollection(objType, len)
			if err != nil {
				return
			}
		} else {
			collection, err = reqCtx.CreateCollection(objType, len)
			if err != nil {
				return
			}
		}
	} else {
		collection = reflect.MakeSlice(collVal.Elem().Type(), len, len)
	}
	//if collVal.Kind() == reflect.Ptr {
	log.Error(ctx, "setting ptr", "objType", objType)
	collVal.Elem().Set(reflect.ValueOf(collection).Elem())
	/*} else {
		if collVal.CanSet() {
			collVal.Set(reflect.ValueOf(collection).Convert(collVal.Type()))
		}
	}*/
	return
}

func (rdr *JsonReader) readArray(ctx ctx.Context, cdc core.Codec, arrV *fastjson.Value, val interface{}) error {
	vArr, err := arrV.Array()
	if err != nil {
		return err
	}
	objType, reg, isPtr, err := rdr.createCollection(ctx, cdc, val, len(vArr))
	if err != nil {
		return err
	}
	log.Error(ctx, "reading array", "objType", objType, "reg", reg, "isPtr", isPtr)
	if !reg {
		byts := arrV.MarshalTo(nil)
		err = json.Unmarshal(byts, val)
		return err
	}
	collVal := reflect.ValueOf(val).Elem()
	for i, v := range vArr {
		log.Error(ctx, "reading array CreateObject", "objType", objType)
		obj, err := ctx.(core.RequestContext).CreateObject(objType)
		if err != nil {
			return err
		}
		err = rdr.readObject(ctx, cdc, v, obj)
		if err != nil {
			return err
		}
		log.Error(ctx, "reading array readObject", "obj", obj)
		iVal := collVal.Index(i)
		log.Error(ctx, "reading array readObject", "iVal", iVal)
		iVal.Set(reflect.ValueOf(obj))
	}
	return nil
}

/*
func (rdr *JsonReader) readMap(ctx ctx.Context, cdc core.Codec, mapV *fastjson.Value, val interface{}) error {

	objVal := reflect.ValueOf(val)
	if objVal.IsZero() {
		return nil
	}
	if objVal.Kind() == reflect.Array {
		return rdr.readArray(ctx, cdc, val, obj)
	}
	srl, ok := obj.(core.Serializable)
	if ok {
		newrdr := &JsonReader{root: val}
		err = srl.ReadAll(ctx, cdc, newrdr)
	} else {
		byts := val.MarshalTo(nil)
		err = json.Unmarshal(byts, obj)
	}
	return nil
}
*/
