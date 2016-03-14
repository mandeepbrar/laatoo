package laatoodata

import (
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	CONF_REF_OPS                = "reference_object_operations"
	CONF_REF_OP                 = "operation"
	CONF_REF_PRIMARY_OBJ        = "primaryobj"
	CONF_REF_PRIMARY_OBJ_METHOD = "method"
	CONF_REF_CD_TARG_OBJ        = "target"
	CONF_REF_CD_TARG_FIELD      = "targetfield"
	CONF_REF_CD_TARG_SOFTDELETE = "softdelete"
)

type refKeyOperation struct {
	operation           string
	primaryObjectName   string
	primaryObjectMethod string
	operationData       map[string]interface{}
	do                  func(ctx core.Context, args ...interface{}) error
}

func buildRefOps(ctx core.Context, conf map[string]interface{}) (map[string][]*refKeyOperation, map[string][]*refKeyOperation, map[string][]*refKeyOperation, map[string][]*refKeyOperation, error) {
	deleteOps := map[string][]*refKeyOperation{}
	refOpsInt, ok := conf[CONF_REF_OPS]
	if ok {
		refOps, _ := refOpsInt.(map[string]interface{})
		if refOps != nil {
			for k, v := range refOps {
				operConf, _ := v.(map[string]interface{})
				methodInt, ok := operConf[CONF_REF_PRIMARY_OBJ_METHOD]
				if !ok {
					return nil, nil, nil, nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", CONF_REF_PRIMARY_OBJ_METHOD, "Operation", k)
				}
				if ok {
					method := methodInt.(string)
					objInt, ok := operConf[CONF_REF_PRIMARY_OBJ]
					if !ok {
						return nil, nil, nil, nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", CONF_REF_PRIMARY_OBJ, "Operation", k)
					}
					if objInt != "" {
						obj := objInt.(string)
						operInt, ok := operConf[CONF_REF_OP]
						if !ok {
							return nil, nil, nil, nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", CONF_REF_OP, "Operation", k)
						}
						if ok {
							oper := operInt.(string)
							opr := &refKeyOperation{operation: oper, primaryObjectName: obj, primaryObjectMethod: method, operationData: operConf}
							function, err := opr.getOperMethod(ctx)
							if err != nil {
								return nil, nil, nil, nil, err
							}
							opr.do = function
							switch method {
							case "Delete":
								opsArr, ok := deleteOps[obj]
								if !ok {
									opsArr = make([]*refKeyOperation, 0, 10)
								}
								deleteOps[obj] = append(opsArr, opr)
							case "Get":
							case "Put":
							case "Update":

							}
						}
					}
				}
			}
		}
	}
	return deleteOps, nil, nil, nil, nil
}

func (oper *refKeyOperation) getOperMethod(ctx core.Context) (func(ctx core.Context, args ...interface{}) error, error) {
	switch oper.operation {
	case "CascadedDelete":
		{
			targetobjInt, ok := oper.operationData[CONF_REF_CD_TARG_OBJ]
			if !ok {
				return nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", CONF_REF_CD_TARG_OBJ, "Operation", "CascadedDelete")
			}
			targetfieldInt, ok := oper.operationData[CONF_REF_CD_TARG_FIELD]
			if !ok {
				return nil, errors.ThrowErrorInCtx(ctx, LOGGING_CONTEXT, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", CONF_REF_CD_TARG_FIELD, "Operation", "CascadedDelete")
			}
			targetSoftDelete := false
			softDelete, ok := oper.operationData[CONF_REF_CD_TARG_SOFTDELETE]
			if ok {
				targetSoftDelete = (softDelete.(string) == "true")
			}
			return func(newctx core.Context, args ...interface{}) error {
				ds := args[0].(data.DataService)
				ids := args[1].([]string)
				return oper.cascadeDelete(newctx, ds, targetobjInt.(string), targetfieldInt.(string), ids, targetSoftDelete)
			}, nil
		}
	}
	return nil, nil
}

func (oper *refKeyOperation) cascadeDelete(ctx core.Context, dataService data.DataService, targetobj string, targetfield string, ids []string, softDelete bool) error {
	if dataService.Supports(data.InQueries) {
		condition, _ := dataService.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, targetfield, ids)
		_, err := dataService.DeleteAll(ctx, targetobj, condition, softDelete)
		return err
	} else {
		for _, id := range ids {
			condition, _ := dataService.CreateCondition(ctx, data.FIELDVALUE, targetfield, id)
			_, err := dataService.DeleteAll(ctx, targetobj, condition, softDelete)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func deleteRefOps(ctx core.Context, svc data.DataService, refops map[string][]*refKeyOperation, objectType string, ids []string) error {
	opers := refops[objectType]
	if opers != nil {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "deleterefops", "refops", refops, "opers", opers)
		for _, oper := range opers {
			log.Logger.Trace(ctx, LOGGING_CONTEXT, "deleterefops", "oper", oper)
			//only cascaded delete ref ops supported for delete
			err := oper.do(ctx, svc, ids)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
