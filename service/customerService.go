package service

import (
	"github.com/RamendraGo/Banking/domain"
	"github.com/RamendraGo/Banking/errs"
)

// CustomerService defines the methods that a customer service should implement.
// It includes all Customer methods.
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService implements the CustomerService interface
// It provides methods to interact with the customer repository.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer retrieves all customers from the repository
// It returns a slice of Customer and an error if any occurs during the retrieval.
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	// Call the repository to get all customers

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

// GetCustomer retrieves a customer by ID from the repository
// It returns a pointer to a Customer and an error if any occurs during the retrieval.
func (s DefaultCustomerService) GetCustomerById(CustomerId string) (*domain.Customer, *errs.AppError) {
	// Call the repository to get all customers
	return s.repo.GetCustomerById(CustomerId)
}

// NewCustomerService creates a new instance of DefaultCustomerService
// It takes a CustomerRepository as an argument and returns a DefaultCustomerService.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
