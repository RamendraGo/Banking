package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll implements CustomerRepository.
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {

	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "1", Name: "Alice", City: "Wonderland", Zipcode: "12345", DateOfBirth: "1990-01-01", Status: "active"},
		{ID: "2", Name: "Bob", City: "Builderland", Zipcode: "54321", DateOfBirth: "1985-05-05", Status: "inactive"},
	}
	return CustomerRepositoryStub{customers: customers}
}
