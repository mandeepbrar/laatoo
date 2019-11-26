package main

/*
import (
	"fmt"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
)

type importGLService struct {
	core.Service
	glAccountSvc data.DataComponent
}

func (svc *importGLService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *importGLService) importAccount(ctx core.RequestContext, accdata map[string]interface{}, acct *GLAccount) error {
	parentAcctId, ok := accdata["parent"]
	var parentAcct *GLAccount
	if ok {
		cond, err := svc.glAccountSvc.CreateCondition(ctx, data.FIELDVALUE, map[string]interface{}{"Code": parentAcctId})
		if err != nil {
			return err
		}
		stordata, _, _, recs, err := svc.glAccountSvc.Get(ctx, cond, -1, -1, "", nil)
		if err != nil {
			return err
		}
		if recs > 0 {
			parentAcct = stordata[0].(*GLAccount)
		}
	}
	accId, ok := accdata["Id"]
	if !ok {
		return errors.MissingArg(ctx, "Id")
	}
	accDesc, _ := accdata["Name"]

	accToCreate := &GLAccount{Code: accId.(string), Description: "accDesc", Title: fmt.Sprintf("%s %s", accId, accDesc)}
	if parentAcct != nil {
		accToCreate.Parent = GLAccount_Ref{Id: parentAcct.GetId(), Name: parentAcct.Title}
	}
	err := svc.glAccountSvc.Save(ctx, accToCreate)
	return err
}
*/
