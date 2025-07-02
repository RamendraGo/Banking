package service

import (
	"github.com/RamendraGo/Banking/domain"
	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/errs"
)

// CustomerService defines the methods that a customer service should implement.
// It includes all Customer methods.
type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService implements the CustomerService interface
// It provides methods to interact with the customer repository.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer retrieves all customers from the repository
// It returns a slice of Customer and an error if any occurs during the retrieval.
func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	// Call the repository to get all customers

	customer, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	responses := make([]dto.CustomerResponse, len(customer))
	for i, c := range customer {
		responses[i] = c.ToDto() // assumes domain.Customer has a ToDto() method returning dto.CustomerResponse
	}
	return responses, nil

}

// GetCustomer retrieves a customer by ID from the repository
// It returns a pointer to a Customer and an error if any occurs during the retrieval.
func (s DefaultCustomerService) GetCustomerById(CustomerId string) (*dto.CustomerResponse, *errs.AppError) {
	// Call the repository to get all customers

	customer, err := s.repo.GetCustomerById(CustomerId)

	if err != nil {
		return nil, err
	}

	response := customer.ToDto()
	return &response, nil

}

// NewCustomerService creates a new instance of DefaultCustomerService
// It takes a CustomerRepository as an argument and returns a DefaultCustomerService.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
