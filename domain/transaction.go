package domain

import (
	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/errs"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   int     `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

type TransactionRepository interface {
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		NewBalance:      t.Amount,
		AccountId:       t.AccountId,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
