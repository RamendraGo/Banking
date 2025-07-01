package domain

import (
	"database/sql"
	"log"

	"github.com/RamendraGo/Banking/errs"
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
		log.Println("Error executing query:", err.Error())
		return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customers", 50)
	}
	defer rows.Close()

	customers := make([]Customer, 0) // Preallocate slice with capacity of 10

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CustomerId, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error scanning row:", err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customers", 50)
		}
		customers = append(customers, c)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err.Error())
		return nil, errs.NewNotFoundError("Customer not found", 404)
	}
	log.Println("Found", len(customers), "customers")
	return customers, nil

}

func (r *CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppError) {
	log.Println("Getting customer by ID:", id)
	findByIdSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = @p1"
	row := DB.QueryRow(findByIdSql, sql.Named("p1", id))

	if row == nil {
		log.Println("No customer found with ID:", id)
		return nil, nil // Return nil if no customer found
	}
	log.Println("Query:", findByIdSql, "with ID:", id)
	log.Println("Query executed successfully")

	var c Customer
	err := row.Scan(&c.CustomerId, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)

	if err == sql.ErrNoRows {
		log.Println("No customer found with ID:", id)
		return nil, errs.NewNotFoundError("Customer not found", 404)
	} else {

		if err != nil {
			log.Println("Error scanning row:", err.Error())
			return nil, errs.NewUnexpectedError("unexpected error occurred while retrieving customer", 50)

		}
	}

	log.Println("Found customer with ID:", id)
	return &c, nil
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {

	// Connect to the database
	Connect()
	if !DBConnected {
		log.Fatal("Database connection failed. Exiting.")
	}

	log.Println("Creating new CustomerRepositoryDb")
	return &CustomerRepositoryDb{}
}
