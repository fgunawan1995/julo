package db

import (
	"github.com/fgunawan1995/julo/wallet/model"
	"github.com/jmoiron/sqlx"
)

func (dal *DAL) InsertWallet(tx *sqlx.Tx, p model.InsertWallet) (model.Wallet, error) {
	var result model.Wallet
	err := tx.Get(&result, `
		INSERT INTO wallet(owned_by,status) 
		values($1,$2)
		RETURNING *`, p.OwnedBy, p.Status)
	return result, err
}

func (dal *DAL) UpdateWalletStatus(tx *sqlx.Tx, p model.UpdateWalletStatus) (model.Wallet, error) {
	var result model.Wallet
	err := tx.Get(&result, `
	UPDATE wallet SET
		status = $1,
		enabled_at = $2,
		disabled_at = $3
	WHERE 
		owned_by=$4
	RETURNING *`, p.Status, p.EnabledAt, p.DisabledAt, p.CustomerID)
	return result, err
}

func (dal *DAL) UpdateWalletBalance(tx *sqlx.Tx, p model.UpdateWalletBalance) (model.Wallet, error) {
	var result model.Wallet
	err := tx.Get(&result, `
	UPDATE wallet SET
		balance = balance + $1
	WHERE 
		owned_by=$2
	RETURNING *`, p.Increment, p.CustomerID)
	return result, err
}

func (dal *DAL) GetWallet(customerID string) (model.Wallet, error) {
	var result model.Wallet
	err := dal.DB.Get(&result, `
		SELECT * FROM wallet 
		WHERE 
			owned_by = $1`, customerID)
	return result, err
}
