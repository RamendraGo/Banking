package domain

import (
	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/errs"
)

type Customer struct {
	CustomerId  string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"full_name" xml:"full_name" db:"name"`
	City        string `json:"city" xml:"city" db:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode" db:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status" db:"status"`
}

func (customer Customer) StatusAsText() string {
	var StatusAsText = "active"

	if customer.Status == "0" {
		StatusAsText = "inactive"
	}

	return StatusAsText

}

func (customer Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		CustomerId:  customer.CustomerId,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.StatusAsText(),
	}

}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	GetCustomerById(id string) (*Customer, *errs.AppError)
}
