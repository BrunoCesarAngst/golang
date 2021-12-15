package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{
		db: db,
	}
}

func (t *TransactionRepositoryDb) Insert(id string, accountId string, amount float64, status string, errorMassage string) error {
	stmt, err := t.db.Prepare(`
	INSERT INTO transactions (id, account_id, amount, status, error_message, create_at, update_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
		accountId,
		amount,
		status,
		errorMassage,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
