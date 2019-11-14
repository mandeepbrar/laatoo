package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/utils"

	//"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type mapToObjectProcessor struct {
	core.Service
	objectFactory core.ObjectFactory
	mappings      config.Config
	fieldmappings map[string]string
	lookups       map[string]utils.LookupFunc
	lookupfields map[string]
}

func (proc *mapToObjectProcessor) Initialize(ctx core.ServerContext, conf config.Config) error {
	object, ok := proc.GetStringConfiguration(ctx, "object")
	if !ok {
		return errors.MissingConf(ctx, "object")
	}
	fact, ok := ctx.GetObjectFactory(object)
	if !ok {
		return errors.BadConf(ctx, "object", "Error", "No object found")
	}
	proc.objectFactory = fact
	mappings, ok := proc.GetMapConfiguration(ctx, "fieldmappings")
	if ok {
		proc.fieldmappings = make(map[string]string)
		proc.lookups = make(map[string]utils.LookupFunc)
		proc.lookupfields = make(map[string]string)
		fields := mappings.AllConfigurations(ctx)
		for _, mapfield := range fields {
			objField, ok := mappings.GetString(ctx, mapfield)
			if ok {
				proc.fieldmappings[mapfield] = objField
				continue
			}
			lookupFieldConf, ok := mappings.GetSubConfig(ctx, mapfield)
			if ok {
				objField, ok := lookupFieldConf.GetString(ctx, "field")
				if ok {
					lookupField, ok := lookupFieldConf.GetString(ctx, "lookupfield")
					if ok {
						proc.lookupfields[objField] = lookupField
					}
					proc.fieldmappings[mapfield] = objField
					lookupSvcName, ok := lookupFieldConf.GetString(ctx, "dataservice")
					if !ok {
						return errors.BadConf(ctx, "fieldmappings", "Message", "Lookup data service incorrect", "mapfield", mapfield, "detail", "missing 'dataservice'")
					} else {
						svc, err := ctx.GetService(lookupSvcName)
						if err != nil {
							return errors.BadConf(ctx, "fieldmappings", "Message", "Lookup data service incorrect", "mapfield", mapfield, "detail", "wrong 'dataservice'")
						}
						dataComp, ok := svc.(data.DataComponent)
						if !ok {
							return errors.BadConf(ctx, "fieldmappings", "Message", "Lookup data service incorrect", "mapfield", mapfield, "detail", "wrong 'dataservice'")
						}
						proc.lookups[mapField] = proc.getLookup(ctx, dataComp)
					}
				} else {
					return errors.BadConf(ctx, "fieldmappings", "Message", "Lookup data service incorrect", "mapfield", mapfield, "detail", "missing 'field'")
				}
			}
		}
	}
	return nil
}

func (proc *mapToObjectProcessor) Transform(ctx core.RequestContext, input interface{}) (interface{}, error) {
	inputMap := utils.CastToStringMap(input)

	if inputMap != nil {
		obj := proc.objectFactory.CreateObject(ctx)

		err := utils.SetObjectFields(ctx, obj, inputMap, proc.fieldmappings, proc.lookups)
		if err != nil {
			return nil, err
		}

		return obj, nil
	} else {
		return nil, errors.BadRequest(ctx, "Error", "Recieved a wrong input map", "input", input)
	}
	return nil, nil
}

func (proc *mapToObjectProcessor) getLookup(ctx core.ServerContext, dataComp data.DataComponent) utils.LookupFunc {
	return func(ctx interface{}, name string, val interface{}) (interface{}, error) {
		strVal := val.(string)
		if strVal == "" {
			return nil, nil
		}
		reqCtx := ctx.(core.RequestContext)
		var stor data.Storable
		var err error
		lookupfield, ok := proc.lookupfields
		if ok {
			cond := dataComp.CreateCondition(ctx, data.FIELDVALUE, lookupfield, strVal)
			stors, _,_, recs, err := dataComp.Get(ctx, cond, -1, -1, ", nil")
			if err == nil {
				if recs > 0 {
					stor = stors[0]
				} 
			}
		} else {
			stor, err = dataComp.GetById(reqCtx, strVal)
		}
		if err != nil {
			return nil, err
		}
		if stor == nil {
			return nil, fmt.Errorf("Lookup resource not found. Id : %s", strVal)
		}
		log.Error(reqCtx, "looked up object", "stor", stor, "err", err, "id", strVal)
		return &data.StorableRef{Id: stor.GetId(), Type: dataComp.GetObject(), Name: stor.GetLabel()}, nil
	}
}
