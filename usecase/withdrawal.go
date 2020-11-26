package usecase

import (
	"time"

	"github.com/fgunawan1995/julo/wallet/dal/db"
	"github.com/fgunawan1995/julo/wallet/model"
)

func WithdrawalBalance(p model.WithdrawalRequest, customerID string) (model.Withdrawal, error) {
	var result model.Withdrawal
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
	_, err = dbDAL.UpdateWalletBalance(tx, model.BuildUpdateWalletBalance(customerID, p.Amount*-1))
	if err != nil {
		return result, err
	}
	result, err = dbDAL.InsertWithdrawal(tx, model.BuildInsertWithdrawal(p, customerID, time.Now().UTC()))
	if err != nil {
		return result, err
	}
	err = tx.Commit()
	if err != nil {
		return result, err
	}
	return result, nil
}
