package db

import (
	"github.com/fgunawan1995/julo/wallet/model"
	"github.com/jmoiron/sqlx"
)

func (dal *DAL) InsertWithdrawal(tx *sqlx.Tx, p model.InsertWithdrawal) (model.Withdrawal, error) {
	var result model.Withdrawal
	err := tx.Get(&result, `
		INSERT INTO withdrawal(withdrawn_by,reference_id,status,amount,withdrawn_at) 
		values($1,$2,$3,$4,$5)
		RETURNING *`, p.WithdrawnBy, p.ReferenceID, p.Status, p.Amount, p.WithdrawnAt)
	return result, err
}
