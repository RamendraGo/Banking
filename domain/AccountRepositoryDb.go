package domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RamendraGo/Banking/errs"
	"github.com/RamendraGo/Banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d *AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := `
INSERT INTO dbo.accounts (customer_id, opening_date, account_type, amount, status)
OUTPUT INSERTED.account_id
VALUES (:customer_id, :opening_date, :account_type, :amount, :status)
`

	logger.Info("Creating account for customer_id: " + a.CustomerId)

	customerIdInt, _ := strconv.Atoi(a.CustomerId)

	fmt.Printf("Creating account for customer_id: %d\n ", customerIdInt)

	params := map[string]interface{}{
		"customer_id":  customerIdInt,
		"opening_date": time.Now().Format("2006-01-02 15:04:05"),
		"account_type": a.AccountType,
		"amount":       a.Amount,
		"status":       a.Status,
	}

	var id int64
	rows, err := d.client.NamedQuery(sqlInsert, params)
	if err != nil {
		logger.Error("Error while creating the account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error occurred while creating account", 50)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			logger.Error("Error while getting the Account ID: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while creating account", 50)
		}
	}

	a.AccountId = int(id)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) *AccountRepositoryDb {

	logger.Info("Creating new AccountRepositoryDb")
	return &AccountRepositoryDb{client: dbClient}
}
