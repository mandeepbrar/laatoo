package main

import (
	"fmt"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type SubAccountCreationService struct {
	core.Service
	GlaccountSvc data.DataComponent
}

func (svc *SubAccountCreationService) Invoke(ctx core.RequestContext) error {
	data, _ := ctx.GetStringMapParam("Data")
	log.Info(ctx, "Created sub accounts", "data", data)
	parentAcct := data["GLAcctId"]
	parentAcctName := data["GLAcctName"]
	log.Info(ctx, "Create sub accounts of parent", "parentAcct", parentAcct)
	createSubAcct := func(entrytype string, entityType string) error {
		selectedEntriesTypeInt, ok := data[entrytype]
		if ok {
			selectedEntriesIntArr, ok := selectedEntriesTypeInt.([]interface{})
			if ok {
				for _, entryInt := range selectedEntriesIntArr {
					entry, ok := entryInt.(map[string]interface{})
					if ok {
						entryId := entry["Id"]
						entryName := entry["Name"]
						subacctInt, _ := ctx.CreateObject("budgetsolnui.GLAccount")
						subacct := subacctInt.(*GLAccount)
						subacct.Title = fmt.Sprint(entryName)
						subacct.LinkedElement.Id = fmt.Sprint(entryId)
						subacct.LinkedElement.Name = fmt.Sprint(entryName)
						subacct.LinkedElement.Type = fmt.Sprint(entityType)
						subacct.Code = fmt.Sprint(entryName)
						subacct.Type = "Customer"
						subacct.Parent.Id = fmt.Sprint(parentAcct)
						subacct.Parent.Name = fmt.Sprint(parentAcctName)
						subacct.Parent.Type = fmt.Sprint("budgetsolnui.GLAccount")
						err := svc.GlaccountSvc.Save(ctx, subacct)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
						log.Info(ctx, "Create sub accounts of parent", "parentAcct", parentAcct, "subacct", subacct)
					}
				}
			}
		}
		return nil
	}

	err := createSubAcct("customers", "budgetsolnui.Customer")
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = createSubAcct("employees", "budgetsolnui.EmpClassification")
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = createSubAcct("suppliers", "budgetsolnui.Supplier")
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	/*budgetId, _ := ctx.GetStringParam("budgetId")
	log.Info(ctx, "Publishing budget", "budgetId", budgetId)
	stor, err := svc.BudgetSvc.GetById(ctx, budgetId)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	budgetToPublish := stor.(*Budget)
	*/

	/*queryCond := svc.glentrySvc.CreateCondition(ctx, data.FIELDVALUE, map[string]interface{}{"Budget": budgetId})
	count := svc.glentrySvc.
	allGlEntriesForBudget, err := svc.glentrySvc.Get(ctx, queryCond, 100, budgetId)

	ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	*/
	/*	if err := svc.createEntries(ctx, budgetToPublish); err != nil {
			return errors.WrapError(ctx, err)
		}

		log.Info(ctx, "Published budget", "budgetId", budgetId)
	*/
	return nil
}
