package utils

import (
	"fmt"
	"math/rand"
	"reflect"

	"laatoo.io/sdk/config"

	"golang.org/x/crypto/bcrypt"
)

type LookupFunc func(interface{}, string, interface{}) (interface{}, error)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Remove(arr []string, elem string) []string {
	for i, v := range arr {
		if v == elem {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func StrContains(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func CastToInterfaceArray(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	itemVal := reflect.ValueOf(val)
	if itemVal.Kind() != reflect.Slice && itemVal.Kind() != reflect.Array {
		return nil
	}
	leng := itemVal.Len()
	res := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		res[i] = itemVal.Index(i).Interface()
	}
	return res
}

/*
func CastToArrayType(val interface{}, typ reflect.Type) reflect.Value {
	log.Println("***************************** casting")
	if val == nil {
		return reflect.MakeSlice(typ, 0, 0)
	}
	itemVal := reflect.ValueOf(val)
	log.Println("***************************** setting field item val", val, itemVal.Kind())
	if itemVal.Kind() != reflect.Slice && itemVal.Kind() != reflect.Array {
		return reflect.MakeSlice(typ, 0, 0)
	}
	leng := itemVal.Len()
	log.Println("***************************** setting field item val", leng)
	res := reflect.MakeSlice(typ, leng, leng)
	log.Println("***************************** setting field")
	for i := 0; i < leng; i++ {
		log.Println("***************************** setting field val", itemVal.Index(i), itemVal.Index(i).Interface())
		obj := reflect.New(res.Index(i).Type())
		obj.SetBytes(itemVal.Index(i).Bytes())
		//converted := itemVal.Index(i).Convert()
		res.Index(i).Set(converted)
	}
	return res
}*/

func CastToStringArray(val interface{}) []string {
	if val == nil {
		return nil
	}
	itemVal := reflect.ValueOf(val)
	if itemVal.Kind() != reflect.Array && itemVal.Kind() != reflect.Slice {
		return nil
	}
	len := itemVal.Len()
	res := make([]string, len)
	for i := 0; i < len; i++ {
		res[i] = itemVal.Index(i).Interface().(string)
	}
	return res
}

func CastToMapArray(val interface{}) []map[string]interface{} {
	if val == nil {
		return nil
	}
	itemVal := reflect.ValueOf(val)
	if itemVal.Kind() != reflect.Array && itemVal.Kind() != reflect.Slice {
		return nil
	}
	len := itemVal.Len()
	res := make([]map[string]interface{}, len)
	for i := 0; i < len; i++ {
		res[i] = CastToStringMap(itemVal.Index(i).Interface())
	}
	return res
}

func CastToConfigArray(val interface{}) []config.Config {
	if val == nil {
		return nil
	}
	itemVal := reflect.ValueOf(val)
	if itemVal.Kind() != reflect.Array && itemVal.Kind() != reflect.Slice {
		return nil
	}
	len := itemVal.Len()
	res := make([]config.Config, len)
	for i := 0; i < len; i++ {
		res[i] = CastToConfig(itemVal.Index(i).Interface())
	}
	return res
}

func CastToConfig(val interface{}) config.Config {
	m := CastToStringMap(val)
	if m != nil {
		return GenericConfig(m)
	}
	return nil
}

func CastToStringMap(val interface{}) map[string]interface{} {
	if val == nil {
		return nil
	}
	itemVal := reflect.ValueOf(val)
	if itemVal.Kind() != reflect.Map {
		return nil
	}
	keys := itemVal.MapKeys()
	res := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		res[key.String()] = itemVal.MapIndex(key).Interface()
	}
	return res
}

func MapKeys(mapToProcess map[string]interface{}) []string {
	maplen := len(mapToProcess)
	if maplen < 1 {
		return []string{}
	}
	retVal := make([]string, maplen)
	i := 0
	for k, _ := range mapToProcess {
		retVal[i] = k
		i++
	}
	return retVal
}

func MapValues(mapToProcess map[string]interface{}) interface{} {
	maplen := len(mapToProcess)
	if maplen < 1 {
		return []interface{}{}
	}
	var arr reflect.Value
	i := 0
	for _, v := range mapToProcess {
		if i == 0 {
			sliceType := reflect.SliceOf(reflect.TypeOf(v))
			arr = reflect.MakeSlice(sliceType, maplen, maplen)
		}
		arr.Index(i).Set(reflect.ValueOf(v))
		i++
	}
	return arr.Interface()
}

func ElementPtr(object interface{}) interface{} {
	return reflect.ValueOf(object).Elem().Interface()
}

func EncryptPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// object: object for which fields are to be set
// newvals: values to be set on the object
// mappings: if fields from the map need to be set to specific fields of the object
// field processor: if values need to be transformed from the map before being set on the object
func SetObjectFields(ctx interface{}, object interface{}, newVals map[string]interface{},
	mappings map[string]string, fieldProcessor map[string]LookupFunc) error {
	entVal := reflect.ValueOf(object).Elem()
	var err error
	for k, v := range newVals {
		objField := k
		objVal := v
		if mappings != nil {
			newfld, ok := mappings[k]
			if ok {
				objField = newfld
			}
		}
		if fieldProcessor != nil {
			tfunc, ok := fieldProcessor[objField]
			if ok {
				objVal, err = tfunc(ctx, objField, objVal)
				if err != nil {
					return err
				}
			}
		}
		if objVal == nil {
			continue
		}
		f := entVal.FieldByName(objField)
		if f.IsValid() {
			if f.CanSet() {
				kind := f.Kind()
				switch kind {
				case reflect.Slice:
					{
						if f.Type().String() == "[]string" {
							arr := CastToStringArray(objVal)
							f.Set(reflect.ValueOf(arr))
						} else if f.Type().String() == "[]config.Config" {
							arr := CastToConfigArray(objVal)
							f.Set(reflect.ValueOf(arr))
						} else if f.Type().String() == "[]map[string]interface{}" {
							arr := CastToMapArray(objVal)
							f.Set(reflect.ValueOf(arr))
						}
						continue
					}
				case reflect.Struct:
					{
						objCreator := ctx.(ObjectCreator)
						objType, isPtr := GetRegisteredName(f.Type())

						structobj, err := objCreator.CreateObject(objType)
						if err != nil {
							return err
						}
						structVals, ok := objVal.(map[string]interface{})
						if ok {
							err = SetObjectFields(ctx, structobj, structVals, mappings, fieldProcessor)
							if err != nil {
								return err
							}
							if isPtr {
								f.Set(reflect.ValueOf(structobj).Convert(f.Type()))
							} else {
								f.Set(reflect.ValueOf(structobj).Elem().Convert(f.Type()))
							}
						}
					}
				default:
					{
						f.Set(reflect.ValueOf(objVal).Convert(f.Type()))
					}
				}
			}
		}
	}
	return nil
}

func GetRegisteredName(typ reflect.Type) (regName string, isptr bool) {
	for {
		kind := typ.Kind()
		if kind == reflect.Ptr {
			typ = typ.Elem()
			isptr = true
			continue
		}
		if kind == reflect.Array || kind == reflect.Slice {
			typ = typ.Elem()
			isptr = false
			continue
		}
		break
	}
	pkg := typ.PkgPath()
	if pkg != "" {
		regName = fmt.Sprintf("%s.%s", typ.PkgPath(), typ.Name())
	} else {
		regName = typ.Name()
	}
	return
}

func GetObjectFields(object interface{}, fields []string) map[string]interface{} {
	entVal := reflect.ValueOf(object).Elem()
	vals := make(map[string]interface{}, len(fields))
	for _, v := range fields {
		f := entVal.FieldByName(v)
		vals[v] = f.Interface()
	}
	return vals
}
