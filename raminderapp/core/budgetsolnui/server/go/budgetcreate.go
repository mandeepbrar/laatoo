package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type BudgetCreationService struct {
	core.Service
	GlAcctNormalItemSvc data.DataComponent
	BudgetSvc           data.DataComponent
	BudgetConfigSvc     data.DataComponent
}

func (svc *BudgetCreationService) Invoke(ctx core.RequestContext) error {
	formdata, _ := ctx.GetStringsMapParam("Data")
	log.Info(ctx, "Created budget", "data", formdata)
	budgetConf, ok := formdata["ConfigId"]
	if !ok {
		return errors.MissingArg(ctx, "ConfigId")
	}
	title, ok := formdata["Title"]
	if !ok {
		return errors.MissingArg(ctx, "Title")
	}
	year, ok := formdata["Year"]
	if !ok {
		return errors.MissingArg(ctx, "Year")
	}

	config, err := svc.BudgetConfigSvc.GetById(ctx, budgetConf)
	if err != nil {
		return err
	}
	budgetConfig := config.(*BudgetConfig)

	budgetInt, err := ctx.CreateObject("budgetsolnui.Budget")
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	budget := budgetInt.(*Budget)
	budget.Year = year
	budget.Title = title
	budget.GLAccounts = make([]*GLAccountLineItem, len(budgetConfig.GLAccounts))

	normalLineItems := make([]data.Storable, 0)

	log.Error(ctx, "adding gl accounts", "accts", budgetConfig.GLAccounts)
	for i, glacct := range budgetConfig.GLAccounts {
		log.Error(ctx, "adding gl account", "acct", glacct)
		glAcctItemInt, err := ctx.CreateObject("budgetsolnui.GLAccountLineItem")
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		glAcctItem := glAcctItemInt.(*GLAccountLineItem)
		glAcctItem.SetId(glacct.GetId())
		glAcctItem.Title = glacct.Title
		glAcctItem.Description = glacct.Description
		glAcctItem.Rollup = glacct.Rollup
		glAcctItem.Type = glacct.Type
		glAcctItem.Code = glacct.Code
		glAcctItem.Parent = glacct.Parent
		glAcctItem.Parent.Type = "budgetsolnui.GLAccountLineItem"
		switch glacct.Type {
			case "Employee": 
			{

			}
			case "Supplier": 
			{
				
			}
			case "Customer": 
			{
				
			}
			default: 
			{
				normalLiteItemInt, _ := ctx.CreateObject("budgetsolnui.GLNormalLineItem")
				normalLineItem := normalLiteItemInt.(*GLNormalLineItem)
				normalLineItem.Title = glacct.Title
				normalLineItem.Budget = budget.GetId()
				normalLineItem.LineItem = glAcctItem.GetId()
				normalLineItems = append(normalLineItems, normalLineItem)
			}
		}
		budget.GLAccounts[i] = glAcctItem
	}
	err = svc.BudgetSvc.Save(ctx, budget)
	if err != nil {
		return errors.WrapError(ctx, err)
	}


	err = svc.GlAcctNormalItemSvc.CreateMulti(ctx, normalLineItems)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}
