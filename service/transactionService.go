package service

import (
	"encoding/json"
	"time"

	"github.com/RamendraGo/Banking/domain"
	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/errs"
	"github.com/RamendraGo/Banking/logger"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo        domain.TransactionRepository
	accountRepo domain.AccountRepository
}

func NewTransactionService(tr domain.TransactionRepository, ar domain.AccountRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: tr, accountRepo: ar}
}

func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()

	if err != nil {
		return nil, err
	}

	//Server side validation for checking the available balance

	if req.IsTransactionTypeWithdrawal() {

		if ok, err := s.accountRepo.CanWithdraw(req.AccountId, req.Amount); !ok {
			if err != nil {
				return nil, err
			}
			return nil, errs.NewValidateError("Insufficient Balance in the Account", 50)
		}

	}
	// ...inside your function...
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	logger.Info("AccountService - Creating transaction for Account Id: " + t.AccountId)

	newTransaction, err := s.repo.SaveTransaction(t)

	if err != nil {
		return nil, err
	}

	response := newTransaction.ToNewTransactionResponseDto()

	respJson, _ := json.Marshal(response)
	logger.Info("AccountService - Created account response: " + string(respJson))

	return &response, nil

}
