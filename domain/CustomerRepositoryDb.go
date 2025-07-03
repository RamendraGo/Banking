package domain

import (
	"database/sql"

	"github.com/RamendraGo/Banking/errs"
	"github.com/RamendraGo/Banking/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

var err error

func (d *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	customers := make([]Customer, 0) // Preallocate slice with capacity of 10

	if status == "" {
		findAllSql := "SELECT customer_id, name,  date_of_birth, city, zipcode, status FROM customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name,  date_of_birth, city, zipcode, status FROM customers WHERE status = @p1"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while query: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customers", 50)
	}
	return customers, nil
}

func (d *CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppError) {
	var customer Customer

	logger.Info("Getting customer by ID: " + id)
	findByIdSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = @p1"
	err = d.client.Get(&customer, findByIdSql, id)

	if err == sql.ErrNoRows {
		logger.Error("No customer found with ID: " + id)
		return nil, errs.NewNotFoundError("Customer not found", 404)
	} else {

		if err != nil {
			logger.Error("Error scanning row: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customer", 50)
		}
	}

	logger.Info("Found customer with ID:" + id)
	return &customer, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) *CustomerRepositoryDb {

	logger.Info("Creating new CustomerRepositoryDb")
	return &CustomerRepositoryDb{client: dbClient}
}
