package service

import (
	"encoding/json"
	"time"

	"github.com/RamendraGo/Banking/domain"
	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/errs"
	"github.com/RamendraGo/Banking/logger"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {

	return DefaultAccountService{repo: repository}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()

	if err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	logger.Info("AccountService - Creating account for customer_id: " + a.CustomerId)

	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	respJson, _ := json.Marshal(response)
	logger.Info("AccountService - Created account response: " + string(respJson))

	return &response, nil

}
