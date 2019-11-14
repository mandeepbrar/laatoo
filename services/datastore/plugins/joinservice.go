package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
)

const (
	CONF_JOIN_OPERATION = "operations"
)

type joinOperation struct {
	targetSvcName string
	targetField   string
	targetSvc     data.DataComponent
}

type joinService struct {
	*data.DataPlugin
	ops []*joinOperation
}

func NewJoinService(ctx core.ServerContext) *joinService {
	return &joinService{DataPlugin: data.NewDataPlugin(ctx)}
}
func NewJoinServiceWithBase(ctx core.ServerContext, base data.DataComponent) *joinService {
	return &joinService{DataPlugin: data.NewDataPluginWithBase(ctx, base)}
}

func (svc *joinService) Initialize(ctx core.ServerContext) error {
	err := svc.DataPlugin.Initialize(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	svc.AddConfigurations(map[string]string{CONF_JOIN_OPERATION: config.CONF_OBJECT_CONFIG})
	return nil
}

func (svc *joinService) Start(ctx core.ServerContext) error {
	err := svc.DataPlugin.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	val, _ := svc.GetConfiguration(CONF_JOIN_OPERATION)
	arr := make([]*joinOperation, 1)
	ops := val.(config.Config)
	joinops := ops.AllConfigurations()
	for _, joinop := range joinops {
		oper, _ := ops.GetSubConfig(joinop)
		targetsvc, ok := oper.GetString(CONF_TARG_SVC)
		if !ok {
			return errors.MissingConf(ctx, CONF_TARG_SVC, "operation", joinop)
		}

		targetfield, _ := oper.GetString(CONF_TARG_FIELD)
		arr = append(arr, &joinOperation{targetSvcName: targetsvc, targetField: targetfield})
	}
	svc.ops = arr

	if svc.ops == nil {
		return nil
	}
	log.Error(ctx, "starting", "ops", svc.ops)
	for _, op := range svc.ops {
		log.Error(ctx, "starting", "op", op)
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

/*func (svc *joinService) Save(ctx core.RequestContext, item data.Storable) error {
	return svc.DataComponent.Save(ctx, item)
}*/

func (svc *joinService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	ctx = ctx.SubContext("Join_GetById")

	stor, err := svc.PluginDataComponent.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res, err := svc.fillJoin(ctx, []string{id}, []data.Storable{stor})
	if err != nil {
		return nil, err
	}

	return res[0], err
}

//Get multiple objects by id
func (svc *joinService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
	ctx = ctx.SubContext("Join_GetMulti")

	res, err := svc.PluginDataComponent.GetMulti(ctx, ids, orderBy)
	if err != nil {
		return nil, err
	}

	res, err = svc.fillJoin(ctx, ids, res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (svc *joinService) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
	ctx = ctx.SubContext("Join_GetMultiHash")

	res, err := svc.PluginDataComponent.GetMultiHash(ctx, ids)
	if err != nil {
		return nil, err
	}

	res, err = svc.fillMapWithJoin(ctx, ids, res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (svc *joinService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return svc.PluginDataComponent.GetList(ctx, pageSize, pageNum, mode, orderBy)
}

func (svc *joinService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy interface{}) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return svc.PluginDataComponent.Get(ctx, queryCond, pageSize, pageNum, mode, orderBy)
}

func (svc *joinService) fillJoin(ctx core.RequestContext, ids []string, inputData []data.Storable) ([]data.Storable, error) {
	for _, op := range svc.ops {
		if op.targetField != "" {
			fieldVals := make([]string, len(inputData))
			for ind, item := range inputData {
				iVal := reflect.ValueOf(item)
				field := iVal.FieldByName(op.targetField)
				if field.IsValid() {
					fieldVals[ind] = field.String()
				} else {
					return nil, errors.BadConf(ctx, CONF_TARG_FIELD)
				}
			}
			ids = fieldVals
		}
		hash, err := op.targetSvc.GetMultiHash(ctx, ids)
		if err != nil {
			return nil, err
		}
		for _, stor := range inputData {
			id := stor.GetId()
			joinedItem, ok := hash[id]
			if ok {
				log.Info(ctx, "Joining item", "item", joinedItem)
				stor.Join(joinedItem)
			}
		}
	}
	return inputData, nil
}

func (svc *joinService) fillMapWithJoin(ctx core.RequestContext, ids []string, inputData map[string]data.Storable) (map[string]data.Storable, error) {
	for _, op := range svc.ops {
		hash, err := op.targetSvc.GetMultiHash(ctx, ids)
		if err != nil {
			return nil, err
		}
		for id, stor := range inputData {
			joinedItem, ok := hash[id]
			if ok {
				stor.Join(joinedItem)
			}
		}
	}
	return inputData, nil
}
