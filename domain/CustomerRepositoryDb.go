package domain

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/RamendraGo/Banking/errs"
	"github.com/RamendraGo/Banking/logger"
)

type CustomerRepositoryDb struct {
}

func (r *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "SELECT customer_id, name,  date_of_birth, city, zipcode, status FROM customers"
		rows, err = DB.Query(findAllSql)
	} else {

		findAllSql := "SELECT customer_id, name,  date_of_birth, city, zipcode, status FROM customers WHERE status = @p1"
		rows, err = DB.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error executing query: " + err.Error())

		return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customers", 50)
	}
	defer rows.Close()

	customers := make([]Customer, 0) // Preallocate slice with capacity of 10

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CustomerId, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			logger.Error("Error scanning row: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customers", 50)
		}
		customers = append(customers, c)
	}
	if err = rows.Err(); err != nil {
		logger.Error("Error iterating over rows: " + err.Error())
		return nil, errs.NewNotFoundError("Customer not found", 404)
	}
	logger.Info("Found " + strconv.Itoa(len(customers)) + " customers")
	return customers, nil

}

func (r *CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppError) {
	logger.Info("Getting customer by ID: " + id)
	findByIdSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = @p1"
	row := DB.QueryRow(findByIdSql, sql.Named("p1", id))

	if row == nil {
		log.Println("No customer found with ID:", id)
		return nil, nil // Return nil if no customer found
	}
	logger.Info("Query: " + findByIdSql + "with ID: " + id)
	logger.Info("Query executed successfully")

	var c Customer
	err := row.Scan(&c.CustomerId, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)

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
	return &c, nil
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {

	// Connect to the database
	Connect()
	if !DBConnected {
		logger.Fatal("Database connection failed. Exiting." + sql.ErrConnDone.Error())
	}

	logger.Info("Creating new CustomerRepositoryDb")
	return &CustomerRepositoryDb{}
}
