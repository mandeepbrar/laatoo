package data

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_REF_OPS                = "reference_operations"
	CONF_REF_OP                 = "operation"
	CONF_REF_TARG_SVC           = "targetsvc"
	CONF_REF_TARG_FIELD         = "targetfield"
	CONF_REF_CD_TARG_SOFTDELETE = "softdelete"
)

type refOperation struct {
	name          string
	targetsvcname string
	targetService data.DataComponent
	data          []interface{}
	do            func(ctx core.RequestContext, args ...interface{}) error
}

func buildRefOps(ctx core.ServerContext, conf config.Config) (deleterefops []*refOperation, updaterefops []*refOperation, saverefops []*refOperation, getrefops []*refOperation, err error) {
	deleteOps := []*refOperation{}
	updaterefOps := []*refOperation{}
	saverefOps := []*refOperation{}
	getrefOps := []*refOperation{}
	refOps, ok := conf.GetSubConfig(CONF_REF_OPS)
	if ok {
		opnames := refOps.AllConfigurations()
		for _, opname := range opnames {
			operConf, _ := refOps.GetSubConfig(opname)
			operation, ok := operConf.GetString(CONF_REF_OP)
			if !ok {
				return nil, nil, nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", CONF_REF_OP, "Operation", opname)
			}
			targetsvc, ok := operConf.GetString(CONF_REF_TARG_SVC)
			if !ok {
				return nil, nil, nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", CONF_REF_TARG_SVC, "Operation", opname)
			}
			switch operation {
			case "CascadedDelete":
				oper, err := buildCascadedDeleteOperation(ctx, conf, opname, targetsvc)
				if err != nil {
					return nil, nil, nil, nil, errors.WrapError(ctx, err)
				}
				deleteOps = append(deleteOps, oper)
			/*case "CompleteEntity":
			oper, err := buildCompleteEntityOperation(ctx, conf, opname, targetsvc)
			if err != nil {
				return nil, nil, nil, nil, errors.WrapError(ctx, err)
			}
			getrefOps = append(getrefOps, oper)*/
			case "Save":
			case "Update":

			}
		}
	}
	return deleteOps, updaterefOps, saverefOps, getrefOps, nil
}

func (oper *refOperation) Initialize(ctx core.ServerContext) error {
	svc, err := ctx.GetService(oper.targetsvcname)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	targetsvc, ok := svc.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF)
	}
	oper.targetService = targetsvc
	return nil
}

func buildCascadedDeleteOperation(ctx core.ServerContext, conf config.Config, opname string, targetsvcname string) (*refOperation, error) {
	targetfield, ok := conf.GetString(CONF_REF_TARG_FIELD)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", CONF_REF_TARG_FIELD, "Operation", opname)
	}
	opr := &refOperation{name: opname, targetsvcname: targetsvcname}
	opr.do = func(newctx core.RequestContext, args ...interface{}) error {
		ids := args[0].([]string)
		return cascadeDelete(newctx, opr.targetService, targetfield, ids)
	}
	return opr, nil
}

func cascadeDelete(ctx core.RequestContext, dataService data.DataComponent, targetfield string, ids []string) error {
	if dataService.Supports(data.InQueries) {
		condition, _ := dataService.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, targetfield, ids)
		_, err := dataService.DeleteAll(ctx, condition)
		return err
	} else {
		for _, id := range ids {
			condition, _ := dataService.CreateCondition(ctx, data.FIELDVALUE, targetfield, id)
			_, err := dataService.DeleteAll(ctx, condition)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func deleteRefOps(ctx core.RequestContext, opers []*refOperation, ids []string) error {
	if opers != nil {
		log.Logger.Trace(ctx, "deleterefops")
		for _, oper := range opers {
			log.Logger.Trace(ctx, "deleterefops", "oper", oper.name)
			err := oper.do(ctx, ids)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
func buildCompleteEntityOperation(ctx core.ServerContext, conf config.Config, opname string, targetsvcname string) (*refOperation, error) {
	targetfield, ok := conf.GetString(CONF_REF_TARG_FIELD)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", CONF_REF_TARG_FIELD, "Operation", opname)
	}
	opr := &refOperation{name: opname, targetsvcname: targetsvcname}
	opr.do = func(newctx core.RequestContext, args ...interface{}) error {
		entities := args[0].([]data.Storable)
		return completeEntity(newctx, opr.targetService, targetfield, entities)
	}
	return opr, nil
}

func completeEntity(ctx core.RequestContext, dataService data.DataService, targetfield string, entities []data.Storable) error {

	condition, _ := dataService.CreateCondition(ctx, data.FIELDVALUE, targetfield, id)
	_, err := dataService.DeleteAll(ctx, condition)
	if err != nil {
		return err
	}
	return nil
}*/
