package db

import (
	"github.com/fgunawan1995/julo/model"
	"github.com/jmoiron/sqlx"
)

func (dal *DAL) InsertDeposit(tx *sqlx.Tx, p model.InsertDeposit) (model.Deposit, error) {
	var result model.Deposit
	err := tx.Get(&result, `
		INSERT INTO deposit(deposited_by,reference_id,status,amount,deposited_at) 
		values($1,$2,$3,$4,$5)
		RETURNING *`, p.DepositedBy, p.ReferenceID, p.Status, p.Amount, p.DepositedAt)
	return result, err
}
