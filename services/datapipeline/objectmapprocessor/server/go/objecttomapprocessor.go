package main

import (
	"laatoo/sdk/common/config"
	//"laatoo/sdk/modules/datapipeline"

	"laatoo/sdk/server/core"
	"reflect"
)

type ObjectToMapProcessor struct {
	core.Service
	//mappings      config.Config
	//fieldmappings map[string]string
}

func (proc *ObjectToMapProcessor) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (proc *ObjectToMapProcessor) Transform(ctx core.RequestContext, input interface{}) (interface{}, error) {
	res := proc.convertStructToMap(input) //for
	return res, nil
}

func (proc *ObjectToMapProcessor) convertStructToMap(input interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	rVal := reflect.ValueOf(input)
	if rVal.Kind() == reflect.Struct {
		typ := reflect.TypeOf(input)
		for i := 0; i < typ.NumField(); i++ {
			fldtyp := typ.Field(i)
			fldVal := rVal.FieldByName(fldtyp.Name)
			if fldVal.IsValid() {
				fldkind := fldtyp.Type.Kind()
				switch fldkind {
				case reflect.Struct:
					{
						val := proc.convertStructToMap(fldVal.Interface())
						res[fldtyp.Name] = val
					}
				case reflect.Ptr:
					{
						ptrfldVal := fldVal.Elem()
						if fldVal.Kind() == reflect.Struct {
							val := proc.convertStructToMap(ptrfldVal.Interface())
							res[fldtyp.Name] = val
						} else {
							res[fldtyp.Name] = ptrfldVal.Interface()
						}
					}
				default:
					{
						res[fldtyp.Name] = fldVal.Interface()
					}
				}
			}
		}

	}
	return res
}
