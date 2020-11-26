package usecase

import (
	"time"

	"github.com/fgunawan1995/julo/dal/cache"
	"github.com/fgunawan1995/julo/dal/db"
	"github.com/fgunawan1995/julo/model"
)

func InitWallet(p model.InitWalletRequest) (model.InitWalletResponse, error) {
	var result model.InitWalletResponse
	dbDAL := db.New(model.GetDB())
	cacheDAL := cache.New(model.GetCache())
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
	_, err = dbDAL.InsertWallet(tx, model.BuildInsertWallet(p.CustomerXID))
	if err != nil {
		return result, err
	}
	result.Token, err = cacheDAL.SetTokenForUser(p.CustomerXID)
	if err != nil {
		return result, err
	}
	err = tx.Commit()
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetWallet(customerID string) (model.Wallet, error) {
	dbDAL := db.New(model.GetDB())
	return dbDAL.GetWallet(customerID)
}

func EnableWallet(customerID string) (model.Wallet, error) {
	var result model.Wallet
	dbDAL := db.New(model.GetDB())
	tx, err := model.GetDB().Beginx()
	if err != nil {
		return result, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	result, err = dbDAL.UpdateWalletStatus(tx, model.BuildUpdateWalletStatusEnabled(customerID, time.Now().UTC()))
	if err != nil {
		return result, err
	}
	err = tx.Commit()
	if err != nil {
		return result, err
	}
	return result, nil
}

func DisableWallet(customerID string) (model.Wallet, error) {
	var result model.Wallet
	dbDAL := db.New(model.GetDB())
	tx, err := model.GetDB().Beginx()
	if err != nil {
		return result, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	result, err = dbDAL.UpdateWalletStatus(tx, model.BuildUpdateWalletStatusDisabled(customerID, time.Now().UTC()))
	if err != nil {
		return result, err
	}
	err = tx.Commit()
	if err != nil {
		return result, err
	}
	return result, nil
}
