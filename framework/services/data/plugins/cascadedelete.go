package plugins

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

const (
	CONF_DELETE_OPERATION = "operations"
)

type cascadeDeleteOperation struct {
	targetSvcName string
	targetField   string
	targetSvc     data.DataComponent
}

type cascadeDelete struct {
	*data.DataPlugin
	ops []*cascadeDeleteOperation
}

func NewCascadeDeleteService(ctx core.ServerContext) *cascadeDelete {
	return &cascadeDelete{DataPlugin: data.NewDataPlugin(ctx)}
}

func (svc *cascadeDelete) Initialize(ctx core.ServerContext, conf config.Config) error {
	err := svc.DataPlugin.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	arr := make([]*cascadeDeleteOperation, 1)
	ops, ok := conf.GetSubConfig(CONF_DELETE_OPERATION)
	if ok {
		delops := ops.AllConfigurations()
		for _, delop := range delops {
			oper, _ := ops.GetSubConfig(delop)
			targetsvc, ok := oper.GetString(CONF_TARG_SVC)
			if !ok {
				return errors.MissingConf(ctx, CONF_TARG_SVC, "operation", delop)
			}

			targetfield, ok := oper.GetString(CONF_TARG_FIELD)
			if !ok {
				return errors.MissingConf(ctx, CONF_TARG_FIELD)
			}
			arr = append(arr, &cascadeDeleteOperation{targetSvcName: targetsvc, targetField: targetfield})
		}
	}
	svc.ops = arr

	return nil
}

func (svc *cascadeDelete) Start(ctx core.ServerContext) error {
	err := svc.DataPlugin.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, op := range svc.ops {
		s, err := ctx.GetService(op.targetSvcName)
		if err != nil {
			return errors.BadConf(ctx, CONF_TARG_SVC)
		}

		dataComponent, ok := s.(data.DataComponent)
		if !ok {
			return errors.BadConf(ctx, CONF_TARG_SVC)
		}
		op.targetSvc = dataComponent
	}
	return nil
}

func (svc *cascadeDelete) cascadeDelete(ctx core.RequestContext, ids []string) error {
	for _, op := range svc.ops {
		if op.targetSvc.Supports(data.InQueries) {
			condition, _ := op.targetSvc.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, op.targetField, ids)
			_, err := op.targetSvc.DeleteAll(ctx, condition)
			return err
		} else {
			for _, id := range ids {
				condition, _ := svc.DataComponent.CreateCondition(ctx, data.FIELDVALUE, map[string]interface{}{op.targetField: id})
				_, err := op.targetSvc.DeleteAll(ctx, condition)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (svc *cascadeDelete) Delete(ctx core.RequestContext, id string) error {
	err := svc.DataComponent.Delete(ctx, id)
	if err == nil {
		err = svc.cascadeDelete(ctx, []string{id})
	}
	return err
}

//Delete object by ids
func (svc *cascadeDelete) DeleteMulti(ctx core.RequestContext, ids []string) error {
	err := svc.DataComponent.DeleteMulti(ctx, ids)
	if err == nil {
		err = svc.cascadeDelete(ctx, ids)
	}
	return err
}

//Delete object by condition
func (svc *cascadeDelete) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ids, err := svc.DataComponent.DeleteAll(ctx, queryCond)
	if err == nil {
		err = svc.cascadeDelete(ctx, ids)
	}
	return ids, err
}
