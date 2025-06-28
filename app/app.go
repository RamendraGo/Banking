package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	// Set up the HTTP server and route
	http.HandleFunc("/", greet)
	http.HandleFunc("/customers", getCustomers)
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
