package usecase

import (
	"time"

	"github.com/fgunawan1995/julo/wallet/dal/db"
	"github.com/fgunawan1995/julo/wallet/model"
)

func DepositBalance(p model.DepositRequest, customerID string) (model.Deposit, error) {
	var result model.Deposit
	dbDAL := db.New(model.GetDB())
	err := model.Validate.Struct(p)
	if err != nil {
		return result, err
	}
	tx, err := model.GetDB().Beginx()
	if err != nil {
		return result, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = dbDAL.UpdateWalletBalance(tx, model.BuildUpdateWalletBalance(customerID, p.Amount))
	if err != nil {
		return result, err
	}
	result, err = dbDAL.InsertDeposit(tx, model.BuildInsertDeposit(p, customerID, time.Now().UTC()))
	if err != nil {
		return result, err
	}
	err = tx.Commit()
	if err != nil {
		return result, err
	}
	return result, nil
}
