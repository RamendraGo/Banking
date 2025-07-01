package domain

import (
	"log"
)

type CustomerRepositoryDb struct {
}

func (r *CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT customer_id, name,  date_of_birth, city, zipcode, status FROM customers"
	rows, err := DB.Query(findAllSql)
	if err != nil {
		log.Println("Error executing query:", err.Error())
		return nil, err
	}
	defer rows.Close()

	customers := make([]Customer, 0) // Preallocate slice with capacity of 10

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CustomerId, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error scanning row:", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err.Error())
		return nil, err
	}
	log.Println("Found", len(customers), "customers")
	return customers, nil

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
