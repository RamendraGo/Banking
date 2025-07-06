package dto

import (
	"strings"

	"github.com/RamendraGo/Banking/errs"
)

type NewTransactionRequest struct {
	CustomerId      string  `json:"-"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {

	if r.Amount < 0 {
		return errs.NewValidateError("Transaction amount must be +ive number", 400)
	}

	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidateError("Transaction Type should be either withdrawal or deposit", 400)
	}

	return nil

}

func (r NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == "withdrawal"
}
