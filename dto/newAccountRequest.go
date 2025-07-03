package dto

import (
	"strings"

	"github.com/RamendraGo/Banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {

	if r.Amount <= 5000 {
		return errs.NewValidateError("Amount must be greeater than 5000", 400)
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidateError("Account Type should be checking or saving", 400)
	}

	return nil

}
