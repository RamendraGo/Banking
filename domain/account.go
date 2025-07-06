package domain

import (
	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/errs"
)

type Account struct {
	AccountId   int    `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
	CanWithdraw(accountId string, amount float64) (bool, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}
