package domain

import (
	"strconv"
	"time"

	"github.com/RamendraGo/Banking/errs"
	"github.com/RamendraGo/Banking/logger"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) *TransactionRepositoryDb {

	logger.Info("Creating new TransactionRepositoryDb")
	return &TransactionRepositoryDb{client: dbClient}
}

func (d *TransactionRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {

	tx, err := d.client.Beginx()

	if err != nil {
		logger.Error("Error while initiating the new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error occurred while initiating the new transaction", 50)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Prepare params for account update
	accountIdInt, _ := strconv.Atoi(t.AccountId)
	updateParams := map[string]interface{}{
		"amount":     t.Amount,
		"account_id": accountIdInt,
	}

	// Choose SQL for withdrawal or deposit
	var sqlUpdate string
	if t.IsWithdrawal() {
		sqlUpdate = `
        UPDATE accounts
        SET amount = amount - :amount 
        OUTPUT INSERTED.account_id, INSERTED.amount
        WHERE account_id = :account_id`
	} else {
		sqlUpdate = `
        UPDATE accounts
        SET amount = amount + :amount 
        OUTPUT INSERTED.account_id, INSERTED.amount
        WHERE account_id = :account_id`
	}
	// Execute account update
	rows, err := tx.NamedQuery(sqlUpdate, updateParams)
	if err != nil {
		logger.Error("Error while updating account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error occurred while updating account", 50)
	}
	defer rows.Close()

	var updatedAccountId int
	var updatedAmount float64
	if rows.Next() {
		err = rows.Scan(&updatedAccountId, &updatedAmount)
		if err != nil {
			logger.Error("Error while scanning updated account: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while updating account", 50)
		}
	} else {
		return nil, errs.NewUnexpectedError("account not found or update failed", 50)
	}

	// Prepare params for transaction insert
	insertParams := map[string]interface{}{
		"account_id":       accountIdInt,
		"transaction_date": time.Now().Format("2006-01-02 15:04:05"),
		"transaction_type": t.TransactionType,
		"amount":           t.Amount,
	}

	// Insert transaction
	sqlInsert := `
    INSERT INTO transactions
        (account_id, amount, transaction_type, transaction_date)
        OUTPUT INSERTED.transaction_id
    VALUES
        (:account_id, :amount, :transaction_type, :transaction_date)`

	insertRows, err := tx.NamedQuery(sqlInsert, insertParams)
	if err != nil {
		logger.Error("Error while creating the transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error occurred while creating the transaction", 50)
	}
	defer insertRows.Close()

	var transactionId int64
	if insertRows.Next() {
		err = insertRows.Scan(&transactionId)
		if err != nil {
			logger.Error("Error while getting the Transaction ID: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while creating transaction", 50)
		}
	}

	t.TransactionId = int(transactionId)
	t.Amount = updatedAmount
	return &t, nil
}
