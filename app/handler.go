package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	Age     int    `json:"age" xml:"age"`
	Address string `json:"address" xml:"address"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	Customers := []Customer{
		{Name: "Alice", Age: 30, Address: "123 Main St"},
		{Name: "Bob", Age: 25, Address: "456 Elm St"},
		{Name: "Charlie", Age: 35, Address: "789 Oak St"},
		{Name: "Diana", Age: 28, Address: "321 Maple Ave"},
		{Name: "Ethan", Age: 40, Address: "654 Pine Rd"},
		{Name: "Fiona", Age: 22, Address: "987 Cedar Blvd"},
		{Name: "George", Age: 45, Address: "159 Spruce St"},
		{Name: "Hannah", Age: 33, Address: "753 Birch Ln"},
		{Name: "Ian", Age: 29, Address: "852 Willow Dr"},
	}

	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customers)
	}

}
